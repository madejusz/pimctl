package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseable instead.
type DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponse struct {
    DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse
}
// NewDirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponse instantiates a new DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponse and sets the default values.
func NewDirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponse()(*DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponse) {
    m := &DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponse{
        DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse: *NewDirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateDirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseable instead.
type DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnResponseable interface {
    DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
