// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WindowsAutopilotSettingsRequestBuilder is request builder for WindowsAutopilotSettings
type WindowsAutopilotSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsAutopilotSettingsRequest
func (b *WindowsAutopilotSettingsRequestBuilder) Request() *WindowsAutopilotSettingsRequest {
	return &WindowsAutopilotSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsAutopilotSettingsRequest is request for WindowsAutopilotSettings
type WindowsAutopilotSettingsRequest struct{ BaseRequest }

// Get performs GET request for WindowsAutopilotSettings
func (r *WindowsAutopilotSettingsRequest) Get(ctx context.Context) (resObj *WindowsAutopilotSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsAutopilotSettings
func (r *WindowsAutopilotSettingsRequest) Update(ctx context.Context, reqObj *WindowsAutopilotSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsAutopilotSettings
func (r *WindowsAutopilotSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}