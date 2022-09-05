package k8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/opencdk8s/cdk8s-external-secrets-go/opencdk8scdk8sexternalsecrets/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/opencdk8s/cdk8s-external-secrets-go/opencdk8scdk8sexternalsecrets/k8s/internal"
)

// ReplicaSet ensures that a specified number of pod replicas are running at any given time.
// Experimental.
type KubeReplicaSet interface {
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

// The jsii proxy struct for KubeReplicaSet
type jsiiProxy_KubeReplicaSet struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_KubeReplicaSet) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeReplicaSet) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeReplicaSet) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeReplicaSet) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeReplicaSet) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeReplicaSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeReplicaSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "io.k8s.api.apps.v1.ReplicaSet" API object.
// Experimental.
func NewKubeReplicaSet(scope constructs.Construct, id *string, props *KubeReplicaSetProps) KubeReplicaSet {
	_init_.Initialize()

	if err := validateNewKubeReplicaSetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubeReplicaSet{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeReplicaSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "io.k8s.api.apps.v1.ReplicaSet" API object.
// Experimental.
func NewKubeReplicaSet_Override(k KubeReplicaSet, scope constructs.Construct, id *string, props *KubeReplicaSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeReplicaSet",
		[]interface{}{scope, id, props},
		k,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func KubeReplicaSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubeReplicaSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeReplicaSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "io.k8s.api.apps.v1.ReplicaSet".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
// Experimental.
func KubeReplicaSet_Manifest(props *KubeReplicaSetProps) interface{} {
	_init_.Initialize()

	if err := validateKubeReplicaSet_ManifestParameters(props); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeReplicaSet",
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
func KubeReplicaSet_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	if err := validateKubeReplicaSet_OfParameters(c); err != nil {
		panic(err)
	}
	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeReplicaSet",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func KubeReplicaSet_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeReplicaSet",
		"GVK",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubeReplicaSet) AddDependency(dependencies ...constructs.IConstruct) {
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

func (k *jsiiProxy_KubeReplicaSet) AddJsonPatch(ops ...cdk8s.JsonPatch) {
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

func (k *jsiiProxy_KubeReplicaSet) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeReplicaSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

