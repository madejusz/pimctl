package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PasswordCredentialConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // String value that indicates the maximum lifetime for password expiration, defined as an ISO 8601 duration. For example, P4DT12H30M5S represents four days, 12 hours, 30 minutes, and five seconds. This property is required when restrictionType is set to passwordLifetime.
    maxLifetime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The OdataType property
    odataType *string
    // Specifies the date from which the policy restriction applies to newly created applications. For existing applications, the enforcement date can be retroactively applied.
    restrictForAppsCreatedAfterDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of restriction being applied. The possible values are: passwordAddition, passwordLifetime, symmetricKeyAddition, symmetricKeyLifetime, customPasswordAddition, and unknownFutureValue. Each value of restrictionType can be used only once per policy.
    restrictionType *AppCredentialRestrictionType
    // The state property
    state *AppManagementRestrictionState
}
// NewPasswordCredentialConfiguration instantiates a new PasswordCredentialConfiguration and sets the default values.
func NewPasswordCredentialConfiguration()(*PasswordCredentialConfiguration) {
    m := &PasswordCredentialConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePasswordCredentialConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePasswordCredentialConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPasswordCredentialConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PasswordCredentialConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PasswordCredentialConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maxLifetime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxLifetime(val)
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
    res["restrictForAppsCreatedAfterDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictForAppsCreatedAfterDateTime(val)
        }
        return nil
    }
    res["restrictionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppCredentialRestrictionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictionType(val.(*AppCredentialRestrictionType))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppManagementRestrictionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AppManagementRestrictionState))
        }
        return nil
    }
    return res
}
// GetMaxLifetime gets the maxLifetime property value. String value that indicates the maximum lifetime for password expiration, defined as an ISO 8601 duration. For example, P4DT12H30M5S represents four days, 12 hours, 30 minutes, and five seconds. This property is required when restrictionType is set to passwordLifetime.
// returns a *ISODuration when successful
func (m *PasswordCredentialConfiguration) GetMaxLifetime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.maxLifetime
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *PasswordCredentialConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetRestrictForAppsCreatedAfterDateTime gets the restrictForAppsCreatedAfterDateTime property value. Specifies the date from which the policy restriction applies to newly created applications. For existing applications, the enforcement date can be retroactively applied.
// returns a *Time when successful
func (m *PasswordCredentialConfiguration) GetRestrictForAppsCreatedAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.restrictForAppsCreatedAfterDateTime
}
// GetRestrictionType gets the restrictionType property value. The type of restriction being applied. The possible values are: passwordAddition, passwordLifetime, symmetricKeyAddition, symmetricKeyLifetime, customPasswordAddition, and unknownFutureValue. Each value of restrictionType can be used only once per policy.
// returns a *AppCredentialRestrictionType when successful
func (m *PasswordCredentialConfiguration) GetRestrictionType()(*AppCredentialRestrictionType) {
    return m.restrictionType
}
// GetState gets the state property value. The state property
// returns a *AppManagementRestrictionState when successful
func (m *PasswordCredentialConfiguration) GetState()(*AppManagementRestrictionState) {
    return m.state
}
// Serialize serializes information the current object
func (m *PasswordCredentialConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteISODurationValue("maxLifetime", m.GetMaxLifetime())
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
        err := writer.WriteTimeValue("restrictForAppsCreatedAfterDateTime", m.GetRestrictForAppsCreatedAfterDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRestrictionType() != nil {
        cast := (*m.GetRestrictionType()).String()
        err := writer.WriteStringValue("restrictionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *PasswordCredentialConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMaxLifetime sets the maxLifetime property value. String value that indicates the maximum lifetime for password expiration, defined as an ISO 8601 duration. For example, P4DT12H30M5S represents four days, 12 hours, 30 minutes, and five seconds. This property is required when restrictionType is set to passwordLifetime.
func (m *PasswordCredentialConfiguration) SetMaxLifetime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.maxLifetime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PasswordCredentialConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRestrictForAppsCreatedAfterDateTime sets the restrictForAppsCreatedAfterDateTime property value. Specifies the date from which the policy restriction applies to newly created applications. For existing applications, the enforcement date can be retroactively applied.
func (m *PasswordCredentialConfiguration) SetRestrictForAppsCreatedAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.restrictForAppsCreatedAfterDateTime = value
}
// SetRestrictionType sets the restrictionType property value. The type of restriction being applied. The possible values are: passwordAddition, passwordLifetime, symmetricKeyAddition, symmetricKeyLifetime, customPasswordAddition, and unknownFutureValue. Each value of restrictionType can be used only once per policy.
func (m *PasswordCredentialConfiguration) SetRestrictionType(value *AppCredentialRestrictionType)() {
    m.restrictionType = value
}
// SetState sets the state property value. The state property
func (m *PasswordCredentialConfiguration) SetState(value *AppManagementRestrictionState)() {
    m.state = value
}
type PasswordCredentialConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxLifetime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetOdataType()(*string)
    GetRestrictForAppsCreatedAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRestrictionType()(*AppCredentialRestrictionType)
    GetState()(*AppManagementRestrictionState)
    SetMaxLifetime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetOdataType(value *string)()
    SetRestrictForAppsCreatedAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRestrictionType(value *AppCredentialRestrictionType)()
    SetState(value *AppManagementRestrictionState)()
}
