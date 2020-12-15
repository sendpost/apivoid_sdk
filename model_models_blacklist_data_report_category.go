/*
 * apivoid
 *
 * apivoids client for sendpost
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ModelsBlacklistDataReportCategory struct {
	IsFreeHosting bool `json:"is_free_hosting,omitempty"`
	IsAnonymizer bool `json:"is_anonymizer,omitempty"`
	IsUrlShortener bool `json:"is_url_shortener,omitempty"`
	IsFreeDynamicDns bool `json:"is_free_dynamic_dns,omitempty"`
}