// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datacloudinitconfig

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataCloudinitConfigPartList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataCloudinitConfigPartList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataCloudinitConfigPartList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataCloudinitConfigPartList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataCloudinitConfigPartList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataCloudinitConfigPartList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataCloudinitConfigPartList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataCloudinitConfigPartListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

