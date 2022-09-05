package k8s


// Projection that may be projected along with other supported volume types.
// Experimental.
type VolumeProjection struct {
	// information about the configMap data to project.
	// Experimental.
	ConfigMap *ConfigMapProjection `field:"optional" json:"configMap" yaml:"configMap"`
	// information about the downwardAPI data to project.
	// Experimental.
	DownwardApi *DownwardApiProjection `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// information about the secret data to project.
	// Experimental.
	Secret *SecretProjection `field:"optional" json:"secret" yaml:"secret"`
	// information about the serviceAccountToken data to project.
	// Experimental.
	ServiceAccountToken *ServiceAccountTokenProjection `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

