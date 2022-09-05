package k8s


// CustomResourceValidation is a list of validation methods for CustomResources.
// Experimental.
type CustomResourceValidation struct {
	// openAPIV3Schema is the OpenAPI v3 schema to use for validation and pruning.
	// Experimental.
	OpenApiv3Schema *JsonSchemaProps `field:"optional" json:"openApiv3Schema" yaml:"openApiv3Schema"`
}

