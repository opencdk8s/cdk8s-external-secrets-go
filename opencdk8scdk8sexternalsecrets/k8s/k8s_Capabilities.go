package k8s


// Adds and removes POSIX capabilities from running containers.
// Experimental.
type Capabilities struct {
	// Added capabilities.
	// Experimental.
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	// Experimental.
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

