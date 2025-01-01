package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomTrainingSetting struct {
    TrainingSetting
    // A user collection that specifies to whom the training should be assigned. Possible values are: none, allUsers, clickedPayload, compromised, reportedPhish, readButNotClicked, didNothing, unknownFutureValue.
    assignedTo []TrainingAssignedTo
    // The description of the custom training setting.
    description *string
    // The display name of the custom training setting.
    displayName *string
    // Training duration.
    durationInMinutes *int32
    // The training URL.
    url *string
}
// NewCustomTrainingSetting instantiates a new CustomTrainingSetting and sets the default values.
func NewCustomTrainingSetting()(*CustomTrainingSetting) {
    m := &CustomTrainingSetting{
        TrainingSetting: *NewTrainingSetting(),
    }
    odataTypeValue := "#microsoft.graph.customTrainingSetting"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomTrainingSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomTrainingSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomTrainingSetting(), nil
}
// GetAssignedTo gets the assignedTo property value. A user collection that specifies to whom the training should be assigned. Possible values are: none, allUsers, clickedPayload, compromised, reportedPhish, readButNotClicked, didNothing, unknownFutureValue.
// returns a []TrainingAssignedTo when successful
func (m *CustomTrainingSetting) GetAssignedTo()([]TrainingAssignedTo) {
    return m.assignedTo
}
// GetDescription gets the description property value. The description of the custom training setting.
// returns a *string when successful
func (m *CustomTrainingSetting) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the custom training setting.
// returns a *string when successful
func (m *CustomTrainingSetting) GetDisplayName()(*string) {
    return m.displayName
}
// GetDurationInMinutes gets the durationInMinutes property value. Training duration.
// returns a *int32 when successful
func (m *CustomTrainingSetting) GetDurationInMinutes()(*int32) {
    return m.durationInMinutes
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomTrainingSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TrainingSetting.GetFieldDeserializers()
    res["assignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseTrainingAssignedTo)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrainingAssignedTo, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*TrainingAssignedTo))
                }
            }
            m.SetAssignedTo(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
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
    res["durationInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInMinutes(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetUrl gets the url property value. The training URL.
// returns a *string when successful
func (m *CustomTrainingSetting) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CustomTrainingSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TrainingSetting.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignedTo() != nil {
        err = writer.WriteCollectionOfStringValues("assignedTo", SerializeTrainingAssignedTo(m.GetAssignedTo()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("durationInMinutes", m.GetDurationInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTo sets the assignedTo property value. A user collection that specifies to whom the training should be assigned. Possible values are: none, allUsers, clickedPayload, compromised, reportedPhish, readButNotClicked, didNothing, unknownFutureValue.
func (m *CustomTrainingSetting) SetAssignedTo(value []TrainingAssignedTo)() {
    m.assignedTo = value
}
// SetDescription sets the description property value. The description of the custom training setting.
func (m *CustomTrainingSetting) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the custom training setting.
func (m *CustomTrainingSetting) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDurationInMinutes sets the durationInMinutes property value. Training duration.
func (m *CustomTrainingSetting) SetDurationInMinutes(value *int32)() {
    m.durationInMinutes = value
}
// SetUrl sets the url property value. The training URL.
func (m *CustomTrainingSetting) SetUrl(value *string)() {
    m.url = value
}
type CustomTrainingSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TrainingSettingable
    GetAssignedTo()([]TrainingAssignedTo)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDurationInMinutes()(*int32)
    GetUrl()(*string)
    SetAssignedTo(value []TrainingAssignedTo)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDurationInMinutes(value *int32)()
    SetUrl(value *string)()
}
