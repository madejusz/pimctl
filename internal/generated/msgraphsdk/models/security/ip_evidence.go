package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IpEvidence struct {
    AlertEvidence
    // The two-letter country code according to ISO 3166 format, for example: US, UK, CA, etc.
    countryLetterCode *string
    // The value of the IP Address, can be either in V4 address or V6 address format.
    ipAddress *string
    // The location property
    location GeoLocationable
    // The stream property
    stream Streamable
}
// NewIpEvidence instantiates a new IpEvidence and sets the default values.
func NewIpEvidence()(*IpEvidence) {
    m := &IpEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.ipEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIpEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIpEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpEvidence(), nil
}
// GetCountryLetterCode gets the countryLetterCode property value. The two-letter country code according to ISO 3166 format, for example: US, UK, CA, etc.
// returns a *string when successful
func (m *IpEvidence) GetCountryLetterCode()(*string) {
    return m.countryLetterCode
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IpEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["countryLetterCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryLetterCode(val)
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGeoLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(GeoLocationable))
        }
        return nil
    }
    res["stream"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStreamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStream(val.(Streamable))
        }
        return nil
    }
    return res
}
// GetIpAddress gets the ipAddress property value. The value of the IP Address, can be either in V4 address or V6 address format.
// returns a *string when successful
func (m *IpEvidence) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetLocation gets the location property value. The location property
// returns a GeoLocationable when successful
func (m *IpEvidence) GetLocation()(GeoLocationable) {
    return m.location
}
// GetStream gets the stream property value. The stream property
// returns a Streamable when successful
func (m *IpEvidence) GetStream()(Streamable) {
    return m.stream
}
// Serialize serializes information the current object
func (m *IpEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("countryLetterCode", m.GetCountryLetterCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("stream", m.GetStream())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCountryLetterCode sets the countryLetterCode property value. The two-letter country code according to ISO 3166 format, for example: US, UK, CA, etc.
func (m *IpEvidence) SetCountryLetterCode(value *string)() {
    m.countryLetterCode = value
}
// SetIpAddress sets the ipAddress property value. The value of the IP Address, can be either in V4 address or V6 address format.
func (m *IpEvidence) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetLocation sets the location property value. The location property
func (m *IpEvidence) SetLocation(value GeoLocationable)() {
    m.location = value
}
// SetStream sets the stream property value. The stream property
func (m *IpEvidence) SetStream(value Streamable)() {
    m.stream = value
}
type IpEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCountryLetterCode()(*string)
    GetIpAddress()(*string)
    GetLocation()(GeoLocationable)
    GetStream()(Streamable)
    SetCountryLetterCode(value *string)()
    SetIpAddress(value *string)()
    SetLocation(value GeoLocationable)()
    SetStream(value Streamable)()
}
