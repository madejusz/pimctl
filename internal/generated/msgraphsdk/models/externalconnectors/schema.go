package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type Schema struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // Must be set to microsoft.graph.externalConnector.externalItem. Required.
    baseType *string
    // The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128.
    properties []Propertyable
}
// NewSchema instantiates a new Schema and sets the default values.
func NewSchema()(*Schema) {
    m := &Schema{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSchema(), nil
}
// GetBaseType gets the baseType property value. Must be set to microsoft.graph.externalConnector.externalItem. Required.
// returns a *string when successful
func (m *Schema) GetBaseType()(*string) {
    return m.baseType
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Schema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["baseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseType(val)
        }
        return nil
    }
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Propertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Propertyable)
                }
            }
            m.SetProperties(res)
        }
        return nil
    }
    return res
}
// GetProperties gets the properties property value. The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128.
// returns a []Propertyable when successful
func (m *Schema) GetProperties()([]Propertyable) {
    return m.properties
}
// Serialize serializes information the current object
func (m *Schema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("baseType", m.GetBaseType())
        if err != nil {
            return err
        }
    }
    if m.GetProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBaseType sets the baseType property value. Must be set to microsoft.graph.externalConnector.externalItem. Required.
func (m *Schema) SetBaseType(value *string)() {
    m.baseType = value
}
// SetProperties sets the properties property value. The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128.
func (m *Schema) SetProperties(value []Propertyable)() {
    m.properties = value
}
type Schemaable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseType()(*string)
    GetProperties()([]Propertyable)
    SetBaseType(value *string)()
    SetProperties(value []Propertyable)()
}
