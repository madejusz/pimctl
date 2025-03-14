package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Identity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The display name of the identity. This property is read-only.
    displayName *string
    // The identifier of the identity. This property is read-only.
    id *string
    // The OdataType property
    odataType *string
}
// NewIdentity instantiates a new Identity and sets the default values.
func NewIdentity()(*Identity) {
    m := &Identity{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.auditUserIdentity":
                        return NewAuditUserIdentity(), nil
                    case "#microsoft.graph.azureCommunicationServicesUserIdentity":
                        return NewAzureCommunicationServicesUserIdentity(), nil
                    case "#microsoft.graph.communicationsApplicationIdentity":
                        return NewCommunicationsApplicationIdentity(), nil
                    case "#microsoft.graph.communicationsApplicationInstanceIdentity":
                        return NewCommunicationsApplicationInstanceIdentity(), nil
                    case "#microsoft.graph.communicationsEncryptedIdentity":
                        return NewCommunicationsEncryptedIdentity(), nil
                    case "#microsoft.graph.communicationsGuestIdentity":
                        return NewCommunicationsGuestIdentity(), nil
                    case "#microsoft.graph.communicationsPhoneIdentity":
                        return NewCommunicationsPhoneIdentity(), nil
                    case "#microsoft.graph.communicationsUserIdentity":
                        return NewCommunicationsUserIdentity(), nil
                    case "#microsoft.graph.emailIdentity":
                        return NewEmailIdentity(), nil
                    case "#microsoft.graph.initiator":
                        return NewInitiator(), nil
                    case "#microsoft.graph.programResource":
                        return NewProgramResource(), nil
                    case "#microsoft.graph.provisionedIdentity":
                        return NewProvisionedIdentity(), nil
                    case "#microsoft.graph.provisioningServicePrincipal":
                        return NewProvisioningServicePrincipal(), nil
                    case "#microsoft.graph.provisioningSystem":
                        return NewProvisioningSystem(), nil
                    case "#microsoft.graph.servicePrincipalIdentity":
                        return NewServicePrincipalIdentity(), nil
                    case "#microsoft.graph.sharePointIdentity":
                        return NewSharePointIdentity(), nil
                    case "#microsoft.graph.sourceProvisionedIdentity":
                        return NewSourceProvisionedIdentity(), nil
                    case "#microsoft.graph.targetProvisionedIdentity":
                        return NewTargetProvisionedIdentity(), nil
                    case "#microsoft.graph.teamworkApplicationIdentity":
                        return NewTeamworkApplicationIdentity(), nil
                    case "#microsoft.graph.teamworkConversationIdentity":
                        return NewTeamworkConversationIdentity(), nil
                    case "#microsoft.graph.teamworkTagIdentity":
                        return NewTeamworkTagIdentity(), nil
                    case "#microsoft.graph.teamworkUserIdentity":
                        return NewTeamworkUserIdentity(), nil
                    case "#microsoft.graph.userIdentity":
                        return NewUserIdentity(), nil
                }
            }
        }
    }
    return NewIdentity(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Identity) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The display name of the identity. This property is read-only.
// returns a *string when successful
func (m *Identity) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Identity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
// GetId gets the id property value. The identifier of the identity. This property is read-only.
// returns a *string when successful
func (m *Identity) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *Identity) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *Identity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *Identity) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The display name of the identity. This property is read-only.
func (m *Identity) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. The identifier of the identity. This property is read-only.
func (m *Identity) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Identity) SetOdataType(value *string)() {
    m.odataType = value
}
type Identityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetId()(*string)
    GetOdataType()(*string)
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetOdataType(value *string)()
}
