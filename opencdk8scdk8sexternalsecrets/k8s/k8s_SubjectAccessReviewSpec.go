package k8s


// SubjectAccessReviewSpec is a description of the access request.
//
// Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set.
// Experimental.
type SubjectAccessReviewSpec struct {
	// Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.
	// Experimental.
	Extra *map[string]*[]*string `field:"optional" json:"extra" yaml:"extra"`
	// Groups is the groups you're testing for.
	// Experimental.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// NonResourceAttributes describes information for a non-resource access request.
	// Experimental.
	NonResourceAttributes *NonResourceAttributes `field:"optional" json:"nonResourceAttributes" yaml:"nonResourceAttributes"`
	// ResourceAuthorizationAttributes describes information for a resource access request.
	// Experimental.
	ResourceAttributes *ResourceAttributes `field:"optional" json:"resourceAttributes" yaml:"resourceAttributes"`
	// UID information about the requesting user.
	// Experimental.
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
	// User is the user you're testing for.
	//
	// If you specify "User" but not "Groups", then is it interpreted as "What if User were not a member of any groups.
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
}

