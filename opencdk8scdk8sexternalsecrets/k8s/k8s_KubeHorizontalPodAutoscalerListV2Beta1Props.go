package k8s


// HorizontalPodAutoscaler is a list of horizontal pod autoscaler objects.
// Experimental.
type KubeHorizontalPodAutoscalerListV2Beta1Props struct {
	// items is the list of horizontal pod autoscaler objects.
	// Experimental.
	Items *[]*KubeHorizontalPodAutoscalerV2Beta1Props `field:"required" json:"items" yaml:"items"`
	// metadata is the standard list metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

