package k8s


// Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
// Experimental.
type CephFsVolumeSource struct {
	// Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	// Experimental.
	Monitors *[]*string `field:"required" json:"monitors" yaml:"monitors"`
	// Optional: Used as the mounted root, rather than the full Ceph tree, default is /.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Optional: Defaults to false (read/write).
	//
	// ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// Experimental.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	// Experimental.
	SecretFile *string `field:"optional" json:"secretFile" yaml:"secretFile"`
	// Optional: SecretRef is reference to the authentication secret for User, default is empty.
	//
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// Experimental.
	SecretRef *LocalObjectReference `field:"optional" json:"secretRef" yaml:"secretRef"`
	// Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
}

