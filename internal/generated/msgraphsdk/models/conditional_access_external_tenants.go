package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConditionalAccessExternalTenants struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The membership kind. Possible values are: all, enumerated, unknownFutureValue. The enumerated member references an conditionalAccessEnumeratedExternalTenants object.
    membershipKind *ConditionalAccessExternalTenantsMembershipKind
    // The OdataType property
    odataType *string
}
// NewConditionalAccessExternalTenants instantiates a new ConditionalAccessExternalTenants and sets the default values.
func NewConditionalAccessExternalTenants()(*ConditionalAccessExternalTenants) {
    m := &ConditionalAccessExternalTenants{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessExternalTenantsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionalAccessExternalTenantsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.conditionalAccessAllExternalTenants":
                        return NewConditionalAccessAllExternalTenants(), nil
                    case "#microsoft.graph.conditionalAccessEnumeratedExternalTenants":
                        return NewConditionalAccessEnumeratedExternalTenants(), nil
                }
            }
        }
    }
    return NewConditionalAccessExternalTenants(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConditionalAccessExternalTenants) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConditionalAccessExternalTenants) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["membershipKind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessExternalTenantsMembershipKind)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipKind(val.(*ConditionalAccessExternalTenantsMembershipKind))
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
// GetMembershipKind gets the membershipKind property value. The membership kind. Possible values are: all, enumerated, unknownFutureValue. The enumerated member references an conditionalAccessEnumeratedExternalTenants object.
// returns a *ConditionalAccessExternalTenantsMembershipKind when successful
func (m *ConditionalAccessExternalTenants) GetMembershipKind()(*ConditionalAccessExternalTenantsMembershipKind) {
    return m.membershipKind
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ConditionalAccessExternalTenants) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ConditionalAccessExternalTenants) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMembershipKind() != nil {
        cast := (*m.GetMembershipKind()).String()
        err := writer.WriteStringValue("membershipKind", &cast)
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
func (m *ConditionalAccessExternalTenants) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMembershipKind sets the membershipKind property value. The membership kind. Possible values are: all, enumerated, unknownFutureValue. The enumerated member references an conditionalAccessEnumeratedExternalTenants object.
func (m *ConditionalAccessExternalTenants) SetMembershipKind(value *ConditionalAccessExternalTenantsMembershipKind)() {
    m.membershipKind = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessExternalTenants) SetOdataType(value *string)() {
    m.odataType = value
}
type ConditionalAccessExternalTenantsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMembershipKind()(*ConditionalAccessExternalTenantsMembershipKind)
    GetOdataType()(*string)
    SetMembershipKind(value *ConditionalAccessExternalTenantsMembershipKind)()
    SetOdataType(value *string)()
}
