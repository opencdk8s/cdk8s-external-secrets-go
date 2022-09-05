package k8s


// ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume.
// Experimental.
type ScaleIoPersistentVolumeSource struct {
	// The host address of the ScaleIO API Gateway.
	// Experimental.
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// SecretRef references to the secret for ScaleIO user and other sensitive information.
	//
	// If this is not provided, Login operation will fail.
	// Experimental.
	SecretRef *SecretReference `field:"required" json:"secretRef" yaml:"secretRef"`
	// The name of the storage system as configured in ScaleIO.
	// Experimental.
	System *string `field:"required" json:"system" yaml:"system"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs"
	// Experimental.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The name of the ScaleIO Protection Domain for the configured storage.
	// Experimental.
	ProtectionDomain *string `field:"optional" json:"protectionDomain" yaml:"protectionDomain"`
	// Defaults to false (read/write).
	//
	// ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// Experimental.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Flag to enable/disable SSL communication with Gateway, default false.
	// Experimental.
	SslEnabled *bool `field:"optional" json:"sslEnabled" yaml:"sslEnabled"`
	// Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.
	//
	// Default is ThinProvisioned.
	// Experimental.
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
	// The ScaleIO Storage Pool associated with the protection domain.
	// Experimental.
	StoragePool *string `field:"optional" json:"storagePool" yaml:"storagePool"`
	// The name of a volume already created in the ScaleIO system that is associated with this volume source.
	// Experimental.
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

