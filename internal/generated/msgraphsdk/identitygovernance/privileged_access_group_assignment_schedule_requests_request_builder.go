package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
    i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0 "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models/odataerrors"
)

// PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder provides operations to manage the assignmentScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderGetQueryParameters get a list of the privilegedAccessGroupAssignmentScheduleRequest objects and their properties.
type PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderGetQueryParameters
}
// PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPrivilegedAccessGroupAssignmentScheduleRequestId provides operations to manage the assignmentScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) ByPrivilegedAccessGroupAssignmentScheduleRequestId(privilegedAccessGroupAssignmentScheduleRequestId string)(*PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if privilegedAccessGroupAssignmentScheduleRequestId != "" {
        urlTplParams["privilegedAccessGroupAssignmentScheduleRequest%2Did"] = privilegedAccessGroupAssignmentScheduleRequestId
    }
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderInternal instantiates a new PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) {
    m := &PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentScheduleRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder instantiates a new PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrivilegedAccessGroupAssignmentScheduleRequestsCountRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) Count()(*PrivilegedAccessGroupAssignmentScheduleRequestsCountRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get get a list of the privilegedAccessGroupAssignmentScheduleRequest objects and their properties.
// returns a PrivilegedAccessGroupAssignmentScheduleRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroup-list-assignmentschedulerequests?view=graph-rest-1.0
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderGetRequestConfiguration)(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreatePrivilegedAccessGroupAssignmentScheduleRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestCollectionResponseable), nil
}
// Post create a new privilegedAccessGroupAssignmentScheduleRequest object.
// returns a PrivilegedAccessGroupAssignmentScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroup-post-assignmentschedulerequests?view=graph-rest-1.0
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) Post(ctx context.Context, body ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestable, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderPostRequestConfiguration)(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreatePrivilegedAccessGroupAssignmentScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestable), nil
}
// ToGetRequestInformation get a list of the privilegedAccessGroupAssignmentScheduleRequest objects and their properties.
// returns a *RequestInformation when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create a new privilegedAccessGroupAssignmentScheduleRequest object.
// returns a *RequestInformation when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestable, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) WithUrl(rawUrl string)(*PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
