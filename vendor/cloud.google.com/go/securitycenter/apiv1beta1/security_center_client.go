// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package securitycenter

import (
	"context"
	"math"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1beta1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateSource               []gax.CallOption
	CreateFinding              []gax.CallOption
	GetIamPolicy               []gax.CallOption
	GetOrganizationSettings    []gax.CallOption
	GetSource                  []gax.CallOption
	GroupAssets                []gax.CallOption
	GroupFindings              []gax.CallOption
	ListAssets                 []gax.CallOption
	ListFindings               []gax.CallOption
	ListSources                []gax.CallOption
	RunAssetDiscovery          []gax.CallOption
	SetFindingState            []gax.CallOption
	SetIamPolicy               []gax.CallOption
	TestIamPermissions         []gax.CallOption
	UpdateFinding              []gax.CallOption
	UpdateOrganizationSettings []gax.CallOption
	UpdateSource               []gax.CallOption
	UpdateSecurityMarks        []gax.CallOption
}

func defaultClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("securitycenter.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultCallOptions() *CallOptions {
	retry := map[[2]string][]gax.CallOption{
		{"default", "idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
	}
	return &CallOptions{
		CreateSource:               retry[[2]string{"default", "non_idempotent"}],
		CreateFinding:              retry[[2]string{"default", "non_idempotent"}],
		GetIamPolicy:               retry[[2]string{"default", "idempotent"}],
		GetOrganizationSettings:    retry[[2]string{"default", "idempotent"}],
		GetSource:                  retry[[2]string{"default", "idempotent"}],
		GroupAssets:                retry[[2]string{"default", "idempotent"}],
		GroupFindings:              retry[[2]string{"default", "idempotent"}],
		ListAssets:                 retry[[2]string{"default", "idempotent"}],
		ListFindings:               retry[[2]string{"default", "idempotent"}],
		ListSources:                retry[[2]string{"default", "idempotent"}],
		RunAssetDiscovery:          retry[[2]string{"default", "non_idempotent"}],
		SetFindingState:            retry[[2]string{"default", "non_idempotent"}],
		SetIamPolicy:               retry[[2]string{"default", "non_idempotent"}],
		TestIamPermissions:         retry[[2]string{"default", "idempotent"}],
		UpdateFinding:              retry[[2]string{"default", "non_idempotent"}],
		UpdateOrganizationSettings: retry[[2]string{"default", "non_idempotent"}],
		UpdateSource:               retry[[2]string{"default", "non_idempotent"}],
		UpdateSecurityMarks:        retry[[2]string{"default", "non_idempotent"}],
	}
}

// Client is a client for interacting with Cloud Security Command Center API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type Client struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	client securitycenterpb.SecurityCenterClient

	// LROClient is used internally to handle longrunning operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

	// The call options for this service.
	CallOptions *CallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new security center client.
//
// V1 Beta APIs for Security Center service.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:        conn,
		CallOptions: defaultCallOptions(),

		client: securitycenterpb.NewSecurityCenterClient(conn),
	}
	c.setGoogleClientInfo()

	c.LROClient, err = lroauto.NewOperationsClient(ctx, option.WithGRPCConn(conn))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection
		// and never actually need to dial.
		// If this does happen, we could leak conn. However, we cannot close conn:
		// If the user invoked the function with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO(pongad): investigate error conditions.
		return nil, err
	}
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *Client) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateSource creates a source.
func (c *Client) CreateSource(ctx context.Context, req *securitycenterpb.CreateSourceRequest, opts ...gax.CallOption) (*securitycenterpb.Source, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.CreateSource[0:len(c.CallOptions.CreateSource):len(c.CallOptions.CreateSource)], opts...)
	var resp *securitycenterpb.Source
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateSource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateFinding creates a finding. The corresponding source must exist for finding creation
// to succeed.
func (c *Client) CreateFinding(ctx context.Context, req *securitycenterpb.CreateFindingRequest, opts ...gax.CallOption) (*securitycenterpb.Finding, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.CreateFinding[0:len(c.CallOptions.CreateFinding):len(c.CallOptions.CreateFinding)], opts...)
	var resp *securitycenterpb.Finding
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateFinding(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetIamPolicy gets the access control policy on the specified Source.
func (c *Client) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GetIamPolicy[0:len(c.CallOptions.GetIamPolicy):len(c.CallOptions.GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetOrganizationSettings gets the settings for an organization.
func (c *Client) GetOrganizationSettings(ctx context.Context, req *securitycenterpb.GetOrganizationSettingsRequest, opts ...gax.CallOption) (*securitycenterpb.OrganizationSettings, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GetOrganizationSettings[0:len(c.CallOptions.GetOrganizationSettings):len(c.CallOptions.GetOrganizationSettings)], opts...)
	var resp *securitycenterpb.OrganizationSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetOrganizationSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSource gets a source.
func (c *Client) GetSource(ctx context.Context, req *securitycenterpb.GetSourceRequest, opts ...gax.CallOption) (*securitycenterpb.Source, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GetSource[0:len(c.CallOptions.GetSource):len(c.CallOptions.GetSource)], opts...)
	var resp *securitycenterpb.Source
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetSource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GroupAssets filters an organization's assets and  groups them by their specified
// properties.
func (c *Client) GroupAssets(ctx context.Context, req *securitycenterpb.GroupAssetsRequest, opts ...gax.CallOption) *GroupResultIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GroupAssets[0:len(c.CallOptions.GroupAssets):len(c.CallOptions.GroupAssets)], opts...)
	it := &GroupResultIterator{}
	req = proto.Clone(req).(*securitycenterpb.GroupAssetsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*securitycenterpb.GroupResult, string, error) {
		var resp *securitycenterpb.GroupAssetsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.GroupAssets(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.GroupByResults, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// GroupFindings filters an organization or source's findings and  groups them by their
// specified properties.
//
// To group across all sources provide a - as the source id.
// Example: /v1beta1/organizations/123/sources/-/findings
func (c *Client) GroupFindings(ctx context.Context, req *securitycenterpb.GroupFindingsRequest, opts ...gax.CallOption) *GroupResultIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GroupFindings[0:len(c.CallOptions.GroupFindings):len(c.CallOptions.GroupFindings)], opts...)
	it := &GroupResultIterator{}
	req = proto.Clone(req).(*securitycenterpb.GroupFindingsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*securitycenterpb.GroupResult, string, error) {
		var resp *securitycenterpb.GroupFindingsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.GroupFindings(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.GroupByResults, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// ListAssets lists an organization's assets.
func (c *Client) ListAssets(ctx context.Context, req *securitycenterpb.ListAssetsRequest, opts ...gax.CallOption) *ListAssetsResponse_ListAssetsResultIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListAssets[0:len(c.CallOptions.ListAssets):len(c.CallOptions.ListAssets)], opts...)
	it := &ListAssetsResponse_ListAssetsResultIterator{}
	req = proto.Clone(req).(*securitycenterpb.ListAssetsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*securitycenterpb.ListAssetsResponse_ListAssetsResult, string, error) {
		var resp *securitycenterpb.ListAssetsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListAssets(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.ListAssetsResults, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// ListFindings lists an organization or source's findings.
//
// To list across all sources provide a - as the source id.
// Example: /v1beta1/organizations/123/sources/-/findings
func (c *Client) ListFindings(ctx context.Context, req *securitycenterpb.ListFindingsRequest, opts ...gax.CallOption) *FindingIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListFindings[0:len(c.CallOptions.ListFindings):len(c.CallOptions.ListFindings)], opts...)
	it := &FindingIterator{}
	req = proto.Clone(req).(*securitycenterpb.ListFindingsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*securitycenterpb.Finding, string, error) {
		var resp *securitycenterpb.ListFindingsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListFindings(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.Findings, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// ListSources lists all sources belonging to an organization.
func (c *Client) ListSources(ctx context.Context, req *securitycenterpb.ListSourcesRequest, opts ...gax.CallOption) *SourceIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListSources[0:len(c.CallOptions.ListSources):len(c.CallOptions.ListSources)], opts...)
	it := &SourceIterator{}
	req = proto.Clone(req).(*securitycenterpb.ListSourcesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*securitycenterpb.Source, string, error) {
		var resp *securitycenterpb.ListSourcesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListSources(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.Sources, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// RunAssetDiscovery runs asset discovery. The discovery is tracked with a long-running
// operation.
//
// This API can only be called with limited frequency for an organization. If
// it is called too frequently the caller will receive a TOO_MANY_REQUESTS
// error.
func (c *Client) RunAssetDiscovery(ctx context.Context, req *securitycenterpb.RunAssetDiscoveryRequest, opts ...gax.CallOption) (*RunAssetDiscoveryOperation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.RunAssetDiscovery[0:len(c.CallOptions.RunAssetDiscovery):len(c.CallOptions.RunAssetDiscovery)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.RunAssetDiscovery(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RunAssetDiscoveryOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// SetFindingState updates the state of a finding.
func (c *Client) SetFindingState(ctx context.Context, req *securitycenterpb.SetFindingStateRequest, opts ...gax.CallOption) (*securitycenterpb.Finding, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.SetFindingState[0:len(c.CallOptions.SetFindingState):len(c.CallOptions.SetFindingState)], opts...)
	var resp *securitycenterpb.Finding
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.SetFindingState(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SetIamPolicy sets the access control policy on the specified Source.
func (c *Client) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.SetIamPolicy[0:len(c.CallOptions.SetIamPolicy):len(c.CallOptions.SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TestIamPermissions returns the permissions that a caller has on the specified source.
func (c *Client) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.TestIamPermissions[0:len(c.CallOptions.TestIamPermissions):len(c.CallOptions.TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateFinding creates or updates a finding. The corresponding source must exist for a
// finding creation to succeed.
func (c *Client) UpdateFinding(ctx context.Context, req *securitycenterpb.UpdateFindingRequest, opts ...gax.CallOption) (*securitycenterpb.Finding, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.UpdateFinding[0:len(c.CallOptions.UpdateFinding):len(c.CallOptions.UpdateFinding)], opts...)
	var resp *securitycenterpb.Finding
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateFinding(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateOrganizationSettings updates an organization's settings.
func (c *Client) UpdateOrganizationSettings(ctx context.Context, req *securitycenterpb.UpdateOrganizationSettingsRequest, opts ...gax.CallOption) (*securitycenterpb.OrganizationSettings, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.UpdateOrganizationSettings[0:len(c.CallOptions.UpdateOrganizationSettings):len(c.CallOptions.UpdateOrganizationSettings)], opts...)
	var resp *securitycenterpb.OrganizationSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateOrganizationSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateSource updates a source.
func (c *Client) UpdateSource(ctx context.Context, req *securitycenterpb.UpdateSourceRequest, opts ...gax.CallOption) (*securitycenterpb.Source, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.UpdateSource[0:len(c.CallOptions.UpdateSource):len(c.CallOptions.UpdateSource)], opts...)
	var resp *securitycenterpb.Source
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateSource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateSecurityMarks updates security marks.
func (c *Client) UpdateSecurityMarks(ctx context.Context, req *securitycenterpb.UpdateSecurityMarksRequest, opts ...gax.CallOption) (*securitycenterpb.SecurityMarks, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.UpdateSecurityMarks[0:len(c.CallOptions.UpdateSecurityMarks):len(c.CallOptions.UpdateSecurityMarks)], opts...)
	var resp *securitycenterpb.SecurityMarks
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateSecurityMarks(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FindingIterator manages a stream of *securitycenterpb.Finding.
type FindingIterator struct {
	items    []*securitycenterpb.Finding
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.Finding, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *FindingIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *FindingIterator) Next() (*securitycenterpb.Finding, error) {
	var item *securitycenterpb.Finding
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *FindingIterator) bufLen() int {
	return len(it.items)
}

func (it *FindingIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// GroupResultIterator manages a stream of *securitycenterpb.GroupResult.
type GroupResultIterator struct {
	items    []*securitycenterpb.GroupResult
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.GroupResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *GroupResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *GroupResultIterator) Next() (*securitycenterpb.GroupResult, error) {
	var item *securitycenterpb.GroupResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *GroupResultIterator) bufLen() int {
	return len(it.items)
}

func (it *GroupResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ListAssetsResponse_ListAssetsResultIterator manages a stream of *securitycenterpb.ListAssetsResponse_ListAssetsResult.
type ListAssetsResponse_ListAssetsResultIterator struct {
	items    []*securitycenterpb.ListAssetsResponse_ListAssetsResult
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.ListAssetsResponse_ListAssetsResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ListAssetsResponse_ListAssetsResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ListAssetsResponse_ListAssetsResultIterator) Next() (*securitycenterpb.ListAssetsResponse_ListAssetsResult, error) {
	var item *securitycenterpb.ListAssetsResponse_ListAssetsResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ListAssetsResponse_ListAssetsResultIterator) bufLen() int {
	return len(it.items)
}

func (it *ListAssetsResponse_ListAssetsResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// SourceIterator manages a stream of *securitycenterpb.Source.
type SourceIterator struct {
	items    []*securitycenterpb.Source
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*securitycenterpb.Source, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SourceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SourceIterator) Next() (*securitycenterpb.Source, error) {
	var item *securitycenterpb.Source
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SourceIterator) bufLen() int {
	return len(it.items)
}

func (it *SourceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// RunAssetDiscoveryOperation manages a long-running operation from RunAssetDiscovery.
type RunAssetDiscoveryOperation struct {
	lro *longrunning.Operation
}

// RunAssetDiscoveryOperation returns a new RunAssetDiscoveryOperation from a given name.
// The name must be that of a previously created RunAssetDiscoveryOperation, possibly from a different process.
func (c *Client) RunAssetDiscoveryOperation(name string) *RunAssetDiscoveryOperation {
	return &RunAssetDiscoveryOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning any error encountered.
//
// See documentation of Poll for error-handling information.
func (op *RunAssetDiscoveryOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, 5000*time.Millisecond, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully, op.Done will return true.
func (op *RunAssetDiscoveryOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Done reports whether the long-running operation has completed.
func (op *RunAssetDiscoveryOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *RunAssetDiscoveryOperation) Name() string {
	return op.lro.Name()
}
