package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type M365AppsInstallationOptions struct {
    Entity
    // The appsForMac property
    appsForMac AppsInstallationOptionsForMacable
    // The appsForWindows property
    appsForWindows AppsInstallationOptionsForWindowsable
    // The updateChannel property
    updateChannel *AppsUpdateChannelType
}
// NewM365AppsInstallationOptions instantiates a new M365AppsInstallationOptions and sets the default values.
func NewM365AppsInstallationOptions()(*M365AppsInstallationOptions) {
    m := &M365AppsInstallationOptions{
        Entity: *NewEntity(),
    }
    return m
}
// CreateM365AppsInstallationOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateM365AppsInstallationOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewM365AppsInstallationOptions(), nil
}
// GetAppsForMac gets the appsForMac property value. The appsForMac property
// returns a AppsInstallationOptionsForMacable when successful
func (m *M365AppsInstallationOptions) GetAppsForMac()(AppsInstallationOptionsForMacable) {
    return m.appsForMac
}
// GetAppsForWindows gets the appsForWindows property value. The appsForWindows property
// returns a AppsInstallationOptionsForWindowsable when successful
func (m *M365AppsInstallationOptions) GetAppsForWindows()(AppsInstallationOptionsForWindowsable) {
    return m.appsForWindows
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *M365AppsInstallationOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appsForMac"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppsInstallationOptionsForMacFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsForMac(val.(AppsInstallationOptionsForMacable))
        }
        return nil
    }
    res["appsForWindows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppsInstallationOptionsForWindowsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsForWindows(val.(AppsInstallationOptionsForWindowsable))
        }
        return nil
    }
    res["updateChannel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppsUpdateChannelType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateChannel(val.(*AppsUpdateChannelType))
        }
        return nil
    }
    return res
}
// GetUpdateChannel gets the updateChannel property value. The updateChannel property
// returns a *AppsUpdateChannelType when successful
func (m *M365AppsInstallationOptions) GetUpdateChannel()(*AppsUpdateChannelType) {
    return m.updateChannel
}
// Serialize serializes information the current object
func (m *M365AppsInstallationOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appsForMac", m.GetAppsForMac())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("appsForWindows", m.GetAppsForWindows())
        if err != nil {
            return err
        }
    }
    if m.GetUpdateChannel() != nil {
        cast := (*m.GetUpdateChannel()).String()
        err = writer.WriteStringValue("updateChannel", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppsForMac sets the appsForMac property value. The appsForMac property
func (m *M365AppsInstallationOptions) SetAppsForMac(value AppsInstallationOptionsForMacable)() {
    m.appsForMac = value
}
// SetAppsForWindows sets the appsForWindows property value. The appsForWindows property
func (m *M365AppsInstallationOptions) SetAppsForWindows(value AppsInstallationOptionsForWindowsable)() {
    m.appsForWindows = value
}
// SetUpdateChannel sets the updateChannel property value. The updateChannel property
func (m *M365AppsInstallationOptions) SetUpdateChannel(value *AppsUpdateChannelType)() {
    m.updateChannel = value
}
type M365AppsInstallationOptionsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsForMac()(AppsInstallationOptionsForMacable)
    GetAppsForWindows()(AppsInstallationOptionsForWindowsable)
    GetUpdateChannel()(*AppsUpdateChannelType)
    SetAppsForMac(value AppsInstallationOptionsForMacable)()
    SetAppsForWindows(value AppsInstallationOptionsForWindowsable)()
    SetUpdateChannel(value *AppsUpdateChannelType)()
}
