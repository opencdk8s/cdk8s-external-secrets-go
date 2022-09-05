package k8s


// EndpointHints provides hints describing how an endpoint should be consumed.
// Experimental.
type EndpointHints struct {
	// forZones indicates the zone(s) this endpoint should be consumed by to enable topology aware routing.
	// Experimental.
	ForZones *[]*ForZone `field:"optional" json:"forZones" yaml:"forZones"`
}

