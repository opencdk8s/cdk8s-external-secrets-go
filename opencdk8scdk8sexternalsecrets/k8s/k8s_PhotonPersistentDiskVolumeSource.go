package k8s


// Represents a Photon Controller persistent disk resource.
// Experimental.
type PhotonPersistentDiskVolumeSource struct {
	// ID that identifies Photon Controller persistent disk.
	// Experimental.
	PdId *string `field:"required" json:"pdId" yaml:"pdId"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// Experimental.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
}

