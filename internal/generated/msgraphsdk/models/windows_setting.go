package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WindowsSetting struct {
    Entity
    // A collection of setting values for a given windowsSetting.
    instances []WindowsSettingInstanceable
    // The type of setting payloads contained in the instances navigation property.
    payloadType *string
    // The settingType property
    settingType *WindowsSettingType
    // A unique identifier for the device the setting might belong to if it is of the settingType backup.
    windowsDeviceId *string
}
// NewWindowsSetting instantiates a new WindowsSetting and sets the default values.
func NewWindowsSetting()(*WindowsSetting) {
    m := &WindowsSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsSetting(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["instances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsSettingInstanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsSettingInstanceable)
                }
            }
            m.SetInstances(res)
        }
        return nil
    }
    res["payloadType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadType(val)
        }
        return nil
    }
    res["settingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsSettingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingType(val.(*WindowsSettingType))
        }
        return nil
    }
    res["windowsDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsDeviceId(val)
        }
        return nil
    }
    return res
}
// GetInstances gets the instances property value. A collection of setting values for a given windowsSetting.
// returns a []WindowsSettingInstanceable when successful
func (m *WindowsSetting) GetInstances()([]WindowsSettingInstanceable) {
    return m.instances
}
// GetPayloadType gets the payloadType property value. The type of setting payloads contained in the instances navigation property.
// returns a *string when successful
func (m *WindowsSetting) GetPayloadType()(*string) {
    return m.payloadType
}
// GetSettingType gets the settingType property value. The settingType property
// returns a *WindowsSettingType when successful
func (m *WindowsSetting) GetSettingType()(*WindowsSettingType) {
    return m.settingType
}
// GetWindowsDeviceId gets the windowsDeviceId property value. A unique identifier for the device the setting might belong to if it is of the settingType backup.
// returns a *string when successful
func (m *WindowsSetting) GetWindowsDeviceId()(*string) {
    return m.windowsDeviceId
}
// Serialize serializes information the current object
func (m *WindowsSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInstances()))
        for i, v := range m.GetInstances() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("instances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payloadType", m.GetPayloadType())
        if err != nil {
            return err
        }
    }
    if m.GetSettingType() != nil {
        cast := (*m.GetSettingType()).String()
        err = writer.WriteStringValue("settingType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("windowsDeviceId", m.GetWindowsDeviceId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInstances sets the instances property value. A collection of setting values for a given windowsSetting.
func (m *WindowsSetting) SetInstances(value []WindowsSettingInstanceable)() {
    m.instances = value
}
// SetPayloadType sets the payloadType property value. The type of setting payloads contained in the instances navigation property.
func (m *WindowsSetting) SetPayloadType(value *string)() {
    m.payloadType = value
}
// SetSettingType sets the settingType property value. The settingType property
func (m *WindowsSetting) SetSettingType(value *WindowsSettingType)() {
    m.settingType = value
}
// SetWindowsDeviceId sets the windowsDeviceId property value. A unique identifier for the device the setting might belong to if it is of the settingType backup.
func (m *WindowsSetting) SetWindowsDeviceId(value *string)() {
    m.windowsDeviceId = value
}
type WindowsSettingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInstances()([]WindowsSettingInstanceable)
    GetPayloadType()(*string)
    GetSettingType()(*WindowsSettingType)
    GetWindowsDeviceId()(*string)
    SetInstances(value []WindowsSettingInstanceable)()
    SetPayloadType(value *string)()
    SetSettingType(value *WindowsSettingType)()
    SetWindowsDeviceId(value *string)()
}
