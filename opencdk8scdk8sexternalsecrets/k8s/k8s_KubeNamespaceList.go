package k8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/opencdk8s/cdk8s-external-secrets-go/opencdk8scdk8sexternalsecrets/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/opencdk8s/cdk8s-external-secrets-go/opencdk8scdk8sexternalsecrets/k8s/internal"
)

// NamespaceList is a list of Namespaces.
// Experimental.
type KubeNamespaceList interface {
	cdk8s.ApiObject
	// The group portion of the API version (e.g. `authorization.k8s.io`).
	// Experimental.
	ApiGroup() *string
	// The object's API version (e.g. `authorization.k8s.io/v1`).
	// Experimental.
	ApiVersion() *string
	// The chart in which this object is defined.
	// Experimental.
	Chart() cdk8s.Chart
	// The object kind.
	// Experimental.
	Kind() *string
	// Metadata associated with this API object.
	// Experimental.
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of the API object.
	//
	// If a name is specified in `metadata.name` this will be the name returned.
	// Otherwise, a name will be generated by calling
	// `Chart.of(this).generatedObjectName(this)`, which by default uses the
	// construct path to generate a DNS-compatible name for the resource.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Create a dependency between this ApiObject and other constructs.
	//
	// These can be other ApiObjects, Charts, or custom.
	// Experimental.
	AddDependency(dependencies ...constructs.IConstruct)
	// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
	//
	// Example:
	//     kubePod.addJsonPatch(JsonPatch.replace('/spec/enableServiceLinks', true));
	//
	// Experimental.
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	// Renders the object to Kubernetes JSON.
	// Experimental.
	ToJson() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubeNamespaceList
type jsiiProxy_KubeNamespaceList struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_KubeNamespaceList) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeNamespaceList) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeNamespaceList) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeNamespaceList) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeNamespaceList) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeNamespaceList) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeNamespaceList) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "io.k8s.api.core.v1.NamespaceList" API object.
// Experimental.
func NewKubeNamespaceList(scope constructs.Construct, id *string, props *KubeNamespaceListProps) KubeNamespaceList {
	_init_.Initialize()

	if err := validateNewKubeNamespaceListParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubeNamespaceList{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeNamespaceList",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "io.k8s.api.core.v1.NamespaceList" API object.
// Experimental.
func NewKubeNamespaceList_Override(k KubeNamespaceList, scope constructs.Construct, id *string, props *KubeNamespaceListProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeNamespaceList",
		[]interface{}{scope, id, props},
		k,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func KubeNamespaceList_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubeNamespaceList_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeNamespaceList",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "io.k8s.api.core.v1.NamespaceList".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
// Experimental.
func KubeNamespaceList_Manifest(props *KubeNamespaceListProps) interface{} {
	_init_.Initialize()

	if err := validateKubeNamespaceList_ManifestParameters(props); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeNamespaceList",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
// Experimental.
func KubeNamespaceList_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	if err := validateKubeNamespaceList_OfParameters(c); err != nil {
		panic(err)
	}
	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeNamespaceList",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func KubeNamespaceList_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeNamespaceList",
		"GVK",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubeNamespaceList) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		k,
		"addDependency",
		args,
	)
}

func (k *jsiiProxy_KubeNamespaceList) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		k,
		"addJsonPatch",
		args,
	)
}

func (k *jsiiProxy_KubeNamespaceList) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeNamespaceList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

