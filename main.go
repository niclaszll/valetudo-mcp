package main

import (
	"os"
	"valetudo-mcp/client"
	"valetudo-mcp/tools"

	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

func registerTool[Args any](
	server *mcp_golang.Server,
	def tools.ToolDefinition[Args],
	handler func(*client.ValetudoClient, Args) (*mcp_golang.ToolResponse, error),
	valetudoClient *client.ValetudoClient,
) error {
	return server.RegisterTool(
		def.Name,
		def.Description,
		func(args Args) (*mcp_golang.ToolResponse, error) {
			return handler(valetudoClient, args)
		},
	)
}

func main() {
	done := make(chan struct{})

	baseURL := os.Getenv("VALETUDO_URL")
	if baseURL == "" {
		panic("VALETUDO_URL environment variable must be set")
	}
	valetudoClient := client.NewValetudoClient(baseURL)

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	if err := registerTool(server, tools.GetRobotStateDefinition, tools.GetRobotState, valetudoClient); err != nil {
		panic(err)
	}

	if err := registerTool(server, tools.GetRobotInfoDefinition, tools.GetRobotInfo, valetudoClient); err != nil {
		panic(err)
	}

	if err := registerTool(server, tools.GetRobotCapabilitiesDefinition, tools.GetRobotCapabilities, valetudoClient); err != nil {
		panic(err)
	}

	if err := registerTool(server, tools.SetLogLevelDefinition, tools.SetLogLevel, valetudoClient); err != nil {
		panic(err)
	}

	if err := server.Serve(); err != nil {
		panic(err)
	}

	<-done
}
