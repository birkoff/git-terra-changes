# git-terra-changes
Generate a list of Terraform &amp; Terragrunt Components to apply based on git path changes

## Usage
```
>$./git-terra-changes <git_changes_file.txt> <mappings_file.json>
```

### Constants
```
logFileName        = "git-terra-changes.log"
componentsFileName = "git-terra-changes-components.txt"
liveComponentsPath = "infrastructure/live"
modulesPath        = "infrastructure/modules"
```
 
### git_changes_file.txt
Kind of Based on https://terragrunt.gruntwork.io/docs/features/keep-your-terragrunt-architecture-dry/
You actually need 2 dirs to state the live dir path
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


# Outputs

## git-terra-changes.log
```
[INFO] 2023/06/22 08:03:13.289813 Processing Git Changes File.
[INFO] 2023/06/22 08:03:13.289862 Generating list of components to be deployed.
[INFO] 2023/06/22 08:03:13.289958 ===> Reading file line:  infrastructure/live/_env/sqs
[INFO] 2023/06/22 08:03:13.289978 the Path is a component:  infrastructure/live/_env/sqs
[INFO] 2023/06/22 08:03:13.289988 ===> Reading file line:  infrastructure/live/dev/eu-west-1/sqs
[INFO] 2023/06/22 08:03:13.289996 ===> Reading file line:  infrastructure/live/prod/eu-west-1/lambda
[INFO] 2023/06/22 08:03:13.290003 the Path is a component:  infrastructure/live/prod/eu-west-1/lambda
[INFO] 2023/06/22 08:03:13.290010 ===> Reading file line:  infrastructure/live/prod/eu-west-1/cognito
[INFO] 2023/06/22 08:03:13.290018 the Path is a component:  infrastructure/live/prod/eu-west-1/cognito
[INFO] 2023/06/22 08:03:13.290026 ===> Reading file line:  infrastructure/live/prod/eu-west-1/sqs
[INFO] 2023/06/22 08:03:13.290061 ===> Reading file line:  infrastructure/live/qa/eu-west-1/sqs
[INFO] 2023/06/22 08:03:13.290069 ===> Reading file line:  infrastructure/live/sandbox/eu-west-1/sqs
[INFO] 2023/06/22 08:03:13.290091 ===> Reading file line:  infrastructure/live/staging/eu-west-1/sqs
[INFO] 2023/06/22 08:03:13.290168 ===> Reading file line:  infrastructure/live/uat/eu-west-1/sqs
[INFO] 2023/06/22 08:03:13.290179 ===> Reading file line:  infrastructure/modules/aws-budgets
[INFO] 2023/06/22 08:03:13.290203 ===> Reading file line:  pipelines/templates
[INFO] 2023/06/22 08:03:13.290234 The path is not in the scope, skipping:  pipelines/templates
[INFO] 2023/06/22 08:03:13.290243 ===> Reading file line:  sdjfhskdjfh
[WARNING] 2023/06/22 08:03:13.290251 The path is not valid, skipping:  sdjfhskdjfh
[INFO] 2023/06/22 08:03:13.290259 ===> Reading file line:  sdhf/sdf/sdf/sdf/sdf/sdf
[INFO] 2023/06/22 08:03:13.290267 The path is not in the scope, skipping:  sdhf/sdf/sdf/sdf/sdf/sdf
[INFO] 2023/06/22 08:03:13.290295 Writing the list of components to a file...
[INFO] 2023/06/22 08:03:13.291448 List of components: 
[INFO] 2023/06/22 08:03:13.291464 sqs
[INFO] 2023/06/22 08:03:13.291530 lambda
[INFO] 2023/06/22 08:03:13.291539 cognito
```

## git-terra-changes-components.txt
```
sqs
lambda
cognito
```
