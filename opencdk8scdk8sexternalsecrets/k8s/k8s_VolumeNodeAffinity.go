package k8s


// VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.
// Experimental.
type VolumeNodeAffinity struct {
	// Required specifies hard node constraints that must be met.
	// Experimental.
	Required *NodeSelector `field:"optional" json:"required" yaml:"required"`
}

