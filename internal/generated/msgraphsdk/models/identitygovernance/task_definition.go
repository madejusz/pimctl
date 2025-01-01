package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type TaskDefinition struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The category property
    category *LifecycleTaskCategory
    // Defines if the workflow will continue if the task has an error.
    continueOnError *bool
    // The description of the taskDefinition.
    description *string
    // The display name of the taskDefinition.Supports $filter(eq, ne) and $orderby.
    displayName *string
    // The parameters that must be supplied when creating a workflow task object.Supports $filter(any).
    parameters []Parameterable
    // The version number of the taskDefinition. New records are pushed when we add support for new parameters.Supports $filter(ge, gt, le, lt, eq, ne) and $orderby.
    version *int32
}
// NewTaskDefinition instantiates a new TaskDefinition and sets the default values.
func NewTaskDefinition()(*TaskDefinition) {
    m := &TaskDefinition{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateTaskDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTaskDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTaskDefinition(), nil
}
// GetCategory gets the category property value. The category property
// returns a *LifecycleTaskCategory when successful
func (m *TaskDefinition) GetCategory()(*LifecycleTaskCategory) {
    return m.category
}
// GetContinueOnError gets the continueOnError property value. Defines if the workflow will continue if the task has an error.
// returns a *bool when successful
func (m *TaskDefinition) GetContinueOnError()(*bool) {
    return m.continueOnError
}
// GetDescription gets the description property value. The description of the taskDefinition.
// returns a *string when successful
func (m *TaskDefinition) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the taskDefinition.Supports $filter(eq, ne) and $orderby.
// returns a *string when successful
func (m *TaskDefinition) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TaskDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLifecycleTaskCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*LifecycleTaskCategory))
        }
        return nil
    }
    res["continueOnError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContinueOnError(val)
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
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Parameterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Parameterable)
                }
            }
            m.SetParameters(res)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters that must be supplied when creating a workflow task object.Supports $filter(any).
// returns a []Parameterable when successful
func (m *TaskDefinition) GetParameters()([]Parameterable) {
    return m.parameters
}
// GetVersion gets the version property value. The version number of the taskDefinition. New records are pushed when we add support for new parameters.Supports $filter(ge, gt, le, lt, eq, ne) and $orderby.
// returns a *int32 when successful
func (m *TaskDefinition) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *TaskDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("continueOnError", m.GetContinueOnError())
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
    if m.GetParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategory sets the category property value. The category property
func (m *TaskDefinition) SetCategory(value *LifecycleTaskCategory)() {
    m.category = value
}
// SetContinueOnError sets the continueOnError property value. Defines if the workflow will continue if the task has an error.
func (m *TaskDefinition) SetContinueOnError(value *bool)() {
    m.continueOnError = value
}
// SetDescription sets the description property value. The description of the taskDefinition.
func (m *TaskDefinition) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the taskDefinition.Supports $filter(eq, ne) and $orderby.
func (m *TaskDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetParameters sets the parameters property value. The parameters that must be supplied when creating a workflow task object.Supports $filter(any).
func (m *TaskDefinition) SetParameters(value []Parameterable)() {
    m.parameters = value
}
// SetVersion sets the version property value. The version number of the taskDefinition. New records are pushed when we add support for new parameters.Supports $filter(ge, gt, le, lt, eq, ne) and $orderby.
func (m *TaskDefinition) SetVersion(value *int32)() {
    m.version = value
}
type TaskDefinitionable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*LifecycleTaskCategory)
    GetContinueOnError()(*bool)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetParameters()([]Parameterable)
    GetVersion()(*int32)
    SetCategory(value *LifecycleTaskCategory)()
    SetContinueOnError(value *bool)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetParameters(value []Parameterable)()
    SetVersion(value *int32)()
}
