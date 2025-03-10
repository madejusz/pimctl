package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageAssignmentRequestorSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // False indicates that the requestor isn't permitted to include a schedule in their request.
    allowCustomAssignmentSchedule *bool
    // True allows on-behalf-of requestors to create a request to add access for another principal.
    enableOnBehalfRequestorsToAddAccess *bool
    // True allows on-behalf-of requestors to create a request to remove access for another principal.
    enableOnBehalfRequestorsToRemoveAccess *bool
    // True allows on-behalf-of requestors to create a request to update access for another principal.
    enableOnBehalfRequestorsToUpdateAccess *bool
    // True allows requestors to create a request to add access for themselves.
    enableTargetsToSelfAddAccess *bool
    // True allows requestors to create a request to remove their access.
    enableTargetsToSelfRemoveAccess *bool
    // True allows requestors to create a request to update their access.
    enableTargetsToSelfUpdateAccess *bool
    // The OdataType property
    odataType *string
    // The principals who can request on-behalf-of others.
    onBehalfRequestors []SubjectSetable
}
// NewAccessPackageAssignmentRequestorSettings instantiates a new AccessPackageAssignmentRequestorSettings and sets the default values.
func NewAccessPackageAssignmentRequestorSettings()(*AccessPackageAssignmentRequestorSettings) {
    m := &AccessPackageAssignmentRequestorSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessPackageAssignmentRequestorSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageAssignmentRequestorSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentRequestorSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccessPackageAssignmentRequestorSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowCustomAssignmentSchedule gets the allowCustomAssignmentSchedule property value. False indicates that the requestor isn't permitted to include a schedule in their request.
// returns a *bool when successful
func (m *AccessPackageAssignmentRequestorSettings) GetAllowCustomAssignmentSchedule()(*bool) {
    return m.allowCustomAssignmentSchedule
}
// GetEnableOnBehalfRequestorsToAddAccess gets the enableOnBehalfRequestorsToAddAccess property value. True allows on-behalf-of requestors to create a request to add access for another principal.
// returns a *bool when successful
func (m *AccessPackageAssignmentRequestorSettings) GetEnableOnBehalfRequestorsToAddAccess()(*bool) {
    return m.enableOnBehalfRequestorsToAddAccess
}
// GetEnableOnBehalfRequestorsToRemoveAccess gets the enableOnBehalfRequestorsToRemoveAccess property value. True allows on-behalf-of requestors to create a request to remove access for another principal.
// returns a *bool when successful
func (m *AccessPackageAssignmentRequestorSettings) GetEnableOnBehalfRequestorsToRemoveAccess()(*bool) {
    return m.enableOnBehalfRequestorsToRemoveAccess
}
// GetEnableOnBehalfRequestorsToUpdateAccess gets the enableOnBehalfRequestorsToUpdateAccess property value. True allows on-behalf-of requestors to create a request to update access for another principal.
// returns a *bool when successful
func (m *AccessPackageAssignmentRequestorSettings) GetEnableOnBehalfRequestorsToUpdateAccess()(*bool) {
    return m.enableOnBehalfRequestorsToUpdateAccess
}
// GetEnableTargetsToSelfAddAccess gets the enableTargetsToSelfAddAccess property value. True allows requestors to create a request to add access for themselves.
// returns a *bool when successful
func (m *AccessPackageAssignmentRequestorSettings) GetEnableTargetsToSelfAddAccess()(*bool) {
    return m.enableTargetsToSelfAddAccess
}
// GetEnableTargetsToSelfRemoveAccess gets the enableTargetsToSelfRemoveAccess property value. True allows requestors to create a request to remove their access.
// returns a *bool when successful
func (m *AccessPackageAssignmentRequestorSettings) GetEnableTargetsToSelfRemoveAccess()(*bool) {
    return m.enableTargetsToSelfRemoveAccess
}
// GetEnableTargetsToSelfUpdateAccess gets the enableTargetsToSelfUpdateAccess property value. True allows requestors to create a request to update their access.
// returns a *bool when successful
func (m *AccessPackageAssignmentRequestorSettings) GetEnableTargetsToSelfUpdateAccess()(*bool) {
    return m.enableTargetsToSelfUpdateAccess
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageAssignmentRequestorSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowCustomAssignmentSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCustomAssignmentSchedule(val)
        }
        return nil
    }
    res["enableOnBehalfRequestorsToAddAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableOnBehalfRequestorsToAddAccess(val)
        }
        return nil
    }
    res["enableOnBehalfRequestorsToRemoveAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableOnBehalfRequestorsToRemoveAccess(val)
        }
        return nil
    }
    res["enableOnBehalfRequestorsToUpdateAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableOnBehalfRequestorsToUpdateAccess(val)
        }
        return nil
    }
    res["enableTargetsToSelfAddAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTargetsToSelfAddAccess(val)
        }
        return nil
    }
    res["enableTargetsToSelfRemoveAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTargetsToSelfRemoveAccess(val)
        }
        return nil
    }
    res["enableTargetsToSelfUpdateAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTargetsToSelfUpdateAccess(val)
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
    res["onBehalfRequestors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectSetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SubjectSetable)
                }
            }
            m.SetOnBehalfRequestors(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AccessPackageAssignmentRequestorSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetOnBehalfRequestors gets the onBehalfRequestors property value. The principals who can request on-behalf-of others.
// returns a []SubjectSetable when successful
func (m *AccessPackageAssignmentRequestorSettings) GetOnBehalfRequestors()([]SubjectSetable) {
    return m.onBehalfRequestors
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentRequestorSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowCustomAssignmentSchedule", m.GetAllowCustomAssignmentSchedule())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableOnBehalfRequestorsToAddAccess", m.GetEnableOnBehalfRequestorsToAddAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableOnBehalfRequestorsToRemoveAccess", m.GetEnableOnBehalfRequestorsToRemoveAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableOnBehalfRequestorsToUpdateAccess", m.GetEnableOnBehalfRequestorsToUpdateAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableTargetsToSelfAddAccess", m.GetEnableTargetsToSelfAddAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableTargetsToSelfRemoveAccess", m.GetEnableTargetsToSelfRemoveAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableTargetsToSelfUpdateAccess", m.GetEnableTargetsToSelfUpdateAccess())
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
    if m.GetOnBehalfRequestors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnBehalfRequestors()))
        for i, v := range m.GetOnBehalfRequestors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("onBehalfRequestors", cast)
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
func (m *AccessPackageAssignmentRequestorSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowCustomAssignmentSchedule sets the allowCustomAssignmentSchedule property value. False indicates that the requestor isn't permitted to include a schedule in their request.
func (m *AccessPackageAssignmentRequestorSettings) SetAllowCustomAssignmentSchedule(value *bool)() {
    m.allowCustomAssignmentSchedule = value
}
// SetEnableOnBehalfRequestorsToAddAccess sets the enableOnBehalfRequestorsToAddAccess property value. True allows on-behalf-of requestors to create a request to add access for another principal.
func (m *AccessPackageAssignmentRequestorSettings) SetEnableOnBehalfRequestorsToAddAccess(value *bool)() {
    m.enableOnBehalfRequestorsToAddAccess = value
}
// SetEnableOnBehalfRequestorsToRemoveAccess sets the enableOnBehalfRequestorsToRemoveAccess property value. True allows on-behalf-of requestors to create a request to remove access for another principal.
func (m *AccessPackageAssignmentRequestorSettings) SetEnableOnBehalfRequestorsToRemoveAccess(value *bool)() {
    m.enableOnBehalfRequestorsToRemoveAccess = value
}
// SetEnableOnBehalfRequestorsToUpdateAccess sets the enableOnBehalfRequestorsToUpdateAccess property value. True allows on-behalf-of requestors to create a request to update access for another principal.
func (m *AccessPackageAssignmentRequestorSettings) SetEnableOnBehalfRequestorsToUpdateAccess(value *bool)() {
    m.enableOnBehalfRequestorsToUpdateAccess = value
}
// SetEnableTargetsToSelfAddAccess sets the enableTargetsToSelfAddAccess property value. True allows requestors to create a request to add access for themselves.
func (m *AccessPackageAssignmentRequestorSettings) SetEnableTargetsToSelfAddAccess(value *bool)() {
    m.enableTargetsToSelfAddAccess = value
}
// SetEnableTargetsToSelfRemoveAccess sets the enableTargetsToSelfRemoveAccess property value. True allows requestors to create a request to remove their access.
func (m *AccessPackageAssignmentRequestorSettings) SetEnableTargetsToSelfRemoveAccess(value *bool)() {
    m.enableTargetsToSelfRemoveAccess = value
}
// SetEnableTargetsToSelfUpdateAccess sets the enableTargetsToSelfUpdateAccess property value. True allows requestors to create a request to update their access.
func (m *AccessPackageAssignmentRequestorSettings) SetEnableTargetsToSelfUpdateAccess(value *bool)() {
    m.enableTargetsToSelfUpdateAccess = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageAssignmentRequestorSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOnBehalfRequestors sets the onBehalfRequestors property value. The principals who can request on-behalf-of others.
func (m *AccessPackageAssignmentRequestorSettings) SetOnBehalfRequestors(value []SubjectSetable)() {
    m.onBehalfRequestors = value
}
type AccessPackageAssignmentRequestorSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowCustomAssignmentSchedule()(*bool)
    GetEnableOnBehalfRequestorsToAddAccess()(*bool)
    GetEnableOnBehalfRequestorsToRemoveAccess()(*bool)
    GetEnableOnBehalfRequestorsToUpdateAccess()(*bool)
    GetEnableTargetsToSelfAddAccess()(*bool)
    GetEnableTargetsToSelfRemoveAccess()(*bool)
    GetEnableTargetsToSelfUpdateAccess()(*bool)
    GetOdataType()(*string)
    GetOnBehalfRequestors()([]SubjectSetable)
    SetAllowCustomAssignmentSchedule(value *bool)()
    SetEnableOnBehalfRequestorsToAddAccess(value *bool)()
    SetEnableOnBehalfRequestorsToRemoveAccess(value *bool)()
    SetEnableOnBehalfRequestorsToUpdateAccess(value *bool)()
    SetEnableTargetsToSelfAddAccess(value *bool)()
    SetEnableTargetsToSelfRemoveAccess(value *bool)()
    SetEnableTargetsToSelfUpdateAccess(value *bool)()
    SetOdataType(value *string)()
    SetOnBehalfRequestors(value []SubjectSetable)()
}
