package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IpAddress struct {
    Host
    // The details about the autonomous system to which this IP address belongs.
    autonomousSystem AutonomousSystemable
    // The country/region for this IP address.
    countryOrRegion *string
    // The hosting company listed for this host.
    hostingProvider *string
    // The block of IP addresses this IP address belongs to.
    netblock *string
}
// NewIpAddress instantiates a new IpAddress and sets the default values.
func NewIpAddress()(*IpAddress) {
    m := &IpAddress{
        Host: *NewHost(),
    }
    odataTypeValue := "#microsoft.graph.security.ipAddress"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIpAddressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIpAddressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpAddress(), nil
}
// GetAutonomousSystem gets the autonomousSystem property value. The details about the autonomous system to which this IP address belongs.
// returns a AutonomousSystemable when successful
func (m *IpAddress) GetAutonomousSystem()(AutonomousSystemable) {
    return m.autonomousSystem
}
// GetCountryOrRegion gets the countryOrRegion property value. The country/region for this IP address.
// returns a *string when successful
func (m *IpAddress) GetCountryOrRegion()(*string) {
    return m.countryOrRegion
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IpAddress) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Host.GetFieldDeserializers()
    res["autonomousSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAutonomousSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutonomousSystem(val.(AutonomousSystemable))
        }
        return nil
    }
    res["countryOrRegion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryOrRegion(val)
        }
        return nil
    }
    res["hostingProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostingProvider(val)
        }
        return nil
    }
    res["netblock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetblock(val)
        }
        return nil
    }
    return res
}
// GetHostingProvider gets the hostingProvider property value. The hosting company listed for this host.
// returns a *string when successful
func (m *IpAddress) GetHostingProvider()(*string) {
    return m.hostingProvider
}
// GetNetblock gets the netblock property value. The block of IP addresses this IP address belongs to.
// returns a *string when successful
func (m *IpAddress) GetNetblock()(*string) {
    return m.netblock
}
// Serialize serializes information the current object
func (m *IpAddress) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Host.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("autonomousSystem", m.GetAutonomousSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countryOrRegion", m.GetCountryOrRegion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hostingProvider", m.GetHostingProvider())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("netblock", m.GetNetblock())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAutonomousSystem sets the autonomousSystem property value. The details about the autonomous system to which this IP address belongs.
func (m *IpAddress) SetAutonomousSystem(value AutonomousSystemable)() {
    m.autonomousSystem = value
}
// SetCountryOrRegion sets the countryOrRegion property value. The country/region for this IP address.
func (m *IpAddress) SetCountryOrRegion(value *string)() {
    m.countryOrRegion = value
}
// SetHostingProvider sets the hostingProvider property value. The hosting company listed for this host.
func (m *IpAddress) SetHostingProvider(value *string)() {
    m.hostingProvider = value
}
// SetNetblock sets the netblock property value. The block of IP addresses this IP address belongs to.
func (m *IpAddress) SetNetblock(value *string)() {
    m.netblock = value
}
type IpAddressable interface {
    Hostable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutonomousSystem()(AutonomousSystemable)
    GetCountryOrRegion()(*string)
    GetHostingProvider()(*string)
    GetNetblock()(*string)
    SetAutonomousSystem(value AutonomousSystemable)()
    SetCountryOrRegion(value *string)()
    SetHostingProvider(value *string)()
    SetNetblock(value *string)()
}
