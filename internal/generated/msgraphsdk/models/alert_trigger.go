package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AlertTrigger struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the property serving as a detection trigger.
    name *string
    // The OdataType property
    odataType *string
    // Type of the property in the key:value pair for interpretation. For example, String, Boolean etc.
    typeEscaped *string
    // Value of the property serving as a detection trigger.
    value *string
}
// NewAlertTrigger instantiates a new AlertTrigger and sets the default values.
func NewAlertTrigger()(*AlertTrigger) {
    m := &AlertTrigger{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAlertTriggerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAlertTriggerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertTrigger(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AlertTrigger) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AlertTrigger) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the property serving as a detection trigger.
// returns a *string when successful
func (m *AlertTrigger) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AlertTrigger) GetOdataType()(*string) {
    return m.odataType
}
// GetTypeEscaped gets the type property value. Type of the property in the key:value pair for interpretation. For example, String, Boolean etc.
// returns a *string when successful
func (m *AlertTrigger) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetValue gets the value property value. Value of the property serving as a detection trigger.
// returns a *string when successful
func (m *AlertTrigger) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *AlertTrigger) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *AlertTrigger) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Name of the property serving as a detection trigger.
func (m *AlertTrigger) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AlertTrigger) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTypeEscaped sets the type property value. Type of the property in the key:value pair for interpretation. For example, String, Boolean etc.
func (m *AlertTrigger) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. Value of the property serving as a detection trigger.
func (m *AlertTrigger) SetValue(value *string)() {
    m.value = value
}
type AlertTriggerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetOdataType()(*string)
    GetTypeEscaped()(*string)
    GetValue()(*string)
    SetName(value *string)()
    SetOdataType(value *string)()
    SetTypeEscaped(value *string)()
    SetValue(value *string)()
}
