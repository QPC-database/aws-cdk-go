package lambdalayerkubectl

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/lambdalayerkubectl/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// An AWS Lambda layer that includes `kubectl` and `helm`.
// Experimental.
type KubectlLayer interface {
	awslambda.LayerVersion
	CompatibleRuntimes() *[]awslambda.Runtime
	Env() *awscdk.ResourceEnvironment
	LayerVersionArn() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	AddPermission(id *string, permission *awslambda.LayerVersionPermission)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for KubectlLayer
type jsiiProxy_KubectlLayer struct {
	internal.Type__awslambdaLayerVersion
}

func (j *jsiiProxy_KubectlLayer) CompatibleRuntimes() *[]awslambda.Runtime {
	var returns *[]awslambda.Runtime
	_jsii_.Get(
		j,
		"compatibleRuntimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlLayer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlLayer) LayerVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlLayer) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlLayer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlLayer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewKubectlLayer(scope constructs.Construct, id *string) KubectlLayer {
	_init_.Initialize()

	j := jsiiProxy_KubectlLayer{}

	_jsii_.Create(
		"monocdk.lambda_layer_kubectl.KubectlLayer",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewKubectlLayer_Override(k KubectlLayer, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.lambda_layer_kubectl.KubectlLayer",
		[]interface{}{scope, id},
		k,
	)
}

// Imports a layer version by ARN.
//
// Assumes it is compatible with all Lambda runtimes.
// Experimental.
func KubectlLayer_FromLayerVersionArn(scope constructs.Construct, id *string, layerVersionArn *string) awslambda.ILayerVersion {
	_init_.Initialize()

	var returns awslambda.ILayerVersion

	_jsii_.StaticInvoke(
		"monocdk.lambda_layer_kubectl.KubectlLayer",
		"fromLayerVersionArn",
		[]interface{}{scope, id, layerVersionArn},
		&returns,
	)

	return returns
}

// Imports a Layer that has been defined externally.
// Experimental.
func KubectlLayer_FromLayerVersionAttributes(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) awslambda.ILayerVersion {
	_init_.Initialize()

	var returns awslambda.ILayerVersion

	_jsii_.StaticInvoke(
		"monocdk.lambda_layer_kubectl.KubectlLayer",
		"fromLayerVersionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func KubectlLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.lambda_layer_kubectl.KubectlLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func KubectlLayer_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.lambda_layer_kubectl.KubectlLayer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add permission for this layer version to specific entities.
//
// Usage within
// the same account where the layer is defined is always allowed and does not
// require calling this method. Note that the principal that creates the
// Lambda function using the layer (for example, a CloudFormation changeset
// execution role) also needs to have the ``lambda:GetLayerVersion``
// permission on the layer version.
// Experimental.
func (k *jsiiProxy_KubectlLayer) AddPermission(id *string, permission *awslambda.LayerVersionPermission) {
	_jsii_.InvokeVoid(
		k,
		"addPermission",
		[]interface{}{id, permission},
	)
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (k *jsiiProxy_KubectlLayer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		k,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (k *jsiiProxy_KubectlLayer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (k *jsiiProxy_KubectlLayer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (k *jsiiProxy_KubectlLayer) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (k *jsiiProxy_KubectlLayer) OnPrepare() {
	_jsii_.InvokeVoid(
		k,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (k *jsiiProxy_KubectlLayer) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (k *jsiiProxy_KubectlLayer) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (k *jsiiProxy_KubectlLayer) Prepare() {
	_jsii_.InvokeVoid(
		k,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (k *jsiiProxy_KubectlLayer) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (k *jsiiProxy_KubectlLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (k *jsiiProxy_KubectlLayer) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

