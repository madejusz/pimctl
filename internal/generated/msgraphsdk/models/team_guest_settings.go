package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TeamGuestSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // If set to true, guests can add and update channels.
    allowCreateUpdateChannels *bool
    // If set to true, guests can delete channels.
    allowDeleteChannels *bool
    // The OdataType property
    odataType *string
}
// NewTeamGuestSettings instantiates a new TeamGuestSettings and sets the default values.
func NewTeamGuestSettings()(*TeamGuestSettings) {
    m := &TeamGuestSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamGuestSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamGuestSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamGuestSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamGuestSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowCreateUpdateChannels gets the allowCreateUpdateChannels property value. If set to true, guests can add and update channels.
// returns a *bool when successful
func (m *TeamGuestSettings) GetAllowCreateUpdateChannels()(*bool) {
    return m.allowCreateUpdateChannels
}
// GetAllowDeleteChannels gets the allowDeleteChannels property value. If set to true, guests can delete channels.
// returns a *bool when successful
func (m *TeamGuestSettings) GetAllowDeleteChannels()(*bool) {
    return m.allowDeleteChannels
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamGuestSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowCreateUpdateChannels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCreateUpdateChannels(val)
        }
        return nil
    }
    res["allowDeleteChannels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeleteChannels(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TeamGuestSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TeamGuestSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowCreateUpdateChannels", m.GetAllowCreateUpdateChannels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowDeleteChannels", m.GetAllowDeleteChannels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamGuestSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowCreateUpdateChannels sets the allowCreateUpdateChannels property value. If set to true, guests can add and update channels.
func (m *TeamGuestSettings) SetAllowCreateUpdateChannels(value *bool)() {
    m.allowCreateUpdateChannels = value
}
// SetAllowDeleteChannels sets the allowDeleteChannels property value. If set to true, guests can delete channels.
func (m *TeamGuestSettings) SetAllowDeleteChannels(value *bool)() {
    m.allowDeleteChannels = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamGuestSettings) SetOdataType(value *string)() {
    m.odataType = value
}
type TeamGuestSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowCreateUpdateChannels()(*bool)
    GetAllowDeleteChannels()(*bool)
    GetOdataType()(*string)
    SetAllowCreateUpdateChannels(value *bool)()
    SetAllowDeleteChannels(value *bool)()
    SetOdataType(value *string)()
}
