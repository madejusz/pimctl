package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MultiTenantOrganizationPartnerConfigurationTemplate struct {
    Entity
    // Determines the partner-specific configuration for automatic user consent settings. Unless configured, the inboundAllowed and outboundAllowed properties are null and inherit from the default settings, which is always false.
    automaticUserConsentSettings InboundOutboundPolicyConfigurationable
    // Defines your partner-specific configuration for users from other organizations accessing your resources via Microsoft Entra B2B collaboration.
    b2bCollaborationInbound CrossTenantAccessPolicyB2BSettingable
    // Defines your partner-specific configuration for users in your organization going outbound to access resources in another organization via Microsoft Entra B2B collaboration.
    b2bCollaborationOutbound CrossTenantAccessPolicyB2BSettingable
    // Defines your partner-specific configuration for users from other organizations accessing your resources via Azure B2B direct connect.
    b2bDirectConnectInbound CrossTenantAccessPolicyB2BSettingable
    // Defines your partner-specific configuration for users in your organization going outbound to access resources in another organization via Microsoft Entra B2B direct connect.
    b2bDirectConnectOutbound CrossTenantAccessPolicyB2BSettingable
    // Determines the partner-specific configuration for trusting other Conditional Access claims from external Microsoft Entra organizations.
    inboundTrust CrossTenantAccessPolicyInboundTrustable
    // The templateApplicationLevel property
    templateApplicationLevel *TemplateApplicationLevel
}
// NewMultiTenantOrganizationPartnerConfigurationTemplate instantiates a new MultiTenantOrganizationPartnerConfigurationTemplate and sets the default values.
func NewMultiTenantOrganizationPartnerConfigurationTemplate()(*MultiTenantOrganizationPartnerConfigurationTemplate) {
    m := &MultiTenantOrganizationPartnerConfigurationTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMultiTenantOrganizationPartnerConfigurationTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMultiTenantOrganizationPartnerConfigurationTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMultiTenantOrganizationPartnerConfigurationTemplate(), nil
}
// GetAutomaticUserConsentSettings gets the automaticUserConsentSettings property value. Determines the partner-specific configuration for automatic user consent settings. Unless configured, the inboundAllowed and outboundAllowed properties are null and inherit from the default settings, which is always false.
// returns a InboundOutboundPolicyConfigurationable when successful
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) GetAutomaticUserConsentSettings()(InboundOutboundPolicyConfigurationable) {
    return m.automaticUserConsentSettings
}
// GetB2bCollaborationInbound gets the b2bCollaborationInbound property value. Defines your partner-specific configuration for users from other organizations accessing your resources via Microsoft Entra B2B collaboration.
// returns a CrossTenantAccessPolicyB2BSettingable when successful
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) GetB2bCollaborationInbound()(CrossTenantAccessPolicyB2BSettingable) {
    return m.b2bCollaborationInbound
}
// GetB2bCollaborationOutbound gets the b2bCollaborationOutbound property value. Defines your partner-specific configuration for users in your organization going outbound to access resources in another organization via Microsoft Entra B2B collaboration.
// returns a CrossTenantAccessPolicyB2BSettingable when successful
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) GetB2bCollaborationOutbound()(CrossTenantAccessPolicyB2BSettingable) {
    return m.b2bCollaborationOutbound
}
// GetB2bDirectConnectInbound gets the b2bDirectConnectInbound property value. Defines your partner-specific configuration for users from other organizations accessing your resources via Azure B2B direct connect.
// returns a CrossTenantAccessPolicyB2BSettingable when successful
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) GetB2bDirectConnectInbound()(CrossTenantAccessPolicyB2BSettingable) {
    return m.b2bDirectConnectInbound
}
// GetB2bDirectConnectOutbound gets the b2bDirectConnectOutbound property value. Defines your partner-specific configuration for users in your organization going outbound to access resources in another organization via Microsoft Entra B2B direct connect.
// returns a CrossTenantAccessPolicyB2BSettingable when successful
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) GetB2bDirectConnectOutbound()(CrossTenantAccessPolicyB2BSettingable) {
    return m.b2bDirectConnectOutbound
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["automaticUserConsentSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInboundOutboundPolicyConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticUserConsentSettings(val.(InboundOutboundPolicyConfigurationable))
        }
        return nil
    }
    res["b2bCollaborationInbound"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bCollaborationInbound(val.(CrossTenantAccessPolicyB2BSettingable))
        }
        return nil
    }
    res["b2bCollaborationOutbound"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bCollaborationOutbound(val.(CrossTenantAccessPolicyB2BSettingable))
        }
        return nil
    }
    res["b2bDirectConnectInbound"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bDirectConnectInbound(val.(CrossTenantAccessPolicyB2BSettingable))
        }
        return nil
    }
    res["b2bDirectConnectOutbound"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bDirectConnectOutbound(val.(CrossTenantAccessPolicyB2BSettingable))
        }
        return nil
    }
    res["inboundTrust"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyInboundTrustFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundTrust(val.(CrossTenantAccessPolicyInboundTrustable))
        }
        return nil
    }
    res["templateApplicationLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTemplateApplicationLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateApplicationLevel(val.(*TemplateApplicationLevel))
        }
        return nil
    }
    return res
}
// GetInboundTrust gets the inboundTrust property value. Determines the partner-specific configuration for trusting other Conditional Access claims from external Microsoft Entra organizations.
// returns a CrossTenantAccessPolicyInboundTrustable when successful
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) GetInboundTrust()(CrossTenantAccessPolicyInboundTrustable) {
    return m.inboundTrust
}
// GetTemplateApplicationLevel gets the templateApplicationLevel property value. The templateApplicationLevel property
// returns a *TemplateApplicationLevel when successful
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) GetTemplateApplicationLevel()(*TemplateApplicationLevel) {
    return m.templateApplicationLevel
}
// Serialize serializes information the current object
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("automaticUserConsentSettings", m.GetAutomaticUserConsentSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("b2bCollaborationInbound", m.GetB2bCollaborationInbound())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("b2bCollaborationOutbound", m.GetB2bCollaborationOutbound())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("b2bDirectConnectInbound", m.GetB2bDirectConnectInbound())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("b2bDirectConnectOutbound", m.GetB2bDirectConnectOutbound())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inboundTrust", m.GetInboundTrust())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateApplicationLevel() != nil {
        cast := (*m.GetTemplateApplicationLevel()).String()
        err = writer.WriteStringValue("templateApplicationLevel", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAutomaticUserConsentSettings sets the automaticUserConsentSettings property value. Determines the partner-specific configuration for automatic user consent settings. Unless configured, the inboundAllowed and outboundAllowed properties are null and inherit from the default settings, which is always false.
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) SetAutomaticUserConsentSettings(value InboundOutboundPolicyConfigurationable)() {
    m.automaticUserConsentSettings = value
}
// SetB2bCollaborationInbound sets the b2bCollaborationInbound property value. Defines your partner-specific configuration for users from other organizations accessing your resources via Microsoft Entra B2B collaboration.
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) SetB2bCollaborationInbound(value CrossTenantAccessPolicyB2BSettingable)() {
    m.b2bCollaborationInbound = value
}
// SetB2bCollaborationOutbound sets the b2bCollaborationOutbound property value. Defines your partner-specific configuration for users in your organization going outbound to access resources in another organization via Microsoft Entra B2B collaboration.
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) SetB2bCollaborationOutbound(value CrossTenantAccessPolicyB2BSettingable)() {
    m.b2bCollaborationOutbound = value
}
// SetB2bDirectConnectInbound sets the b2bDirectConnectInbound property value. Defines your partner-specific configuration for users from other organizations accessing your resources via Azure B2B direct connect.
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) SetB2bDirectConnectInbound(value CrossTenantAccessPolicyB2BSettingable)() {
    m.b2bDirectConnectInbound = value
}
// SetB2bDirectConnectOutbound sets the b2bDirectConnectOutbound property value. Defines your partner-specific configuration for users in your organization going outbound to access resources in another organization via Microsoft Entra B2B direct connect.
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) SetB2bDirectConnectOutbound(value CrossTenantAccessPolicyB2BSettingable)() {
    m.b2bDirectConnectOutbound = value
}
// SetInboundTrust sets the inboundTrust property value. Determines the partner-specific configuration for trusting other Conditional Access claims from external Microsoft Entra organizations.
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) SetInboundTrust(value CrossTenantAccessPolicyInboundTrustable)() {
    m.inboundTrust = value
}
// SetTemplateApplicationLevel sets the templateApplicationLevel property value. The templateApplicationLevel property
func (m *MultiTenantOrganizationPartnerConfigurationTemplate) SetTemplateApplicationLevel(value *TemplateApplicationLevel)() {
    m.templateApplicationLevel = value
}
type MultiTenantOrganizationPartnerConfigurationTemplateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutomaticUserConsentSettings()(InboundOutboundPolicyConfigurationable)
    GetB2bCollaborationInbound()(CrossTenantAccessPolicyB2BSettingable)
    GetB2bCollaborationOutbound()(CrossTenantAccessPolicyB2BSettingable)
    GetB2bDirectConnectInbound()(CrossTenantAccessPolicyB2BSettingable)
    GetB2bDirectConnectOutbound()(CrossTenantAccessPolicyB2BSettingable)
    GetInboundTrust()(CrossTenantAccessPolicyInboundTrustable)
    GetTemplateApplicationLevel()(*TemplateApplicationLevel)
    SetAutomaticUserConsentSettings(value InboundOutboundPolicyConfigurationable)()
    SetB2bCollaborationInbound(value CrossTenantAccessPolicyB2BSettingable)()
    SetB2bCollaborationOutbound(value CrossTenantAccessPolicyB2BSettingable)()
    SetB2bDirectConnectInbound(value CrossTenantAccessPolicyB2BSettingable)()
    SetB2bDirectConnectOutbound(value CrossTenantAccessPolicyB2BSettingable)()
    SetInboundTrust(value CrossTenantAccessPolicyInboundTrustable)()
    SetTemplateApplicationLevel(value *TemplateApplicationLevel)()
}
