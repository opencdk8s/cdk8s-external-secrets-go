package k8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/opencdk8s/cdk8s-external-secrets-go/opencdk8scdk8sexternalsecrets/jsii"
)

// Experimental.
type IntOrString interface {
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for IntOrString
type jsiiProxy_IntOrString struct {
	_ byte // padding
}

func (j *jsiiProxy_IntOrString) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func IntOrString_FromNumber(value *float64) IntOrString {
	_init_.Initialize()

	if err := validateIntOrString_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns IntOrString

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.IntOrString",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func IntOrString_FromString(value *string) IntOrString {
	_init_.Initialize()

	if err := validateIntOrString_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns IntOrString

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.IntOrString",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

