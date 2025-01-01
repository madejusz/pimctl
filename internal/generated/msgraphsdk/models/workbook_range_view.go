package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookRangeView struct {
    Entity
    // The cell addresses.
    cellAddresses i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // The number of visible columns. Read-only.
    columnCount *int32
    // The formula in A1-style notation.
    formulas i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // The formula in A1-style notation, in the user's language and number-formatting locale. For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
    formulasLocal i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // Represents the formula in R1C1-style notation.
    formulasR1C1 i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // The index of the range.
    index *int32
    // Excel's number format code for the given cell. Read-only.
    numberFormat i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // The number of visible rows. Read-only.
    rowCount *int32
    // The collection of range views associated with the range. Read-only.Read-only.
    rows []WorkbookRangeViewable
    // The text values of the specified range. The Text value won't depend on the cell width. The # sign substitution that happens in Excel UI won't affect the text value returned by the API. Read-only.
    text i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // The raw values of the specified range view. The data returned could be of type string, number, or a Boolean. Cell that contains an error returns the error string.
    values i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // The type of data of each cell. Read-only. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error.
    valueTypes i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
}
// NewWorkbookRangeView instantiates a new WorkbookRangeView and sets the default values.
func NewWorkbookRangeView()(*WorkbookRangeView) {
    m := &WorkbookRangeView{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookRangeViewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookRangeViewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookRangeView(), nil
}
// GetCellAddresses gets the cellAddresses property value. The cell addresses.
// returns a UntypedNodeable when successful
func (m *WorkbookRangeView) GetCellAddresses()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.cellAddresses
}
// GetColumnCount gets the columnCount property value. The number of visible columns. Read-only.
// returns a *int32 when successful
func (m *WorkbookRangeView) GetColumnCount()(*int32) {
    return m.columnCount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookRangeView) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cellAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellAddresses(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["columnCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnCount(val)
        }
        return nil
    }
    res["formulas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulas(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["formulasLocal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulasLocal(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["formulasR1C1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulasR1C1(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    res["numberFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberFormat(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["rowCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowCount(val)
        }
        return nil
    }
    res["rows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookRangeViewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookRangeViewable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkbookRangeViewable)
                }
            }
            m.SetRows(res)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["valueTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueTypes(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    return res
}
// GetFormulas gets the formulas property value. The formula in A1-style notation.
// returns a UntypedNodeable when successful
func (m *WorkbookRangeView) GetFormulas()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.formulas
}
// GetFormulasLocal gets the formulasLocal property value. The formula in A1-style notation, in the user's language and number-formatting locale. For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
// returns a UntypedNodeable when successful
func (m *WorkbookRangeView) GetFormulasLocal()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.formulasLocal
}
// GetFormulasR1C1 gets the formulasR1C1 property value. Represents the formula in R1C1-style notation.
// returns a UntypedNodeable when successful
func (m *WorkbookRangeView) GetFormulasR1C1()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.formulasR1C1
}
// GetIndex gets the index property value. The index of the range.
// returns a *int32 when successful
func (m *WorkbookRangeView) GetIndex()(*int32) {
    return m.index
}
// GetNumberFormat gets the numberFormat property value. Excel's number format code for the given cell. Read-only.
// returns a UntypedNodeable when successful
func (m *WorkbookRangeView) GetNumberFormat()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.numberFormat
}
// GetRowCount gets the rowCount property value. The number of visible rows. Read-only.
// returns a *int32 when successful
func (m *WorkbookRangeView) GetRowCount()(*int32) {
    return m.rowCount
}
// GetRows gets the rows property value. The collection of range views associated with the range. Read-only.Read-only.
// returns a []WorkbookRangeViewable when successful
func (m *WorkbookRangeView) GetRows()([]WorkbookRangeViewable) {
    return m.rows
}
// GetText gets the text property value. The text values of the specified range. The Text value won't depend on the cell width. The # sign substitution that happens in Excel UI won't affect the text value returned by the API. Read-only.
// returns a UntypedNodeable when successful
func (m *WorkbookRangeView) GetText()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.text
}
// GetValues gets the values property value. The raw values of the specified range view. The data returned could be of type string, number, or a Boolean. Cell that contains an error returns the error string.
// returns a UntypedNodeable when successful
func (m *WorkbookRangeView) GetValues()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.values
}
// GetValueTypes gets the valueTypes property value. The type of data of each cell. Read-only. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error.
// returns a UntypedNodeable when successful
func (m *WorkbookRangeView) GetValueTypes()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.valueTypes
}
// Serialize serializes information the current object
func (m *WorkbookRangeView) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("cellAddresses", m.GetCellAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("columnCount", m.GetColumnCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("formulas", m.GetFormulas())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("formulasLocal", m.GetFormulasLocal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("formulasR1C1", m.GetFormulasR1C1())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("numberFormat", m.GetNumberFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("rowCount", m.GetRowCount())
        if err != nil {
            return err
        }
    }
    if m.GetRows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRows()))
        for i, v := range m.GetRows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("rows", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("valueTypes", m.GetValueTypes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCellAddresses sets the cellAddresses property value. The cell addresses.
func (m *WorkbookRangeView) SetCellAddresses(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.cellAddresses = value
}
// SetColumnCount sets the columnCount property value. The number of visible columns. Read-only.
func (m *WorkbookRangeView) SetColumnCount(value *int32)() {
    m.columnCount = value
}
// SetFormulas sets the formulas property value. The formula in A1-style notation.
func (m *WorkbookRangeView) SetFormulas(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.formulas = value
}
// SetFormulasLocal sets the formulasLocal property value. The formula in A1-style notation, in the user's language and number-formatting locale. For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
func (m *WorkbookRangeView) SetFormulasLocal(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.formulasLocal = value
}
// SetFormulasR1C1 sets the formulasR1C1 property value. Represents the formula in R1C1-style notation.
func (m *WorkbookRangeView) SetFormulasR1C1(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.formulasR1C1 = value
}
// SetIndex sets the index property value. The index of the range.
func (m *WorkbookRangeView) SetIndex(value *int32)() {
    m.index = value
}
// SetNumberFormat sets the numberFormat property value. Excel's number format code for the given cell. Read-only.
func (m *WorkbookRangeView) SetNumberFormat(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.numberFormat = value
}
// SetRowCount sets the rowCount property value. The number of visible rows. Read-only.
func (m *WorkbookRangeView) SetRowCount(value *int32)() {
    m.rowCount = value
}
// SetRows sets the rows property value. The collection of range views associated with the range. Read-only.Read-only.
func (m *WorkbookRangeView) SetRows(value []WorkbookRangeViewable)() {
    m.rows = value
}
// SetText sets the text property value. The text values of the specified range. The Text value won't depend on the cell width. The # sign substitution that happens in Excel UI won't affect the text value returned by the API. Read-only.
func (m *WorkbookRangeView) SetText(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.text = value
}
// SetValues sets the values property value. The raw values of the specified range view. The data returned could be of type string, number, or a Boolean. Cell that contains an error returns the error string.
func (m *WorkbookRangeView) SetValues(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.values = value
}
// SetValueTypes sets the valueTypes property value. The type of data of each cell. Read-only. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error.
func (m *WorkbookRangeView) SetValueTypes(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.valueTypes = value
}
type WorkbookRangeViewable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCellAddresses()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetColumnCount()(*int32)
    GetFormulas()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetFormulasLocal()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetFormulasR1C1()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetIndex()(*int32)
    GetNumberFormat()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetRowCount()(*int32)
    GetRows()([]WorkbookRangeViewable)
    GetText()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetValues()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetValueTypes()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    SetCellAddresses(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetColumnCount(value *int32)()
    SetFormulas(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetFormulasLocal(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetFormulasR1C1(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetIndex(value *int32)()
    SetNumberFormat(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetRowCount(value *int32)()
    SetRows(value []WorkbookRangeViewable)()
    SetText(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetValues(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetValueTypes(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
}
