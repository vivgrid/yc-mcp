package main

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/vivgrid/yc"
	"github.com/yomorun/yomo/serverless"
)

const (
	CmdYC      = "yc"
	CmdCreate  = "create"
	CmdDeploy  = "deploy"
	CmdLogs    = "logs"
	CmdRemove  = "remove"
	CmdStatus  = "status"
	CmdUpload  = "upload"
	CmdVersion = "version"
)

// Init is an optional function invoked during the initialization phase of the sfn instance
func Init() error {
	return nil
}

// Description returns a description of the function
func Description() string {
	return `yc is the command-line interface (CLI) tool for deploying your serverless functions on Vivgrid. It provides a convenient way to manage your deployments and serverless instances globally.

	To install yc, you can use the following command:
	'curl "https://bina.egoist.dev/vivgrid/yc" | sh'

	- You can create a 'yc.yml' file in the root of your project to configure the deployment settings.
	- Return the usage method based on the input command
	`
}

// InputSchema defines the argument structure for LLM Function Calling. It
// utilizes jsonschema tags to detail the definition. For jsonschema in Go,
// see https://github.com/invopop/jsonschema.
func InputSchema() any {
	return &LLMArguments{}
}

// LLMArguments defines the arguments for the LLM Function Calling. These
// arguments are combined to form a prompt automatically.
type LLMArguments struct {
	Command string `json:"command" jsonschema:"title=Command,description=The command to execute,example=yc deploy"`
}

// Handler orchestrates the core processing logic of this function.
// - ctx.ReadLLMArguments() parses LLM Function Calling Arguments (skip if none).
// - ctx.WriteLLMResult() sends the retrieval result back to LLM.
func Handler(ctx serverless.Context) {
	var p LLMArguments
	ctx.ReadLLMArguments(&p)

	cmd := p.Command
	switch {
	// create
	case strings.Contains(cmd, CmdCreate):
		cmd = CmdCreate
	// deploy this is an alias of chaining commands (upload -> remove -> create)
	case strings.Contains(cmd, CmdDeploy):
		cmd = CmdDeploy
	// logs
	case strings.Contains(cmd, CmdLogs):
		cmd = CmdLogs
	// remove
	case strings.Contains(cmd, CmdRemove):
		cmd = CmdRemove
	// status
	case strings.Contains(cmd, CmdStatus):
		cmd = CmdStatus
	// upload
	case strings.Contains(cmd, CmdUpload):
		cmd = CmdUpload
	// version
	case strings.Contains(cmd, CmdVersion):
		cmd = CmdVersion
	default:
		cmd = CmdYC
	}

	doc, err := yc.Doc(cmd)
	if err != nil {
		slog.Error("yc-mcp", "command", p.Command, "error", err)
		ctx.WriteLLMResult(fmt.Sprintf("Error retrieving documentation for command %s: %v", cmd, err))
		return
	}

	ctx.WriteLLMResult(doc)
	slog.Info("yc-mcp", "command", p.Command, "doc", doc)
}
