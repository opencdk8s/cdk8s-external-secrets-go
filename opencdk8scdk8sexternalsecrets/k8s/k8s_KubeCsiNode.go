package k8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/opencdk8s/cdk8s-external-secrets-go/opencdk8scdk8sexternalsecrets/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/opencdk8s/cdk8s-external-secrets-go/opencdk8scdk8sexternalsecrets/k8s/internal"
)

// CSINode holds information about all CSI drivers installed on a node.
//
// CSI drivers do not need to create the CSINode object directly. As long as they use the node-driver-registrar sidecar container, the kubelet will automatically populate the CSINode object for the CSI driver as part of kubelet plugin registration. CSINode has the same name as a node. If the object is missing, it means either there are no CSI Drivers available on the node, or the Kubelet version is low enough that it doesn't create this object. CSINode has an OwnerReference that points to the corresponding node object.
// Experimental.
type KubeCsiNode interface {
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

// The jsii proxy struct for KubeCsiNode
type jsiiProxy_KubeCsiNode struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_KubeCsiNode) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeCsiNode) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeCsiNode) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeCsiNode) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeCsiNode) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeCsiNode) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeCsiNode) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "io.k8s.api.storage.v1.CSINode" API object.
// Experimental.
func NewKubeCsiNode(scope constructs.Construct, id *string, props *KubeCsiNodeProps) KubeCsiNode {
	_init_.Initialize()

	if err := validateNewKubeCsiNodeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubeCsiNode{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeCsiNode",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "io.k8s.api.storage.v1.CSINode" API object.
// Experimental.
func NewKubeCsiNode_Override(k KubeCsiNode, scope constructs.Construct, id *string, props *KubeCsiNodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeCsiNode",
		[]interface{}{scope, id, props},
		k,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func KubeCsiNode_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubeCsiNode_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeCsiNode",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "io.k8s.api.storage.v1.CSINode".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
// Experimental.
func KubeCsiNode_Manifest(props *KubeCsiNodeProps) interface{} {
	_init_.Initialize()

	if err := validateKubeCsiNode_ManifestParameters(props); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeCsiNode",
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
func KubeCsiNode_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	if err := validateKubeCsiNode_OfParameters(c); err != nil {
		panic(err)
	}
	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeCsiNode",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func KubeCsiNode_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-external-secrets.k8s.KubeCsiNode",
		"GVK",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubeCsiNode) AddDependency(dependencies ...constructs.IConstruct) {
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

func (k *jsiiProxy_KubeCsiNode) AddJsonPatch(ops ...cdk8s.JsonPatch) {
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

func (k *jsiiProxy_KubeCsiNode) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeCsiNode) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

