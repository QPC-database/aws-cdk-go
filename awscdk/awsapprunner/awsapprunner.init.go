package awsapprunner

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_apprunner.CfnService",
		reflect.TypeOf((*CfnService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceArn", GoGetter: "AttrServiceArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceId", GoGetter: "AttrServiceId"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceUrl", GoGetter: "AttrServiceUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingConfigurationArn", GoGetter: "AutoScalingConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckConfiguration", GoGetter: "HealthCheckConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceConfiguration", GoGetter: "InstanceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceConfiguration", GoGetter: "SourceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnService{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.AuthenticationConfigurationProperty",
		reflect.TypeOf((*CfnService_AuthenticationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.CodeConfigurationProperty",
		reflect.TypeOf((*CfnService_CodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.CodeConfigurationValuesProperty",
		reflect.TypeOf((*CfnService_CodeConfigurationValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.CodeRepositoryProperty",
		reflect.TypeOf((*CfnService_CodeRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnService_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.HealthCheckConfigurationProperty",
		reflect.TypeOf((*CfnService_HealthCheckConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.ImageConfigurationProperty",
		reflect.TypeOf((*CfnService_ImageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.ImageRepositoryProperty",
		reflect.TypeOf((*CfnService_ImageRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.InstanceConfigurationProperty",
		reflect.TypeOf((*CfnService_InstanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.KeyValuePairProperty",
		reflect.TypeOf((*CfnService_KeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.SourceCodeVersionProperty",
		reflect.TypeOf((*CfnService_SourceCodeVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.SourceConfigurationProperty",
		reflect.TypeOf((*CfnService_SourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnServiceProps",
		reflect.TypeOf((*CfnServiceProps)(nil)).Elem(),
	)
}
