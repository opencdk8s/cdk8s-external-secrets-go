package k8s


// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
// Experimental.
type AzureFilePersistentVolumeSource struct {
	// the name of secret that contains Azure Storage Account Name and Key.
	// Experimental.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Share Name.
	// Experimental.
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Defaults to false (read/write).
	//
	// ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// Experimental.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod.
	// Experimental.
	SecretNamespace *string `field:"optional" json:"secretNamespace" yaml:"secretNamespace"`
}

