package k8s


// VolumeAttachmentList is a collection of VolumeAttachment objects.
// Experimental.
type KubeVolumeAttachmentListV1Alpha1Props struct {
	// Items is the list of VolumeAttachments.
	// Experimental.
	Items *[]*KubeVolumeAttachmentV1Alpha1Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

