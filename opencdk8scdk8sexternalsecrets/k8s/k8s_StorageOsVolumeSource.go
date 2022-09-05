package k8s


// Represents a StorageOS persistent volume resource.
// Experimental.
type StorageOsVolumeSource struct {
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// Experimental.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Defaults to false (read/write).
	//
	// ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// Experimental.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// SecretRef specifies the secret to use for obtaining the StorageOS API credentials.
	//
	// If not specified, default values will be attempted.
	// Experimental.
	SecretRef *LocalObjectReference `field:"optional" json:"secretRef" yaml:"secretRef"`
	// VolumeName is the human-readable name of the StorageOS volume.
	//
	// Volume names are only unique within a namespace.
	// Experimental.
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
	// VolumeNamespace specifies the scope of the volume within StorageOS.
	//
	// If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
	// Experimental.
	VolumeNamespace *string `field:"optional" json:"volumeNamespace" yaml:"volumeNamespace"`
}

