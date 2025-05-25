package tools

import (
	"fmt"
	"valetudo-mcp/client"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

var GetRobotCapabilitiesDefinition = ToolDefinition{
	Name:        "get_robot_capabilities",
	Description: "Get the list of capabilities supported by the robot",
}

type GetRobotCapabilitiesArgs struct{}

func GetRobotCapabilities(valetudoClient *client.ValetudoClient, args GetRobotCapabilitiesArgs) (*mcp_golang.ToolResponse, error) {
	responseBody, err := valetudoClient.MakeRequest("GET", "/api/v2/robot/capabilities", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get robot capabilities: %v", err)
	}

	formattedJSON, err := client.FormatJSONResponse(responseBody)
	if err != nil {
		return nil, fmt.Errorf("failed to format robot capabilities response: %v", err)
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(formattedJSON)), nil
}
