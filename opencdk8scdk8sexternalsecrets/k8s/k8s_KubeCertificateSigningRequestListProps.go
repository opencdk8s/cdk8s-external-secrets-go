package k8s


// CertificateSigningRequestList is a collection of CertificateSigningRequest objects.
// Experimental.
type KubeCertificateSigningRequestListProps struct {
	// items is a collection of CertificateSigningRequest objects.
	// Experimental.
	Items *[]*KubeCertificateSigningRequestProps `field:"required" json:"items" yaml:"items"`
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

