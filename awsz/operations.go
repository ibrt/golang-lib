package awsz

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cf "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	cft "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	awsecr "github.com/aws/aws-sdk-go-v2/service/ecr"
	awskms "github.com/aws/aws-sdk-go-v2/service/kms"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
	"github.com/cheggaaa/pb/v3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/rdegges/go-ipify"

	"github.com/ibrt/golang-lib/cfz"
	"github.com/ibrt/golang-lib/consolez"
	"github.com/ibrt/golang-lib/errorz"
	"github.com/ibrt/golang-lib/filez"
	"github.com/ibrt/golang-lib/memz"
	"github.com/ibrt/golang-lib/shellz"
)

// Operations provides AWS utils.
type Operations struct {
	appPrefix             string
	stagePrefix           string
	awsRegion             string
	namespace             string
	tplUtils              *cfz.TemplateUtils
	awsCF                 *cf.Client
	awsEC2                *awsec2.Client
	awsECR                *awsecr.Client
	awsKMS                *awskms.Client
	awsS3                 *awss3.Client
	isDockerAuthenticated bool
}

// MustNewOperations initializes a new Operations.
func MustNewOperations(appPrefix, stagePrefix, awsRegion string) *Operations {
	namespace := strings.TrimSuffix(fmt.Sprintf("%v-%v", appPrefix, stagePrefix), "-")

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(awsRegion),
		config.WithSharedConfigProfile(appPrefix))
	errorz.MaybeMustWrap(err)

	return &Operations{
		appPrefix:             appPrefix,
		stagePrefix:           stagePrefix,
		awsRegion:             awsRegion,
		namespace:             namespace,
		tplUtils:              cfz.NewTemplateUtils(namespace),
		awsCF:                 cf.NewFromConfig(cfg),
		awsEC2:                awsec2.NewFromConfig(cfg),
		awsECR:                awsecr.NewFromConfig(cfg),
		awsKMS:                awskms.NewFromConfig(cfg),
		awsS3:                 awss3.NewFromConfig(cfg),
		isDockerAuthenticated: false,
	}
}

// GetAppPrefix returns the app prefix.
func (o *Operations) GetAppPrefix() string {
	return o.appPrefix
}

// GetStagePrefix gets the stage prefix.
func (o *Operations) GetStagePrefix() string {
	return o.stagePrefix
}

// GetAWSRegion returns the AWS region.
func (o *Operations) GetAWSRegion() string {
	return o.awsRegion
}

// GetNamespace returns the namespace.
func (o *Operations) GetNamespace() string {
	return o.namespace
}

// GetTemplateUtils returns the *cfz.TemplateUtils.
func (o *Operations) GetTemplateUtils() *cfz.TemplateUtils {
	return o.tplUtils
}

// MustUploadS3 uploads a file to S3.
func (o *Operations) MustUploadS3(bucketName, key, inFilePath string) {
	consolez.DefaultCLI.Notice("aws-operations", "uploading to S3...", bucketName, key, inFilePath)

	fullContentType, err := mimetype.DetectFile(inFilePath)
	errorz.MaybeMustWrap(err)

	buf := filez.MustReadFile(inFilePath)

	bar := pb.New(len(buf)).
		SetTemplate(pb.Full).
		Set(pb.Bytes, true).
		SetRefreshRate(50 * time.Millisecond).
		Start()
	r := bar.NewProxyReader(bytes.NewReader(buf))
	defer bar.Finish()

	_, err = o.awsS3.PutObject(context.Background(), &awss3.PutObjectInput{
		Bucket:        memz.Ptr(bucketName),
		Key:           memz.Ptr(key),
		Body:          r,
		ContentLength: memz.Ptr(int64(len(buf))),
		ContentType:   memz.Ptr(fullContentType.String()),
	})
	errorz.MaybeMustWrap(err)
}

