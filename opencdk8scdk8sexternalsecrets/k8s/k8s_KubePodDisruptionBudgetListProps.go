package k8s


// PodDisruptionBudgetList is a collection of PodDisruptionBudgets.
// Experimental.
type KubePodDisruptionBudgetListProps struct {
	// Items is a list of PodDisruptionBudgets.
	// Experimental.
	Items *[]*KubePodDisruptionBudgetProps `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}
