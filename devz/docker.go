package devz

import (
	"bytes"
	"os"

	ct "github.com/compose-spec/compose-go/types"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"

	"github.com/ibrt/golang-lib/errorz"
	"github.com/ibrt/golang-lib/shellz"
)

// DockerCompose helps operate a Docker Compose config.
type DockerCompose struct {
	cfg *ct.Config
}

// NewDockerCompose initializes a new Docker Compose.
func NewDockerCompose(cfg *ct.Config) *DockerCompose {
	return &DockerCompose{
		cfg: cfg,
	}
}

// GetConfig returns the Docker Compose config.
func (d *DockerCompose) GetConfig() *ct.Config {
	return d.cfg
}

// GetMarshaledConfig returns the Docker Compose config marshaled to YAML.
func (d *DockerCompose) GetMarshaledConfig() []byte {
	buf, err := yaml.Marshal(d.cfg)
	errorz.MaybeMustWrap(err)
	return buf
}

// GetDownCommand returns a pre-configured "docker compose down" command.
func (d *DockerCompose) GetDownCommand() *shellz.Command {
	return d.GetCommand().
		AddParams("down", "--volumes", "--rmi=local", "--remove-orphans", "--timeout", "5")
}

// GetUpCommand returns a pre-configured "docker compose up" command.
func (d *DockerCompose) GetUpCommand() *shellz.Command {
	return d.GetCommand().
		AddParams("up", "--detach", "--build", "--pull=always", "--force-recreate", "--remove-orphans", "--wait")
}

// GetCommand returns a pre-configured "docker compose" command.
func (d *DockerCompose) GetCommand() *shellz.Command {
	cmd := shellz.NewCommand("docker", "compose")

	if color.NoColor || os.Getenv("CI") != "" {
		cmd = cmd.AddParams("--ansi", "never", "--progress", "plain")
	}

	return cmd.
		AddParams("-p", d.cfg.Name).
		AddParams("-f", "-").
		SetIn(bytes.NewReader(d.GetMarshaledConfig()))
}
