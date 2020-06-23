/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type ResourceQuota struct {
	MsgRateIn    float64 `json:"msgRateIn,omitempty"`
	MsgRateOut   float64 `json:"msgRateOut,omitempty"`
	BandwidthIn  float64 `json:"bandwidthIn,omitempty"`
	BandwidthOut float64 `json:"bandwidthOut,omitempty"`
	Memory       float64 `json:"memory,omitempty"`
	Dynamic      bool    `json:"dynamic,omitempty"`
}
