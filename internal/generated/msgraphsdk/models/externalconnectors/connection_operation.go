package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type ConnectionOperation struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // If status is failed, provides more information about the error that caused the failure.
    error ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable
    // Indicates the status of the asynchronous operation. Possible values are: unspecified, inprogress, completed, failed, unknownFutureValue.
    status *ConnectionOperationStatus
}
// NewConnectionOperation instantiates a new ConnectionOperation and sets the default values.
func NewConnectionOperation()(*ConnectionOperation) {
    m := &ConnectionOperation{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateConnectionOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConnectionOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnectionOperation(), nil
}
// GetError gets the error property value. If status is failed, provides more information about the error that caused the failure.
// returns a PublicErrorable when successful
func (m *ConnectionOperation) GetError()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConnectionOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectionOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ConnectionOperationStatus))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. Indicates the status of the asynchronous operation. Possible values are: unspecified, inprogress, completed, failed, unknownFutureValue.
// returns a *ConnectionOperationStatus when successful
func (m *ConnectionOperation) GetStatus()(*ConnectionOperationStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *ConnectionOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. If status is failed, provides more information about the error that caused the failure.
func (m *ConnectionOperation) SetError(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable)() {
    m.error = value
}
// SetStatus sets the status property value. Indicates the status of the asynchronous operation. Possible values are: unspecified, inprogress, completed, failed, unknownFutureValue.
func (m *ConnectionOperation) SetStatus(value *ConnectionOperationStatus)() {
    m.status = value
}
type ConnectionOperationable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable)
    GetStatus()(*ConnectionOperationStatus)
    SetError(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable)()
    SetStatus(value *ConnectionOperationStatus)()
}
