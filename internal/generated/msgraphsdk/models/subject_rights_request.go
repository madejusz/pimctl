package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubjectRightsRequest struct {
    Entity
    // Collection of users who can approve the request. Currently only supported for requests of type delete.
    approvers []Userable
    // Identity that the request is assigned to.
    assignedTo Identityable
    // The date and time when the request was closed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    closedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Collection of users who can collaborate on the request.
    collaborators []Userable
    // KQL based content query that should be used for search. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
    contentQuery *string
    // Identity information for the entity that created the request.
    createdBy IdentitySetable
    // The date and time when the request was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Information about the data subject.
    dataSubject DataSubjectable
    // The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
    dataSubjectType *DataSubjectType
    // Description for the request.
    description *string
    // The name of the request.
    displayName *string
    // The external ID for the request that is immutable after creation and is used for tracking the request for the external system. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
    externalId *string
    // Collection of history change events.
    history []SubjectRightsRequestHistoryable
    // Include all versions of the documents. By default, the current copies of the documents are returned. If SharePoint sites have versioning enabled, including all versions includes the historical copies of the documents. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
    includeAllVersions *bool
    // Include content authored by the data subject. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
    includeAuthoredContent *bool
    // Insight about the request.
    insight SubjectRightsRequestDetailable
    // The date and time when the request is internally due. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    internalDueDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Identity information for the entity that last modified the request.
    lastModifiedBy IdentitySetable
    // The date and time when the request was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The mailbox locations that should be searched. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
    mailboxLocations SubjectRightsRequestMailboxLocationable
    // List of notes associated with the request.
    notes []AuthoredNoteable
    // Pause the request after estimate has finished. By default, the data estimate runs and then pauses, allowing you to preview results and then select the option to retrieve data in the UI. You can set this property to false if you want it to perform the estimate and then automatically begin with the retrieval of the content. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
    pauseAfterEstimate *bool
    // List of regulations that this request fulfill.
    regulations []string
    // The SharePoint and OneDrive site locations that should be searched. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
    siteLocations SubjectRightsRequestSiteLocationable
    // Information about the different stages for the request.
    stages []SubjectRightsRequestStageDetailable
    // The status of the request. Possible values are: active, closed, unknownFutureValue.
    status *SubjectRightsRequestStatus
    // Information about the Microsoft Teams team that was created for the request.
    team Teamable
    // The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue.
    typeEscaped *SubjectRightsRequestType
}
// NewSubjectRightsRequest instantiates a new SubjectRightsRequest and sets the default values.
func NewSubjectRightsRequest()(*SubjectRightsRequest) {
    m := &SubjectRightsRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSubjectRightsRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubjectRightsRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubjectRightsRequest(), nil
}
// GetApprovers gets the approvers property value. Collection of users who can approve the request. Currently only supported for requests of type delete.
// returns a []Userable when successful
func (m *SubjectRightsRequest) GetApprovers()([]Userable) {
    return m.approvers
}
// GetAssignedTo gets the assignedTo property value. Identity that the request is assigned to.
// returns a Identityable when successful
func (m *SubjectRightsRequest) GetAssignedTo()(Identityable) {
    return m.assignedTo
}
// GetClosedDateTime gets the closedDateTime property value. The date and time when the request was closed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *SubjectRightsRequest) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.closedDateTime
}
// GetCollaborators gets the collaborators property value. Collection of users who can collaborate on the request.
// returns a []Userable when successful
func (m *SubjectRightsRequest) GetCollaborators()([]Userable) {
    return m.collaborators
}
// GetContentQuery gets the contentQuery property value. KQL based content query that should be used for search. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
// returns a *string when successful
func (m *SubjectRightsRequest) GetContentQuery()(*string) {
    return m.contentQuery
}
// GetCreatedBy gets the createdBy property value. Identity information for the entity that created the request.
// returns a IdentitySetable when successful
func (m *SubjectRightsRequest) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the request was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *SubjectRightsRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDataSubject gets the dataSubject property value. Information about the data subject.
// returns a DataSubjectable when successful
func (m *SubjectRightsRequest) GetDataSubject()(DataSubjectable) {
    return m.dataSubject
}
// GetDataSubjectType gets the dataSubjectType property value. The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
// returns a *DataSubjectType when successful
func (m *SubjectRightsRequest) GetDataSubjectType()(*DataSubjectType) {
    return m.dataSubjectType
}
// GetDescription gets the description property value. Description for the request.
// returns a *string when successful
func (m *SubjectRightsRequest) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of the request.
// returns a *string when successful
func (m *SubjectRightsRequest) GetDisplayName()(*string) {
    return m.displayName
}
// GetExternalId gets the externalId property value. The external ID for the request that is immutable after creation and is used for tracking the request for the external system. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
// returns a *string when successful
func (m *SubjectRightsRequest) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubjectRightsRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetApprovers(res)
        }
        return nil
    }
    res["assignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val.(Identityable))
        }
        return nil
    }
    res["closedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedDateTime(val)
        }
        return nil
    }
    res["collaborators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetCollaborators(res)
        }
        return nil
    }
    res["contentQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentQuery(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
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
    res["dataSubject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDataSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSubject(val.(DataSubjectable))
        }
        return nil
    }
    res["dataSubjectType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataSubjectType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSubjectType(val.(*DataSubjectType))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["history"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectRightsRequestHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectRightsRequestHistoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SubjectRightsRequestHistoryable)
                }
            }
            m.SetHistory(res)
        }
        return nil
    }
    res["includeAllVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeAllVersions(val)
        }
        return nil
    }
    res["includeAuthoredContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeAuthoredContent(val)
        }
        return nil
    }
    res["insight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubjectRightsRequestDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsight(val.(SubjectRightsRequestDetailable))
        }
        return nil
    }
    res["internalDueDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalDueDateTime(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["mailboxLocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubjectRightsRequestMailboxLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailboxLocations(val.(SubjectRightsRequestMailboxLocationable))
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthoredNoteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthoredNoteable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthoredNoteable)
                }
            }
            m.SetNotes(res)
        }
        return nil
    }
    res["pauseAfterEstimate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPauseAfterEstimate(val)
        }
        return nil
    }
    res["regulations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRegulations(res)
        }
        return nil
    }
    res["siteLocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubjectRightsRequestSiteLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteLocations(val.(SubjectRightsRequestSiteLocationable))
        }
        return nil
    }
    res["stages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectRightsRequestStageDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectRightsRequestStageDetailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SubjectRightsRequestStageDetailable)
                }
            }
            m.SetStages(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SubjectRightsRequestStatus))
        }
        return nil
    }
    res["team"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeam(val.(Teamable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*SubjectRightsRequestType))
        }
        return nil
    }
    return res
}
// GetHistory gets the history property value. Collection of history change events.
// returns a []SubjectRightsRequestHistoryable when successful
func (m *SubjectRightsRequest) GetHistory()([]SubjectRightsRequestHistoryable) {
    return m.history
}
// GetIncludeAllVersions gets the includeAllVersions property value. Include all versions of the documents. By default, the current copies of the documents are returned. If SharePoint sites have versioning enabled, including all versions includes the historical copies of the documents. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
// returns a *bool when successful
func (m *SubjectRightsRequest) GetIncludeAllVersions()(*bool) {
    return m.includeAllVersions
}
// GetIncludeAuthoredContent gets the includeAuthoredContent property value. Include content authored by the data subject. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
// returns a *bool when successful
func (m *SubjectRightsRequest) GetIncludeAuthoredContent()(*bool) {
    return m.includeAuthoredContent
}
// GetInsight gets the insight property value. Insight about the request.
// returns a SubjectRightsRequestDetailable when successful
func (m *SubjectRightsRequest) GetInsight()(SubjectRightsRequestDetailable) {
    return m.insight
}
// GetInternalDueDateTime gets the internalDueDateTime property value. The date and time when the request is internally due. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *SubjectRightsRequest) GetInternalDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.internalDueDateTime
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity information for the entity that last modified the request.
// returns a IdentitySetable when successful
func (m *SubjectRightsRequest) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the request was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *SubjectRightsRequest) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetMailboxLocations gets the mailboxLocations property value. The mailbox locations that should be searched. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
// returns a SubjectRightsRequestMailboxLocationable when successful
func (m *SubjectRightsRequest) GetMailboxLocations()(SubjectRightsRequestMailboxLocationable) {
    return m.mailboxLocations
}
// GetNotes gets the notes property value. List of notes associated with the request.
// returns a []AuthoredNoteable when successful
func (m *SubjectRightsRequest) GetNotes()([]AuthoredNoteable) {
    return m.notes
}
// GetPauseAfterEstimate gets the pauseAfterEstimate property value. Pause the request after estimate has finished. By default, the data estimate runs and then pauses, allowing you to preview results and then select the option to retrieve data in the UI. You can set this property to false if you want it to perform the estimate and then automatically begin with the retrieval of the content. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
// returns a *bool when successful
func (m *SubjectRightsRequest) GetPauseAfterEstimate()(*bool) {
    return m.pauseAfterEstimate
}
// GetRegulations gets the regulations property value. List of regulations that this request fulfill.
// returns a []string when successful
func (m *SubjectRightsRequest) GetRegulations()([]string) {
    return m.regulations
}
// GetSiteLocations gets the siteLocations property value. The SharePoint and OneDrive site locations that should be searched. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
// returns a SubjectRightsRequestSiteLocationable when successful
func (m *SubjectRightsRequest) GetSiteLocations()(SubjectRightsRequestSiteLocationable) {
    return m.siteLocations
}
// GetStages gets the stages property value. Information about the different stages for the request.
// returns a []SubjectRightsRequestStageDetailable when successful
func (m *SubjectRightsRequest) GetStages()([]SubjectRightsRequestStageDetailable) {
    return m.stages
}
// GetStatus gets the status property value. The status of the request. Possible values are: active, closed, unknownFutureValue.
// returns a *SubjectRightsRequestStatus when successful
func (m *SubjectRightsRequest) GetStatus()(*SubjectRightsRequestStatus) {
    return m.status
}
// GetTeam gets the team property value. Information about the Microsoft Teams team that was created for the request.
// returns a Teamable when successful
func (m *SubjectRightsRequest) GetTeam()(Teamable) {
    return m.team
}
// GetTypeEscaped gets the type property value. The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue.
// returns a *SubjectRightsRequestType when successful
func (m *SubjectRightsRequest) GetTypeEscaped()(*SubjectRightsRequestType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *SubjectRightsRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApprovers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApprovers()))
        for i, v := range m.GetApprovers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("approvers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("closedDateTime", m.GetClosedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetCollaborators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCollaborators()))
        for i, v := range m.GetCollaborators() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("collaborators", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentQuery", m.GetContentQuery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteObjectValue("dataSubject", m.GetDataSubject())
        if err != nil {
            return err
        }
    }
    if m.GetDataSubjectType() != nil {
        cast := (*m.GetDataSubjectType()).String()
        err = writer.WriteStringValue("dataSubjectType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    if m.GetHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHistory()))
        for i, v := range m.GetHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("history", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeAllVersions", m.GetIncludeAllVersions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeAuthoredContent", m.GetIncludeAuthoredContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("insight", m.GetInsight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("internalDueDateTime", m.GetInternalDueDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mailboxLocations", m.GetMailboxLocations())
        if err != nil {
            return err
        }
    }
    if m.GetNotes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotes()))
        for i, v := range m.GetNotes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("notes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("pauseAfterEstimate", m.GetPauseAfterEstimate())
        if err != nil {
            return err
        }
    }
    if m.GetRegulations() != nil {
        err = writer.WriteCollectionOfStringValues("regulations", m.GetRegulations())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("siteLocations", m.GetSiteLocations())
        if err != nil {
            return err
        }
    }
    if m.GetStages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStages()))
        for i, v := range m.GetStages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("stages", cast)
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
    {
        err = writer.WriteObjectValue("team", m.GetTeam())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApprovers sets the approvers property value. Collection of users who can approve the request. Currently only supported for requests of type delete.
func (m *SubjectRightsRequest) SetApprovers(value []Userable)() {
    m.approvers = value
}
// SetAssignedTo sets the assignedTo property value. Identity that the request is assigned to.
func (m *SubjectRightsRequest) SetAssignedTo(value Identityable)() {
    m.assignedTo = value
}
// SetClosedDateTime sets the closedDateTime property value. The date and time when the request was closed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closedDateTime = value
}
// SetCollaborators sets the collaborators property value. Collection of users who can collaborate on the request.
func (m *SubjectRightsRequest) SetCollaborators(value []Userable)() {
    m.collaborators = value
}
// SetContentQuery sets the contentQuery property value. KQL based content query that should be used for search. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) SetContentQuery(value *string)() {
    m.contentQuery = value
}
// SetCreatedBy sets the createdBy property value. Identity information for the entity that created the request.
func (m *SubjectRightsRequest) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the request was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDataSubject sets the dataSubject property value. Information about the data subject.
func (m *SubjectRightsRequest) SetDataSubject(value DataSubjectable)() {
    m.dataSubject = value
}
// SetDataSubjectType sets the dataSubjectType property value. The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
func (m *SubjectRightsRequest) SetDataSubjectType(value *DataSubjectType)() {
    m.dataSubjectType = value
}
// SetDescription sets the description property value. Description for the request.
func (m *SubjectRightsRequest) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of the request.
func (m *SubjectRightsRequest) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExternalId sets the externalId property value. The external ID for the request that is immutable after creation and is used for tracking the request for the external system. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) SetExternalId(value *string)() {
    m.externalId = value
}
// SetHistory sets the history property value. Collection of history change events.
func (m *SubjectRightsRequest) SetHistory(value []SubjectRightsRequestHistoryable)() {
    m.history = value
}
// SetIncludeAllVersions sets the includeAllVersions property value. Include all versions of the documents. By default, the current copies of the documents are returned. If SharePoint sites have versioning enabled, including all versions includes the historical copies of the documents. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) SetIncludeAllVersions(value *bool)() {
    m.includeAllVersions = value
}
// SetIncludeAuthoredContent sets the includeAuthoredContent property value. Include content authored by the data subject. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) SetIncludeAuthoredContent(value *bool)() {
    m.includeAuthoredContent = value
}
// SetInsight sets the insight property value. Insight about the request.
func (m *SubjectRightsRequest) SetInsight(value SubjectRightsRequestDetailable)() {
    m.insight = value
}
// SetInternalDueDateTime sets the internalDueDateTime property value. The date and time when the request is internally due. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetInternalDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.internalDueDateTime = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity information for the entity that last modified the request.
func (m *SubjectRightsRequest) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the request was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetMailboxLocations sets the mailboxLocations property value. The mailbox locations that should be searched. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) SetMailboxLocations(value SubjectRightsRequestMailboxLocationable)() {
    m.mailboxLocations = value
}
// SetNotes sets the notes property value. List of notes associated with the request.
func (m *SubjectRightsRequest) SetNotes(value []AuthoredNoteable)() {
    m.notes = value
}
// SetPauseAfterEstimate sets the pauseAfterEstimate property value. Pause the request after estimate has finished. By default, the data estimate runs and then pauses, allowing you to preview results and then select the option to retrieve data in the UI. You can set this property to false if you want it to perform the estimate and then automatically begin with the retrieval of the content. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) SetPauseAfterEstimate(value *bool)() {
    m.pauseAfterEstimate = value
}
// SetRegulations sets the regulations property value. List of regulations that this request fulfill.
func (m *SubjectRightsRequest) SetRegulations(value []string)() {
    m.regulations = value
}
// SetSiteLocations sets the siteLocations property value. The SharePoint and OneDrive site locations that should be searched. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) SetSiteLocations(value SubjectRightsRequestSiteLocationable)() {
    m.siteLocations = value
}
// SetStages sets the stages property value. Information about the different stages for the request.
func (m *SubjectRightsRequest) SetStages(value []SubjectRightsRequestStageDetailable)() {
    m.stages = value
}
// SetStatus sets the status property value. The status of the request. Possible values are: active, closed, unknownFutureValue.
func (m *SubjectRightsRequest) SetStatus(value *SubjectRightsRequestStatus)() {
    m.status = value
}
// SetTeam sets the team property value. Information about the Microsoft Teams team that was created for the request.
func (m *SubjectRightsRequest) SetTeam(value Teamable)() {
    m.team = value
}
// SetTypeEscaped sets the type property value. The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue.
func (m *SubjectRightsRequest) SetTypeEscaped(value *SubjectRightsRequestType)() {
    m.typeEscaped = value
}
type SubjectRightsRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprovers()([]Userable)
    GetAssignedTo()(Identityable)
    GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCollaborators()([]Userable)
    GetContentQuery()(*string)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDataSubject()(DataSubjectable)
    GetDataSubjectType()(*DataSubjectType)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetHistory()([]SubjectRightsRequestHistoryable)
    GetIncludeAllVersions()(*bool)
    GetIncludeAuthoredContent()(*bool)
    GetInsight()(SubjectRightsRequestDetailable)
    GetInternalDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMailboxLocations()(SubjectRightsRequestMailboxLocationable)
    GetNotes()([]AuthoredNoteable)
    GetPauseAfterEstimate()(*bool)
    GetRegulations()([]string)
    GetSiteLocations()(SubjectRightsRequestSiteLocationable)
    GetStages()([]SubjectRightsRequestStageDetailable)
    GetStatus()(*SubjectRightsRequestStatus)
    GetTeam()(Teamable)
    GetTypeEscaped()(*SubjectRightsRequestType)
    SetApprovers(value []Userable)()
    SetAssignedTo(value Identityable)()
    SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCollaborators(value []Userable)()
    SetContentQuery(value *string)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDataSubject(value DataSubjectable)()
    SetDataSubjectType(value *DataSubjectType)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetHistory(value []SubjectRightsRequestHistoryable)()
    SetIncludeAllVersions(value *bool)()
    SetIncludeAuthoredContent(value *bool)()
    SetInsight(value SubjectRightsRequestDetailable)()
    SetInternalDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMailboxLocations(value SubjectRightsRequestMailboxLocationable)()
    SetNotes(value []AuthoredNoteable)()
    SetPauseAfterEstimate(value *bool)()
    SetRegulations(value []string)()
    SetSiteLocations(value SubjectRightsRequestSiteLocationable)()
    SetStages(value []SubjectRightsRequestStageDetailable)()
    SetStatus(value *SubjectRightsRequestStatus)()
    SetTeam(value Teamable)()
    SetTypeEscaped(value *SubjectRightsRequestType)()
}
