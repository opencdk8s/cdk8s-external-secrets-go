package k8s


// Volume represents a named volume in a pod that may be accessed by any container in the pod.
// Experimental.
type Volume struct {
	// Volume's name.
	//
	// Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// Experimental.
	AwsElasticBlockStore *AwsElasticBlockStoreVolumeSource `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	// Experimental.
	AzureDisk *AzureDiskVolumeSource `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	// Experimental.
	AzureFile *AzureFileVolumeSource `field:"optional" json:"azureFile" yaml:"azureFile"`
	// CephFS represents a Ceph FS mount on the host that shares a pod's lifetime.
	// Experimental.
	Cephfs *CephFsVolumeSource `field:"optional" json:"cephfs" yaml:"cephfs"`
	// Cinder represents a cinder volume attached and mounted on kubelets host machine.
	//
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// Experimental.
	Cinder *CinderVolumeSource `field:"optional" json:"cinder" yaml:"cinder"`
	// ConfigMap represents a configMap that should populate this volume.
	// Experimental.
	ConfigMap *ConfigMapVolumeSource `field:"optional" json:"configMap" yaml:"configMap"`
	// CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).
	// Experimental.
	Csi *CsiVolumeSource `field:"optional" json:"csi" yaml:"csi"`
	// DownwardAPI represents downward API about the pod that should populate this volume.
	// Experimental.
	DownwardApi *DownwardApiVolumeSource `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// EmptyDir represents a temporary directory that shares a pod's lifetime.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	// Experimental.
	EmptyDir *EmptyDirVolumeSource `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// Ephemeral represents a volume that is handled by a cluster storage driver.
	//
	// The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.
	//
	// Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity
	// tracking are needed,
	// c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through
	// a PersistentVolumeClaim (see EphemeralVolumeSource for more
	// information on the connection between this volume type
	// and PersistentVolumeClaim).
	//
	// Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.
	//
	// Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.
	//
	// A pod can use both types of ephemeral volumes and persistent volumes at the same time.
	//
	// This is a beta feature and only available when the GenericEphemeralVolume feature gate is enabled.
	// Experimental.
	Ephemeral *EphemeralVolumeSource `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	// FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	// Experimental.
	Fc *FcVolumeSource `field:"optional" json:"fc" yaml:"fc"`
	// FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
	// Experimental.
	FlexVolume *FlexVolumeSource `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	// Flocker represents a Flocker volume attached to a kubelet's host machine.
	//
	// This depends on the Flocker control service being running.
	// Experimental.
	Flocker *FlockerVolumeSource `field:"optional" json:"flocker" yaml:"flocker"`
	// GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// Experimental.
	GcePersistentDisk *GcePersistentDiskVolumeSource `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	// GitRepo represents a git repository at a particular revision.
	//
	// DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.
	// Experimental.
	GitRepo *GitRepoVolumeSource `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	// Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.
	//
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md
	// Experimental.
	Glusterfs *GlusterfsVolumeSource `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	// HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container.
	//
	// This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// Experimental.
	HostPath *HostPathVolumeSource `field:"optional" json:"hostPath" yaml:"hostPath"`
	// ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod.
	//
	// More info: https://examples.k8s.io/volumes/iscsi/README.md
	// Experimental.
	Iscsi *IscsiVolumeSource `field:"optional" json:"iscsi" yaml:"iscsi"`
	// NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs.
	// Experimental.
	Nfs *NfsVolumeSource `field:"optional" json:"nfs" yaml:"nfs"`
	// PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// Experimental.
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	// PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine.
	// Experimental.
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSource `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	// PortworxVolume represents a portworx volume attached and mounted on kubelets host machine.
	// Experimental.
	PortworxVolume *PortworxVolumeSource `field:"optional" json:"portworxVolume" yaml:"portworxVolume"`
	// Items for all in one resources secrets, configmaps, and downward API.
	// Experimental.
	Projected *ProjectedVolumeSource `field:"optional" json:"projected" yaml:"projected"`
	// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime.
	// Experimental.
	Quobyte *QuobyteVolumeSource `field:"optional" json:"quobyte" yaml:"quobyte"`
	// RBD represents a Rados Block Device mount on the host that shares a pod's lifetime.
	//
	// More info: https://examples.k8s.io/volumes/rbd/README.md
	// Experimental.
	Rbd *RbdVolumeSource `field:"optional" json:"rbd" yaml:"rbd"`
	// ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	// Experimental.
	ScaleIo *ScaleIoVolumeSource `field:"optional" json:"scaleIo" yaml:"scaleIo"`
	// Secret represents a secret that should populate this volume.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	// Experimental.
	Secret *SecretVolumeSource `field:"optional" json:"secret" yaml:"secret"`
	// StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.
	// Experimental.
	Storageos *StorageOsVolumeSource `field:"optional" json:"storageos" yaml:"storageos"`
	// VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine.
	// Experimental.
	VsphereVolume *VsphereVirtualDiskVolumeSource `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

