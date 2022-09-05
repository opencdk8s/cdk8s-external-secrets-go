package k8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/opencdk8s/cdk8s-external-secrets-go/opencdk8scdk8sexternalsecrets/jsii"
)

// Experimental.
type Quantity interface {
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for Quantity
type jsiiProxy_Quantity struct {
	_ byte // padding
}

func (j *jsiiProxy_Quantity) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func Quantity_FromNumber(value *float64) Quantity {
	_init_.Initialize()

	if err := validateQuantity_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns Quantity

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.Quantity",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func Quantity_FromString(value *string) Quantity {
	_init_.Initialize()

	if err := validateQuantity_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns Quantity

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-external-secrets.k8s.Quantity",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

