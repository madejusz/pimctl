package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedPCAccountManagerPolicy sharedPC Account Manager Policy. Only applies when the account manager is enabled.
type SharedPCAccountManagerPolicy struct {
    // Possible values for when accounts are deleted on a shared PC.
    accountDeletionPolicy *SharedPCAccountDeletionPolicyType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Sets the percentage of available disk space a PC should have before it stops deleting cached shared PC accounts. Only applies when AccountDeletionPolicy is DiskSpaceThreshold or DiskSpaceThresholdOrInactiveThreshold. Valid values 0 to 100
    cacheAccountsAboveDiskFreePercentage *int32
    // Specifies when the accounts will start being deleted when they have not been logged on during the specified period, given as number of days. Only applies when AccountDeletionPolicy is DiskSpaceThreshold or DiskSpaceThresholdOrInactiveThreshold.
    inactiveThresholdDays *int32
    // The OdataType property
    odataType *string
    // Sets the percentage of disk space remaining on a PC before cached accounts will be deleted to free disk space. Accounts that have been inactive the longest will be deleted first. Only applies when AccountDeletionPolicy is DiskSpaceThresholdOrInactiveThreshold. Valid values 0 to 100
    removeAccountsBelowDiskFreePercentage *int32
}
// NewSharedPCAccountManagerPolicy instantiates a new SharedPCAccountManagerPolicy and sets the default values.
func NewSharedPCAccountManagerPolicy()(*SharedPCAccountManagerPolicy) {
    m := &SharedPCAccountManagerPolicy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSharedPCAccountManagerPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSharedPCAccountManagerPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharedPCAccountManagerPolicy(), nil
}
// GetAccountDeletionPolicy gets the accountDeletionPolicy property value. Possible values for when accounts are deleted on a shared PC.
// returns a *SharedPCAccountDeletionPolicyType when successful
func (m *SharedPCAccountManagerPolicy) GetAccountDeletionPolicy()(*SharedPCAccountDeletionPolicyType) {
    return m.accountDeletionPolicy
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SharedPCAccountManagerPolicy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCacheAccountsAboveDiskFreePercentage gets the cacheAccountsAboveDiskFreePercentage property value. Sets the percentage of available disk space a PC should have before it stops deleting cached shared PC accounts. Only applies when AccountDeletionPolicy is DiskSpaceThreshold or DiskSpaceThresholdOrInactiveThreshold. Valid values 0 to 100
// returns a *int32 when successful
func (m *SharedPCAccountManagerPolicy) GetCacheAccountsAboveDiskFreePercentage()(*int32) {
    return m.cacheAccountsAboveDiskFreePercentage
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SharedPCAccountManagerPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountDeletionPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSharedPCAccountDeletionPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountDeletionPolicy(val.(*SharedPCAccountDeletionPolicyType))
        }
        return nil
    }
    res["cacheAccountsAboveDiskFreePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheAccountsAboveDiskFreePercentage(val)
        }
        return nil
    }
    res["inactiveThresholdDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInactiveThresholdDays(val)
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
    res["removeAccountsBelowDiskFreePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveAccountsBelowDiskFreePercentage(val)
        }
        return nil
    }
    return res
}
// GetInactiveThresholdDays gets the inactiveThresholdDays property value. Specifies when the accounts will start being deleted when they have not been logged on during the specified period, given as number of days. Only applies when AccountDeletionPolicy is DiskSpaceThreshold or DiskSpaceThresholdOrInactiveThreshold.
// returns a *int32 when successful
func (m *SharedPCAccountManagerPolicy) GetInactiveThresholdDays()(*int32) {
    return m.inactiveThresholdDays
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *SharedPCAccountManagerPolicy) GetOdataType()(*string) {
    return m.odataType
}
// GetRemoveAccountsBelowDiskFreePercentage gets the removeAccountsBelowDiskFreePercentage property value. Sets the percentage of disk space remaining on a PC before cached accounts will be deleted to free disk space. Accounts that have been inactive the longest will be deleted first. Only applies when AccountDeletionPolicy is DiskSpaceThresholdOrInactiveThreshold. Valid values 0 to 100
// returns a *int32 when successful
func (m *SharedPCAccountManagerPolicy) GetRemoveAccountsBelowDiskFreePercentage()(*int32) {
    return m.removeAccountsBelowDiskFreePercentage
}
// Serialize serializes information the current object
func (m *SharedPCAccountManagerPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccountDeletionPolicy() != nil {
        cast := (*m.GetAccountDeletionPolicy()).String()
        err := writer.WriteStringValue("accountDeletionPolicy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("cacheAccountsAboveDiskFreePercentage", m.GetCacheAccountsAboveDiskFreePercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inactiveThresholdDays", m.GetInactiveThresholdDays())
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
        err := writer.WriteInt32Value("removeAccountsBelowDiskFreePercentage", m.GetRemoveAccountsBelowDiskFreePercentage())
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
// SetAccountDeletionPolicy sets the accountDeletionPolicy property value. Possible values for when accounts are deleted on a shared PC.
func (m *SharedPCAccountManagerPolicy) SetAccountDeletionPolicy(value *SharedPCAccountDeletionPolicyType)() {
    m.accountDeletionPolicy = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharedPCAccountManagerPolicy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCacheAccountsAboveDiskFreePercentage sets the cacheAccountsAboveDiskFreePercentage property value. Sets the percentage of available disk space a PC should have before it stops deleting cached shared PC accounts. Only applies when AccountDeletionPolicy is DiskSpaceThreshold or DiskSpaceThresholdOrInactiveThreshold. Valid values 0 to 100
func (m *SharedPCAccountManagerPolicy) SetCacheAccountsAboveDiskFreePercentage(value *int32)() {
    m.cacheAccountsAboveDiskFreePercentage = value
}
// SetInactiveThresholdDays sets the inactiveThresholdDays property value. Specifies when the accounts will start being deleted when they have not been logged on during the specified period, given as number of days. Only applies when AccountDeletionPolicy is DiskSpaceThreshold or DiskSpaceThresholdOrInactiveThreshold.
func (m *SharedPCAccountManagerPolicy) SetInactiveThresholdDays(value *int32)() {
    m.inactiveThresholdDays = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SharedPCAccountManagerPolicy) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRemoveAccountsBelowDiskFreePercentage sets the removeAccountsBelowDiskFreePercentage property value. Sets the percentage of disk space remaining on a PC before cached accounts will be deleted to free disk space. Accounts that have been inactive the longest will be deleted first. Only applies when AccountDeletionPolicy is DiskSpaceThresholdOrInactiveThreshold. Valid values 0 to 100
func (m *SharedPCAccountManagerPolicy) SetRemoveAccountsBelowDiskFreePercentage(value *int32)() {
    m.removeAccountsBelowDiskFreePercentage = value
}
type SharedPCAccountManagerPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountDeletionPolicy()(*SharedPCAccountDeletionPolicyType)
    GetCacheAccountsAboveDiskFreePercentage()(*int32)
    GetInactiveThresholdDays()(*int32)
    GetOdataType()(*string)
    GetRemoveAccountsBelowDiskFreePercentage()(*int32)
    SetAccountDeletionPolicy(value *SharedPCAccountDeletionPolicyType)()
    SetCacheAccountsAboveDiskFreePercentage(value *int32)()
    SetInactiveThresholdDays(value *int32)()
    SetOdataType(value *string)()
    SetRemoveAccountsBelowDiskFreePercentage(value *int32)()
}
