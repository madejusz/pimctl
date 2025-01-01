package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EdiscoveryExportOperation struct {
    CaseOperation
    // The description provided for the export.
    description *string
    // Contains the properties for an export file metadata, including downloadUrl, fileName, and size.
    exportFileMetadata []ExportFileMetadataable
    // The options provided for the export. For more information, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement,  tags.
    exportOptions *ExportOptions
    // The options that specify the structure of the export. For more information, see reviewSet: export. Possible values are: none, directory, pst.
    exportStructure *ExportFileStructure
    // The name provided for the export.
    outputName *string
    // Review set from where documents are exported.
    reviewSet EdiscoveryReviewSetable
    // The review set query that is used to filter the documents for export.
    reviewSetQuery EdiscoveryReviewSetQueryable
}
// NewEdiscoveryExportOperation instantiates a new EdiscoveryExportOperation and sets the default values.
func NewEdiscoveryExportOperation()(*EdiscoveryExportOperation) {
    m := &EdiscoveryExportOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEdiscoveryExportOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEdiscoveryExportOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryExportOperation(), nil
}
// GetDescription gets the description property value. The description provided for the export.
// returns a *string when successful
func (m *EdiscoveryExportOperation) GetDescription()(*string) {
    return m.description
}
// GetExportFileMetadata gets the exportFileMetadata property value. Contains the properties for an export file metadata, including downloadUrl, fileName, and size.
// returns a []ExportFileMetadataable when successful
func (m *EdiscoveryExportOperation) GetExportFileMetadata()([]ExportFileMetadataable) {
    return m.exportFileMetadata
}
// GetExportOptions gets the exportOptions property value. The options provided for the export. For more information, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement,  tags.
// returns a *ExportOptions when successful
func (m *EdiscoveryExportOperation) GetExportOptions()(*ExportOptions) {
    return m.exportOptions
}
// GetExportStructure gets the exportStructure property value. The options that specify the structure of the export. For more information, see reviewSet: export. Possible values are: none, directory, pst.
// returns a *ExportFileStructure when successful
func (m *EdiscoveryExportOperation) GetExportStructure()(*ExportFileStructure) {
    return m.exportStructure
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EdiscoveryExportOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
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
    res["exportFileMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExportFileMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExportFileMetadataable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExportFileMetadataable)
                }
            }
            m.SetExportFileMetadata(res)
        }
        return nil
    }
    res["exportOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExportOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportOptions(val.(*ExportOptions))
        }
        return nil
    }
    res["exportStructure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExportFileStructure)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportStructure(val.(*ExportFileStructure))
        }
        return nil
    }
    res["outputName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputName(val)
        }
        return nil
    }
    res["reviewSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryReviewSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSet(val.(EdiscoveryReviewSetable))
        }
        return nil
    }
    res["reviewSetQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryReviewSetQueryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSetQuery(val.(EdiscoveryReviewSetQueryable))
        }
        return nil
    }
    return res
}
// GetOutputName gets the outputName property value. The name provided for the export.
// returns a *string when successful
func (m *EdiscoveryExportOperation) GetOutputName()(*string) {
    return m.outputName
}
// GetReviewSet gets the reviewSet property value. Review set from where documents are exported.
// returns a EdiscoveryReviewSetable when successful
func (m *EdiscoveryExportOperation) GetReviewSet()(EdiscoveryReviewSetable) {
    return m.reviewSet
}
// GetReviewSetQuery gets the reviewSetQuery property value. The review set query that is used to filter the documents for export.
// returns a EdiscoveryReviewSetQueryable when successful
func (m *EdiscoveryExportOperation) GetReviewSetQuery()(EdiscoveryReviewSetQueryable) {
    return m.reviewSetQuery
}
// Serialize serializes information the current object
func (m *EdiscoveryExportOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetExportFileMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExportFileMetadata()))
        for i, v := range m.GetExportFileMetadata() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("exportFileMetadata", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportOptions() != nil {
        cast := (*m.GetExportOptions()).String()
        err = writer.WriteStringValue("exportOptions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportStructure() != nil {
        cast := (*m.GetExportStructure()).String()
        err = writer.WriteStringValue("exportStructure", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("outputName", m.GetOutputName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewSet", m.GetReviewSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewSetQuery", m.GetReviewSetQuery())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description provided for the export.
func (m *EdiscoveryExportOperation) SetDescription(value *string)() {
    m.description = value
}
// SetExportFileMetadata sets the exportFileMetadata property value. Contains the properties for an export file metadata, including downloadUrl, fileName, and size.
func (m *EdiscoveryExportOperation) SetExportFileMetadata(value []ExportFileMetadataable)() {
    m.exportFileMetadata = value
}
// SetExportOptions sets the exportOptions property value. The options provided for the export. For more information, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement,  tags.
func (m *EdiscoveryExportOperation) SetExportOptions(value *ExportOptions)() {
    m.exportOptions = value
}
// SetExportStructure sets the exportStructure property value. The options that specify the structure of the export. For more information, see reviewSet: export. Possible values are: none, directory, pst.
func (m *EdiscoveryExportOperation) SetExportStructure(value *ExportFileStructure)() {
    m.exportStructure = value
}
// SetOutputName sets the outputName property value. The name provided for the export.
func (m *EdiscoveryExportOperation) SetOutputName(value *string)() {
    m.outputName = value
}
// SetReviewSet sets the reviewSet property value. Review set from where documents are exported.
func (m *EdiscoveryExportOperation) SetReviewSet(value EdiscoveryReviewSetable)() {
    m.reviewSet = value
}
// SetReviewSetQuery sets the reviewSetQuery property value. The review set query that is used to filter the documents for export.
func (m *EdiscoveryExportOperation) SetReviewSetQuery(value EdiscoveryReviewSetQueryable)() {
    m.reviewSetQuery = value
}
type EdiscoveryExportOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetExportFileMetadata()([]ExportFileMetadataable)
    GetExportOptions()(*ExportOptions)
    GetExportStructure()(*ExportFileStructure)
    GetOutputName()(*string)
    GetReviewSet()(EdiscoveryReviewSetable)
    GetReviewSetQuery()(EdiscoveryReviewSetQueryable)
    SetDescription(value *string)()
    SetExportFileMetadata(value []ExportFileMetadataable)()
    SetExportOptions(value *ExportOptions)()
    SetExportStructure(value *ExportFileStructure)()
    SetOutputName(value *string)()
    SetReviewSet(value EdiscoveryReviewSetable)()
    SetReviewSetQuery(value EdiscoveryReviewSetQueryable)()
}
