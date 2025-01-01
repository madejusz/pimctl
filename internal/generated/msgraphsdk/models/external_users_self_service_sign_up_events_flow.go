package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExternalUsersSelfServiceSignUpEventsFlow struct {
    AuthenticationEventsFlow
    // The configuration for what to invoke when attributes are ready to be collected from the user.
    onAttributeCollection OnAttributeCollectionHandlerable
    // Required. The configuration for what to invoke when authentication methods are ready to be presented to the user. Must have at least one identity provider linked.  Supports $filter (eq). See support for filtering on user flows for syntax information.
    onAuthenticationMethodLoadStart OnAuthenticationMethodLoadStartHandlerable
    // Required. The configuration for what to invoke when an authentication flow is ready to be initiated.
    onInteractiveAuthFlowStart OnInteractiveAuthFlowStartHandlerable
    // The configuration for what to invoke during user creation.
    onUserCreateStart OnUserCreateStartHandlerable
}
// NewExternalUsersSelfServiceSignUpEventsFlow instantiates a new ExternalUsersSelfServiceSignUpEventsFlow and sets the default values.
func NewExternalUsersSelfServiceSignUpEventsFlow()(*ExternalUsersSelfServiceSignUpEventsFlow) {
    m := &ExternalUsersSelfServiceSignUpEventsFlow{
        AuthenticationEventsFlow: *NewAuthenticationEventsFlow(),
    }
    odataTypeValue := "#microsoft.graph.externalUsersSelfServiceSignUpEventsFlow"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateExternalUsersSelfServiceSignUpEventsFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalUsersSelfServiceSignUpEventsFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalUsersSelfServiceSignUpEventsFlow(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalUsersSelfServiceSignUpEventsFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationEventsFlow.GetFieldDeserializers()
    res["onAttributeCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnAttributeCollectionHandlerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnAttributeCollection(val.(OnAttributeCollectionHandlerable))
        }
        return nil
    }
    res["onAuthenticationMethodLoadStart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnAuthenticationMethodLoadStartHandlerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnAuthenticationMethodLoadStart(val.(OnAuthenticationMethodLoadStartHandlerable))
        }
        return nil
    }
    res["onInteractiveAuthFlowStart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnInteractiveAuthFlowStartHandlerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnInteractiveAuthFlowStart(val.(OnInteractiveAuthFlowStartHandlerable))
        }
        return nil
    }
    res["onUserCreateStart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnUserCreateStartHandlerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnUserCreateStart(val.(OnUserCreateStartHandlerable))
        }
        return nil
    }
    return res
}
// GetOnAttributeCollection gets the onAttributeCollection property value. The configuration for what to invoke when attributes are ready to be collected from the user.
// returns a OnAttributeCollectionHandlerable when successful
func (m *ExternalUsersSelfServiceSignUpEventsFlow) GetOnAttributeCollection()(OnAttributeCollectionHandlerable) {
    return m.onAttributeCollection
}
// GetOnAuthenticationMethodLoadStart gets the onAuthenticationMethodLoadStart property value. Required. The configuration for what to invoke when authentication methods are ready to be presented to the user. Must have at least one identity provider linked.  Supports $filter (eq). See support for filtering on user flows for syntax information.
// returns a OnAuthenticationMethodLoadStartHandlerable when successful
func (m *ExternalUsersSelfServiceSignUpEventsFlow) GetOnAuthenticationMethodLoadStart()(OnAuthenticationMethodLoadStartHandlerable) {
    return m.onAuthenticationMethodLoadStart
}
// GetOnInteractiveAuthFlowStart gets the onInteractiveAuthFlowStart property value. Required. The configuration for what to invoke when an authentication flow is ready to be initiated.
// returns a OnInteractiveAuthFlowStartHandlerable when successful
func (m *ExternalUsersSelfServiceSignUpEventsFlow) GetOnInteractiveAuthFlowStart()(OnInteractiveAuthFlowStartHandlerable) {
    return m.onInteractiveAuthFlowStart
}
// GetOnUserCreateStart gets the onUserCreateStart property value. The configuration for what to invoke during user creation.
// returns a OnUserCreateStartHandlerable when successful
func (m *ExternalUsersSelfServiceSignUpEventsFlow) GetOnUserCreateStart()(OnUserCreateStartHandlerable) {
    return m.onUserCreateStart
}
// Serialize serializes information the current object
func (m *ExternalUsersSelfServiceSignUpEventsFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationEventsFlow.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("onAttributeCollection", m.GetOnAttributeCollection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onAuthenticationMethodLoadStart", m.GetOnAuthenticationMethodLoadStart())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onInteractiveAuthFlowStart", m.GetOnInteractiveAuthFlowStart())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onUserCreateStart", m.GetOnUserCreateStart())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOnAttributeCollection sets the onAttributeCollection property value. The configuration for what to invoke when attributes are ready to be collected from the user.
func (m *ExternalUsersSelfServiceSignUpEventsFlow) SetOnAttributeCollection(value OnAttributeCollectionHandlerable)() {
    m.onAttributeCollection = value
}
// SetOnAuthenticationMethodLoadStart sets the onAuthenticationMethodLoadStart property value. Required. The configuration for what to invoke when authentication methods are ready to be presented to the user. Must have at least one identity provider linked.  Supports $filter (eq). See support for filtering on user flows for syntax information.
func (m *ExternalUsersSelfServiceSignUpEventsFlow) SetOnAuthenticationMethodLoadStart(value OnAuthenticationMethodLoadStartHandlerable)() {
    m.onAuthenticationMethodLoadStart = value
}
// SetOnInteractiveAuthFlowStart sets the onInteractiveAuthFlowStart property value. Required. The configuration for what to invoke when an authentication flow is ready to be initiated.
func (m *ExternalUsersSelfServiceSignUpEventsFlow) SetOnInteractiveAuthFlowStart(value OnInteractiveAuthFlowStartHandlerable)() {
    m.onInteractiveAuthFlowStart = value
}
// SetOnUserCreateStart sets the onUserCreateStart property value. The configuration for what to invoke during user creation.
func (m *ExternalUsersSelfServiceSignUpEventsFlow) SetOnUserCreateStart(value OnUserCreateStartHandlerable)() {
    m.onUserCreateStart = value
}
type ExternalUsersSelfServiceSignUpEventsFlowable interface {
    AuthenticationEventsFlowable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOnAttributeCollection()(OnAttributeCollectionHandlerable)
    GetOnAuthenticationMethodLoadStart()(OnAuthenticationMethodLoadStartHandlerable)
    GetOnInteractiveAuthFlowStart()(OnInteractiveAuthFlowStartHandlerable)
    GetOnUserCreateStart()(OnUserCreateStartHandlerable)
    SetOnAttributeCollection(value OnAttributeCollectionHandlerable)()
    SetOnAuthenticationMethodLoadStart(value OnAuthenticationMethodLoadStartHandlerable)()
    SetOnInteractiveAuthFlowStart(value OnInteractiveAuthFlowStartHandlerable)()
    SetOnUserCreateStart(value OnUserCreateStartHandlerable)()
}
