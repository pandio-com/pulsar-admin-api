/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type RetentionPolicies struct {
	RetentionTimeInMinutes int32 `json:"retentionTimeInMinutes,omitempty"`
	RetentionSizeInMB      int64 `json:"retentionSizeInMB,omitempty"`
}
