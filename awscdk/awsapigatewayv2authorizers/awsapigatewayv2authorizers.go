package awsapigatewayv2authorizers

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2authorizers/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Authorize Http Api routes on whether the requester is registered as part of an AWS Cognito user pool.
// Experimental.
type HttpJwtAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpJwtAuthorizer
type jsiiProxy_HttpJwtAuthorizer struct {
	internal.Type__awsapigatewayv2IHttpRouteAuthorizer
}

// Experimental.
func NewHttpJwtAuthorizer(props *HttpJwtAuthorizerProps) HttpJwtAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpJwtAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpJwtAuthorizer",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpJwtAuthorizer_Override(h HttpJwtAuthorizer, props *HttpJwtAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpJwtAuthorizer",
		[]interface{}{props},
		h,
	)
}

// Bind this authorizer to a specified Http route.
// Experimental.
func (h *jsiiProxy_HttpJwtAuthorizer) Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
	var returns *awsapigatewayv2.HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Properties to initialize HttpJwtAuthorizer.
// Experimental.
type HttpJwtAuthorizerProps struct {
	// A list of the intended recipients of the JWT.
	//
	// A valid JWT must provide an aud that matches at least one entry in this list.
	// Experimental.
	JwtAudience *[]*string `json:"jwtAudience"`
	// The base domain of the identity provider that issues JWT.
	// Experimental.
	JwtIssuer *string `json:"jwtIssuer"`
	// The name of the authorizer.
	// Experimental.
	AuthorizerName *string `json:"authorizerName"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `json:"identitySource"`
}

// Authorize Http Api routes via a lambda function.
// Experimental.
type HttpLambdaAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpLambdaAuthorizer
type jsiiProxy_HttpLambdaAuthorizer struct {
	internal.Type__awsapigatewayv2IHttpRouteAuthorizer
}

// Experimental.
func NewHttpLambdaAuthorizer(props *HttpLambdaAuthorizerProps) HttpLambdaAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpLambdaAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpLambdaAuthorizer",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpLambdaAuthorizer_Override(h HttpLambdaAuthorizer, props *HttpLambdaAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpLambdaAuthorizer",
		[]interface{}{props},
		h,
	)
}

// Bind this authorizer to a specified Http route.
// Experimental.
func (h *jsiiProxy_HttpLambdaAuthorizer) Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
	var returns *awsapigatewayv2.HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Properties to initialize HttpTokenAuthorizer.
// Experimental.
type HttpLambdaAuthorizerProps struct {
	// The name of the authorizer.
	// Experimental.
	AuthorizerName *string `json:"authorizerName"`
	// The lambda function used for authorization.
	// Experimental.
	Handler awslambda.IFunction `json:"handler"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `json:"identitySource"`
	// The types of responses the lambda can return.
	//
	// If HttpLambdaResponseType.SIMPLE is included then
	// response format 2.0 will be used.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html#http-api-lambda-authorizer.payload-format-response
	//
	// Experimental.
	ResponseTypes *[]HttpLambdaResponseType `json:"responseTypes"`
	// How long APIGateway should cache the results.
	//
	// Max 1 hour.
	// Disable caching by setting this to `Duration.seconds(0)`.
	// Experimental.
	ResultsCacheTtl awscdk.Duration `json:"resultsCacheTtl"`
}

// Specifies the type responses the lambda returns.
// Experimental.
type HttpLambdaResponseType string

const (
	HttpLambdaResponseType_SIMPLE HttpLambdaResponseType = "SIMPLE"
	HttpLambdaResponseType_IAM HttpLambdaResponseType = "IAM"
)

// Authorize Http Api routes on whether the requester is registered as part of an AWS Cognito user pool.
// Experimental.
type HttpUserPoolAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpUserPoolAuthorizer
type jsiiProxy_HttpUserPoolAuthorizer struct {
	internal.Type__awsapigatewayv2IHttpRouteAuthorizer
}

// Experimental.
func NewHttpUserPoolAuthorizer(props *UserPoolAuthorizerProps) HttpUserPoolAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpUserPoolAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpUserPoolAuthorizer",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpUserPoolAuthorizer_Override(h HttpUserPoolAuthorizer, props *UserPoolAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpUserPoolAuthorizer",
		[]interface{}{props},
		h,
	)
}

// Bind this authorizer to a specified Http route.
// Experimental.
func (h *jsiiProxy_HttpUserPoolAuthorizer) Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
	var returns *awsapigatewayv2.HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Properties to initialize UserPoolAuthorizer.
// Experimental.
type UserPoolAuthorizerProps struct {
	// The associated user pool.
	// Experimental.
	UserPool awscognito.IUserPool `json:"userPool"`
	// The user pool client that should be used to authorize requests with the user pool.
	// Experimental.
	UserPoolClient awscognito.IUserPoolClient `json:"userPoolClient"`
	// The name of the authorizer.
	// Experimental.
	AuthorizerName *string `json:"authorizerName"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `json:"identitySource"`
	// The AWS region in which the user pool is present.
	// Experimental.
	UserPoolRegion *string `json:"userPoolRegion"`
}

