package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type Task struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // Arguments included within the task.  For guidance to configure this property, see Configure the arguments for built-in Lifecycle Workflow tasks. Required.
    arguments []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.KeyValuePairable
    // The category property
    category *LifecycleTaskCategory
    // A Boolean value that specifies whether, if this task fails, the workflow stops, and subsequent tasks aren't run. Optional.
    continueOnError *bool
    // A string that describes the purpose of the task for administrative use. Optional.
    description *string
    // A unique string that identifies the task. Required.Supports $filter(eq, ne) and orderBy.
    displayName *string
    // An integer that states in what order the task runs in a workflow.Supports $orderby.
    executionSequence *int32
    // A Boolean value that denotes whether the task is set to run or not. Optional.Supports $filter(eq, ne) and orderBy.
    isEnabled *bool
    // A unique template identifier for the task. For more information about the tasks that Lifecycle Workflows currently supports and their unique identifiers, see Configure the arguments for built-in Lifecycle Workflow tasks. Required.Supports $filter(eq, ne).
    taskDefinitionId *string
    // The result of processing the task.
    taskProcessingResults []TaskProcessingResultable
}
// NewTask instantiates a new Task and sets the default values.
func NewTask()(*Task) {
    m := &Task{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTask(), nil
}
// GetArguments gets the arguments property value. Arguments included within the task.  For guidance to configure this property, see Configure the arguments for built-in Lifecycle Workflow tasks. Required.
// returns a []KeyValuePairable when successful
func (m *Task) GetArguments()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.KeyValuePairable) {
    return m.arguments
}
// GetCategory gets the category property value. The category property
// returns a *LifecycleTaskCategory when successful
func (m *Task) GetCategory()(*LifecycleTaskCategory) {
    return m.category
}
// GetContinueOnError gets the continueOnError property value. A Boolean value that specifies whether, if this task fails, the workflow stops, and subsequent tasks aren't run. Optional.
// returns a *bool when successful
func (m *Task) GetContinueOnError()(*bool) {
    return m.continueOnError
}
// GetDescription gets the description property value. A string that describes the purpose of the task for administrative use. Optional.
// returns a *string when successful
func (m *Task) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. A unique string that identifies the task. Required.Supports $filter(eq, ne) and orderBy.
// returns a *string when successful
func (m *Task) GetDisplayName()(*string) {
    return m.displayName
}
// GetExecutionSequence gets the executionSequence property value. An integer that states in what order the task runs in a workflow.Supports $orderby.
// returns a *int32 when successful
func (m *Task) GetExecutionSequence()(*int32) {
    return m.executionSequence
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Task) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["arguments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.KeyValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.KeyValuePairable)
                }
            }
            m.SetArguments(res)
        }
        return nil
    }
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
    res["executionSequence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecutionSequence(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["taskDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaskDefinitionId(val)
        }
        return nil
    }
    res["taskProcessingResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaskProcessingResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaskProcessingResultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TaskProcessingResultable)
                }
            }
            m.SetTaskProcessingResults(res)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. A Boolean value that denotes whether the task is set to run or not. Optional.Supports $filter(eq, ne) and orderBy.
// returns a *bool when successful
func (m *Task) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetTaskDefinitionId gets the taskDefinitionId property value. A unique template identifier for the task. For more information about the tasks that Lifecycle Workflows currently supports and their unique identifiers, see Configure the arguments for built-in Lifecycle Workflow tasks. Required.Supports $filter(eq, ne).
// returns a *string when successful
func (m *Task) GetTaskDefinitionId()(*string) {
    return m.taskDefinitionId
}
// GetTaskProcessingResults gets the taskProcessingResults property value. The result of processing the task.
// returns a []TaskProcessingResultable when successful
func (m *Task) GetTaskProcessingResults()([]TaskProcessingResultable) {
    return m.taskProcessingResults
}
// Serialize serializes information the current object
func (m *Task) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetArguments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetArguments()))
        for i, v := range m.GetArguments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("arguments", cast)
        if err != nil {
            return err
        }
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
    {
        err = writer.WriteInt32Value("executionSequence", m.GetExecutionSequence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taskDefinitionId", m.GetTaskDefinitionId())
        if err != nil {
            return err
        }
    }
    if m.GetTaskProcessingResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaskProcessingResults()))
        for i, v := range m.GetTaskProcessingResults() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("taskProcessingResults", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArguments sets the arguments property value. Arguments included within the task.  For guidance to configure this property, see Configure the arguments for built-in Lifecycle Workflow tasks. Required.
func (m *Task) SetArguments(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.KeyValuePairable)() {
    m.arguments = value
}
// SetCategory sets the category property value. The category property
func (m *Task) SetCategory(value *LifecycleTaskCategory)() {
    m.category = value
}
// SetContinueOnError sets the continueOnError property value. A Boolean value that specifies whether, if this task fails, the workflow stops, and subsequent tasks aren't run. Optional.
func (m *Task) SetContinueOnError(value *bool)() {
    m.continueOnError = value
}
// SetDescription sets the description property value. A string that describes the purpose of the task for administrative use. Optional.
func (m *Task) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. A unique string that identifies the task. Required.Supports $filter(eq, ne) and orderBy.
func (m *Task) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExecutionSequence sets the executionSequence property value. An integer that states in what order the task runs in a workflow.Supports $orderby.
func (m *Task) SetExecutionSequence(value *int32)() {
    m.executionSequence = value
}
// SetIsEnabled sets the isEnabled property value. A Boolean value that denotes whether the task is set to run or not. Optional.Supports $filter(eq, ne) and orderBy.
func (m *Task) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetTaskDefinitionId sets the taskDefinitionId property value. A unique template identifier for the task. For more information about the tasks that Lifecycle Workflows currently supports and their unique identifiers, see Configure the arguments for built-in Lifecycle Workflow tasks. Required.Supports $filter(eq, ne).
func (m *Task) SetTaskDefinitionId(value *string)() {
    m.taskDefinitionId = value
}
// SetTaskProcessingResults sets the taskProcessingResults property value. The result of processing the task.
func (m *Task) SetTaskProcessingResults(value []TaskProcessingResultable)() {
    m.taskProcessingResults = value
}
type Taskable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArguments()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.KeyValuePairable)
    GetCategory()(*LifecycleTaskCategory)
    GetContinueOnError()(*bool)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExecutionSequence()(*int32)
    GetIsEnabled()(*bool)
    GetTaskDefinitionId()(*string)
    GetTaskProcessingResults()([]TaskProcessingResultable)
    SetArguments(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.KeyValuePairable)()
    SetCategory(value *LifecycleTaskCategory)()
    SetContinueOnError(value *bool)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExecutionSequence(value *int32)()
    SetIsEnabled(value *bool)()
    SetTaskDefinitionId(value *string)()
    SetTaskProcessingResults(value []TaskProcessingResultable)()
}
