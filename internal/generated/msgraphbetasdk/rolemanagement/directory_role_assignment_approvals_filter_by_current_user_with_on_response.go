package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponseable instead.
type DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnResponse struct {
    DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse
}
// NewDirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnResponse instantiates a new DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnResponse and sets the default values.
func NewDirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnResponse()(*DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnResponse) {
    m := &DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnResponse{
        DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse: *NewDirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateDirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponseable instead.
type DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnResponseable interface {
    DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
