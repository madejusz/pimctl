package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConditionalAccessClientApplications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Service principal IDs excluded from the policy scope.
    excludeServicePrincipals []string
    // Service principal IDs included in the policy scope, or ServicePrincipalsInMyTenant.
    includeServicePrincipals []string
    // The OdataType property
    odataType *string
    // Filter that defines the dynamic-servicePrincipal-syntax rule to include/exclude service principals. A filter can use custom security attributes to include/exclude service principals.
    servicePrincipalFilter ConditionalAccessFilterable
}
// NewConditionalAccessClientApplications instantiates a new ConditionalAccessClientApplications and sets the default values.
func NewConditionalAccessClientApplications()(*ConditionalAccessClientApplications) {
    m := &ConditionalAccessClientApplications{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessClientApplicationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionalAccessClientApplicationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessClientApplications(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConditionalAccessClientApplications) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExcludeServicePrincipals gets the excludeServicePrincipals property value. Service principal IDs excluded from the policy scope.
// returns a []string when successful
func (m *ConditionalAccessClientApplications) GetExcludeServicePrincipals()([]string) {
    return m.excludeServicePrincipals
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConditionalAccessClientApplications) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeServicePrincipals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetExcludeServicePrincipals(res)
        }
        return nil
    }
    res["includeServicePrincipals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetIncludeServicePrincipals(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["servicePrincipalFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalFilter(val.(ConditionalAccessFilterable))
        }
        return nil
    }
    return res
}
// GetIncludeServicePrincipals gets the includeServicePrincipals property value. Service principal IDs included in the policy scope, or ServicePrincipalsInMyTenant.
// returns a []string when successful
func (m *ConditionalAccessClientApplications) GetIncludeServicePrincipals()([]string) {
    return m.includeServicePrincipals
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ConditionalAccessClientApplications) GetOdataType()(*string) {
    return m.odataType
}
// GetServicePrincipalFilter gets the servicePrincipalFilter property value. Filter that defines the dynamic-servicePrincipal-syntax rule to include/exclude service principals. A filter can use custom security attributes to include/exclude service principals.
// returns a ConditionalAccessFilterable when successful
func (m *ConditionalAccessClientApplications) GetServicePrincipalFilter()(ConditionalAccessFilterable) {
    return m.servicePrincipalFilter
}
// Serialize serializes information the current object
func (m *ConditionalAccessClientApplications) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExcludeServicePrincipals() != nil {
        err := writer.WriteCollectionOfStringValues("excludeServicePrincipals", m.GetExcludeServicePrincipals())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeServicePrincipals() != nil {
        err := writer.WriteCollectionOfStringValues("includeServicePrincipals", m.GetIncludeServicePrincipals())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("servicePrincipalFilter", m.GetServicePrincipalFilter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessClientApplications) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExcludeServicePrincipals sets the excludeServicePrincipals property value. Service principal IDs excluded from the policy scope.
func (m *ConditionalAccessClientApplications) SetExcludeServicePrincipals(value []string)() {
    m.excludeServicePrincipals = value
}
// SetIncludeServicePrincipals sets the includeServicePrincipals property value. Service principal IDs included in the policy scope, or ServicePrincipalsInMyTenant.
func (m *ConditionalAccessClientApplications) SetIncludeServicePrincipals(value []string)() {
    m.includeServicePrincipals = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessClientApplications) SetOdataType(value *string)() {
    m.odataType = value
}
// SetServicePrincipalFilter sets the servicePrincipalFilter property value. Filter that defines the dynamic-servicePrincipal-syntax rule to include/exclude service principals. A filter can use custom security attributes to include/exclude service principals.
func (m *ConditionalAccessClientApplications) SetServicePrincipalFilter(value ConditionalAccessFilterable)() {
    m.servicePrincipalFilter = value
}
type ConditionalAccessClientApplicationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcludeServicePrincipals()([]string)
    GetIncludeServicePrincipals()([]string)
    GetOdataType()(*string)
    GetServicePrincipalFilter()(ConditionalAccessFilterable)
    SetExcludeServicePrincipals(value []string)()
    SetIncludeServicePrincipals(value []string)()
    SetOdataType(value *string)()
    SetServicePrincipalFilter(value ConditionalAccessFilterable)()
}
