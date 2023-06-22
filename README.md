# git-terra-changes
Generate a list of Terraform &amp; Terragrunt Components to apply based on git path changes

## Usage
```
>$./git-terra-changes <git_changes_file.txt> <mappings_file.json>
```

logFileName        = "git-terra-changes.log"
componentsFileName = "git-terra-changes-components.txt"
liveComponentsPath = "infrastructure/live"
modulesPath        = "infrastructure/modules"

 
### git_changes_file.txt
Based on https://terragrunt.gruntwork.io/docs/features/keep-your-terragrunt-architecture-dry/
```
infrastructure/live/_env/sqs
infrastructure/live/dev/eu-west-1/sqs
infrastructure/live/prod/eu-west-1/lambda
infrastructure/live/prod/eu-west-1/cognito
infrastructure/live/prod/eu-west-1/sqs
infrastructure/live/qa/eu-west-1/sqs
infrastructure/live/sandbox/eu-west-1/sqs
infrastructure/live/staging/eu-west-1/sqs
infrastructure/live/uat/eu-west-1/sqs
infrastructure/modules/aws-budgets
pipelines/templates
sdjfhskdjfh
sdhf/sdf/sdf/sdf/sdf/sdf
```

### mappings_file.json
module_name = component_name
```
{
	"api-gateway-domain-name": "api-gateway-cuustom-domains",
	"cognito": "cognito",
	"eventbridge": "eventbridge",
	"frontend": "frontend"
}
```
