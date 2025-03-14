package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CommunicationsApplicationInstanceIdentity struct {
    Identity
    // True if the participant shouldn't be shown in other participants' rosters.
    hidden *bool
    // The tenant ID of the application.
    tenantId *string
}
// NewCommunicationsApplicationInstanceIdentity instantiates a new CommunicationsApplicationInstanceIdentity and sets the default values.
func NewCommunicationsApplicationInstanceIdentity()(*CommunicationsApplicationInstanceIdentity) {
    m := &CommunicationsApplicationInstanceIdentity{
        Identity: *NewIdentity(),
    }
    odataTypeValue := "#microsoft.graph.communicationsApplicationInstanceIdentity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCommunicationsApplicationInstanceIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommunicationsApplicationInstanceIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsApplicationInstanceIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CommunicationsApplicationInstanceIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["hidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHidden(val)
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
    return res
}
// GetHidden gets the hidden property value. True if the participant shouldn't be shown in other participants' rosters.
// returns a *bool when successful
func (m *CommunicationsApplicationInstanceIdentity) GetHidden()(*bool) {
    return m.hidden
}
// GetTenantId gets the tenantId property value. The tenant ID of the application.
// returns a *string when successful
func (m *CommunicationsApplicationInstanceIdentity) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *CommunicationsApplicationInstanceIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("hidden", m.GetHidden())
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
    return nil
}
// SetHidden sets the hidden property value. True if the participant shouldn't be shown in other participants' rosters.
func (m *CommunicationsApplicationInstanceIdentity) SetHidden(value *bool)() {
    m.hidden = value
}
// SetTenantId sets the tenantId property value. The tenant ID of the application.
func (m *CommunicationsApplicationInstanceIdentity) SetTenantId(value *string)() {
    m.tenantId = value
}
type CommunicationsApplicationInstanceIdentityable interface {
    Identityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHidden()(*bool)
    GetTenantId()(*string)
    SetHidden(value *bool)()
    SetTenantId(value *string)()
}
