package k8s


// Represents a Fibre Channel volume.
//
// Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
// Experimental.
type FcVolumeSource struct {
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// Experimental.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Optional: FC target lun number.
	// Experimental.
	Lun *float64 `field:"optional" json:"lun" yaml:"lun"`
	// Optional: Defaults to false (read/write).
	//
	// ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// Experimental.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Optional: FC target worldwide names (WWNs).
	// Experimental.
	TargetWwNs *[]*string `field:"optional" json:"targetWwNs" yaml:"targetWwNs"`
	// Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
	// Experimental.
	Wwids *[]*string `field:"optional" json:"wwids" yaml:"wwids"`
}

