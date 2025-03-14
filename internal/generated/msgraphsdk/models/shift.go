package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Shift struct {
    ChangeTrackedEntity
    // Draft changes in the shift. Draft changes are only visible to managers. The changes are visible to employees when they are shared, which copies the changes from the draftShift to the sharedShift property.
    draftShift ShiftItemable
    // The shift is marked for deletion, a process that is finalized when the schedule is shared.
    isStagedForDeletion *bool
    // ID of the scheduling group the shift is part of. Required.
    schedulingGroupId *string
    // The shared version of this shift that is viewable by both employees and managers. Updates to the sharedShift property send notifications to users in the Teams client.
    sharedShift ShiftItemable
    // ID of the user assigned to the shift. Required.
    userId *string
}
// NewShift instantiates a new Shift and sets the default values.
func NewShift()(*Shift) {
    m := &Shift{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.shift"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateShiftFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateShiftFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewShift(), nil
}
// GetDraftShift gets the draftShift property value. Draft changes in the shift. Draft changes are only visible to managers. The changes are visible to employees when they are shared, which copies the changes from the draftShift to the sharedShift property.
// returns a ShiftItemable when successful
func (m *Shift) GetDraftShift()(ShiftItemable) {
    return m.draftShift
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Shift) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftShift"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateShiftItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraftShift(val.(ShiftItemable))
        }
        return nil
    }
    res["isStagedForDeletion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsStagedForDeletion(val)
        }
        return nil
    }
    res["schedulingGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedulingGroupId(val)
        }
        return nil
    }
    res["sharedShift"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateShiftItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedShift(val.(ShiftItemable))
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetIsStagedForDeletion gets the isStagedForDeletion property value. The shift is marked for deletion, a process that is finalized when the schedule is shared.
// returns a *bool when successful
func (m *Shift) GetIsStagedForDeletion()(*bool) {
    return m.isStagedForDeletion
}
// GetSchedulingGroupId gets the schedulingGroupId property value. ID of the scheduling group the shift is part of. Required.
// returns a *string when successful
func (m *Shift) GetSchedulingGroupId()(*string) {
    return m.schedulingGroupId
}
// GetSharedShift gets the sharedShift property value. The shared version of this shift that is viewable by both employees and managers. Updates to the sharedShift property send notifications to users in the Teams client.
// returns a ShiftItemable when successful
func (m *Shift) GetSharedShift()(ShiftItemable) {
    return m.sharedShift
}
// GetUserId gets the userId property value. ID of the user assigned to the shift. Required.
// returns a *string when successful
func (m *Shift) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *Shift) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftShift", m.GetDraftShift())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isStagedForDeletion", m.GetIsStagedForDeletion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("schedulingGroupId", m.GetSchedulingGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedShift", m.GetSharedShift())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDraftShift sets the draftShift property value. Draft changes in the shift. Draft changes are only visible to managers. The changes are visible to employees when they are shared, which copies the changes from the draftShift to the sharedShift property.
func (m *Shift) SetDraftShift(value ShiftItemable)() {
    m.draftShift = value
}
// SetIsStagedForDeletion sets the isStagedForDeletion property value. The shift is marked for deletion, a process that is finalized when the schedule is shared.
func (m *Shift) SetIsStagedForDeletion(value *bool)() {
    m.isStagedForDeletion = value
}
// SetSchedulingGroupId sets the schedulingGroupId property value. ID of the scheduling group the shift is part of. Required.
func (m *Shift) SetSchedulingGroupId(value *string)() {
    m.schedulingGroupId = value
}
// SetSharedShift sets the sharedShift property value. The shared version of this shift that is viewable by both employees and managers. Updates to the sharedShift property send notifications to users in the Teams client.
func (m *Shift) SetSharedShift(value ShiftItemable)() {
    m.sharedShift = value
}
// SetUserId sets the userId property value. ID of the user assigned to the shift. Required.
func (m *Shift) SetUserId(value *string)() {
    m.userId = value
}
type Shiftable interface {
    ChangeTrackedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDraftShift()(ShiftItemable)
    GetIsStagedForDeletion()(*bool)
    GetSchedulingGroupId()(*string)
    GetSharedShift()(ShiftItemable)
    GetUserId()(*string)
    SetDraftShift(value ShiftItemable)()
    SetIsStagedForDeletion(value *bool)()
    SetSchedulingGroupId(value *string)()
    SetSharedShift(value ShiftItemable)()
    SetUserId(value *string)()
}
