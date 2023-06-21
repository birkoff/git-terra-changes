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
	setComponent("infrastructure/live/_env/sqs", &components, mappings)
	setComponent("infrastructure/live/dev/eu-west-1/sqs", &components, mappings)
	setComponent("/sqs", &components, mappings)
	setComponent("", &components, mappings)
	setComponent("/", &components, mappings)
	setComponent("infrastructure/", &components, mappings)
	setComponent("infrastructure-live-dev-eu-west-1-sqs", &components, mappings)
	setComponent("infrastructure/live/prod/eu-west-1/lambda", &components, mappings)
	setComponent("infrastructure/live/prod/eu-west-1/cognito", &components, mappings)
	setComponent("infrastructure/live/prod/eu-west-1/sqs", &components, mappings)
	setComponent("infrastructure/live/qa/eu-west-1/sqs", &components, mappings)
	setComponent("infrastructure/live/sandbox/eu-west-1/sqs", &components, mappings)
	setComponent("infrastructure/live/staging/eu-west-1/sqs", &components, mappings)
	setComponent("infrastructure/live/uat/eu-west-1/vpc", &components, mappings)
	setComponent("infrastructure/modules/aws-budgets", &components, mappings)

	if len(components) != 4 {
		t.Errorf("Expected 4 components to be set, got %d", len(components))
	}

	if components[0] != "sqs" {
		t.Errorf("Expected first component to be sqs, got %s", components[0])
	}

	if components[1] != "lambda" {
		t.Errorf("Expected second component to be lambda, got %s", components[1])
	}

	if components[2] != "cognito" {
		t.Errorf("Expected third component to be cognito, got %s", components[2])
	}

	if components[3] != "vpc" {
		t.Errorf("Expected third component to be vpc, got %s", components[2])
	}
}
