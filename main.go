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
	def tools.ToolDefinition,
	handler func(*client.ValetudoClient, Args) (*mcp_golang.ToolResponse, error),
	valetudoClient *client.ValetudoClient,
) {
	err := server.RegisterTool(
		def.Name,
		def.Description,
		func(args Args) (*mcp_golang.ToolResponse, error) {
			return handler(valetudoClient, args)
		},
	)

	if err != nil {
		panic(err)
	}
}

func main() {
	done := make(chan struct{})

	baseURL := os.Getenv("VALETUDO_URL")
	if baseURL == "" {
		panic("VALETUDO_URL environment variable must be set")
	}
	valetudoClient := client.NewValetudoClient(baseURL)

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	registerTool(server, tools.BasicControlDefinition, tools.BasicControl, valetudoClient)
	registerTool(server, tools.GetRobotCapabilitiesDefinition, tools.GetRobotCapabilities, valetudoClient)
	registerTool(server, tools.GetRobotInfoDefinition, tools.GetRobotInfo, valetudoClient)
	registerTool(server, tools.GetRobotStateDefinition, tools.GetRobotState, valetudoClient)
	registerTool(server, tools.LocateDefinition, tools.Locate, valetudoClient)
	registerTool(server, tools.SetLogLevelDefinition, tools.SetLogLevel, valetudoClient)
	registerTool(server, tools.ZoneCleaningDefinition, tools.ZoneCleaning, valetudoClient)

	if err := server.Serve(); err != nil {
		panic(err)
	}

	<-done
}
