package main

import (
	"testing"
)

func TestSetComponent(t *testing.T) {
	mappings := make(map[string]string)
	mappings["api-gateway-domain-name"] = "api-gateway-cuustom-domains"
	mappings["cognito"] = "cognito"
	mappings["eventbridge"] = "eventbridge"
	mappings["frontend"] = "frontend"

	var components []string
	setComponent("infrastructure/live/_env/sqs", &components, mappings, "infrastructure/live", "infrastructure/modules")                //  1
	setComponent("infrastructure/live/dev/eu-west-1/event-bus", &components, mappings, "infrastructure/live", "infrastructure/modules") // 2
	setComponent("/sqs", &components, mappings, "infrastructure/live", "infrastructure/modules")
	setComponent("", &components, mappings, "infrastructure/live", "infrastructure/modules")
	setComponent("/", &components, mappings, "infrastructure/live", "infrastructure/modules")
	setComponent("infrastructure/", &components, mappings, "infrastructure/live", "infrastructure/modules")
	setComponent("infrastructure-live-dev-eu-west-1-sqs", &components, mappings, "infrastructure/live", "infrastructure/modules")
	setComponent("infrastructure/live/prod/eu-west-1/lambda", &components, mappings, "infrastructure/live", "infrastructure/modules")       // 3
	setComponent("infrastructure/live/prod/eu-west-1/cognito", &components, mappings, "infrastructure/live", "infrastructure/modules")      // 4
	setComponent("infrastructure/live/prod/eu-west-1/user-service", &components, mappings, "infrastructure/live", "infrastructure/modules") // 5
	setComponent("live/qa/eu-west-1/logs", &components, mappings, "live", "infrastructure/modules")                                         // 6
	setComponent("live/sandbox/eu-west-1/account-service", &components, mappings, "live", "infrastructure/modules")                         // 7
	setComponent("live/staging/eu-west-1/database-stack", &components, mappings, "live", "infrastructure/modules")                          // 8
	setComponent("live/uat/eu-west-1/vpc", &components, mappings, "live", "infrastructure/modules")                                         // 9
	setComponent("modules/aws-budgets", &components, mappings, "live", "infrastructure/modules")

	if len(components) != 9 {
		t.Errorf("Expected 4 components to be set, got %d", len(components))
	}

	if components[0] != "sqs" {
		t.Errorf("Expected first component to be sqs, got %s", components[0])
	}

	if components[1] != "event-bus" {
		t.Errorf("Expected second component to be lambda, got %s", components[1])
	}

	if components[2] != "lambda" {
		t.Errorf("Expected third component to be cognito, got %s", components[2])
	}

	if components[3] != "cognito" {
		t.Errorf("Expected third component to be vpc, got %s", components[3])
	}

	if components[4] != "user-service" {
		t.Errorf("Expected third component to be vpc, got %s", components[2])
	}

	if components[5] != "logs" {
		t.Errorf("Expected third component to be vpc, got %s", components[2])
	}

	if components[6] != "account-service" {
		t.Errorf("Expected third component to be vpc, got %s", components[2])
	}

	if components[7] != "database-stack" {
		t.Errorf("Expected third component to be vpc, got %s", components[2])
	}

	if components[8] != "vpc" {
		t.Errorf("Expected third component to be vpc, got %s", components[2])
	}
}
