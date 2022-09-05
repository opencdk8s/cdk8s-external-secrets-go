package k8s


// FlowSchemaList is a list of FlowSchema objects.
// Experimental.
type KubeFlowSchemaListV1Beta1Props struct {
	// `items` is a list of FlowSchemas.
	// Experimental.
	Items *[]*KubeFlowSchemaV1Beta1Props `field:"required" json:"items" yaml:"items"`
	// `metadata` is the standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

