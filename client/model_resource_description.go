/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type ResourceDescription struct {
	UsagePct      int32                    `json:"usagePct,omitempty"`
	ResourceUsage map[string]ResourceUsage `json:"resourceUsage,omitempty"`
}
