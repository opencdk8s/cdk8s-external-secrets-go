package k8s


// list of horizontal pod autoscaler objects.
// Experimental.
type KubeHorizontalPodAutoscalerListProps struct {
	// list of horizontal pod autoscaler objects.
	// Experimental.
	Items *[]*KubeHorizontalPodAutoscalerProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

