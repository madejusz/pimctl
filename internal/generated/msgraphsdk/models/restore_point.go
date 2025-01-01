package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RestorePoint struct {
    Entity
    // Expiration date time of the restore point.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date time when the restore point was created.
    protectionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The site, drive, or mailbox units that are protected under a protection policy.
    protectionUnit ProtectionUnitBaseable
    // The type of the restore point. The possible values are: none, fastRestore, unknownFutureValue.
    tags *RestorePointTags
}
// NewRestorePoint instantiates a new RestorePoint and sets the default values.
func NewRestorePoint()(*RestorePoint) {
    m := &RestorePoint{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRestorePointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRestorePointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRestorePoint(), nil
}
// GetExpirationDateTime gets the expirationDateTime property value. Expiration date time of the restore point.
// returns a *Time when successful
func (m *RestorePoint) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RestorePoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["protectionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionDateTime(val)
        }
        return nil
    }
    res["protectionUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectionUnitBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionUnit(val.(ProtectionUnitBaseable))
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRestorePointTags)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTags(val.(*RestorePointTags))
        }
        return nil
    }
    return res
}
// GetProtectionDateTime gets the protectionDateTime property value. Date time when the restore point was created.
// returns a *Time when successful
func (m *RestorePoint) GetProtectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.protectionDateTime
}
// GetProtectionUnit gets the protectionUnit property value. The site, drive, or mailbox units that are protected under a protection policy.
// returns a ProtectionUnitBaseable when successful
func (m *RestorePoint) GetProtectionUnit()(ProtectionUnitBaseable) {
    return m.protectionUnit
}
// GetTags gets the tags property value. The type of the restore point. The possible values are: none, fastRestore, unknownFutureValue.
// returns a *RestorePointTags when successful
func (m *RestorePoint) GetTags()(*RestorePointTags) {
    return m.tags
}
// Serialize serializes information the current object
func (m *RestorePoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("protectionDateTime", m.GetProtectionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("protectionUnit", m.GetProtectionUnit())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        cast := (*m.GetTags()).String()
        err = writer.WriteStringValue("tags", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpirationDateTime sets the expirationDateTime property value. Expiration date time of the restore point.
func (m *RestorePoint) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetProtectionDateTime sets the protectionDateTime property value. Date time when the restore point was created.
func (m *RestorePoint) SetProtectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.protectionDateTime = value
}
// SetProtectionUnit sets the protectionUnit property value. The site, drive, or mailbox units that are protected under a protection policy.
func (m *RestorePoint) SetProtectionUnit(value ProtectionUnitBaseable)() {
    m.protectionUnit = value
}
// SetTags sets the tags property value. The type of the restore point. The possible values are: none, fastRestore, unknownFutureValue.
func (m *RestorePoint) SetTags(value *RestorePointTags)() {
    m.tags = value
}
type RestorePointable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProtectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProtectionUnit()(ProtectionUnitBaseable)
    GetTags()(*RestorePointTags)
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProtectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProtectionUnit(value ProtectionUnitBaseable)()
    SetTags(value *RestorePointTags)()
}
