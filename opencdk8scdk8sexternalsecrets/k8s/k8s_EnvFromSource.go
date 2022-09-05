package k8s


// EnvFromSource represents the source of a set of ConfigMaps.
// Experimental.
type EnvFromSource struct {
	// The ConfigMap to select from.
	// Experimental.
	ConfigMapRef *ConfigMapEnvSource `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	// An optional identifier to prepend to each key in the ConfigMap.
	//
	// Must be a C_IDENTIFIER.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The Secret to select from.
	// Experimental.
	SecretRef *SecretEnvSource `field:"optional" json:"secretRef" yaml:"secretRef"`
}

