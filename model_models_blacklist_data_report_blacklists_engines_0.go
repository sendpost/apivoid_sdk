/*
 * apivoid
 *
 * apivoids client for sendpost
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ModelsBlacklistDataReportBlacklistsEngines0 struct {
	Engine string `json:"engine,omitempty"`
	Detected bool `json:"detected,omitempty"`
	Reference string `json:"reference,omitempty"`
	Confidence string `json:"confidence,omitempty"`
	Elapsed string `json:"elapsed,omitempty"`
}