// MustDownloadS3 downloads a file from S3.
func (o *Operations) MustDownloadS3(bucketName, key, outFilePath string) {
	consolez.DefaultCLI.Notice("aws-operations", "downloading from S3...", bucketName, key, outFilePath)

	obj, err := o.awsS3.GetObject(context.Background(), &awss3.GetObjectInput{
		Bucket: memz.Ptr(bucketName),
		Key:    memz.Ptr(key),
	})
	errorz.MaybeMustWrap(err)
	defer errorz.IgnoreClose(obj.Body)

	errorz.MaybeMustWrap(os.MkdirAll(filepath.Dir(outFilePath), 0777))

	fd, err := os.Create(outFilePath)
	errorz.MaybeMustWrap(err)
	defer errorz.IgnoreClose(fd)

	bar := pb.New64(*obj.ContentLength).
		SetTemplate(pb.Full).
		Set(pb.Bytes, true).
		SetRefreshRate(50 * time.Millisecond).
		Start()
	r := bar.NewProxyReader(obj.Body)
	defer bar.Finish()

	_, err = io.Copy(fd, r)
	errorz.MaybeMustWrap(err)
}

// MustDecryptKMS decrypts some data using a KMS key.
func (o *Operations) MustDecryptKMS(keyID string, ciphertext []byte) []byte {
	consolez.DefaultCLI.Notice("aws-operations", "decrypting with KMS...", keyID)

	resp, err := o.awsKMS.Decrypt(context.Background(), &awskms.DecryptInput{
		KeyId:          memz.Ptr(keyID),
		CiphertextBlob: ciphertext,
	})
	errorz.MaybeMustWrap(err)
	return resp.Plaintext
}

// MustEncryptKMS encrypts some data using a KMS key.
func (o *Operations) MustEncryptKMS(keyID string, plaintext []byte) []byte {
	consolez.DefaultCLI.Notice("aws-operations", "encrypting with KMS...", keyID)

	resp, err := o.awsKMS.Encrypt(context.Background(), &awskms.EncryptInput{
		KeyId:     memz.Ptr(keyID),
		Plaintext: plaintext,
	})
	errorz.MaybeMustWrap(err)
	return resp.CiphertextBlob
}

// MustAuthorizeSecurityGroupIngress authorizes ingress on the given security group and port from the current public IP.
func (o *Operations) MustAuthorizeSecurityGroupIngress(securityGroupID string, port uint32) {
	consolez.DefaultCLI.Notice("aws-operations", "authorizing EC2 security group ingress...", fmt.Sprintf("%v", port))

	ip, err := ipify.GetIp()
	errorz.MaybeMustWrap(err)

	_, err = o.awsEC2.AuthorizeSecurityGroupIngress(context.Background(), &awsec2.AuthorizeSecurityGroupIngressInput{
		CidrIp:     memz.Ptr(fmt.Sprintf("%v/32", ip)),
		FromPort:   memz.Ptr(int32(port)),
		GroupId:    memz.Ptr(securityGroupID),
		IpProtocol: memz.Ptr("tcp"),
		ToPort:     memz.Ptr(int32(port)),
	})

	if err != nil {
		if eErr := (smithy.APIError)(nil); !errors.As(err, &eErr) || eErr.ErrorCode() != "InvalidPermission.Duplicate" {
			errorz.MustWrap(err)
		}
	}
}

// MustCreateStack creates a CloudFormation stack.
func (o *Operations) MustCreateStack(name string, templateBody string, tagsMap map[string]string) *cft.Stack {
	consolez.DefaultCLI.Notice("aws-operations", "creating CloudFormation stack...", name)

	_, err := o.awsCF.CreateStack(context.Background(), &cf.CreateStackInput{
		Capabilities: []cft.Capability{
			cft.CapabilityCapabilityIam,
			cft.CapabilityCapabilityNamedIam,
		},
		EnableTerminationProtection: aws.Bool(false),
		OnFailure:                   cft.OnFailureRollback,
		StackName:                   memz.Ptr(name),
		Tags: func() []cft.Tag {
			tags := make([]cft.Tag, 0)
			for k, v := range tagsMap {
				tags = append(tags, cft.Tag{
					Key:   memz.Ptr(k),
					Value: memz.Ptr(v),
				})
			}
			return tags
		}(),
		TemplateBody:     memz.Ptr(templateBody),
		TimeoutInMinutes: aws.Int32(30),
	})
	errorz.MaybeMustWrap(err)

	errorz.MaybeMustWrap(
		cf.NewStackCreateCompleteWaiter(o.awsCF).
			Wait(
				context.Background(),
				&cf.DescribeStacksInput{
					StackName: memz.Ptr(name),
				},
				30*time.Minute))

	return o.MustDescribeStack(name)
}

