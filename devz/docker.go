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

// Up brings up the services.
func (d *DockerCompose) Up() {
	d.GetCommand().
		AddParams("up", "--detach", "--build", "--pull=always", "--force-recreate", "--remove-orphans", "--wait").
		MustRun()
}

// Down takes down the services.
func (d *DockerCompose) Down() {
	d.GetCommand().
		AddParams("down", "--volumes", "--rmi=local", "--remove-orphans", "--timeout", "5").
		MustRun()
}

// GetCommand returns a pre-configured "docker compose" command.
func (d *DockerCompose) GetCommand() *shellz.Command {
	buf, err := yaml.Marshal(d.cfg)
	errorz.MaybeMustWrap(err)

	cmd := shellz.NewCommand("docker", "compose")

	if color.NoColor || os.Getenv("CI") != "" {
		cmd = cmd.AddParams("--ansi", "never", "--progress", "plain")
	}

	return cmd.
		AddParams("-p", d.cfg.Name).
		AddParams("-f", "-").
		SetIn(bytes.NewReader(buf))
}
