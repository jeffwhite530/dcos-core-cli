/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type MetronomeV1JobRunPlacementConstraints struct {
	// The attribute name for this constraint.
	Attribute string `json:"attribute"`
	// The operator for this constraint.
	Operator string `json:"operator"`
	// The value for this constraint.
	Value string `json:"value,omitempty"`
}