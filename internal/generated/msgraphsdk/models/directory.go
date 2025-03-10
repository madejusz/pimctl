package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Directory struct {
    Entity
    // Conceptual container for user and group directory objects.
    administrativeUnits []AdministrativeUnitable
    // Group of related custom security attribute definitions.
    attributeSets []AttributeSetable
    // Schema of a custom security attributes (key-value pairs).
    customSecurityAttributeDefinitions []CustomSecurityAttributeDefinitionable
    // Recently deleted items. Read-only. Nullable.
    deletedItems []DirectoryObjectable
    // The credentials of the device's local administrator account backed up to Microsoft Entra ID.
    deviceLocalCredentials []DeviceLocalCredentialInfoable
    // Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
    federationConfigurations []IdentityProviderBaseable
    // A container for on-premises directory synchronization functionalities that are available for the organization.
    onPremisesSynchronization []OnPremisesDirectorySynchronizationable
    // List of commercial subscriptions that an organization acquired.
    subscriptions []CompanySubscriptionable
}
// NewDirectory instantiates a new Directory and sets the default values.
func NewDirectory()(*Directory) {
    m := &Directory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectory(), nil
}
// GetAdministrativeUnits gets the administrativeUnits property value. Conceptual container for user and group directory objects.
// returns a []AdministrativeUnitable when successful
func (m *Directory) GetAdministrativeUnits()([]AdministrativeUnitable) {
    return m.administrativeUnits
}
// GetAttributeSets gets the attributeSets property value. Group of related custom security attribute definitions.
// returns a []AttributeSetable when successful
func (m *Directory) GetAttributeSets()([]AttributeSetable) {
    return m.attributeSets
}
// GetCustomSecurityAttributeDefinitions gets the customSecurityAttributeDefinitions property value. Schema of a custom security attributes (key-value pairs).
// returns a []CustomSecurityAttributeDefinitionable when successful
func (m *Directory) GetCustomSecurityAttributeDefinitions()([]CustomSecurityAttributeDefinitionable) {
    return m.customSecurityAttributeDefinitions
}
// GetDeletedItems gets the deletedItems property value. Recently deleted items. Read-only. Nullable.
// returns a []DirectoryObjectable when successful
func (m *Directory) GetDeletedItems()([]DirectoryObjectable) {
    return m.deletedItems
}
// GetDeviceLocalCredentials gets the deviceLocalCredentials property value. The credentials of the device's local administrator account backed up to Microsoft Entra ID.
// returns a []DeviceLocalCredentialInfoable when successful
func (m *Directory) GetDeviceLocalCredentials()([]DeviceLocalCredentialInfoable) {
    return m.deviceLocalCredentials
}
// GetFederationConfigurations gets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
// returns a []IdentityProviderBaseable when successful
func (m *Directory) GetFederationConfigurations()([]IdentityProviderBaseable) {
    return m.federationConfigurations
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Directory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["administrativeUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAdministrativeUnitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AdministrativeUnitable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AdministrativeUnitable)
                }
            }
            m.SetAdministrativeUnits(res)
        }
        return nil
    }
    res["attributeSets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttributeSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttributeSetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AttributeSetable)
                }
            }
            m.SetAttributeSets(res)
        }
        return nil
    }
    res["customSecurityAttributeDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomSecurityAttributeDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomSecurityAttributeDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomSecurityAttributeDefinitionable)
                }
            }
            m.SetCustomSecurityAttributeDefinitions(res)
        }
        return nil
    }
    res["deletedItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DirectoryObjectable)
                }
            }
            m.SetDeletedItems(res)
        }
        return nil
    }
    res["deviceLocalCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceLocalCredentialInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceLocalCredentialInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceLocalCredentialInfoable)
                }
            }
            m.SetDeviceLocalCredentials(res)
        }
        return nil
    }
    res["federationConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityProviderBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IdentityProviderBaseable)
                }
            }
            m.SetFederationConfigurations(res)
        }
        return nil
    }
    res["onPremisesSynchronization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesDirectorySynchronizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesDirectorySynchronizationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OnPremisesDirectorySynchronizationable)
                }
            }
            m.SetOnPremisesSynchronization(res)
        }
        return nil
    }
    res["subscriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCompanySubscriptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CompanySubscriptionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CompanySubscriptionable)
                }
            }
            m.SetSubscriptions(res)
        }
        return nil
    }
    return res
}
// GetOnPremisesSynchronization gets the onPremisesSynchronization property value. A container for on-premises directory synchronization functionalities that are available for the organization.
// returns a []OnPremisesDirectorySynchronizationable when successful
func (m *Directory) GetOnPremisesSynchronization()([]OnPremisesDirectorySynchronizationable) {
    return m.onPremisesSynchronization
}
// GetSubscriptions gets the subscriptions property value. List of commercial subscriptions that an organization acquired.
// returns a []CompanySubscriptionable when successful
func (m *Directory) GetSubscriptions()([]CompanySubscriptionable) {
    return m.subscriptions
}
// Serialize serializes information the current object
func (m *Directory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdministrativeUnits() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdministrativeUnits()))
        for i, v := range m.GetAdministrativeUnits() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("administrativeUnits", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAttributeSets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttributeSets()))
        for i, v := range m.GetAttributeSets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("attributeSets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomSecurityAttributeDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomSecurityAttributeDefinitions()))
        for i, v := range m.GetCustomSecurityAttributeDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customSecurityAttributeDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedItems()))
        for i, v := range m.GetDeletedItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deletedItems", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceLocalCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceLocalCredentials()))
        for i, v := range m.GetDeviceLocalCredentials() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceLocalCredentials", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFederationConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFederationConfigurations()))
        for i, v := range m.GetFederationConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("federationConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOnPremisesSynchronization() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnPremisesSynchronization()))
        for i, v := range m.GetOnPremisesSynchronization() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("onPremisesSynchronization", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubscriptions()))
        for i, v := range m.GetSubscriptions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("subscriptions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdministrativeUnits sets the administrativeUnits property value. Conceptual container for user and group directory objects.
func (m *Directory) SetAdministrativeUnits(value []AdministrativeUnitable)() {
    m.administrativeUnits = value
}
// SetAttributeSets sets the attributeSets property value. Group of related custom security attribute definitions.
func (m *Directory) SetAttributeSets(value []AttributeSetable)() {
    m.attributeSets = value
}
// SetCustomSecurityAttributeDefinitions sets the customSecurityAttributeDefinitions property value. Schema of a custom security attributes (key-value pairs).
func (m *Directory) SetCustomSecurityAttributeDefinitions(value []CustomSecurityAttributeDefinitionable)() {
    m.customSecurityAttributeDefinitions = value
}
// SetDeletedItems sets the deletedItems property value. Recently deleted items. Read-only. Nullable.
func (m *Directory) SetDeletedItems(value []DirectoryObjectable)() {
    m.deletedItems = value
}
// SetDeviceLocalCredentials sets the deviceLocalCredentials property value. The credentials of the device's local administrator account backed up to Microsoft Entra ID.
func (m *Directory) SetDeviceLocalCredentials(value []DeviceLocalCredentialInfoable)() {
    m.deviceLocalCredentials = value
}
// SetFederationConfigurations sets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
func (m *Directory) SetFederationConfigurations(value []IdentityProviderBaseable)() {
    m.federationConfigurations = value
}
// SetOnPremisesSynchronization sets the onPremisesSynchronization property value. A container for on-premises directory synchronization functionalities that are available for the organization.
func (m *Directory) SetOnPremisesSynchronization(value []OnPremisesDirectorySynchronizationable)() {
    m.onPremisesSynchronization = value
}
// SetSubscriptions sets the subscriptions property value. List of commercial subscriptions that an organization acquired.
func (m *Directory) SetSubscriptions(value []CompanySubscriptionable)() {
    m.subscriptions = value
}
type Directoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdministrativeUnits()([]AdministrativeUnitable)
    GetAttributeSets()([]AttributeSetable)
    GetCustomSecurityAttributeDefinitions()([]CustomSecurityAttributeDefinitionable)
    GetDeletedItems()([]DirectoryObjectable)
    GetDeviceLocalCredentials()([]DeviceLocalCredentialInfoable)
    GetFederationConfigurations()([]IdentityProviderBaseable)
    GetOnPremisesSynchronization()([]OnPremisesDirectorySynchronizationable)
    GetSubscriptions()([]CompanySubscriptionable)
    SetAdministrativeUnits(value []AdministrativeUnitable)()
    SetAttributeSets(value []AttributeSetable)()
    SetCustomSecurityAttributeDefinitions(value []CustomSecurityAttributeDefinitionable)()
    SetDeletedItems(value []DirectoryObjectable)()
    SetDeviceLocalCredentials(value []DeviceLocalCredentialInfoable)()
    SetFederationConfigurations(value []IdentityProviderBaseable)()
    SetOnPremisesSynchronization(value []OnPremisesDirectorySynchronizationable)()
    SetSubscriptions(value []CompanySubscriptionable)()
}
