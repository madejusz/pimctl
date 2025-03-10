package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MultiTenantOrganizationMember struct {
    DirectoryObject
    // Tenant ID of the tenant that added the tenant to the multitenant organization. Read-only.
    addedByTenantId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Date and time when the tenant was added to the multitenant organization. Read-only.
    addedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Display name of the tenant added to the multitenant organization.
    displayName *string
    // Date and time when the tenant joined the multitenant organization. Read-only.
    joinedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Role of the tenant in the multitenant organization. The possible values are: owner, member (default), unknownFutureValue. Tenants with the owner role can manage the multitenant organization but tenants with the member role can only participate in a multitenant organization. There can be multiple tenants with the owner role in a multitenant organization.
    role *MultiTenantOrganizationMemberRole
    // State of the tenant in the multitenant organization. The possible values are: pending, active, removed, unknownFutureValue. Tenants in the pending state must join the multitenant organization to participate in the multitenant organization. Tenants in the active state can participate in the multitenant organization. Tenants in the removed state are in the process of being removed from the multitenant organization. Read-only.
    state *MultiTenantOrganizationMemberState
    // Tenant ID of the Microsoft Entra tenant added to the multitenant organization. Set at the time tenant is added.Supports $filter. Key.
    tenantId *string
    // Details of the processing status for a tenant in a multitenant organization. Read-only. Nullable.
    transitionDetails MultiTenantOrganizationMemberTransitionDetailsable
}
// NewMultiTenantOrganizationMember instantiates a new MultiTenantOrganizationMember and sets the default values.
func NewMultiTenantOrganizationMember()(*MultiTenantOrganizationMember) {
    m := &MultiTenantOrganizationMember{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.multiTenantOrganizationMember"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMultiTenantOrganizationMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMultiTenantOrganizationMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMultiTenantOrganizationMember(), nil
}
// GetAddedByTenantId gets the addedByTenantId property value. Tenant ID of the tenant that added the tenant to the multitenant organization. Read-only.
// returns a *UUID when successful
func (m *MultiTenantOrganizationMember) GetAddedByTenantId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.addedByTenantId
}
// GetAddedDateTime gets the addedDateTime property value. Date and time when the tenant was added to the multitenant organization. Read-only.
// returns a *Time when successful
func (m *MultiTenantOrganizationMember) GetAddedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.addedDateTime
}
// GetDisplayName gets the displayName property value. Display name of the tenant added to the multitenant organization.
// returns a *string when successful
func (m *MultiTenantOrganizationMember) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MultiTenantOrganizationMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["addedByTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedByTenantId(val)
        }
        return nil
    }
    res["addedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedDateTime(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["joinedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinedDateTime(val)
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMultiTenantOrganizationMemberRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*MultiTenantOrganizationMemberRole))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMultiTenantOrganizationMemberState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*MultiTenantOrganizationMemberState))
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["transitionDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMultiTenantOrganizationMemberTransitionDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransitionDetails(val.(MultiTenantOrganizationMemberTransitionDetailsable))
        }
        return nil
    }
    return res
}
// GetJoinedDateTime gets the joinedDateTime property value. Date and time when the tenant joined the multitenant organization. Read-only.
// returns a *Time when successful
func (m *MultiTenantOrganizationMember) GetJoinedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.joinedDateTime
}
// GetRole gets the role property value. Role of the tenant in the multitenant organization. The possible values are: owner, member (default), unknownFutureValue. Tenants with the owner role can manage the multitenant organization but tenants with the member role can only participate in a multitenant organization. There can be multiple tenants with the owner role in a multitenant organization.
// returns a *MultiTenantOrganizationMemberRole when successful
func (m *MultiTenantOrganizationMember) GetRole()(*MultiTenantOrganizationMemberRole) {
    return m.role
}
// GetState gets the state property value. State of the tenant in the multitenant organization. The possible values are: pending, active, removed, unknownFutureValue. Tenants in the pending state must join the multitenant organization to participate in the multitenant organization. Tenants in the active state can participate in the multitenant organization. Tenants in the removed state are in the process of being removed from the multitenant organization. Read-only.
// returns a *MultiTenantOrganizationMemberState when successful
func (m *MultiTenantOrganizationMember) GetState()(*MultiTenantOrganizationMemberState) {
    return m.state
}
// GetTenantId gets the tenantId property value. Tenant ID of the Microsoft Entra tenant added to the multitenant organization. Set at the time tenant is added.Supports $filter. Key.
// returns a *string when successful
func (m *MultiTenantOrganizationMember) GetTenantId()(*string) {
    return m.tenantId
}
// GetTransitionDetails gets the transitionDetails property value. Details of the processing status for a tenant in a multitenant organization. Read-only. Nullable.
// returns a MultiTenantOrganizationMemberTransitionDetailsable when successful
func (m *MultiTenantOrganizationMember) GetTransitionDetails()(MultiTenantOrganizationMemberTransitionDetailsable) {
    return m.transitionDetails
}
// Serialize serializes information the current object
func (m *MultiTenantOrganizationMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteUUIDValue("addedByTenantId", m.GetAddedByTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("addedDateTime", m.GetAddedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("joinedDateTime", m.GetJoinedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err = writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("transitionDetails", m.GetTransitionDetails())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddedByTenantId sets the addedByTenantId property value. Tenant ID of the tenant that added the tenant to the multitenant organization. Read-only.
func (m *MultiTenantOrganizationMember) SetAddedByTenantId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.addedByTenantId = value
}
// SetAddedDateTime sets the addedDateTime property value. Date and time when the tenant was added to the multitenant organization. Read-only.
func (m *MultiTenantOrganizationMember) SetAddedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.addedDateTime = value
}
// SetDisplayName sets the displayName property value. Display name of the tenant added to the multitenant organization.
func (m *MultiTenantOrganizationMember) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetJoinedDateTime sets the joinedDateTime property value. Date and time when the tenant joined the multitenant organization. Read-only.
func (m *MultiTenantOrganizationMember) SetJoinedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.joinedDateTime = value
}
// SetRole sets the role property value. Role of the tenant in the multitenant organization. The possible values are: owner, member (default), unknownFutureValue. Tenants with the owner role can manage the multitenant organization but tenants with the member role can only participate in a multitenant organization. There can be multiple tenants with the owner role in a multitenant organization.
func (m *MultiTenantOrganizationMember) SetRole(value *MultiTenantOrganizationMemberRole)() {
    m.role = value
}
// SetState sets the state property value. State of the tenant in the multitenant organization. The possible values are: pending, active, removed, unknownFutureValue. Tenants in the pending state must join the multitenant organization to participate in the multitenant organization. Tenants in the active state can participate in the multitenant organization. Tenants in the removed state are in the process of being removed from the multitenant organization. Read-only.
func (m *MultiTenantOrganizationMember) SetState(value *MultiTenantOrganizationMemberState)() {
    m.state = value
}
// SetTenantId sets the tenantId property value. Tenant ID of the Microsoft Entra tenant added to the multitenant organization. Set at the time tenant is added.Supports $filter. Key.
func (m *MultiTenantOrganizationMember) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetTransitionDetails sets the transitionDetails property value. Details of the processing status for a tenant in a multitenant organization. Read-only. Nullable.
func (m *MultiTenantOrganizationMember) SetTransitionDetails(value MultiTenantOrganizationMemberTransitionDetailsable)() {
    m.transitionDetails = value
}
type MultiTenantOrganizationMemberable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedByTenantId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetAddedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetJoinedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRole()(*MultiTenantOrganizationMemberRole)
    GetState()(*MultiTenantOrganizationMemberState)
    GetTenantId()(*string)
    GetTransitionDetails()(MultiTenantOrganizationMemberTransitionDetailsable)
    SetAddedByTenantId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetAddedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetJoinedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRole(value *MultiTenantOrganizationMemberRole)()
    SetState(value *MultiTenantOrganizationMemberState)()
    SetTenantId(value *string)()
    SetTransitionDetails(value MultiTenantOrganizationMemberTransitionDetailsable)()
}
