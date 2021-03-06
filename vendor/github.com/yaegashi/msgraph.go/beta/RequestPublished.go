// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PublishedResourceRequestBuilder is request builder for PublishedResource
type PublishedResourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns PublishedResourceRequest
func (b *PublishedResourceRequestBuilder) Request() *PublishedResourceRequest {
	return &PublishedResourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PublishedResourceRequest is request for PublishedResource
type PublishedResourceRequest struct{ BaseRequest }

// Get performs GET request for PublishedResource
func (r *PublishedResourceRequest) Get(ctx context.Context) (resObj *PublishedResource, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PublishedResource
func (r *PublishedResourceRequest) Update(ctx context.Context, reqObj *PublishedResource) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PublishedResource
func (r *PublishedResourceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
