package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttackSimulationOperation the status of a long-running operation.
type AttackSimulationOperation struct {
    LongRunningOperation
    // Percentage of completion of the respective operation.
    percentageCompleted *int32
    // Tenant identifier.
    tenantId *string
    // The attack simulation operation type. Possible values are: createSimulation, updateSimulation, unknownFutureValue.
    typeEscaped *AttackSimulationOperationType
}
// NewAttackSimulationOperation instantiates a new AttackSimulationOperation and sets the default values.
func NewAttackSimulationOperation()(*AttackSimulationOperation) {
    m := &AttackSimulationOperation{
        LongRunningOperation: *NewLongRunningOperation(),
    }
    return m
}
// CreateAttackSimulationOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAttackSimulationOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttackSimulationOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AttackSimulationOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LongRunningOperation.GetFieldDeserializers()
    res["percentageCompleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentageCompleted(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttackSimulationOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*AttackSimulationOperationType))
        }
        return nil
    }
    return res
}
// GetPercentageCompleted gets the percentageCompleted property value. Percentage of completion of the respective operation.
// returns a *int32 when successful
func (m *AttackSimulationOperation) GetPercentageCompleted()(*int32) {
    return m.percentageCompleted
}
// GetTenantId gets the tenantId property value. Tenant identifier.
// returns a *string when successful
func (m *AttackSimulationOperation) GetTenantId()(*string) {
    return m.tenantId
}
// GetTypeEscaped gets the type property value. The attack simulation operation type. Possible values are: createSimulation, updateSimulation, unknownFutureValue.
// returns a *AttackSimulationOperationType when successful
func (m *AttackSimulationOperation) GetTypeEscaped()(*AttackSimulationOperationType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *AttackSimulationOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LongRunningOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("percentageCompleted", m.GetPercentageCompleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPercentageCompleted sets the percentageCompleted property value. Percentage of completion of the respective operation.
func (m *AttackSimulationOperation) SetPercentageCompleted(value *int32)() {
    m.percentageCompleted = value
}
// SetTenantId sets the tenantId property value. Tenant identifier.
func (m *AttackSimulationOperation) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetTypeEscaped sets the type property value. The attack simulation operation type. Possible values are: createSimulation, updateSimulation, unknownFutureValue.
func (m *AttackSimulationOperation) SetTypeEscaped(value *AttackSimulationOperationType)() {
    m.typeEscaped = value
}
type AttackSimulationOperationable interface {
    LongRunningOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPercentageCompleted()(*int32)
    GetTenantId()(*string)
    GetTypeEscaped()(*AttackSimulationOperationType)
    SetPercentageCompleted(value *int32)()
    SetTenantId(value *string)()
    SetTypeEscaped(value *AttackSimulationOperationType)()
}
