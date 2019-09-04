/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore
// TypeHolderExample struct for TypeHolderExample
type TypeHolderExample struct {
	StringItem string `json:"string_item" xml:"string_item"`
	NumberItem float32 `json:"number_item" xml:"number_item"`
	IntegerItem int32 `json:"integer_item" xml:"integer_item"`
	BoolItem bool `json:"bool_item" xml:"bool_item"`
	ArrayItem []int32 `json:"array_item" xml:"array_item"`
}
