package tools

import (
	"fmt"
	"valetudo-mcp/client"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

var ZoneCleaningDefinition = ToolDefinition[ZoneCleaningArgs]{
	Name:        "zone_cleaning",
	Description: "Start zone cleaning with specified zones and iterations",
}

type ZoneCleaningArgs struct {
	Zones      []Zone `json:"zones" jsonschema:"required,description=List of zones to clean"`
	Iterations int    `json:"iterations" jsonschema:"description=Number of cleaning iterations per zone (default: 1)"`
}

type Zone struct {
	Points ZonePoints `json:"points" jsonschema:"required,description=Points defining the zone"`
}

type ZonePoints struct {
	PA Point `json:"pA" jsonschema:"required,description=First point of the zone"`
	PB Point `json:"pB" jsonschema:"required,description=Second point of the zone"`
	PC Point `json:"pC" jsonschema:"required,description=Third point of the zone"`
	PD Point `json:"pD" jsonschema:"required,description=Fourth point of the zone"`
}

type Point struct {
	X float64 `json:"x" jsonschema:"required,description=X coordinate"`
	Y float64 `json:"y" jsonschema:"required,description=Y coordinate"`
}

func ZoneCleaning(valetudoClient *client.ValetudoClient, args ZoneCleaningArgs) (*mcp_golang.ToolResponse, error) {
	if len(args.Zones) == 0 {
		return nil, fmt.Errorf("at least one zone must be specified")
	}

	if args.Iterations < 1 {
		args.Iterations = 1
	}

	payload := map[string]interface{}{
		"action":     "clean",
		"zones":      args.Zones,
		"iterations": args.Iterations,
	}

	responseBody, err := valetudoClient.MakeRequest("PUT", "/api/v2/robot/capabilities/ZoneCleaningCapability", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to start zone cleaning: %v", err)
	}

	var response string
	if string(responseBody) == "OK" {
		response = fmt.Sprintf("Successfully started zone cleaning with %d zones and %d iterations", len(args.Zones), args.Iterations)
	} else {
		response = fmt.Sprintf("Unexpected response: %s", string(responseBody))
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(response)), nil
}
