package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualEventRegistration struct {
    Entity
    // Date and time when the registrant cancels their registration for the virtual event. Only appears when applicable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    cancelationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Email address of the registrant.
    email *string
    // First name of the registrant.
    firstName *string
    // Last name of the registrant.
    lastName *string
    // The registrant's preferred language.
    preferredLanguage *string
    // The registrant's time zone details.
    preferredTimezone *string
    // Date and time when the registrant registers for the virtual event. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    registrationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The registrant's answer to the registration questions.
    registrationQuestionAnswers []VirtualEventRegistrationQuestionAnswerable
    // Sessions for a registration.
    sessions []VirtualEventSessionable
    // Registration status of the registrant. Read-only. Possible values are registered, canceled, waitlisted, pendingApproval, rejectedByOrganizer, and unknownFutureValue.
    status *VirtualEventAttendeeRegistrationStatus
    // The registrant's ID in Microsoft Entra ID. Only appears when the registrant is registered in Microsoft Entra ID.
    userId *string
}
// NewVirtualEventRegistration instantiates a new VirtualEventRegistration and sets the default values.
func NewVirtualEventRegistration()(*VirtualEventRegistration) {
    m := &VirtualEventRegistration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualEventRegistrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEventRegistrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventRegistration(), nil
}
// GetCancelationDateTime gets the cancelationDateTime property value. Date and time when the registrant cancels their registration for the virtual event. Only appears when applicable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *VirtualEventRegistration) GetCancelationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.cancelationDateTime
}
// GetEmail gets the email property value. Email address of the registrant.
// returns a *string when successful
func (m *VirtualEventRegistration) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualEventRegistration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cancelationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCancelationDateTime(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["firstName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstName(val)
        }
        return nil
    }
    res["lastName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastName(val)
        }
        return nil
    }
    res["preferredLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguage(val)
        }
        return nil
    }
    res["preferredTimezone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredTimezone(val)
        }
        return nil
    }
    res["registrationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationDateTime(val)
        }
        return nil
    }
    res["registrationQuestionAnswers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventRegistrationQuestionAnswerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventRegistrationQuestionAnswerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventRegistrationQuestionAnswerable)
                }
            }
            m.SetRegistrationQuestionAnswers(res)
        }
        return nil
    }
    res["sessions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventSessionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventSessionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventSessionable)
                }
            }
            m.SetSessions(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVirtualEventAttendeeRegistrationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*VirtualEventAttendeeRegistrationStatus))
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
// GetFirstName gets the firstName property value. First name of the registrant.
// returns a *string when successful
func (m *VirtualEventRegistration) GetFirstName()(*string) {
    return m.firstName
}
// GetLastName gets the lastName property value. Last name of the registrant.
// returns a *string when successful
func (m *VirtualEventRegistration) GetLastName()(*string) {
    return m.lastName
}
// GetPreferredLanguage gets the preferredLanguage property value. The registrant's preferred language.
// returns a *string when successful
func (m *VirtualEventRegistration) GetPreferredLanguage()(*string) {
    return m.preferredLanguage
}
// GetPreferredTimezone gets the preferredTimezone property value. The registrant's time zone details.
// returns a *string when successful
func (m *VirtualEventRegistration) GetPreferredTimezone()(*string) {
    return m.preferredTimezone
}
// GetRegistrationDateTime gets the registrationDateTime property value. Date and time when the registrant registers for the virtual event. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *VirtualEventRegistration) GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.registrationDateTime
}
// GetRegistrationQuestionAnswers gets the registrationQuestionAnswers property value. The registrant's answer to the registration questions.
// returns a []VirtualEventRegistrationQuestionAnswerable when successful
func (m *VirtualEventRegistration) GetRegistrationQuestionAnswers()([]VirtualEventRegistrationQuestionAnswerable) {
    return m.registrationQuestionAnswers
}
// GetSessions gets the sessions property value. Sessions for a registration.
// returns a []VirtualEventSessionable when successful
func (m *VirtualEventRegistration) GetSessions()([]VirtualEventSessionable) {
    return m.sessions
}
// GetStatus gets the status property value. Registration status of the registrant. Read-only. Possible values are registered, canceled, waitlisted, pendingApproval, rejectedByOrganizer, and unknownFutureValue.
// returns a *VirtualEventAttendeeRegistrationStatus when successful
func (m *VirtualEventRegistration) GetStatus()(*VirtualEventAttendeeRegistrationStatus) {
    return m.status
}
// GetUserId gets the userId property value. The registrant's ID in Microsoft Entra ID. Only appears when the registrant is registered in Microsoft Entra ID.
// returns a *string when successful
func (m *VirtualEventRegistration) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *VirtualEventRegistration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("cancelationDateTime", m.GetCancelationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("firstName", m.GetFirstName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastName", m.GetLastName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredLanguage", m.GetPreferredLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredTimezone", m.GetPreferredTimezone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("registrationDateTime", m.GetRegistrationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRegistrationQuestionAnswers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegistrationQuestionAnswers()))
        for i, v := range m.GetRegistrationQuestionAnswers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("registrationQuestionAnswers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSessions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSessions()))
        for i, v := range m.GetSessions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sessions", cast)
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
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCancelationDateTime sets the cancelationDateTime property value. Date and time when the registrant cancels their registration for the virtual event. Only appears when applicable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *VirtualEventRegistration) SetCancelationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.cancelationDateTime = value
}
// SetEmail sets the email property value. Email address of the registrant.
func (m *VirtualEventRegistration) SetEmail(value *string)() {
    m.email = value
}
// SetFirstName sets the firstName property value. First name of the registrant.
func (m *VirtualEventRegistration) SetFirstName(value *string)() {
    m.firstName = value
}
// SetLastName sets the lastName property value. Last name of the registrant.
func (m *VirtualEventRegistration) SetLastName(value *string)() {
    m.lastName = value
}
// SetPreferredLanguage sets the preferredLanguage property value. The registrant's preferred language.
func (m *VirtualEventRegistration) SetPreferredLanguage(value *string)() {
    m.preferredLanguage = value
}
// SetPreferredTimezone sets the preferredTimezone property value. The registrant's time zone details.
func (m *VirtualEventRegistration) SetPreferredTimezone(value *string)() {
    m.preferredTimezone = value
}
// SetRegistrationDateTime sets the registrationDateTime property value. Date and time when the registrant registers for the virtual event. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *VirtualEventRegistration) SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.registrationDateTime = value
}
// SetRegistrationQuestionAnswers sets the registrationQuestionAnswers property value. The registrant's answer to the registration questions.
func (m *VirtualEventRegistration) SetRegistrationQuestionAnswers(value []VirtualEventRegistrationQuestionAnswerable)() {
    m.registrationQuestionAnswers = value
}
// SetSessions sets the sessions property value. Sessions for a registration.
func (m *VirtualEventRegistration) SetSessions(value []VirtualEventSessionable)() {
    m.sessions = value
}
// SetStatus sets the status property value. Registration status of the registrant. Read-only. Possible values are registered, canceled, waitlisted, pendingApproval, rejectedByOrganizer, and unknownFutureValue.
func (m *VirtualEventRegistration) SetStatus(value *VirtualEventAttendeeRegistrationStatus)() {
    m.status = value
}
// SetUserId sets the userId property value. The registrant's ID in Microsoft Entra ID. Only appears when the registrant is registered in Microsoft Entra ID.
func (m *VirtualEventRegistration) SetUserId(value *string)() {
    m.userId = value
}
type VirtualEventRegistrationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCancelationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmail()(*string)
    GetFirstName()(*string)
    GetLastName()(*string)
    GetPreferredLanguage()(*string)
    GetPreferredTimezone()(*string)
    GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRegistrationQuestionAnswers()([]VirtualEventRegistrationQuestionAnswerable)
    GetSessions()([]VirtualEventSessionable)
    GetStatus()(*VirtualEventAttendeeRegistrationStatus)
    GetUserId()(*string)
    SetCancelationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmail(value *string)()
    SetFirstName(value *string)()
    SetLastName(value *string)()
    SetPreferredLanguage(value *string)()
    SetPreferredTimezone(value *string)()
    SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRegistrationQuestionAnswers(value []VirtualEventRegistrationQuestionAnswerable)()
    SetSessions(value []VirtualEventSessionable)()
    SetStatus(value *VirtualEventAttendeeRegistrationStatus)()
    SetUserId(value *string)()
}
