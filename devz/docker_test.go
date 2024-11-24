package devz_test

import (
	"testing"
	"time"

	ct "github.com/compose-spec/compose-go/types"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"github.com/ibrt/golang-lib/devz"
	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/memz"
)

type DockerSuite struct {
	// intentionally empty
}

func TestDockerSuite(t *testing.T) {
	fixturez.RunSuite(t, &DockerSuite{})
}

func (*DockerSuite) TestDockerCompose(g *WithT) {
	dc := devz.NewDockerCompose(&ct.Config{Name: "test"})

	g.Expect(dc.GetProjectName()).To(BeEmpty())
	g.Expect(dc.GetProfiles()).To(BeEmpty())
	g.Expect(dc.GetConfig()).To(Equal(&ct.Config{Name: "test"}))
	g.Expect(string(dc.GetMarshaledConfig())).To(Equal("name: test\nservices: {}\n"))

	g.Expect(dc.GetUpCommand().GetParams()).To(HaveExactElements(
		"compose", "--ansi", "never", "--progress", "plain", "-f", "-",
		"up", "--detach", "--build", "--pull", "always", "--force-recreate", "--remove-orphans", "--wait"))
	g.Eventually(gbytes.BufferReader(dc.GetUpCommand().GetIn())).Should(gbytes.Say("name: test\nservices: {}\n"))

	g.Expect(dc.GetDownCommand().GetParams()).To(HaveExactElements(
		"compose", "--ansi", "never", "--progress", "plain", "-f", "-",
		"down", "--volumes", "--rmi", "local", "--remove-orphans", "--timeout", "5"))
	g.Eventually(gbytes.BufferReader(dc.GetDownCommand().GetIn())).Should(gbytes.Say("name: test\nservices: {}\n"))

	g.Expect(dc.GetPSCommand().GetParams()).To(HaveExactElements(
		"compose", "--ansi", "never", "--progress", "plain", "-f", "-",
		"ps", "--all", "--orphans"))
	g.Eventually(gbytes.BufferReader(dc.GetPSCommand().GetIn())).Should(gbytes.Say("name: test\nservices: {}\n"))

	g.Expect(dc.GetCommand().GetParams()).To(HaveExactElements(
		"compose", "--ansi", "never", "--progress", "plain", "-f", "-"))
	g.Eventually(gbytes.BufferReader(dc.GetCommand().GetIn())).Should(gbytes.Say("name: test\nservices: {}\n"))

	dc = dc.WithProjectName("projectName")
	g.Expect(dc.GetProjectName()).To(Equal("projectName"))
	g.Expect(dc.GetProfiles()).To(BeEmpty())

	dc = dc.WithProfiles("a", "b")
	g.Expect(dc.GetProjectName()).To(Equal("projectName"))
	g.Expect(dc.GetProfiles()).To(HaveExactElements("a", "b"))

	g.Expect(dc.GetCommand().GetParams()).To(HaveExactElements(
		"compose", "--project-name", "projectName", "--profile", "a", "--profile", "b",
		"--ansi", "never", "--progress", "plain", "-f", "-"))
}

func (*DockerSuite) TestNewDockerComposeConfigDeploy(g *WithT) {
	g.Expect(devz.NewDockerComposeConfigDeploy(-1, 1)).To(Equal(
		&ct.DeployConfig{
			RestartPolicy: &ct.RestartPolicy{
				Condition:   "on-failure",
				Delay:       memz.Ptr(ct.Duration(10 * time.Second)),
				MaxAttempts: memz.Ptr[uint64](3),
			},
		}))

	g.Expect(devz.NewDockerComposeConfigDeploy(512, 2)).To(Equal(
		&ct.DeployConfig{
			EndpointMode: "dnsrr",
			Replicas:     memz.Ptr[uint64](2),
			Resources: ct.Resources{
				Limits: &ct.Resource{
					MemoryBytes: ct.UnitBytes(512 * 1024 * 1024),
				},
			},
			RestartPolicy: &ct.RestartPolicy{
				Condition:   "on-failure",
				Delay:       memz.Ptr(ct.Duration(10 * time.Second)),
				MaxAttempts: memz.Ptr[uint64](3),
			},
		}))
}

func (*DockerSuite) TestNewDockerComposeConfigExtraHosts(g *WithT) {
	defer devz.RestoreDefaultRuntimeGOOS()

	devz.DefaultRuntimeGOOS = "linux"

	g.Expect(devz.NewDockerComposeConfigExtraHosts(
		map[string]string{"k1": "v1", "k2": "v2"},
		map[string]string{"k3": "v3"})).
		To(Equal(ct.HostsList{
			"k1":                   "v1",
			"k2":                   "v2",
			"k3":                   "v3",
			"host.docker.internal": "host-gateway",
		}))

	devz.DefaultRuntimeGOOS = "darwin"

	g.Expect(devz.NewDockerComposeConfigExtraHosts(
		map[string]string{"k1": "v1", "k2": "v2"},
		map[string]string{"k3": "v3"})).
		To(Equal(ct.HostsList{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		}))
}

func (*DockerSuite) TestNewDockerComposeConfigHealthCheck(g *WithT) {
	g.Expect(devz.NewDockerComposeConfigHealthCheckShell("cat %v", 1)).To(Equal(
		&ct.HealthCheckConfig{
			StartPeriod: memz.Ptr(ct.Duration(30 * time.Second)),
			Interval:    memz.Ptr(ct.Duration(5 * time.Second)),
			Timeout:     memz.Ptr(ct.Duration(3 * time.Second)),
			Retries:     memz.Ptr(uint64(3)),
			Test: ct.HealthCheckTest{
				"CMD-SHELL",
				"cat 1",
			},
		}))
}
