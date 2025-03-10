package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LearningProvider struct {
    Entity
    // The display name that appears in Viva Learning. Required.
    displayName *string
    // Indicates whether a provider can ingest learning course activity records. The default value is false. Set to true to make learningCourseActivities available for this provider.
    isCourseActivitySyncEnabled *bool
    // Learning catalog items for the provider.
    learningContents []LearningContentable
    // The learningCourseActivities property
    learningCourseActivities []LearningCourseActivityable
    // Authentication URL to access the courses for the provider. Optional.
    loginWebUrl *string
    // The long logo URL for the dark mode that needs to be a publicly accessible image. This image would be saved to the blob storage of Viva Learning for rendering within the Viva Learning app. Required.
    longLogoWebUrlForDarkTheme *string
    // The long logo URL for the light mode that needs to be a publicly accessible image. This image would be saved to the blob storage of Viva Learning for rendering within the Viva Learning app. Required.
    longLogoWebUrlForLightTheme *string
    // The square logo URL for the dark mode that needs to be a publicly accessible image. This image would be saved to the blob storage of Viva Learning for rendering within the Viva Learning app. Required.
    squareLogoWebUrlForDarkTheme *string
    // The square logo URL for the light mode that needs to be a publicly accessible image. This image would be saved to the blob storage of Viva Learning for rendering within the Viva Learning app. Required.
    squareLogoWebUrlForLightTheme *string
}
// NewLearningProvider instantiates a new LearningProvider and sets the default values.
func NewLearningProvider()(*LearningProvider) {
    m := &LearningProvider{
        Entity: *NewEntity(),
    }
    return m
}
// CreateLearningProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLearningProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLearningProvider(), nil
}
// GetDisplayName gets the displayName property value. The display name that appears in Viva Learning. Required.
// returns a *string when successful
func (m *LearningProvider) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LearningProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["isCourseActivitySyncEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCourseActivitySyncEnabled(val)
        }
        return nil
    }
    res["learningContents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLearningContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LearningContentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LearningContentable)
                }
            }
            m.SetLearningContents(res)
        }
        return nil
    }
    res["learningCourseActivities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLearningCourseActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LearningCourseActivityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LearningCourseActivityable)
                }
            }
            m.SetLearningCourseActivities(res)
        }
        return nil
    }
    res["loginWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginWebUrl(val)
        }
        return nil
    }
    res["longLogoWebUrlForDarkTheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongLogoWebUrlForDarkTheme(val)
        }
        return nil
    }
    res["longLogoWebUrlForLightTheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongLogoWebUrlForLightTheme(val)
        }
        return nil
    }
    res["squareLogoWebUrlForDarkTheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogoWebUrlForDarkTheme(val)
        }
        return nil
    }
    res["squareLogoWebUrlForLightTheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogoWebUrlForLightTheme(val)
        }
        return nil
    }
    return res
}
// GetIsCourseActivitySyncEnabled gets the isCourseActivitySyncEnabled property value. Indicates whether a provider can ingest learning course activity records. The default value is false. Set to true to make learningCourseActivities available for this provider.
// returns a *bool when successful
func (m *LearningProvider) GetIsCourseActivitySyncEnabled()(*bool) {
    return m.isCourseActivitySyncEnabled
}
// GetLearningContents gets the learningContents property value. Learning catalog items for the provider.
// returns a []LearningContentable when successful
func (m *LearningProvider) GetLearningContents()([]LearningContentable) {
    return m.learningContents
}
// GetLearningCourseActivities gets the learningCourseActivities property value. The learningCourseActivities property
// returns a []LearningCourseActivityable when successful
func (m *LearningProvider) GetLearningCourseActivities()([]LearningCourseActivityable) {
    return m.learningCourseActivities
}
// GetLoginWebUrl gets the loginWebUrl property value. Authentication URL to access the courses for the provider. Optional.
// returns a *string when successful
func (m *LearningProvider) GetLoginWebUrl()(*string) {
    return m.loginWebUrl
}
// GetLongLogoWebUrlForDarkTheme gets the longLogoWebUrlForDarkTheme property value. The long logo URL for the dark mode that needs to be a publicly accessible image. This image would be saved to the blob storage of Viva Learning for rendering within the Viva Learning app. Required.
// returns a *string when successful
func (m *LearningProvider) GetLongLogoWebUrlForDarkTheme()(*string) {
    return m.longLogoWebUrlForDarkTheme
}
// GetLongLogoWebUrlForLightTheme gets the longLogoWebUrlForLightTheme property value. The long logo URL for the light mode that needs to be a publicly accessible image. This image would be saved to the blob storage of Viva Learning for rendering within the Viva Learning app. Required.
// returns a *string when successful
func (m *LearningProvider) GetLongLogoWebUrlForLightTheme()(*string) {
    return m.longLogoWebUrlForLightTheme
}
// GetSquareLogoWebUrlForDarkTheme gets the squareLogoWebUrlForDarkTheme property value. The square logo URL for the dark mode that needs to be a publicly accessible image. This image would be saved to the blob storage of Viva Learning for rendering within the Viva Learning app. Required.
// returns a *string when successful
func (m *LearningProvider) GetSquareLogoWebUrlForDarkTheme()(*string) {
    return m.squareLogoWebUrlForDarkTheme
}
// GetSquareLogoWebUrlForLightTheme gets the squareLogoWebUrlForLightTheme property value. The square logo URL for the light mode that needs to be a publicly accessible image. This image would be saved to the blob storage of Viva Learning for rendering within the Viva Learning app. Required.
// returns a *string when successful
func (m *LearningProvider) GetSquareLogoWebUrlForLightTheme()(*string) {
    return m.squareLogoWebUrlForLightTheme
}
// Serialize serializes information the current object
func (m *LearningProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCourseActivitySyncEnabled", m.GetIsCourseActivitySyncEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetLearningContents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLearningContents()))
        for i, v := range m.GetLearningContents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("learningContents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLearningCourseActivities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLearningCourseActivities()))
        for i, v := range m.GetLearningCourseActivities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("learningCourseActivities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("loginWebUrl", m.GetLoginWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("longLogoWebUrlForDarkTheme", m.GetLongLogoWebUrlForDarkTheme())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("longLogoWebUrlForLightTheme", m.GetLongLogoWebUrlForLightTheme())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("squareLogoWebUrlForDarkTheme", m.GetSquareLogoWebUrlForDarkTheme())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("squareLogoWebUrlForLightTheme", m.GetSquareLogoWebUrlForLightTheme())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name that appears in Viva Learning. Required.
func (m *LearningProvider) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsCourseActivitySyncEnabled sets the isCourseActivitySyncEnabled property value. Indicates whether a provider can ingest learning course activity records. The default value is false. Set to true to make learningCourseActivities available for this provider.
func (m *LearningProvider) SetIsCourseActivitySyncEnabled(value *bool)() {
    m.isCourseActivitySyncEnabled = value
}
// SetLearningContents sets the learningContents property value. Learning catalog items for the provider.
func (m *LearningProvider) SetLearningContents(value []LearningContentable)() {
    m.learningContents = value
}
// SetLearningCourseActivities sets the learningCourseActivities property value. The learningCourseActivities property
func (m *LearningProvider) SetLearningCourseActivities(value []LearningCourseActivityable)() {
    m.learningCourseActivities = value
}
// SetLoginWebUrl sets the loginWebUrl property value. Authentication URL to access the courses for the provider. Optional.
func (m *LearningProvider) SetLoginWebUrl(value *string)() {
    m.loginWebUrl = value
}
// SetLongLogoWebUrlForDarkTheme sets the longLogoWebUrlForDarkTheme property value. The long logo URL for the dark mode that needs to be a publicly accessible image. This image would be saved to the blob storage of Viva Learning for rendering within the Viva Learning app. Required.
func (m *LearningProvider) SetLongLogoWebUrlForDarkTheme(value *string)() {
    m.longLogoWebUrlForDarkTheme = value
}
// SetLongLogoWebUrlForLightTheme sets the longLogoWebUrlForLightTheme property value. The long logo URL for the light mode that needs to be a publicly accessible image. This image would be saved to the blob storage of Viva Learning for rendering within the Viva Learning app. Required.
func (m *LearningProvider) SetLongLogoWebUrlForLightTheme(value *string)() {
    m.longLogoWebUrlForLightTheme = value
}
// SetSquareLogoWebUrlForDarkTheme sets the squareLogoWebUrlForDarkTheme property value. The square logo URL for the dark mode that needs to be a publicly accessible image. This image would be saved to the blob storage of Viva Learning for rendering within the Viva Learning app. Required.
func (m *LearningProvider) SetSquareLogoWebUrlForDarkTheme(value *string)() {
    m.squareLogoWebUrlForDarkTheme = value
}
// SetSquareLogoWebUrlForLightTheme sets the squareLogoWebUrlForLightTheme property value. The square logo URL for the light mode that needs to be a publicly accessible image. This image would be saved to the blob storage of Viva Learning for rendering within the Viva Learning app. Required.
func (m *LearningProvider) SetSquareLogoWebUrlForLightTheme(value *string)() {
    m.squareLogoWebUrlForLightTheme = value
}
type LearningProviderable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetIsCourseActivitySyncEnabled()(*bool)
    GetLearningContents()([]LearningContentable)
    GetLearningCourseActivities()([]LearningCourseActivityable)
    GetLoginWebUrl()(*string)
    GetLongLogoWebUrlForDarkTheme()(*string)
    GetLongLogoWebUrlForLightTheme()(*string)
    GetSquareLogoWebUrlForDarkTheme()(*string)
    GetSquareLogoWebUrlForLightTheme()(*string)
    SetDisplayName(value *string)()
    SetIsCourseActivitySyncEnabled(value *bool)()
    SetLearningContents(value []LearningContentable)()
    SetLearningCourseActivities(value []LearningCourseActivityable)()
    SetLoginWebUrl(value *string)()
    SetLongLogoWebUrlForDarkTheme(value *string)()
    SetLongLogoWebUrlForLightTheme(value *string)()
    SetSquareLogoWebUrlForDarkTheme(value *string)()
    SetSquareLogoWebUrlForLightTheme(value *string)()
}
