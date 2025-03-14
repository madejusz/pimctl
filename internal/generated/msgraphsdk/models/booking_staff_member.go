package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingStaffMember represents a staff member who provides services in a business.
type BookingStaffMember struct {
    BookingStaffMemberBase
    // True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
    availabilityIsAffectedByPersonalCalendar *bool
    // The date, time, and time zone when the staff member was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the staff member, as displayed to customers. Required.
    displayName *string
    // The email address of the staff member. This email address can be in the same Microsoft 365 tenant as the business, or in a different email domain. This email address can be used if the sendConfirmationsToOwner property is set to true in the scheduling policy of the business. Required.
    emailAddress *string
    // Indicates that a staff member is notified via email when a booking assigned to them is created or changed. The default value is true.
    isEmailNotificationEnabled *bool
    // The date, time, and time zone when the staff member was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The membershipStatus property
    membershipStatus *BookingStaffMembershipStatus
    // The role property
    role *BookingStaffRole
    // The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
    timeZone *string
    // True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
    useBusinessHours *bool
    // The range of hours each day of the week that the staff member is available for booking. By default, they're initialized to be the same as the businessHours property of the business.
    workingHours []BookingWorkHoursable
}
// NewBookingStaffMember instantiates a new BookingStaffMember and sets the default values.
func NewBookingStaffMember()(*BookingStaffMember) {
    m := &BookingStaffMember{
        BookingStaffMemberBase: *NewBookingStaffMemberBase(),
    }
    odataTypeValue := "#microsoft.graph.bookingStaffMember"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateBookingStaffMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBookingStaffMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingStaffMember(), nil
}
// GetAvailabilityIsAffectedByPersonalCalendar gets the availabilityIsAffectedByPersonalCalendar property value. True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
// returns a *bool when successful
func (m *BookingStaffMember) GetAvailabilityIsAffectedByPersonalCalendar()(*bool) {
    return m.availabilityIsAffectedByPersonalCalendar
}
// GetCreatedDateTime gets the createdDateTime property value. The date, time, and time zone when the staff member was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *BookingStaffMember) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The name of the staff member, as displayed to customers. Required.
// returns a *string when successful
func (m *BookingStaffMember) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmailAddress gets the emailAddress property value. The email address of the staff member. This email address can be in the same Microsoft 365 tenant as the business, or in a different email domain. This email address can be used if the sendConfirmationsToOwner property is set to true in the scheduling policy of the business. Required.
// returns a *string when successful
func (m *BookingStaffMember) GetEmailAddress()(*string) {
    return m.emailAddress
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BookingStaffMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BookingStaffMemberBase.GetFieldDeserializers()
    res["availabilityIsAffectedByPersonalCalendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityIsAffectedByPersonalCalendar(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["isEmailNotificationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEmailNotificationEnabled(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["membershipStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingStaffMembershipStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipStatus(val.(*BookingStaffMembershipStatus))
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingStaffRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*BookingStaffRole))
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    res["useBusinessHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseBusinessHours(val)
        }
        return nil
    }
    res["workingHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingWorkHoursFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingWorkHoursable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingWorkHoursable)
                }
            }
            m.SetWorkingHours(res)
        }
        return nil
    }
    return res
}
// GetIsEmailNotificationEnabled gets the isEmailNotificationEnabled property value. Indicates that a staff member is notified via email when a booking assigned to them is created or changed. The default value is true.
// returns a *bool when successful
func (m *BookingStaffMember) GetIsEmailNotificationEnabled()(*bool) {
    return m.isEmailNotificationEnabled
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The date, time, and time zone when the staff member was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *BookingStaffMember) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetMembershipStatus gets the membershipStatus property value. The membershipStatus property
// returns a *BookingStaffMembershipStatus when successful
func (m *BookingStaffMember) GetMembershipStatus()(*BookingStaffMembershipStatus) {
    return m.membershipStatus
}
// GetRole gets the role property value. The role property
// returns a *BookingStaffRole when successful
func (m *BookingStaffMember) GetRole()(*BookingStaffRole) {
    return m.role
}
// GetTimeZone gets the timeZone property value. The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
// returns a *string when successful
func (m *BookingStaffMember) GetTimeZone()(*string) {
    return m.timeZone
}
// GetUseBusinessHours gets the useBusinessHours property value. True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
// returns a *bool when successful
func (m *BookingStaffMember) GetUseBusinessHours()(*bool) {
    return m.useBusinessHours
}
// GetWorkingHours gets the workingHours property value. The range of hours each day of the week that the staff member is available for booking. By default, they're initialized to be the same as the businessHours property of the business.
// returns a []BookingWorkHoursable when successful
func (m *BookingStaffMember) GetWorkingHours()([]BookingWorkHoursable) {
    return m.workingHours
}
// Serialize serializes information the current object
func (m *BookingStaffMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BookingStaffMemberBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("availabilityIsAffectedByPersonalCalendar", m.GetAvailabilityIsAffectedByPersonalCalendar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEmailNotificationEnabled", m.GetIsEmailNotificationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMembershipStatus() != nil {
        cast := (*m.GetMembershipStatus()).String()
        err = writer.WriteStringValue("membershipStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err = writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useBusinessHours", m.GetUseBusinessHours())
        if err != nil {
            return err
        }
    }
    if m.GetWorkingHours() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkingHours()))
        for i, v := range m.GetWorkingHours() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("workingHours", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAvailabilityIsAffectedByPersonalCalendar sets the availabilityIsAffectedByPersonalCalendar property value. True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
func (m *BookingStaffMember) SetAvailabilityIsAffectedByPersonalCalendar(value *bool)() {
    m.availabilityIsAffectedByPersonalCalendar = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date, time, and time zone when the staff member was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *BookingStaffMember) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The name of the staff member, as displayed to customers. Required.
func (m *BookingStaffMember) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmailAddress sets the emailAddress property value. The email address of the staff member. This email address can be in the same Microsoft 365 tenant as the business, or in a different email domain. This email address can be used if the sendConfirmationsToOwner property is set to true in the scheduling policy of the business. Required.
func (m *BookingStaffMember) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// SetIsEmailNotificationEnabled sets the isEmailNotificationEnabled property value. Indicates that a staff member is notified via email when a booking assigned to them is created or changed. The default value is true.
func (m *BookingStaffMember) SetIsEmailNotificationEnabled(value *bool)() {
    m.isEmailNotificationEnabled = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The date, time, and time zone when the staff member was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *BookingStaffMember) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetMembershipStatus sets the membershipStatus property value. The membershipStatus property
func (m *BookingStaffMember) SetMembershipStatus(value *BookingStaffMembershipStatus)() {
    m.membershipStatus = value
}
// SetRole sets the role property value. The role property
func (m *BookingStaffMember) SetRole(value *BookingStaffRole)() {
    m.role = value
}
// SetTimeZone sets the timeZone property value. The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
func (m *BookingStaffMember) SetTimeZone(value *string)() {
    m.timeZone = value
}
// SetUseBusinessHours sets the useBusinessHours property value. True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
func (m *BookingStaffMember) SetUseBusinessHours(value *bool)() {
    m.useBusinessHours = value
}
// SetWorkingHours sets the workingHours property value. The range of hours each day of the week that the staff member is available for booking. By default, they're initialized to be the same as the businessHours property of the business.
func (m *BookingStaffMember) SetWorkingHours(value []BookingWorkHoursable)() {
    m.workingHours = value
}
type BookingStaffMemberable interface {
    BookingStaffMemberBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityIsAffectedByPersonalCalendar()(*bool)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetEmailAddress()(*string)
    GetIsEmailNotificationEnabled()(*bool)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMembershipStatus()(*BookingStaffMembershipStatus)
    GetRole()(*BookingStaffRole)
    GetTimeZone()(*string)
    GetUseBusinessHours()(*bool)
    GetWorkingHours()([]BookingWorkHoursable)
    SetAvailabilityIsAffectedByPersonalCalendar(value *bool)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetEmailAddress(value *string)()
    SetIsEmailNotificationEnabled(value *bool)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMembershipStatus(value *BookingStaffMembershipStatus)()
    SetRole(value *BookingStaffRole)()
    SetTimeZone(value *string)()
    SetUseBusinessHours(value *bool)()
    SetWorkingHours(value []BookingWorkHoursable)()
}
