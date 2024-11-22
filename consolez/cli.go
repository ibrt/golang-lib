package consolez

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/alecthomas/kong"
	"github.com/rodaine/table"

	"github.com/ibrt/golang-lib/errorz"
	"github.com/ibrt/golang-lib/filez"
	"github.com/ibrt/golang-lib/outz"
	"github.com/ibrt/golang-lib/stringz"
)

// Known icons.
const (
	IconRocket                     = "\U0001F680"
	IconHighVoltage                = "\U000026A1"
	IconBackhandIndexPointingRight = "\U0001F449"
	IconRunner                     = "\U0001F3C3"
	IconCollision                  = "\U0001F4A5"
)

var (
	// DefaultCLI is a default, shared instance of CLI.
	DefaultCLI = NewCLI(true, os.Exit)
)

// CLI provides some utilities for printing messages in CLI tools.
type CLI struct {
	m          *sync.Mutex
	hL         int
	addSpacing bool
	exit       func(code int)
}

// NewCLI initializes a new CLI.
func NewCLI(addSpacing bool, exit func(int)) *CLI {
	return &CLI{
		m:          &sync.Mutex{},
		hL:         0,
		addSpacing: addSpacing,
		exit:       exit,
	}
}

// Tool introduces a command line tool.
func (c *CLI) Tool(toolName string, k *kong.Context) {
	commandParts := make([]string, 0)
	options := make([][]string, 0)

	for _, p := range k.Path {
		if p.Command != nil {
			commandParts = append(commandParts, p.Command.Name)
		} else if p.Flag != nil {
			options = append(options, []string{
				p.Flag.Summary(),
				fmt.Sprintf("%v", p.Flag.Target.Interface()),
			})
		} else if p.Positional != nil {
			options = append(options, []string{
				p.Positional.Summary(),
				fmt.Sprintf("%v", p.Positional.Target.Interface()),
			})
		}
	}

	c.Banner(toolName, strings.Join(commandParts, " "))

	if len(options) > 0 {
		fmt.Println()
		c.NewTable("Input", "Value").SetRows(options).Print()
	}
}

// Banner prints a banner.
func (c *CLI) Banner(title, tagLine string) {
	c.m.Lock()
	defer c.m.Unlock()

	fmt.Print("┌", strings.Repeat("─", len(title)+len(tagLine)+6), "┐\n")
	fmt.Print("│ ", IconRocket, " ")
	_, _ = outz.GetColorHighlight().Print(title)
	fmt.Print(" ")
	fmt.Print(tagLine)
	fmt.Print(" │\n")
	fmt.Print("└", strings.Repeat("─", len(title)+len(tagLine)+6), "┘\n")
}

// Header prints a header based on a nesting hierarchy.
// Always call the returned function to close the scope, for example by deferring it.
func (c *CLI) Header(format string, a ...any) func() {
	c.m.Lock()
	defer c.m.Unlock()

	if c.hL < 2 && c.addSpacing {
		fmt.Println()
	}

	switch c.hL {
	case 0:
		fmt.Print(IconHighVoltage)
		fmt.Print(" ")
		_, _ = outz.GetColorHighlight().Printf(format, a...)
		fmt.Println()
	case 1:
		fmt.Print(IconBackhandIndexPointingRight)
		fmt.Print(" ")
		fmt.Printf(format, a...)
		fmt.Println()
	default:
		_, _ = outz.GetColorSecondaryHighlight().Print("—— ")
		_, _ = outz.GetColorSecondaryHighlight().Printf(format, a...)
		fmt.Println()
	}

	c.hL++
	isClosed := false

	return func() {
		c.m.Lock()
		defer c.m.Unlock()

		if !isClosed {
			isClosed = true
			c.hL--
		}
	}
}

// Notice prints a notice.
func (c *CLI) Notice(scope string, highlight string, secondary ...string) {
	c.m.Lock()
	defer c.m.Unlock()

	_, _ = outz.GetColorSecondary().Printf("[%v]", stringz.AlignRight(scope, 24))
	_, _ = outz.GetColorDefault().Print(" ", highlight)

	for _, v := range secondary {
		_, _ = outz.GetColorSecondary().Print(" ", v)
	}

	fmt.Println()
}

// Command prints a command.
func (c *CLI) Command(cmd string, params ...string) {
	c.m.Lock()
	defer c.m.Unlock()

	if filepath.IsAbs(cmd) && strings.HasPrefix(cmd, stringz.EnsureSuffix(filez.MustGetwd(), string(os.PathSeparator))) {
		cmd = filez.MustRel(filez.MustGetwd(), cmd)
	}

	fmt.Print(IconRunner)
	fmt.Printf(" %v ", cmd)
	_, _ = outz.GetColorSecondary().Print(strings.Join(params, " "))
	fmt.Println()
}

// NewTable creates a new table.
func (c *CLI) NewTable(columnHeaders ...any) table.Table {
	return table.
		New(columnHeaders...).
		WithHeaderFormatter(outz.GetColorHighlight().SprintfFunc()).
		WithFirstColumnFormatter(outz.GetColorWarning().SprintfFunc())
}

// Error prints an error.
func (c *CLI) Error(err error, debug bool) {
	c.m.Lock()
	defer c.m.Unlock()

	if c.addSpacing {
		fmt.Println()
	}

	fmt.Print(IconCollision)
	fmt.Print(" ")
	_, _ = outz.GetColorHighlight().Println("Error")
	_, _ = outz.GetColorError().Println(err.Error())

	if debug {
		fmt.Println(errorz.SDump(err))
	}
}
