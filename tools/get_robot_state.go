package tools

import (
	"fmt"
	"valetudo-mcp/client"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

var GetRobotStateDefinition = ToolDefinition{
	Name:        "get_robot_state",
	Description: "Get the current state of the robot including attributes and map",
}

type GetRobotStateArgs struct{}

func GetRobotState(valetudoClient *client.ValetudoClient, args GetRobotStateArgs) (*mcp_golang.ToolResponse, error) {
	responseBody, err := valetudoClient.MakeRequest("GET", "/api/v2/robot/state", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get robot state: %v", err)
	}

	formattedJSON, err := client.FormatJSONResponse(responseBody)
	if err != nil {
		return nil, fmt.Errorf("failed to format robot state response: %v", err)
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(formattedJSON)), nil
}
