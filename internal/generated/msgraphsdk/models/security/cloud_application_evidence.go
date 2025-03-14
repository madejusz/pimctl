package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudApplicationEvidence struct {
    AlertEvidence
    // Unique identifier of the application.
    appId *int64
    // Name of the application.
    displayName *string
    // Identifier of the instance of the Software as a Service (SaaS) application.
    instanceId *int64
    // Name of the instance of the SaaS application.
    instanceName *string
    // The identifier of the SaaS application.
    saasAppId *int64
    // The stream property
    stream Streamable
}
// NewCloudApplicationEvidence instantiates a new CloudApplicationEvidence and sets the default values.
func NewCloudApplicationEvidence()(*CloudApplicationEvidence) {
    m := &CloudApplicationEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.cloudApplicationEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudApplicationEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudApplicationEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudApplicationEvidence(), nil
}
// GetAppId gets the appId property value. Unique identifier of the application.
// returns a *int64 when successful
func (m *CloudApplicationEvidence) GetAppId()(*int64) {
    return m.appId
}
// GetDisplayName gets the displayName property value. Name of the application.
// returns a *string when successful
func (m *CloudApplicationEvidence) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudApplicationEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
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
    res["instanceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceId(val)
        }
        return nil
    }
    res["instanceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceName(val)
        }
        return nil
    }
    res["saasAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSaasAppId(val)
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
// GetInstanceId gets the instanceId property value. Identifier of the instance of the Software as a Service (SaaS) application.
// returns a *int64 when successful
func (m *CloudApplicationEvidence) GetInstanceId()(*int64) {
    return m.instanceId
}
// GetInstanceName gets the instanceName property value. Name of the instance of the SaaS application.
// returns a *string when successful
func (m *CloudApplicationEvidence) GetInstanceName()(*string) {
    return m.instanceName
}
// GetSaasAppId gets the saasAppId property value. The identifier of the SaaS application.
// returns a *int64 when successful
func (m *CloudApplicationEvidence) GetSaasAppId()(*int64) {
    return m.saasAppId
}
// GetStream gets the stream property value. The stream property
// returns a Streamable when successful
func (m *CloudApplicationEvidence) GetStream()(Streamable) {
    return m.stream
}
// Serialize serializes information the current object
func (m *CloudApplicationEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("appId", m.GetAppId())
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
        err = writer.WriteInt64Value("instanceId", m.GetInstanceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("instanceName", m.GetInstanceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("saasAppId", m.GetSaasAppId())
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
// SetAppId sets the appId property value. Unique identifier of the application.
func (m *CloudApplicationEvidence) SetAppId(value *int64)() {
    m.appId = value
}
// SetDisplayName sets the displayName property value. Name of the application.
func (m *CloudApplicationEvidence) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetInstanceId sets the instanceId property value. Identifier of the instance of the Software as a Service (SaaS) application.
func (m *CloudApplicationEvidence) SetInstanceId(value *int64)() {
    m.instanceId = value
}
// SetInstanceName sets the instanceName property value. Name of the instance of the SaaS application.
func (m *CloudApplicationEvidence) SetInstanceName(value *string)() {
    m.instanceName = value
}
// SetSaasAppId sets the saasAppId property value. The identifier of the SaaS application.
func (m *CloudApplicationEvidence) SetSaasAppId(value *int64)() {
    m.saasAppId = value
}
// SetStream sets the stream property value. The stream property
func (m *CloudApplicationEvidence) SetStream(value Streamable)() {
    m.stream = value
}
type CloudApplicationEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*int64)
    GetDisplayName()(*string)
    GetInstanceId()(*int64)
    GetInstanceName()(*string)
    GetSaasAppId()(*int64)
    GetStream()(Streamable)
    SetAppId(value *int64)()
    SetDisplayName(value *string)()
    SetInstanceId(value *int64)()
    SetInstanceName(value *string)()
    SetSaasAppId(value *int64)()
    SetStream(value Streamable)()
}
