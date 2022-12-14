package k8s

import (
	"time"
)

// ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.
// Experimental.
type ManagedFieldsEntry struct {
	// APIVersion defines the version of this resource that this field set applies to.
	//
	// The format is "group/version" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted.
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// FieldsType is the discriminator for the different fields format and version.
	//
	// There is currently only one possible value: "FieldsV1".
	// Experimental.
	FieldsType *string `field:"optional" json:"fieldsType" yaml:"fieldsType"`
	// FieldsV1 holds the first JSON version format as described in the "FieldsV1" type.
	// Experimental.
	FieldsV1 interface{} `field:"optional" json:"fieldsV1" yaml:"fieldsV1"`
	// Manager is an identifier of the workflow managing these fields.
	// Experimental.
	Manager *string `field:"optional" json:"manager" yaml:"manager"`
	// Operation is the type of operation which lead to this ManagedFieldsEntry being created.
	//
	// The only valid values for this field are 'Apply' and 'Update'.
	// Experimental.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource.
	//
	// The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource.
	// Experimental.
	Subresource *string `field:"optional" json:"subresource" yaml:"subresource"`
	// Time is timestamp of when these fields were set.
	//
	// It should always be empty if Operation is 'Apply'.
	// Experimental.
	Time *time.Time `field:"optional" json:"time" yaml:"time"`
}

