package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OpenShift struct {
    ChangeTrackedEntity
    // Draft changes in the openShift are only visible to managers until they're shared.
    draftOpenShift OpenShiftItemable
    // The openShift is marked for deletion, a process that is finalized when the schedule is shared.
    isStagedForDeletion *bool
    // The ID of the schedulingGroup that contains the openShift.
    schedulingGroupId *string
    // The shared version of this openShift that is viewable by both employees and managers.
    sharedOpenShift OpenShiftItemable
}
// NewOpenShift instantiates a new OpenShift and sets the default values.
func NewOpenShift()(*OpenShift) {
    m := &OpenShift{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.openShift"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOpenShiftFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOpenShiftFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOpenShift(), nil
}
// GetDraftOpenShift gets the draftOpenShift property value. Draft changes in the openShift are only visible to managers until they're shared.
// returns a OpenShiftItemable when successful
func (m *OpenShift) GetDraftOpenShift()(OpenShiftItemable) {
    return m.draftOpenShift
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OpenShift) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftOpenShift"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOpenShiftItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraftOpenShift(val.(OpenShiftItemable))
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
    res["sharedOpenShift"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOpenShiftItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedOpenShift(val.(OpenShiftItemable))
        }
        return nil
    }
    return res
}
// GetIsStagedForDeletion gets the isStagedForDeletion property value. The openShift is marked for deletion, a process that is finalized when the schedule is shared.
// returns a *bool when successful
func (m *OpenShift) GetIsStagedForDeletion()(*bool) {
    return m.isStagedForDeletion
}
// GetSchedulingGroupId gets the schedulingGroupId property value. The ID of the schedulingGroup that contains the openShift.
// returns a *string when successful
func (m *OpenShift) GetSchedulingGroupId()(*string) {
    return m.schedulingGroupId
}
// GetSharedOpenShift gets the sharedOpenShift property value. The shared version of this openShift that is viewable by both employees and managers.
// returns a OpenShiftItemable when successful
func (m *OpenShift) GetSharedOpenShift()(OpenShiftItemable) {
    return m.sharedOpenShift
}
// Serialize serializes information the current object
func (m *OpenShift) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftOpenShift", m.GetDraftOpenShift())
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
        err = writer.WriteObjectValue("sharedOpenShift", m.GetSharedOpenShift())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDraftOpenShift sets the draftOpenShift property value. Draft changes in the openShift are only visible to managers until they're shared.
func (m *OpenShift) SetDraftOpenShift(value OpenShiftItemable)() {
    m.draftOpenShift = value
}
// SetIsStagedForDeletion sets the isStagedForDeletion property value. The openShift is marked for deletion, a process that is finalized when the schedule is shared.
func (m *OpenShift) SetIsStagedForDeletion(value *bool)() {
    m.isStagedForDeletion = value
}
// SetSchedulingGroupId sets the schedulingGroupId property value. The ID of the schedulingGroup that contains the openShift.
func (m *OpenShift) SetSchedulingGroupId(value *string)() {
    m.schedulingGroupId = value
}
// SetSharedOpenShift sets the sharedOpenShift property value. The shared version of this openShift that is viewable by both employees and managers.
func (m *OpenShift) SetSharedOpenShift(value OpenShiftItemable)() {
    m.sharedOpenShift = value
}
type OpenShiftable interface {
    ChangeTrackedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDraftOpenShift()(OpenShiftItemable)
    GetIsStagedForDeletion()(*bool)
    GetSchedulingGroupId()(*string)
    GetSharedOpenShift()(OpenShiftItemable)
    SetDraftOpenShift(value OpenShiftItemable)()
    SetIsStagedForDeletion(value *bool)()
    SetSchedulingGroupId(value *string)()
    SetSharedOpenShift(value OpenShiftItemable)()
}
