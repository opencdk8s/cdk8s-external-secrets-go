package k8s


// PodDisruptionBudgetList is a collection of PodDisruptionBudgets.
// Experimental.
type KubePodDisruptionBudgetListV1Beta1Props struct {
	// items list individual PodDisruptionBudget objects.
	// Experimental.
	Items *[]*KubePodDisruptionBudgetV1Beta1Props `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