// MustDescribeStack describes a CloudFormation stack.
func (o *Operations) MustDescribeStack(name string) *cft.Stack {
	out, err := o.awsCF.DescribeStacks(context.Background(), &cf.DescribeStacksInput{
		StackName: memz.Ptr(name),
	})
	if err != nil {
		// TODO(ibrt): Better error handling.
		if strings.Contains(err.Error(), "does not exist") {
			return nil
		}
		errorz.MaybeMustWrap(err, fmt.Errorf("while describing stack '%v'", name))
	}

	errorz.Assertf(len(out.Stacks) == 1, "unexpected number of stacks")
	return &out.Stacks[0]
}

// MustUpdateStack updates a CloudFormation stack.
func (o *Operations) MustUpdateStack(name string, templateBody string, tagsMap map[string]string) *cft.Stack {
	consolez.DefaultCLI.Notice("aws-operations", "updating CloudFormation stack...", name)

	_, err := o.awsCF.UpdateStack(context.Background(), &cf.UpdateStackInput{
		Capabilities: []cft.Capability{
			cft.CapabilityCapabilityIam,
			cft.CapabilityCapabilityNamedIam,
		},
		StackName: memz.Ptr(name),
		Tags: func() []cft.Tag {
			tags := make([]cft.Tag, 0)
			for k, v := range tagsMap {
				tags = append(tags, cft.Tag{
					Key:   memz.Ptr(k),
					Value: memz.Ptr(v),
				})
			}
			return tags
		}(),
		TemplateBody: memz.Ptr(templateBody),
	})
	if err != nil {
		if strings.Contains(err.Error(), "No updates are to be performed") {
			return o.MustDescribeStack(name)
		}
		errorz.MaybeMustWrap(err)
	}

	errorz.MaybeMustWrap(
		cf.NewStackUpdateCompleteWaiter(o.awsCF).
			Wait(
				context.Background(),
				&cf.DescribeStacksInput{StackName: memz.Ptr(name)},
				30*time.Minute))

	return o.MustDescribeStack(name)
}

// MustUpsertStack creates or updates a CloudFormation stack.
func (o *Operations) MustUpsertStack(name string, templateBody string, tagsMap map[string]string) *cft.Stack {
	if o.MustDescribeStack(name) == nil {
		return o.MustCreateStack(name, templateBody, tagsMap)
	}
	return o.MustUpdateStack(name, templateBody, tagsMap)
}

// MustAuthenticateDockerECR runs "docker login" with credentials that allow access to ECR image repositories.
func (o *Operations) MustAuthenticateDockerECR() {
	consolez.DefaultCLI.Notice("aws-operations", "authenticating Docker against ECR...")

	out, err := o.awsECR.GetAuthorizationToken(context.Background(), &awsecr.GetAuthorizationTokenInput{})
	errorz.MaybeMustWrap(err)

	buf, err := base64.StdEncoding.DecodeString(*out.AuthorizationData[0].AuthorizationToken)
	errorz.MaybeMustWrap(err)

	userPass := strings.SplitN(string(buf), ":", 2)
	errorz.Assertf(len(userPass) == 2, "malformed authorization data")

	shellz.NewCommand("docker", "login",
		"--username", userPass[0],
		"--password-stdin",
		strings.TrimPrefix(*out.AuthorizationData[0].ProxyEndpoint, "https://")).
		SetIn(strings.NewReader(userPass[1])).
		MustRun()
}
