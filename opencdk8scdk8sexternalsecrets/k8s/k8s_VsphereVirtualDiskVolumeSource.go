package k8s


// Represents a vSphere volume resource.
// Experimental.
type VsphereVirtualDiskVolumeSource struct {
	// Path that identifies vSphere volume vmdk.
	// Experimental.
	VolumePath *string `field:"required" json:"volumePath" yaml:"volumePath"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// Experimental.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
	// Experimental.
	StoragePolicyId *string `field:"optional" json:"storagePolicyId" yaml:"storagePolicyId"`
	// Storage Policy Based Management (SPBM) profile name.
	// Experimental.
	StoragePolicyName *string `field:"optional" json:"storagePolicyName" yaml:"storagePolicyName"`
}

