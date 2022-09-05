package k8s


// SecretList is a list of Secret.
// Experimental.
type KubeSecretListProps struct {
	// Items is a list of secret objects.
	//
	// More info: https://kubernetes.io/docs/concepts/configuration/secret
	// Experimental.
	Items *[]*KubeSecretProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

