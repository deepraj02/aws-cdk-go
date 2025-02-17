package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ServerlessHwStackProps struct {
	awscdk.StackProps
}

func NewServerlessHwStack(scope constructs.Construct, id string, props *ServerlessHwStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("ServerlessHwQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	/// Define a new lambda function resource
	helloWorldFunction := awslambda.NewFunction(stack, jsii.String("HelloWorldFunction"), &awslambda.FunctionProps{
		/// Choosing the Runtime for the Lambda function
		Runtime: awslambda.Runtime_NODEJS_20_X(),
		// Points to the lambda directory
		Code:    awslambda.Code_FromAsset(jsii.String("lambda"), &awss3assets.AssetOptions{}),
		Handler: jsii.String("hello.handler"),
	})

	/// Define a API Gateway source

	api := awsapigateway.NewLambdaRestApi(stack, jsii.String("HelloWorldAPI"), &awsapigateway.LambdaRestApiProps{
		Handler: helloWorldFunction,
		Proxy:   jsii.Bool(true),
	})
	/// Add a '/hello' resource with a GET method
	helloResource := api.Root().AddResource(jsii.String("hello"), &awsapigateway.ResourceOptions{})
	helloResource.AddMethod(jsii.String("GET"), awsapigateway.NewLambdaIntegration(helloWorldFunction, &awsapigateway.LambdaIntegrationOptions{}), &awsapigateway.MethodOptions{})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewServerlessHwStack(app, "ServerlessHwStack", &ServerlessHwStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
