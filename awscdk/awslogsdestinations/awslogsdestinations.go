package awslogsdestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awslogsdestinations/internal"
)

// Use a Kinesis stream as the destination for a log subscription.
// Experimental.
type KinesisDestination interface {
	awslogs.ILogSubscriptionDestination
	Bind(scope awscdk.Construct, _sourceLogGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig
}

// The jsii proxy struct for KinesisDestination
type jsiiProxy_KinesisDestination struct {
	internal.Type__awslogsILogSubscriptionDestination
}

// Experimental.
func NewKinesisDestination(stream awskinesis.IStream) KinesisDestination {
	_init_.Initialize()

	j := jsiiProxy_KinesisDestination{}

	_jsii_.Create(
		"monocdk.aws_logs_destinations.KinesisDestination",
		[]interface{}{stream},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisDestination_Override(k KinesisDestination, stream awskinesis.IStream) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs_destinations.KinesisDestination",
		[]interface{}{stream},
		k,
	)
}

// Return the properties required to send subscription events to this destination.
//
// If necessary, the destination can use the properties of the SubscriptionFilter
// object itself to configure its permissions to allow the subscription to write
// to it.
//
// The destination may reconfigure its own permissions in response to this
// function call.
// Experimental.
func (k *jsiiProxy_KinesisDestination) Bind(scope awscdk.Construct, _sourceLogGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig {
	var returns *awslogs.LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{scope, _sourceLogGroup},
		&returns,
	)

	return returns
}

// Use a Lambda Function as the destination for a log subscription.
// Experimental.
type LambdaDestination interface {
	awslogs.ILogSubscriptionDestination
	Bind(scope awscdk.Construct, logGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig
}

// The jsii proxy struct for LambdaDestination
type jsiiProxy_LambdaDestination struct {
	internal.Type__awslogsILogSubscriptionDestination
}

// LambdaDestinationOptions.
// Experimental.
func NewLambdaDestination(fn awslambda.IFunction, options *LambdaDestinationOptions) LambdaDestination {
	_init_.Initialize()

	j := jsiiProxy_LambdaDestination{}

	_jsii_.Create(
		"monocdk.aws_logs_destinations.LambdaDestination",
		[]interface{}{fn, options},
		&j,
	)

	return &j
}

// LambdaDestinationOptions.
// Experimental.
func NewLambdaDestination_Override(l LambdaDestination, fn awslambda.IFunction, options *LambdaDestinationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs_destinations.LambdaDestination",
		[]interface{}{fn, options},
		l,
	)
}

// Return the properties required to send subscription events to this destination.
//
// If necessary, the destination can use the properties of the SubscriptionFilter
// object itself to configure its permissions to allow the subscription to write
// to it.
//
// The destination may reconfigure its own permissions in response to this
// function call.
// Experimental.
func (l *jsiiProxy_LambdaDestination) Bind(scope awscdk.Construct, logGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig {
	var returns *awslogs.LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, logGroup},
		&returns,
	)

	return returns
}

// Options that may be provided to LambdaDestination.
// Experimental.
type LambdaDestinationOptions struct {
	// Whether or not to add Lambda Permissions.
	// Experimental.
	AddPermissions *bool `json:"addPermissions"`
}

