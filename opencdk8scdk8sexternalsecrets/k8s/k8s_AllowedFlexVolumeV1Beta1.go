package k8s


// AllowedFlexVolume represents a single Flexvolume that is allowed to be used.
// Experimental.
type AllowedFlexVolumeV1Beta1 struct {
	// driver is the name of the Flexvolume driver.
	// Experimental.
	Driver *string `field:"required" json:"driver" yaml:"driver"`
}

