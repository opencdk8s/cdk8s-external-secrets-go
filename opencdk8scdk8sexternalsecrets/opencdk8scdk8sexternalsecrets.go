package opencdk8scdk8sexternalsecrets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.ExternalSecretData",
		reflect.TypeOf((*ExternalSecretData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.ExternalSecretDataFromRemoteRef",
		reflect.TypeOf((*ExternalSecretDataFromRemoteRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.ExternalSecretDataRemoteRef",
		reflect.TypeOf((*ExternalSecretDataRemoteRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.ExternalSecretFind",
		reflect.TypeOf((*ExternalSecretFind)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.ExternalSecretProps",
		reflect.TypeOf((*ExternalSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.ExternalSecretRewrite",
		reflect.TypeOf((*ExternalSecretRewrite)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.ExternalSecretRewriteRegexp",
		reflect.TypeOf((*ExternalSecretRewriteRegexp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.ExternalSecretSpec",
		reflect.TypeOf((*ExternalSecretSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.ExternalSecretTarget",
		reflect.TypeOf((*ExternalSecretTarget)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.ExternalSecretTemplate",
		reflect.TypeOf((*ExternalSecretTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.ExternalSecretTemplateMetadata",
		reflect.TypeOf((*ExternalSecretTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-external-secrets.ExternalSecretV1Beta1",
		reflect.TypeOf((*ExternalSecretV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ExternalSecretV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.FindName",
		reflect.TypeOf((*FindName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.SecretStoreRef",
		reflect.TypeOf((*SecretStoreRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.TemplateFrom",
		reflect.TypeOf((*TemplateFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.TemplateRef",
		reflect.TypeOf((*TemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-external-secrets.TemplateRefItem",
		reflect.TypeOf((*TemplateRefItem)(nil)).Elem(),
	)
}
