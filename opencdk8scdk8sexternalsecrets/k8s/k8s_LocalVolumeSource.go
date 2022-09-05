package k8s


// Local represents directly-attached storage with node affinity (Beta feature).
// Experimental.
type LocalVolumeSource struct {
	// The full path to the volume on the node.
	//
	// It can be either a directory or block device (disk, partition, ...).
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Filesystem type to mount.
	//
	// It applies only when the Path is a block device. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default value is to auto-select a fileystem if unspecified.
	// Experimental.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
}

