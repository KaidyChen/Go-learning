/*
 * A Bit of Everything
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Contact: none@example.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package abe

type RuntimeError struct {
	Error_ string `json:"error,omitempty"`
	Code int32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Details []ProtobufAny `json:"details,omitempty"`
}
