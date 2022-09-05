package k8s


// HorizontalPodAutoscalerList is a list of horizontal pod autoscaler objects.
// Experimental.
type KubeHorizontalPodAutoscalerListV2Beta2Props struct {
	// items is the list of horizontal pod autoscaler objects.
	// Experimental.
	Items *[]*KubeHorizontalPodAutoscalerV2Beta2Props `field:"required" json:"items" yaml:"items"`
	// metadata is the standard list metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

