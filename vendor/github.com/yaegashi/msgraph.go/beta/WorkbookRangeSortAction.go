// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookRangeSortApplyRequestParameter undocumented
type WorkbookRangeSortApplyRequestParameter struct {
	// Fields undocumented
	Fields []WorkbookSortField `json:"fields,omitempty"`
	// MatchCase undocumented
	MatchCase *bool `json:"matchCase,omitempty"`
	// HasHeaders undocumented
	HasHeaders *bool `json:"hasHeaders,omitempty"`
	// Orientation undocumented
	Orientation *string `json:"orientation,omitempty"`
	// Method undocumented
	Method *string `json:"method,omitempty"`
}

//
type WorkbookRangeSortApplyRequestBuilder struct{ BaseRequestBuilder }

// Apply action undocumented
func (b *WorkbookRangeSortRequestBuilder) Apply(reqObj *WorkbookRangeSortApplyRequestParameter) *WorkbookRangeSortApplyRequestBuilder {
	bb := &WorkbookRangeSortApplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/apply"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookRangeSortApplyRequest struct{ BaseRequest }

//
func (b *WorkbookRangeSortApplyRequestBuilder) Request() *WorkbookRangeSortApplyRequest {
	return &WorkbookRangeSortApplyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookRangeSortApplyRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}