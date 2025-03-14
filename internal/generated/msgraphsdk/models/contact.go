package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Contact struct {
    OutlookItem
    // The name of the contact's assistant.
    assistantName *string
    // The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    birthday *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The contact's business address.
    businessAddress PhysicalAddressable
    // The business home page of the contact.
    businessHomePage *string
    // The contact's business phone numbers.
    businessPhones []string
    // The names of the contact's children.
    children []string
    // The name of the contact's company.
    companyName *string
    // The contact's department.
    department *string
    // The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
    displayName *string
    // The contact's email addresses.
    emailAddresses []EmailAddressable
    // The collection of open extensions defined for the contact. Read-only. Nullable.
    extensions []Extensionable
    // The name the contact is filed under.
    fileAs *string
    // The contact's suffix.
    generation *string
    // The contact's given name.
    givenName *string
    // The contact's home address.
    homeAddress PhysicalAddressable
    // The contact's home phone numbers.
    homePhones []string
    // The contact's instant messaging (IM) addresses.
    imAddresses []string
    // The contact's initials.
    initials *string
    // The contact’s job title.
    jobTitle *string
    // The name of the contact's manager.
    manager *string
    // The contact's middle name.
    middleName *string
    // The contact's mobile phone number.
    mobilePhone *string
    // The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The contact's nickname.
    nickName *string
    // The location of the contact's office.
    officeLocation *string
    // Other addresses for the contact.
    otherAddress PhysicalAddressable
    // The ID of the contact's parent folder.
    parentFolderId *string
    // The user's notes about the contact.
    personalNotes *string
    // Optional contact picture. You can get or set a photo for a contact.
    photo ProfilePhotoable
    // The contact's profession.
    profession *string
    // The collection of single-value extended properties defined for the contact. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
    // The name of the contact's spouse/partner.
    spouseName *string
    // The contact's surname.
    surname *string
    // The contact's title.
    title *string
    // The phonetic Japanese company name of the contact.
    yomiCompanyName *string
    // The phonetic Japanese given name (first name) of the contact.
    yomiGivenName *string
    // The phonetic Japanese surname (last name)  of the contact.
    yomiSurname *string
}
// NewContact instantiates a new Contact and sets the default values.
func NewContact()(*Contact) {
    m := &Contact{
        OutlookItem: *NewOutlookItem(),
    }
    odataTypeValue := "#microsoft.graph.contact"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateContactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContact(), nil
}
// GetAssistantName gets the assistantName property value. The name of the contact's assistant.
// returns a *string when successful
func (m *Contact) GetAssistantName()(*string) {
    return m.assistantName
}
// GetBirthday gets the birthday property value. The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *Contact) GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.birthday
}
// GetBusinessAddress gets the businessAddress property value. The contact's business address.
// returns a PhysicalAddressable when successful
func (m *Contact) GetBusinessAddress()(PhysicalAddressable) {
    return m.businessAddress
}
// GetBusinessHomePage gets the businessHomePage property value. The business home page of the contact.
// returns a *string when successful
func (m *Contact) GetBusinessHomePage()(*string) {
    return m.businessHomePage
}
// GetBusinessPhones gets the businessPhones property value. The contact's business phone numbers.
// returns a []string when successful
func (m *Contact) GetBusinessPhones()([]string) {
    return m.businessPhones
}
// GetChildren gets the children property value. The names of the contact's children.
// returns a []string when successful
func (m *Contact) GetChildren()([]string) {
    return m.children
}
// GetCompanyName gets the companyName property value. The name of the contact's company.
// returns a *string when successful
func (m *Contact) GetCompanyName()(*string) {
    return m.companyName
}
// GetDepartment gets the department property value. The contact's department.
// returns a *string when successful
func (m *Contact) GetDepartment()(*string) {
    return m.department
}
// GetDisplayName gets the displayName property value. The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
// returns a *string when successful
func (m *Contact) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmailAddresses gets the emailAddresses property value. The contact's email addresses.
// returns a []EmailAddressable when successful
func (m *Contact) GetEmailAddresses()([]EmailAddressable) {
    return m.emailAddresses
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the contact. Read-only. Nullable.
// returns a []Extensionable when successful
func (m *Contact) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Contact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["assistantName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssistantName(val)
        }
        return nil
    }
    res["birthday"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBirthday(val)
        }
        return nil
    }
    res["businessAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["businessHomePage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessHomePage(val)
        }
        return nil
    }
    res["businessPhones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetBusinessPhones(res)
        }
        return nil
    }
    res["children"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetChildren(res)
        }
        return nil
    }
    res["companyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyName(val)
        }
        return nil
    }
    res["department"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val)
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
    res["emailAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmailAddressable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EmailAddressable)
                }
            }
            m.SetEmailAddresses(res)
        }
        return nil
    }
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Extensionable)
                }
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["fileAs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileAs(val)
        }
        return nil
    }
    res["generation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneration(val)
        }
        return nil
    }
    res["givenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["homeAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["homePhones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetHomePhones(res)
        }
        return nil
    }
    res["imAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetImAddresses(res)
        }
        return nil
    }
    res["initials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitials(val)
        }
        return nil
    }
    res["jobTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
        }
        return nil
    }
    res["manager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val)
        }
        return nil
    }
    res["middleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
        }
        return nil
    }
    res["mobilePhone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobilePhone(val)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MultiValueLegacyExtendedPropertyable)
                }
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["nickName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNickName(val)
        }
        return nil
    }
    res["officeLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeLocation(val)
        }
        return nil
    }
    res["otherAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOtherAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["parentFolderId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentFolderId(val)
        }
        return nil
    }
    res["personalNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalNotes(val)
        }
        return nil
    }
    res["photo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfilePhotoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val.(ProfilePhotoable))
        }
        return nil
    }
    res["profession"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfession(val)
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SingleValueLegacyExtendedPropertyable)
                }
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["spouseName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpouseName(val)
        }
        return nil
    }
    res["surname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["yomiCompanyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYomiCompanyName(val)
        }
        return nil
    }
    res["yomiGivenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYomiGivenName(val)
        }
        return nil
    }
    res["yomiSurname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYomiSurname(val)
        }
        return nil
    }
    return res
}
// GetFileAs gets the fileAs property value. The name the contact is filed under.
// returns a *string when successful
func (m *Contact) GetFileAs()(*string) {
    return m.fileAs
}
// GetGeneration gets the generation property value. The contact's suffix.
// returns a *string when successful
func (m *Contact) GetGeneration()(*string) {
    return m.generation
}
// GetGivenName gets the givenName property value. The contact's given name.
// returns a *string when successful
func (m *Contact) GetGivenName()(*string) {
    return m.givenName
}
// GetHomeAddress gets the homeAddress property value. The contact's home address.
// returns a PhysicalAddressable when successful
func (m *Contact) GetHomeAddress()(PhysicalAddressable) {
    return m.homeAddress
}
// GetHomePhones gets the homePhones property value. The contact's home phone numbers.
// returns a []string when successful
func (m *Contact) GetHomePhones()([]string) {
    return m.homePhones
}
// GetImAddresses gets the imAddresses property value. The contact's instant messaging (IM) addresses.
// returns a []string when successful
func (m *Contact) GetImAddresses()([]string) {
    return m.imAddresses
}
// GetInitials gets the initials property value. The contact's initials.
// returns a *string when successful
func (m *Contact) GetInitials()(*string) {
    return m.initials
}
// GetJobTitle gets the jobTitle property value. The contact’s job title.
// returns a *string when successful
func (m *Contact) GetJobTitle()(*string) {
    return m.jobTitle
}
// GetManager gets the manager property value. The name of the contact's manager.
// returns a *string when successful
func (m *Contact) GetManager()(*string) {
    return m.manager
}
// GetMiddleName gets the middleName property value. The contact's middle name.
// returns a *string when successful
func (m *Contact) GetMiddleName()(*string) {
    return m.middleName
}
// GetMobilePhone gets the mobilePhone property value. The contact's mobile phone number.
// returns a *string when successful
func (m *Contact) GetMobilePhone()(*string) {
    return m.mobilePhone
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
// returns a []MultiValueLegacyExtendedPropertyable when successful
func (m *Contact) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    return m.multiValueExtendedProperties
}
// GetNickName gets the nickName property value. The contact's nickname.
// returns a *string when successful
func (m *Contact) GetNickName()(*string) {
    return m.nickName
}
// GetOfficeLocation gets the officeLocation property value. The location of the contact's office.
// returns a *string when successful
func (m *Contact) GetOfficeLocation()(*string) {
    return m.officeLocation
}
// GetOtherAddress gets the otherAddress property value. Other addresses for the contact.
// returns a PhysicalAddressable when successful
func (m *Contact) GetOtherAddress()(PhysicalAddressable) {
    return m.otherAddress
}
// GetParentFolderId gets the parentFolderId property value. The ID of the contact's parent folder.
// returns a *string when successful
func (m *Contact) GetParentFolderId()(*string) {
    return m.parentFolderId
}
// GetPersonalNotes gets the personalNotes property value. The user's notes about the contact.
// returns a *string when successful
func (m *Contact) GetPersonalNotes()(*string) {
    return m.personalNotes
}
// GetPhoto gets the photo property value. Optional contact picture. You can get or set a photo for a contact.
// returns a ProfilePhotoable when successful
func (m *Contact) GetPhoto()(ProfilePhotoable) {
    return m.photo
}
// GetProfession gets the profession property value. The contact's profession.
// returns a *string when successful
func (m *Contact) GetProfession()(*string) {
    return m.profession
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contact. Read-only. Nullable.
// returns a []SingleValueLegacyExtendedPropertyable when successful
func (m *Contact) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    return m.singleValueExtendedProperties
}
// GetSpouseName gets the spouseName property value. The name of the contact's spouse/partner.
// returns a *string when successful
func (m *Contact) GetSpouseName()(*string) {
    return m.spouseName
}
// GetSurname gets the surname property value. The contact's surname.
// returns a *string when successful
func (m *Contact) GetSurname()(*string) {
    return m.surname
}
// GetTitle gets the title property value. The contact's title.
// returns a *string when successful
func (m *Contact) GetTitle()(*string) {
    return m.title
}
// GetYomiCompanyName gets the yomiCompanyName property value. The phonetic Japanese company name of the contact.
// returns a *string when successful
func (m *Contact) GetYomiCompanyName()(*string) {
    return m.yomiCompanyName
}
// GetYomiGivenName gets the yomiGivenName property value. The phonetic Japanese given name (first name) of the contact.
// returns a *string when successful
func (m *Contact) GetYomiGivenName()(*string) {
    return m.yomiGivenName
}
// GetYomiSurname gets the yomiSurname property value. The phonetic Japanese surname (last name)  of the contact.
// returns a *string when successful
func (m *Contact) GetYomiSurname()(*string) {
    return m.yomiSurname
}
// Serialize serializes information the current object
func (m *Contact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assistantName", m.GetAssistantName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("birthday", m.GetBirthday())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("businessAddress", m.GetBusinessAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("businessHomePage", m.GetBusinessHomePage())
        if err != nil {
            return err
        }
    }
    if m.GetBusinessPhones() != nil {
        err = writer.WriteCollectionOfStringValues("businessPhones", m.GetBusinessPhones())
        if err != nil {
            return err
        }
    }
    if m.GetChildren() != nil {
        err = writer.WriteCollectionOfStringValues("children", m.GetChildren())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyName", m.GetCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("department", m.GetDepartment())
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
    if m.GetEmailAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmailAddresses()))
        for i, v := range m.GetEmailAddresses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("emailAddresses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileAs", m.GetFileAs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("generation", m.GetGeneration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("givenName", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("homeAddress", m.GetHomeAddress())
        if err != nil {
            return err
        }
    }
    if m.GetHomePhones() != nil {
        err = writer.WriteCollectionOfStringValues("homePhones", m.GetHomePhones())
        if err != nil {
            return err
        }
    }
    if m.GetImAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("imAddresses", m.GetImAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initials", m.GetInitials())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jobTitle", m.GetJobTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("middleName", m.GetMiddleName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobilePhone", m.GetMobilePhone())
        if err != nil {
            return err
        }
    }
    if m.GetMultiValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nickName", m.GetNickName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("otherAddress", m.GetOtherAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("personalNotes", m.GetPersonalNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("photo", m.GetPhoto())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profession", m.GetProfession())
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("spouseName", m.GetSpouseName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("surname", m.GetSurname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiCompanyName", m.GetYomiCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiGivenName", m.GetYomiGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiSurname", m.GetYomiSurname())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssistantName sets the assistantName property value. The name of the contact's assistant.
func (m *Contact) SetAssistantName(value *string)() {
    m.assistantName = value
}
// SetBirthday sets the birthday property value. The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Contact) SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.birthday = value
}
// SetBusinessAddress sets the businessAddress property value. The contact's business address.
func (m *Contact) SetBusinessAddress(value PhysicalAddressable)() {
    m.businessAddress = value
}
// SetBusinessHomePage sets the businessHomePage property value. The business home page of the contact.
func (m *Contact) SetBusinessHomePage(value *string)() {
    m.businessHomePage = value
}
// SetBusinessPhones sets the businessPhones property value. The contact's business phone numbers.
func (m *Contact) SetBusinessPhones(value []string)() {
    m.businessPhones = value
}
// SetChildren sets the children property value. The names of the contact's children.
func (m *Contact) SetChildren(value []string)() {
    m.children = value
}
// SetCompanyName sets the companyName property value. The name of the contact's company.
func (m *Contact) SetCompanyName(value *string)() {
    m.companyName = value
}
// SetDepartment sets the department property value. The contact's department.
func (m *Contact) SetDepartment(value *string)() {
    m.department = value
}
// SetDisplayName sets the displayName property value. The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
func (m *Contact) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmailAddresses sets the emailAddresses property value. The contact's email addresses.
func (m *Contact) SetEmailAddresses(value []EmailAddressable)() {
    m.emailAddresses = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the contact. Read-only. Nullable.
func (m *Contact) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetFileAs sets the fileAs property value. The name the contact is filed under.
func (m *Contact) SetFileAs(value *string)() {
    m.fileAs = value
}
// SetGeneration sets the generation property value. The contact's suffix.
func (m *Contact) SetGeneration(value *string)() {
    m.generation = value
}
// SetGivenName sets the givenName property value. The contact's given name.
func (m *Contact) SetGivenName(value *string)() {
    m.givenName = value
}
// SetHomeAddress sets the homeAddress property value. The contact's home address.
func (m *Contact) SetHomeAddress(value PhysicalAddressable)() {
    m.homeAddress = value
}
// SetHomePhones sets the homePhones property value. The contact's home phone numbers.
func (m *Contact) SetHomePhones(value []string)() {
    m.homePhones = value
}
// SetImAddresses sets the imAddresses property value. The contact's instant messaging (IM) addresses.
func (m *Contact) SetImAddresses(value []string)() {
    m.imAddresses = value
}
// SetInitials sets the initials property value. The contact's initials.
func (m *Contact) SetInitials(value *string)() {
    m.initials = value
}
// SetJobTitle sets the jobTitle property value. The contact’s job title.
func (m *Contact) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// SetManager sets the manager property value. The name of the contact's manager.
func (m *Contact) SetManager(value *string)() {
    m.manager = value
}
// SetMiddleName sets the middleName property value. The contact's middle name.
func (m *Contact) SetMiddleName(value *string)() {
    m.middleName = value
}
// SetMobilePhone sets the mobilePhone property value. The contact's mobile phone number.
func (m *Contact) SetMobilePhone(value *string)() {
    m.mobilePhone = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
func (m *Contact) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    m.multiValueExtendedProperties = value
}
// SetNickName sets the nickName property value. The contact's nickname.
func (m *Contact) SetNickName(value *string)() {
    m.nickName = value
}
// SetOfficeLocation sets the officeLocation property value. The location of the contact's office.
func (m *Contact) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// SetOtherAddress sets the otherAddress property value. Other addresses for the contact.
func (m *Contact) SetOtherAddress(value PhysicalAddressable)() {
    m.otherAddress = value
}
// SetParentFolderId sets the parentFolderId property value. The ID of the contact's parent folder.
func (m *Contact) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// SetPersonalNotes sets the personalNotes property value. The user's notes about the contact.
func (m *Contact) SetPersonalNotes(value *string)() {
    m.personalNotes = value
}
// SetPhoto sets the photo property value. Optional contact picture. You can get or set a photo for a contact.
func (m *Contact) SetPhoto(value ProfilePhotoable)() {
    m.photo = value
}
// SetProfession sets the profession property value. The contact's profession.
func (m *Contact) SetProfession(value *string)() {
    m.profession = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contact. Read-only. Nullable.
func (m *Contact) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    m.singleValueExtendedProperties = value
}
// SetSpouseName sets the spouseName property value. The name of the contact's spouse/partner.
func (m *Contact) SetSpouseName(value *string)() {
    m.spouseName = value
}
// SetSurname sets the surname property value. The contact's surname.
func (m *Contact) SetSurname(value *string)() {
    m.surname = value
}
// SetTitle sets the title property value. The contact's title.
func (m *Contact) SetTitle(value *string)() {
    m.title = value
}
// SetYomiCompanyName sets the yomiCompanyName property value. The phonetic Japanese company name of the contact.
func (m *Contact) SetYomiCompanyName(value *string)() {
    m.yomiCompanyName = value
}
// SetYomiGivenName sets the yomiGivenName property value. The phonetic Japanese given name (first name) of the contact.
func (m *Contact) SetYomiGivenName(value *string)() {
    m.yomiGivenName = value
}
// SetYomiSurname sets the yomiSurname property value. The phonetic Japanese surname (last name)  of the contact.
func (m *Contact) SetYomiSurname(value *string)() {
    m.yomiSurname = value
}
type Contactable interface {
    OutlookItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssistantName()(*string)
    GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetBusinessAddress()(PhysicalAddressable)
    GetBusinessHomePage()(*string)
    GetBusinessPhones()([]string)
    GetChildren()([]string)
    GetCompanyName()(*string)
    GetDepartment()(*string)
    GetDisplayName()(*string)
    GetEmailAddresses()([]EmailAddressable)
    GetExtensions()([]Extensionable)
    GetFileAs()(*string)
    GetGeneration()(*string)
    GetGivenName()(*string)
    GetHomeAddress()(PhysicalAddressable)
    GetHomePhones()([]string)
    GetImAddresses()([]string)
    GetInitials()(*string)
    GetJobTitle()(*string)
    GetManager()(*string)
    GetMiddleName()(*string)
    GetMobilePhone()(*string)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetNickName()(*string)
    GetOfficeLocation()(*string)
    GetOtherAddress()(PhysicalAddressable)
    GetParentFolderId()(*string)
    GetPersonalNotes()(*string)
    GetPhoto()(ProfilePhotoable)
    GetProfession()(*string)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    GetSpouseName()(*string)
    GetSurname()(*string)
    GetTitle()(*string)
    GetYomiCompanyName()(*string)
    GetYomiGivenName()(*string)
    GetYomiSurname()(*string)
    SetAssistantName(value *string)()
    SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetBusinessAddress(value PhysicalAddressable)()
    SetBusinessHomePage(value *string)()
    SetBusinessPhones(value []string)()
    SetChildren(value []string)()
    SetCompanyName(value *string)()
    SetDepartment(value *string)()
    SetDisplayName(value *string)()
    SetEmailAddresses(value []EmailAddressable)()
    SetExtensions(value []Extensionable)()
    SetFileAs(value *string)()
    SetGeneration(value *string)()
    SetGivenName(value *string)()
    SetHomeAddress(value PhysicalAddressable)()
    SetHomePhones(value []string)()
    SetImAddresses(value []string)()
    SetInitials(value *string)()
    SetJobTitle(value *string)()
    SetManager(value *string)()
    SetMiddleName(value *string)()
    SetMobilePhone(value *string)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetNickName(value *string)()
    SetOfficeLocation(value *string)()
    SetOtherAddress(value PhysicalAddressable)()
    SetParentFolderId(value *string)()
    SetPersonalNotes(value *string)()
    SetPhoto(value ProfilePhotoable)()
    SetProfession(value *string)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
    SetSpouseName(value *string)()
    SetSurname(value *string)()
    SetTitle(value *string)()
    SetYomiCompanyName(value *string)()
    SetYomiGivenName(value *string)()
    SetYomiSurname(value *string)()
}
