package k8s


// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
// Experimental.
type HostAlias struct {
	// Hostnames for the above IP address.
	// Experimental.
	Hostnames *[]*string `field:"optional" json:"hostnames" yaml:"hostnames"`
	// IP address of the host file entry.
	// Experimental.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
}

