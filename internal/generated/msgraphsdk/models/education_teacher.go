package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EducationTeacher struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // ID of the teacher in the source system.
    externalId *string
    // The OdataType property
    odataType *string
    // Teacher number.
    teacherNumber *string
}
// NewEducationTeacher instantiates a new EducationTeacher and sets the default values.
func NewEducationTeacher()(*EducationTeacher) {
    m := &EducationTeacher{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEducationTeacherFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationTeacherFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationTeacher(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EducationTeacher) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExternalId gets the externalId property value. ID of the teacher in the source system.
// returns a *string when successful
func (m *EducationTeacher) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationTeacher) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
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
    res["teacherNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeacherNumber(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *EducationTeacher) GetOdataType()(*string) {
    return m.odataType
}
// GetTeacherNumber gets the teacherNumber property value. Teacher number.
// returns a *string when successful
func (m *EducationTeacher) GetTeacherNumber()(*string) {
    return m.teacherNumber
}
// Serialize serializes information the current object
func (m *EducationTeacher) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
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
        err := writer.WriteStringValue("teacherNumber", m.GetTeacherNumber())
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
func (m *EducationTeacher) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExternalId sets the externalId property value. ID of the teacher in the source system.
func (m *EducationTeacher) SetExternalId(value *string)() {
    m.externalId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationTeacher) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTeacherNumber sets the teacherNumber property value. Teacher number.
func (m *EducationTeacher) SetTeacherNumber(value *string)() {
    m.teacherNumber = value
}
type EducationTeacherable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExternalId()(*string)
    GetOdataType()(*string)
    GetTeacherNumber()(*string)
    SetExternalId(value *string)()
    SetOdataType(value *string)()
    SetTeacherNumber(value *string)()
}
