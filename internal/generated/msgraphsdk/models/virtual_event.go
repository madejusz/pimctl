package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualEvent struct {
    Entity
    // The identity information for the creator of the virtual event. Inherited from virtualEvent.
    createdBy CommunicationsIdentitySetable
    // A description of the virtual event.
    description ItemBodyable
    // The display name of the virtual event.
    displayName *string
    // The end time of the virtual event. The timeZone property can be set to any of the time zones currently supported by Windows. For details on how to get all available time zones using PowerShell, see Get-TimeZone.
    endDateTime DateTimeTimeZoneable
    // The virtual event presenters.
    presenters []VirtualEventPresenterable
    // The sessions for the virtual event.
    sessions []VirtualEventSessionable
    // The virtual event settings.
    settings VirtualEventSettingsable
    // Start time of the virtual event. The timeZone property can be set to any of the time zones currently supported by Windows. For details on how to get all available time zones using PowerShell, see Get-TimeZone.
    startDateTime DateTimeTimeZoneable
    // The status of the virtual event. The possible values are: draft, published, canceled, and unknownFutureValue.
    status *VirtualEventStatus
}
// NewVirtualEvent instantiates a new VirtualEvent and sets the default values.
func NewVirtualEvent()(*VirtualEvent) {
    m := &VirtualEvent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.virtualEventTownhall":
                        return NewVirtualEventTownhall(), nil
                    case "#microsoft.graph.virtualEventWebinar":
                        return NewVirtualEventWebinar(), nil
                }
            }
        }
    }
    return NewVirtualEvent(), nil
}
// GetCreatedBy gets the createdBy property value. The identity information for the creator of the virtual event. Inherited from virtualEvent.
// returns a CommunicationsIdentitySetable when successful
func (m *VirtualEvent) GetCreatedBy()(CommunicationsIdentitySetable) {
    return m.createdBy
}
// GetDescription gets the description property value. A description of the virtual event.
// returns a ItemBodyable when successful
func (m *VirtualEvent) GetDescription()(ItemBodyable) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the virtual event.
// returns a *string when successful
func (m *VirtualEvent) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndDateTime gets the endDateTime property value. The end time of the virtual event. The timeZone property can be set to any of the time zones currently supported by Windows. For details on how to get all available time zones using PowerShell, see Get-TimeZone.
// returns a DateTimeTimeZoneable when successful
func (m *VirtualEvent) GetEndDateTime()(DateTimeTimeZoneable) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommunicationsIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(CommunicationsIdentitySetable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val.(ItemBodyable))
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
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["presenters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventPresenterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventPresenterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventPresenterable)
                }
            }
            m.SetPresenters(res)
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
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualEventSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(VirtualEventSettingsable))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVirtualEventStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*VirtualEventStatus))
        }
        return nil
    }
    return res
}
// GetPresenters gets the presenters property value. The virtual event presenters.
// returns a []VirtualEventPresenterable when successful
func (m *VirtualEvent) GetPresenters()([]VirtualEventPresenterable) {
    return m.presenters
}
// GetSessions gets the sessions property value. The sessions for the virtual event.
// returns a []VirtualEventSessionable when successful
func (m *VirtualEvent) GetSessions()([]VirtualEventSessionable) {
    return m.sessions
}
// GetSettings gets the settings property value. The virtual event settings.
// returns a VirtualEventSettingsable when successful
func (m *VirtualEvent) GetSettings()(VirtualEventSettingsable) {
    return m.settings
}
// GetStartDateTime gets the startDateTime property value. Start time of the virtual event. The timeZone property can be set to any of the time zones currently supported by Windows. For details on how to get all available time zones using PowerShell, see Get-TimeZone.
// returns a DateTimeTimeZoneable when successful
func (m *VirtualEvent) GetStartDateTime()(DateTimeTimeZoneable) {
    return m.startDateTime
}
// GetStatus gets the status property value. The status of the virtual event. The possible values are: draft, published, canceled, and unknownFutureValue.
// returns a *VirtualEventStatus when successful
func (m *VirtualEvent) GetStatus()(*VirtualEventStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *VirtualEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("description", m.GetDescription())
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
        err = writer.WriteObjectValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPresenters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPresenters()))
        for i, v := range m.GetPresenters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("presenters", cast)
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
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
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
// SetCreatedBy sets the createdBy property value. The identity information for the creator of the virtual event. Inherited from virtualEvent.
func (m *VirtualEvent) SetCreatedBy(value CommunicationsIdentitySetable)() {
    m.createdBy = value
}
// SetDescription sets the description property value. A description of the virtual event.
func (m *VirtualEvent) SetDescription(value ItemBodyable)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the virtual event.
func (m *VirtualEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndDateTime sets the endDateTime property value. The end time of the virtual event. The timeZone property can be set to any of the time zones currently supported by Windows. For details on how to get all available time zones using PowerShell, see Get-TimeZone.
func (m *VirtualEvent) SetEndDateTime(value DateTimeTimeZoneable)() {
    m.endDateTime = value
}
// SetPresenters sets the presenters property value. The virtual event presenters.
func (m *VirtualEvent) SetPresenters(value []VirtualEventPresenterable)() {
    m.presenters = value
}
// SetSessions sets the sessions property value. The sessions for the virtual event.
func (m *VirtualEvent) SetSessions(value []VirtualEventSessionable)() {
    m.sessions = value
}
// SetSettings sets the settings property value. The virtual event settings.
func (m *VirtualEvent) SetSettings(value VirtualEventSettingsable)() {
    m.settings = value
}
// SetStartDateTime sets the startDateTime property value. Start time of the virtual event. The timeZone property can be set to any of the time zones currently supported by Windows. For details on how to get all available time zones using PowerShell, see Get-TimeZone.
func (m *VirtualEvent) SetStartDateTime(value DateTimeTimeZoneable)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. The status of the virtual event. The possible values are: draft, published, canceled, and unknownFutureValue.
func (m *VirtualEvent) SetStatus(value *VirtualEventStatus)() {
    m.status = value
}
type VirtualEventable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedBy()(CommunicationsIdentitySetable)
    GetDescription()(ItemBodyable)
    GetDisplayName()(*string)
    GetEndDateTime()(DateTimeTimeZoneable)
    GetPresenters()([]VirtualEventPresenterable)
    GetSessions()([]VirtualEventSessionable)
    GetSettings()(VirtualEventSettingsable)
    GetStartDateTime()(DateTimeTimeZoneable)
    GetStatus()(*VirtualEventStatus)
    SetCreatedBy(value CommunicationsIdentitySetable)()
    SetDescription(value ItemBodyable)()
    SetDisplayName(value *string)()
    SetEndDateTime(value DateTimeTimeZoneable)()
    SetPresenters(value []VirtualEventPresenterable)()
    SetSessions(value []VirtualEventSessionable)()
    SetSettings(value VirtualEventSettingsable)()
    SetStartDateTime(value DateTimeTimeZoneable)()
    SetStatus(value *VirtualEventStatus)()
}
