package k8s


// volumeDevice describes a mapping of a raw block device within a container.
// Experimental.
type VolumeDevice struct {
	// devicePath is the path inside of the container that the device will be mapped to.
	// Experimental.
	DevicePath *string `field:"required" json:"devicePath" yaml:"devicePath"`
	// name must match the name of a persistentVolumeClaim in the pod.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

