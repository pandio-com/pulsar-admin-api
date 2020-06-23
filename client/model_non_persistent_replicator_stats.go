/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type NonPersistentReplicatorStats struct {
	MsgRateIn                 float64 `json:"msgRateIn,omitempty"`
	MsgThroughputIn           float64 `json:"msgThroughputIn,omitempty"`
	MsgRateOut                float64 `json:"msgRateOut,omitempty"`
	MsgThroughputOut          float64 `json:"msgThroughputOut,omitempty"`
	MsgRateExpired            float64 `json:"msgRateExpired,omitempty"`
	ReplicationBacklog        int64   `json:"replicationBacklog,omitempty"`
	Connected                 bool    `json:"connected,omitempty"`
	ReplicationDelayInSeconds int64   `json:"replicationDelayInSeconds,omitempty"`
	InboundConnection         string  `json:"inboundConnection,omitempty"`
	InboundConnectedSince     string  `json:"inboundConnectedSince,omitempty"`
	OutboundConnection        string  `json:"outboundConnection,omitempty"`
	OutboundConnectedSince    string  `json:"outboundConnectedSince,omitempty"`
	MsgDropRate               float64 `json:"msgDropRate,omitempty"`
}
