package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EndUserNotificationDetail struct {
    Entity
    // Email HTML content.
    emailContent *string
    // Indicates whether this language is default.
    isDefaultLangauge *bool
    // Notification language.
    language *string
    // Notification locale.
    locale *string
    // The sentFrom property
    sentFrom EmailIdentityable
    // Mail subject.
    subject *string
}
// NewEndUserNotificationDetail instantiates a new EndUserNotificationDetail and sets the default values.
func NewEndUserNotificationDetail()(*EndUserNotificationDetail) {
    m := &EndUserNotificationDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEndUserNotificationDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEndUserNotificationDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEndUserNotificationDetail(), nil
}
// GetEmailContent gets the emailContent property value. Email HTML content.
// returns a *string when successful
func (m *EndUserNotificationDetail) GetEmailContent()(*string) {
    return m.emailContent
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EndUserNotificationDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emailContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailContent(val)
        }
        return nil
    }
    res["isDefaultLangauge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultLangauge(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
        }
        return nil
    }
    res["sentFrom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentFrom(val.(EmailIdentityable))
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
// GetIsDefaultLangauge gets the isDefaultLangauge property value. Indicates whether this language is default.
// returns a *bool when successful
func (m *EndUserNotificationDetail) GetIsDefaultLangauge()(*bool) {
    return m.isDefaultLangauge
}
// GetLanguage gets the language property value. Notification language.
// returns a *string when successful
func (m *EndUserNotificationDetail) GetLanguage()(*string) {
    return m.language
}
// GetLocale gets the locale property value. Notification locale.
// returns a *string when successful
func (m *EndUserNotificationDetail) GetLocale()(*string) {
    return m.locale
}
// GetSentFrom gets the sentFrom property value. The sentFrom property
// returns a EmailIdentityable when successful
func (m *EndUserNotificationDetail) GetSentFrom()(EmailIdentityable) {
    return m.sentFrom
}
// GetSubject gets the subject property value. Mail subject.
// returns a *string when successful
func (m *EndUserNotificationDetail) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *EndUserNotificationDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("emailContent", m.GetEmailContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultLangauge", m.GetIsDefaultLangauge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("locale", m.GetLocale())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sentFrom", m.GetSentFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmailContent sets the emailContent property value. Email HTML content.
func (m *EndUserNotificationDetail) SetEmailContent(value *string)() {
    m.emailContent = value
}
// SetIsDefaultLangauge sets the isDefaultLangauge property value. Indicates whether this language is default.
func (m *EndUserNotificationDetail) SetIsDefaultLangauge(value *bool)() {
    m.isDefaultLangauge = value
}
// SetLanguage sets the language property value. Notification language.
func (m *EndUserNotificationDetail) SetLanguage(value *string)() {
    m.language = value
}
// SetLocale sets the locale property value. Notification locale.
func (m *EndUserNotificationDetail) SetLocale(value *string)() {
    m.locale = value
}
// SetSentFrom sets the sentFrom property value. The sentFrom property
func (m *EndUserNotificationDetail) SetSentFrom(value EmailIdentityable)() {
    m.sentFrom = value
}
// SetSubject sets the subject property value. Mail subject.
func (m *EndUserNotificationDetail) SetSubject(value *string)() {
    m.subject = value
}
type EndUserNotificationDetailable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailContent()(*string)
    GetIsDefaultLangauge()(*bool)
    GetLanguage()(*string)
    GetLocale()(*string)
    GetSentFrom()(EmailIdentityable)
    GetSubject()(*string)
    SetEmailContent(value *string)()
    SetIsDefaultLangauge(value *bool)()
    SetLanguage(value *string)()
    SetLocale(value *string)()
    SetSentFrom(value EmailIdentityable)()
    SetSubject(value *string)()
}
