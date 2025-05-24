package tools

import (
	"fmt"
	"valetudo-mcp/client"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

var GetRobotInfoDefinition = ToolDefinition[GetRobotInfoArgs]{
	Name:        "get_robot_info",
	Description: "Get basic robot information including manufacturer, model, and implementation",
}

type GetRobotInfoArgs struct{}

func GetRobotInfo(valetudoClient *client.ValetudoClient, args GetRobotInfoArgs) (*mcp_golang.ToolResponse, error) {
	responseBody, err := valetudoClient.MakeRequest("GET", "/api/v2/robot", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get robot info: %v", err)
	}

	formattedJSON, err := client.FormatJSONResponse(responseBody)
	if err != nil {
		return nil, fmt.Errorf("failed to format robot info response: %v", err)
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(formattedJSON)), nil
}
