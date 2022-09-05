package k8s


// ISCSIPersistentVolumeSource represents an ISCSI disk.
//
// ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
// Experimental.
type IscsiPersistentVolumeSource struct {
	// Target iSCSI Qualified Name.
	// Experimental.
	Iqn *string `field:"required" json:"iqn" yaml:"iqn"`
	// iSCSI Target Lun number.
	// Experimental.
	Lun *float64 `field:"required" json:"lun" yaml:"lun"`
	// iSCSI Target Portal.
	//
	// The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	// Experimental.
	TargetPortal *string `field:"required" json:"targetPortal" yaml:"targetPortal"`
	// whether support iSCSI Discovery CHAP authentication.
	// Experimental.
	ChapAuthDiscovery *bool `field:"optional" json:"chapAuthDiscovery" yaml:"chapAuthDiscovery"`
	// whether support iSCSI Session CHAP authentication.
	// Experimental.
	ChapAuthSession *bool `field:"optional" json:"chapAuthSession" yaml:"chapAuthSession"`
	// Filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	// Experimental.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Custom iSCSI Initiator Name.
	//
	// If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
	// Experimental.
	InitiatorName *string `field:"optional" json:"initiatorName" yaml:"initiatorName"`
	// iSCSI Interface Name that uses an iSCSI transport.
	//
	// Defaults to 'default' (tcp).
	// Experimental.
	IscsiInterface *string `field:"optional" json:"iscsiInterface" yaml:"iscsiInterface"`
	// iSCSI Target Portal List.
	//
	// The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	// Experimental.
	Portals *[]*string `field:"optional" json:"portals" yaml:"portals"`
	// ReadOnly here will force the ReadOnly setting in VolumeMounts.
	//
	// Defaults to false.
	// Experimental.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// CHAP Secret for iSCSI target and initiator authentication.
	// Experimental.
	SecretRef *SecretReference `field:"optional" json:"secretRef" yaml:"secretRef"`
}

