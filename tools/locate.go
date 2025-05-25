package tools

import (
	"fmt"
	"valetudo-mcp/client"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

var LocateDefinition = ToolDefinition{
	Name:        "locate",
	Description: "Locate the robot",
}

type LocateArgs struct{}

func Locate(valetudoClient *client.ValetudoClient, args LocateArgs) (*mcp_golang.ToolResponse, error) {
	payload := map[string]string{
		"action": "locate",
	}

	responseBody, err := valetudoClient.MakeRequest("PUT", "/api/v2/robot/capabilities/LocateCapability", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to execute locate action: %v", err)
	}

	var response string
	if string(responseBody) == "OK" {
		response = "Successfully executed locate action"
	} else {
		response = fmt.Sprintf("Unexpected response: %s", string(responseBody))
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(response)), nil
}
