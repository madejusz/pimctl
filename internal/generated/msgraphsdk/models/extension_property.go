package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExtensionProperty struct {
    DirectoryObject
    // Display name of the application object on which this extension property is defined. Read-only.
    appDisplayName *string
    // Specifies the data type of the value the extension property can hold. Following values are supported. Binary - 256 bytes maximumBooleanDateTime - Must be specified in ISO 8601 format. Will be stored in UTC.Integer - 32-bit value.LargeInteger - 64-bit value.String - 256 characters maximumNot nullable. For multivalued directory extensions, these limits apply per value in the collection.
    dataType *string
    // Defines the directory extension as a multi-valued property. When true, the directory extension property can store a collection of objects of the dataType; for example, a collection of string types such as 'extensionb7b1c57b532f40b8b5ed4b7a7ba67401jobGroupTracker': ['String 1', 'String 2']. The default value is false. Supports $filter (eq).
    isMultiValued *bool
    // Indicates if this extension property was synced from on-premises active directory using Microsoft Entra Connect. Read-only.
    isSyncedFromOnPremises *bool
    // Name of the extension property. Not nullable. Supports $filter (eq).
    name *string
    // Following values are supported. Not nullable. UserGroupAdministrativeUnitApplicationDeviceOrganization
    targetObjects []string
}
// NewExtensionProperty instantiates a new ExtensionProperty and sets the default values.
func NewExtensionProperty()(*ExtensionProperty) {
    m := &ExtensionProperty{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.extensionProperty"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateExtensionPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExtensionPropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExtensionProperty(), nil
}
// GetAppDisplayName gets the appDisplayName property value. Display name of the application object on which this extension property is defined. Read-only.
// returns a *string when successful
func (m *ExtensionProperty) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetDataType gets the dataType property value. Specifies the data type of the value the extension property can hold. Following values are supported. Binary - 256 bytes maximumBooleanDateTime - Must be specified in ISO 8601 format. Will be stored in UTC.Integer - 32-bit value.LargeInteger - 64-bit value.String - 256 characters maximumNot nullable. For multivalued directory extensions, these limits apply per value in the collection.
// returns a *string when successful
func (m *ExtensionProperty) GetDataType()(*string) {
    return m.dataType
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExtensionProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["dataType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataType(val)
        }
        return nil
    }
    res["isMultiValued"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMultiValued(val)
        }
        return nil
    }
    res["isSyncedFromOnPremises"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSyncedFromOnPremises(val)
        }
        return nil
    }
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
    res["targetObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetTargetObjects(res)
        }
        return nil
    }
    return res
}
// GetIsMultiValued gets the isMultiValued property value. Defines the directory extension as a multi-valued property. When true, the directory extension property can store a collection of objects of the dataType; for example, a collection of string types such as 'extensionb7b1c57b532f40b8b5ed4b7a7ba67401jobGroupTracker': ['String 1', 'String 2']. The default value is false. Supports $filter (eq).
// returns a *bool when successful
func (m *ExtensionProperty) GetIsMultiValued()(*bool) {
    return m.isMultiValued
}
// GetIsSyncedFromOnPremises gets the isSyncedFromOnPremises property value. Indicates if this extension property was synced from on-premises active directory using Microsoft Entra Connect. Read-only.
// returns a *bool when successful
func (m *ExtensionProperty) GetIsSyncedFromOnPremises()(*bool) {
    return m.isSyncedFromOnPremises
}
// GetName gets the name property value. Name of the extension property. Not nullable. Supports $filter (eq).
// returns a *string when successful
func (m *ExtensionProperty) GetName()(*string) {
    return m.name
}
// GetTargetObjects gets the targetObjects property value. Following values are supported. Not nullable. UserGroupAdministrativeUnitApplicationDeviceOrganization
// returns a []string when successful
func (m *ExtensionProperty) GetTargetObjects()([]string) {
    return m.targetObjects
}
// Serialize serializes information the current object
func (m *ExtensionProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("dataType", m.GetDataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMultiValued", m.GetIsMultiValued())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSyncedFromOnPremises", m.GetIsSyncedFromOnPremises())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetTargetObjects() != nil {
        err = writer.WriteCollectionOfStringValues("targetObjects", m.GetTargetObjects())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDisplayName sets the appDisplayName property value. Display name of the application object on which this extension property is defined. Read-only.
func (m *ExtensionProperty) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetDataType sets the dataType property value. Specifies the data type of the value the extension property can hold. Following values are supported. Binary - 256 bytes maximumBooleanDateTime - Must be specified in ISO 8601 format. Will be stored in UTC.Integer - 32-bit value.LargeInteger - 64-bit value.String - 256 characters maximumNot nullable. For multivalued directory extensions, these limits apply per value in the collection.
func (m *ExtensionProperty) SetDataType(value *string)() {
    m.dataType = value
}
// SetIsMultiValued sets the isMultiValued property value. Defines the directory extension as a multi-valued property. When true, the directory extension property can store a collection of objects of the dataType; for example, a collection of string types such as 'extensionb7b1c57b532f40b8b5ed4b7a7ba67401jobGroupTracker': ['String 1', 'String 2']. The default value is false. Supports $filter (eq).
func (m *ExtensionProperty) SetIsMultiValued(value *bool)() {
    m.isMultiValued = value
}
// SetIsSyncedFromOnPremises sets the isSyncedFromOnPremises property value. Indicates if this extension property was synced from on-premises active directory using Microsoft Entra Connect. Read-only.
func (m *ExtensionProperty) SetIsSyncedFromOnPremises(value *bool)() {
    m.isSyncedFromOnPremises = value
}
// SetName sets the name property value. Name of the extension property. Not nullable. Supports $filter (eq).
func (m *ExtensionProperty) SetName(value *string)() {
    m.name = value
}
// SetTargetObjects sets the targetObjects property value. Following values are supported. Not nullable. UserGroupAdministrativeUnitApplicationDeviceOrganization
func (m *ExtensionProperty) SetTargetObjects(value []string)() {
    m.targetObjects = value
}
type ExtensionPropertyable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDisplayName()(*string)
    GetDataType()(*string)
    GetIsMultiValued()(*bool)
    GetIsSyncedFromOnPremises()(*bool)
    GetName()(*string)
    GetTargetObjects()([]string)
    SetAppDisplayName(value *string)()
    SetDataType(value *string)()
    SetIsMultiValued(value *bool)()
    SetIsSyncedFromOnPremises(value *bool)()
    SetName(value *string)()
    SetTargetObjects(value []string)()
}
