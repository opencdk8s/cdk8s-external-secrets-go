package k8s


// SessionAffinityConfig represents the configurations of session affinity.
// Experimental.
type SessionAffinityConfig struct {
	// clientIP contains the configurations of Client IP based session affinity.
	// Experimental.
	ClientIp *ClientIpConfig `field:"optional" json:"clientIp" yaml:"clientIp"`
}

