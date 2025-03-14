package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DnsEvidence struct {
    AlertEvidence
    // The dnsServerIp property
    dnsServerIp IpEvidenceable
    // The domainName property
    domainName *string
    // The hostIpAddress property
    hostIpAddress IpEvidenceable
    // The ipAddresses property
    ipAddresses []IpEvidenceable
}
// NewDnsEvidence instantiates a new DnsEvidence and sets the default values.
func NewDnsEvidence()(*DnsEvidence) {
    m := &DnsEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.dnsEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDnsEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDnsEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDnsEvidence(), nil
}
// GetDnsServerIp gets the dnsServerIp property value. The dnsServerIp property
// returns a IpEvidenceable when successful
func (m *DnsEvidence) GetDnsServerIp()(IpEvidenceable) {
    return m.dnsServerIp
}
// GetDomainName gets the domainName property value. The domainName property
// returns a *string when successful
func (m *DnsEvidence) GetDomainName()(*string) {
    return m.domainName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DnsEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["dnsServerIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIpEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDnsServerIp(val.(IpEvidenceable))
        }
        return nil
    }
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["hostIpAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIpEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostIpAddress(val.(IpEvidenceable))
        }
        return nil
    }
    res["ipAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIpEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpEvidenceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IpEvidenceable)
                }
            }
            m.SetIpAddresses(res)
        }
        return nil
    }
    return res
}
// GetHostIpAddress gets the hostIpAddress property value. The hostIpAddress property
// returns a IpEvidenceable when successful
func (m *DnsEvidence) GetHostIpAddress()(IpEvidenceable) {
    return m.hostIpAddress
}
// GetIpAddresses gets the ipAddresses property value. The ipAddresses property
// returns a []IpEvidenceable when successful
func (m *DnsEvidence) GetIpAddresses()([]IpEvidenceable) {
    return m.ipAddresses
}
// Serialize serializes information the current object
func (m *DnsEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("dnsServerIp", m.GetDnsServerIp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("hostIpAddress", m.GetHostIpAddress())
        if err != nil {
            return err
        }
    }
    if m.GetIpAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIpAddresses()))
        for i, v := range m.GetIpAddresses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("ipAddresses", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDnsServerIp sets the dnsServerIp property value. The dnsServerIp property
func (m *DnsEvidence) SetDnsServerIp(value IpEvidenceable)() {
    m.dnsServerIp = value
}
// SetDomainName sets the domainName property value. The domainName property
func (m *DnsEvidence) SetDomainName(value *string)() {
    m.domainName = value
}
// SetHostIpAddress sets the hostIpAddress property value. The hostIpAddress property
func (m *DnsEvidence) SetHostIpAddress(value IpEvidenceable)() {
    m.hostIpAddress = value
}
// SetIpAddresses sets the ipAddresses property value. The ipAddresses property
func (m *DnsEvidence) SetIpAddresses(value []IpEvidenceable)() {
    m.ipAddresses = value
}
type DnsEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDnsServerIp()(IpEvidenceable)
    GetDomainName()(*string)
    GetHostIpAddress()(IpEvidenceable)
    GetIpAddresses()([]IpEvidenceable)
    SetDnsServerIp(value IpEvidenceable)()
    SetDomainName(value *string)()
    SetHostIpAddress(value IpEvidenceable)()
    SetIpAddresses(value []IpEvidenceable)()
}
