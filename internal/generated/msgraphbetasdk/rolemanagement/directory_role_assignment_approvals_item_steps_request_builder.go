package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc "github.com/co-native-ab/pimctl/internal/generated/msgraphbetasdk/models"
    i3c348b42299dea02992e75ce229fbe66a0442ea71b3fb1c422dfbe0720f96f97 "github.com/co-native-ab/pimctl/internal/generated/msgraphbetasdk/models/odataerrors"
)

// DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder provides operations to manage the steps property of the microsoft.graph.approval entity.
type DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleAssignmentApprovalsItemStepsRequestBuilderGetQueryParameters used to represent the decision associated with a single step in the approval process configured in approvalStage.
type DirectoryRoleAssignmentApprovalsItemStepsRequestBuilderGetQueryParameters struct {
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
// DirectoryRoleAssignmentApprovalsItemStepsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentApprovalsItemStepsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilderGetQueryParameters
}
// DirectoryRoleAssignmentApprovalsItemStepsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentApprovalsItemStepsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByApprovalStepId provides operations to manage the steps property of the microsoft.graph.approval entity.
// returns a *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder when successful
func (m *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder) ByApprovalStepId(approvalStepId string)(*DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if approvalStepId != "" {
        urlTplParams["approvalStep%2Did"] = approvalStepId
    }
    return NewDirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoleAssignmentApprovalsItemStepsRequestBuilderInternal instantiates a new DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentApprovalsItemStepsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder) {
    m := &DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentApprovals/{approval%2Did}/steps{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDirectoryRoleAssignmentApprovalsItemStepsRequestBuilder instantiates a new DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentApprovalsItemStepsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleAssignmentApprovalsItemStepsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder when successful
func (m *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder) Count()(*DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder) {
    return NewDirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get used to represent the decision associated with a single step in the approval process configured in approvalStage.
// returns a ApprovalStepCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilderGetRequestConfiguration)(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.ApprovalStepCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i3c348b42299dea02992e75ce229fbe66a0442ea71b3fb1c422dfbe0720f96f97.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.CreateApprovalStepCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.ApprovalStepCollectionResponseable), nil
}
// Post create new navigation property to steps for roleManagement
// returns a ApprovalStepable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder) Post(ctx context.Context, body ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.ApprovalStepable, requestConfiguration *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilderPostRequestConfiguration)(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.ApprovalStepable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i3c348b42299dea02992e75ce229fbe66a0442ea71b3fb1c422dfbe0720f96f97.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.CreateApprovalStepFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.ApprovalStepable), nil
}
// ToGetRequestInformation used to represent the decision associated with a single step in the approval process configured in approvalStage.
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to steps for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.ApprovalStepable, requestConfiguration *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder when successful
func (m *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder) {
    return NewDirectoryRoleAssignmentApprovalsItemStepsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
