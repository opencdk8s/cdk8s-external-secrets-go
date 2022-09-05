//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// @opencdk8s/cdk8s-external-secrets
package opencdk8scdk8sexternalsecrets

// Building without runtime type checking enabled, so all the below just return nil

func validateExternalSecretV1Beta1_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExternalSecretV1Beta1_ManifestParameters(props *ExternalSecretProps) error {
	return nil
}

func validateExternalSecretV1Beta1_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewExternalSecretV1Beta1Parameters(scope constructs.Construct, id *string, props *ExternalSecretProps) error {
	return nil
}

