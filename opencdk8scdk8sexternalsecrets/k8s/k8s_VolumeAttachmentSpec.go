package k8s


// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
// Experimental.
type VolumeAttachmentSpec struct {
	// Attacher indicates the name of the volume driver that MUST handle this request.
	//
	// This is the name returned by GetPluginName().
	// Experimental.
	Attacher *string `field:"required" json:"attacher" yaml:"attacher"`
	// The node that the volume should be attached to.
	// Experimental.
	NodeName *string `field:"required" json:"nodeName" yaml:"nodeName"`
	// Source represents the volume that should be attached.
	// Experimental.
	Source *VolumeAttachmentSource `field:"required" json:"source" yaml:"source"`
}

