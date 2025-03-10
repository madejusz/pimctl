package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EducationSubmission struct {
    Entity
    // The user that marked the submission as excused.
    excusedBy IdentitySetable
    // The time that the submission was excused. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    excusedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The outcomes property
    outcomes []EducationOutcomeable
    // User who moved the status of this submission to reassigned.
    reassignedBy IdentitySetable
    // Moment in time when the submission was reassigned. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    reassignedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Who this submission is assigned to.
    recipient EducationSubmissionRecipientable
    // The resources property
    resources []EducationSubmissionResourceable
    // Folder where all file resources for this submission need to be stored.
    resourcesFolderUrl *string
    // User who moved the status of this submission to returned.
    returnedBy IdentitySetable
    // Moment in time when the submission was returned. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    returnedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Read-only. Possible values are: excused, reassigned, returned, submitted and working. You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: excused and reassigned.
    status *EducationSubmissionStatus
    // User who moved the resource into the submitted state.
    submittedBy IdentitySetable
    // Moment in time when the submission was moved into the submitted state. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    submittedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The submittedResources property
    submittedResources []EducationSubmissionResourceable
    // User who moved the resource from submitted into the working state.
    unsubmittedBy IdentitySetable
    // Moment in time when the submission was moved from submitted into the working state. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    unsubmittedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The deep link URL for the given submission.
    webUrl *string
}
// NewEducationSubmission instantiates a new EducationSubmission and sets the default values.
func NewEducationSubmission()(*EducationSubmission) {
    m := &EducationSubmission{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationSubmissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationSubmissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSubmission(), nil
}
// GetExcusedBy gets the excusedBy property value. The user that marked the submission as excused.
// returns a IdentitySetable when successful
func (m *EducationSubmission) GetExcusedBy()(IdentitySetable) {
    return m.excusedBy
}
// GetExcusedDateTime gets the excusedDateTime property value. The time that the submission was excused. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *EducationSubmission) GetExcusedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.excusedDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationSubmission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["excusedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcusedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["excusedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcusedDateTime(val)
        }
        return nil
    }
    res["outcomes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationOutcomeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationOutcomeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EducationOutcomeable)
                }
            }
            m.SetOutcomes(res)
        }
        return nil
    }
    res["reassignedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReassignedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["reassignedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReassignedDateTime(val)
        }
        return nil
    }
    res["recipient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSubmissionRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipient(val.(EducationSubmissionRecipientable))
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSubmissionResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSubmissionResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EducationSubmissionResourceable)
                }
            }
            m.SetResources(res)
        }
        return nil
    }
    res["resourcesFolderUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourcesFolderUrl(val)
        }
        return nil
    }
    res["returnedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReturnedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["returnedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReturnedDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationSubmissionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*EducationSubmissionStatus))
        }
        return nil
    }
    res["submittedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmittedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["submittedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmittedDateTime(val)
        }
        return nil
    }
    res["submittedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSubmissionResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSubmissionResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EducationSubmissionResourceable)
                }
            }
            m.SetSubmittedResources(res)
        }
        return nil
    }
    res["unsubmittedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnsubmittedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["unsubmittedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnsubmittedDateTime(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetOutcomes gets the outcomes property value. The outcomes property
// returns a []EducationOutcomeable when successful
func (m *EducationSubmission) GetOutcomes()([]EducationOutcomeable) {
    return m.outcomes
}
// GetReassignedBy gets the reassignedBy property value. User who moved the status of this submission to reassigned.
// returns a IdentitySetable when successful
func (m *EducationSubmission) GetReassignedBy()(IdentitySetable) {
    return m.reassignedBy
}
// GetReassignedDateTime gets the reassignedDateTime property value. Moment in time when the submission was reassigned. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *EducationSubmission) GetReassignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.reassignedDateTime
}
// GetRecipient gets the recipient property value. Who this submission is assigned to.
// returns a EducationSubmissionRecipientable when successful
func (m *EducationSubmission) GetRecipient()(EducationSubmissionRecipientable) {
    return m.recipient
}
// GetResources gets the resources property value. The resources property
// returns a []EducationSubmissionResourceable when successful
func (m *EducationSubmission) GetResources()([]EducationSubmissionResourceable) {
    return m.resources
}
// GetResourcesFolderUrl gets the resourcesFolderUrl property value. Folder where all file resources for this submission need to be stored.
// returns a *string when successful
func (m *EducationSubmission) GetResourcesFolderUrl()(*string) {
    return m.resourcesFolderUrl
}
// GetReturnedBy gets the returnedBy property value. User who moved the status of this submission to returned.
// returns a IdentitySetable when successful
func (m *EducationSubmission) GetReturnedBy()(IdentitySetable) {
    return m.returnedBy
}
// GetReturnedDateTime gets the returnedDateTime property value. Moment in time when the submission was returned. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *EducationSubmission) GetReturnedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.returnedDateTime
}
// GetStatus gets the status property value. Read-only. Possible values are: excused, reassigned, returned, submitted and working. You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: excused and reassigned.
// returns a *EducationSubmissionStatus when successful
func (m *EducationSubmission) GetStatus()(*EducationSubmissionStatus) {
    return m.status
}
// GetSubmittedBy gets the submittedBy property value. User who moved the resource into the submitted state.
// returns a IdentitySetable when successful
func (m *EducationSubmission) GetSubmittedBy()(IdentitySetable) {
    return m.submittedBy
}
// GetSubmittedDateTime gets the submittedDateTime property value. Moment in time when the submission was moved into the submitted state. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *EducationSubmission) GetSubmittedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.submittedDateTime
}
// GetSubmittedResources gets the submittedResources property value. The submittedResources property
// returns a []EducationSubmissionResourceable when successful
func (m *EducationSubmission) GetSubmittedResources()([]EducationSubmissionResourceable) {
    return m.submittedResources
}
// GetUnsubmittedBy gets the unsubmittedBy property value. User who moved the resource from submitted into the working state.
// returns a IdentitySetable when successful
func (m *EducationSubmission) GetUnsubmittedBy()(IdentitySetable) {
    return m.unsubmittedBy
}
// GetUnsubmittedDateTime gets the unsubmittedDateTime property value. Moment in time when the submission was moved from submitted into the working state. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *EducationSubmission) GetUnsubmittedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.unsubmittedDateTime
}
// GetWebUrl gets the webUrl property value. The deep link URL for the given submission.
// returns a *string when successful
func (m *EducationSubmission) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *EducationSubmission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOutcomes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOutcomes()))
        for i, v := range m.GetOutcomes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("outcomes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recipient", m.GetRecipient())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubmittedResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubmittedResources()))
        for i, v := range m.GetSubmittedResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("submittedResources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExcusedBy sets the excusedBy property value. The user that marked the submission as excused.
func (m *EducationSubmission) SetExcusedBy(value IdentitySetable)() {
    m.excusedBy = value
}
// SetExcusedDateTime sets the excusedDateTime property value. The time that the submission was excused. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *EducationSubmission) SetExcusedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.excusedDateTime = value
}
// SetOutcomes sets the outcomes property value. The outcomes property
func (m *EducationSubmission) SetOutcomes(value []EducationOutcomeable)() {
    m.outcomes = value
}
// SetReassignedBy sets the reassignedBy property value. User who moved the status of this submission to reassigned.
func (m *EducationSubmission) SetReassignedBy(value IdentitySetable)() {
    m.reassignedBy = value
}
// SetReassignedDateTime sets the reassignedDateTime property value. Moment in time when the submission was reassigned. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *EducationSubmission) SetReassignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reassignedDateTime = value
}
// SetRecipient sets the recipient property value. Who this submission is assigned to.
func (m *EducationSubmission) SetRecipient(value EducationSubmissionRecipientable)() {
    m.recipient = value
}
// SetResources sets the resources property value. The resources property
func (m *EducationSubmission) SetResources(value []EducationSubmissionResourceable)() {
    m.resources = value
}
// SetResourcesFolderUrl sets the resourcesFolderUrl property value. Folder where all file resources for this submission need to be stored.
func (m *EducationSubmission) SetResourcesFolderUrl(value *string)() {
    m.resourcesFolderUrl = value
}
// SetReturnedBy sets the returnedBy property value. User who moved the status of this submission to returned.
func (m *EducationSubmission) SetReturnedBy(value IdentitySetable)() {
    m.returnedBy = value
}
// SetReturnedDateTime sets the returnedDateTime property value. Moment in time when the submission was returned. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *EducationSubmission) SetReturnedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.returnedDateTime = value
}
// SetStatus sets the status property value. Read-only. Possible values are: excused, reassigned, returned, submitted and working. You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: excused and reassigned.
func (m *EducationSubmission) SetStatus(value *EducationSubmissionStatus)() {
    m.status = value
}
// SetSubmittedBy sets the submittedBy property value. User who moved the resource into the submitted state.
func (m *EducationSubmission) SetSubmittedBy(value IdentitySetable)() {
    m.submittedBy = value
}
// SetSubmittedDateTime sets the submittedDateTime property value. Moment in time when the submission was moved into the submitted state. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *EducationSubmission) SetSubmittedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.submittedDateTime = value
}
// SetSubmittedResources sets the submittedResources property value. The submittedResources property
func (m *EducationSubmission) SetSubmittedResources(value []EducationSubmissionResourceable)() {
    m.submittedResources = value
}
// SetUnsubmittedBy sets the unsubmittedBy property value. User who moved the resource from submitted into the working state.
func (m *EducationSubmission) SetUnsubmittedBy(value IdentitySetable)() {
    m.unsubmittedBy = value
}
// SetUnsubmittedDateTime sets the unsubmittedDateTime property value. Moment in time when the submission was moved from submitted into the working state. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *EducationSubmission) SetUnsubmittedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.unsubmittedDateTime = value
}
// SetWebUrl sets the webUrl property value. The deep link URL for the given submission.
func (m *EducationSubmission) SetWebUrl(value *string)() {
    m.webUrl = value
}
type EducationSubmissionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcusedBy()(IdentitySetable)
    GetExcusedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOutcomes()([]EducationOutcomeable)
    GetReassignedBy()(IdentitySetable)
    GetReassignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecipient()(EducationSubmissionRecipientable)
    GetResources()([]EducationSubmissionResourceable)
    GetResourcesFolderUrl()(*string)
    GetReturnedBy()(IdentitySetable)
    GetReturnedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*EducationSubmissionStatus)
    GetSubmittedBy()(IdentitySetable)
    GetSubmittedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubmittedResources()([]EducationSubmissionResourceable)
    GetUnsubmittedBy()(IdentitySetable)
    GetUnsubmittedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetWebUrl()(*string)
    SetExcusedBy(value IdentitySetable)()
    SetExcusedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOutcomes(value []EducationOutcomeable)()
    SetReassignedBy(value IdentitySetable)()
    SetReassignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecipient(value EducationSubmissionRecipientable)()
    SetResources(value []EducationSubmissionResourceable)()
    SetResourcesFolderUrl(value *string)()
    SetReturnedBy(value IdentitySetable)()
    SetReturnedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *EducationSubmissionStatus)()
    SetSubmittedBy(value IdentitySetable)()
    SetSubmittedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubmittedResources(value []EducationSubmissionResourceable)()
    SetUnsubmittedBy(value IdentitySetable)()
    SetUnsubmittedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetWebUrl(value *string)()
}
