package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookRangeFormat struct {
    Entity
    // Collection of border objects that apply to the overall range selected Read-only.
    borders []WorkbookRangeBorderable
    // The width of all columns within the range. If the column widths aren't uniform, null will be returned.
    columnWidth *float64
    // Returns the fill object defined on the overall range. Read-only.
    fill WorkbookRangeFillable
    // Returns the font object defined on the overall range selected Read-only.
    font WorkbookRangeFontable
    // The horizontal alignment for the specified object. Possible values are: General, Left, Center, Right, Fill, Justify, CenterAcrossSelection, Distributed.
    horizontalAlignment *string
    // Returns the format protection object for a range. Read-only.
    protection WorkbookFormatProtectionable
    // The height of all rows in the range. If the row heights aren't uniform null will be returned.
    rowHeight *float64
    // The vertical alignment for the specified object. Possible values are: Top, Center, Bottom, Justify, Distributed.
    verticalAlignment *string
    // Indicates whether Excel wraps the text in the object. A null value indicates that the entire range doesn't have a uniform wrap setting.
    wrapText *bool
}
// NewWorkbookRangeFormat instantiates a new WorkbookRangeFormat and sets the default values.
func NewWorkbookRangeFormat()(*WorkbookRangeFormat) {
    m := &WorkbookRangeFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookRangeFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookRangeFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookRangeFormat(), nil
}
// GetBorders gets the borders property value. Collection of border objects that apply to the overall range selected Read-only.
// returns a []WorkbookRangeBorderable when successful
func (m *WorkbookRangeFormat) GetBorders()([]WorkbookRangeBorderable) {
    return m.borders
}
// GetColumnWidth gets the columnWidth property value. The width of all columns within the range. If the column widths aren't uniform, null will be returned.
// returns a *float64 when successful
func (m *WorkbookRangeFormat) GetColumnWidth()(*float64) {
    return m.columnWidth
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookRangeFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["borders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookRangeBorderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookRangeBorderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkbookRangeBorderable)
                }
            }
            m.SetBorders(res)
        }
        return nil
    }
    res["columnWidth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnWidth(val)
        }
        return nil
    }
    res["fill"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookRangeFillFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFill(val.(WorkbookRangeFillable))
        }
        return nil
    }
    res["font"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookRangeFontFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFont(val.(WorkbookRangeFontable))
        }
        return nil
    }
    res["horizontalAlignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHorizontalAlignment(val)
        }
        return nil
    }
    res["protection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookFormatProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtection(val.(WorkbookFormatProtectionable))
        }
        return nil
    }
    res["rowHeight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowHeight(val)
        }
        return nil
    }
    res["verticalAlignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerticalAlignment(val)
        }
        return nil
    }
    res["wrapText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWrapText(val)
        }
        return nil
    }
    return res
}
// GetFill gets the fill property value. Returns the fill object defined on the overall range. Read-only.
// returns a WorkbookRangeFillable when successful
func (m *WorkbookRangeFormat) GetFill()(WorkbookRangeFillable) {
    return m.fill
}
// GetFont gets the font property value. Returns the font object defined on the overall range selected Read-only.
// returns a WorkbookRangeFontable when successful
func (m *WorkbookRangeFormat) GetFont()(WorkbookRangeFontable) {
    return m.font
}
// GetHorizontalAlignment gets the horizontalAlignment property value. The horizontal alignment for the specified object. Possible values are: General, Left, Center, Right, Fill, Justify, CenterAcrossSelection, Distributed.
// returns a *string when successful
func (m *WorkbookRangeFormat) GetHorizontalAlignment()(*string) {
    return m.horizontalAlignment
}
// GetProtection gets the protection property value. Returns the format protection object for a range. Read-only.
// returns a WorkbookFormatProtectionable when successful
func (m *WorkbookRangeFormat) GetProtection()(WorkbookFormatProtectionable) {
    return m.protection
}
// GetRowHeight gets the rowHeight property value. The height of all rows in the range. If the row heights aren't uniform null will be returned.
// returns a *float64 when successful
func (m *WorkbookRangeFormat) GetRowHeight()(*float64) {
    return m.rowHeight
}
// GetVerticalAlignment gets the verticalAlignment property value. The vertical alignment for the specified object. Possible values are: Top, Center, Bottom, Justify, Distributed.
// returns a *string when successful
func (m *WorkbookRangeFormat) GetVerticalAlignment()(*string) {
    return m.verticalAlignment
}
// GetWrapText gets the wrapText property value. Indicates whether Excel wraps the text in the object. A null value indicates that the entire range doesn't have a uniform wrap setting.
// returns a *bool when successful
func (m *WorkbookRangeFormat) GetWrapText()(*bool) {
    return m.wrapText
}
// Serialize serializes information the current object
func (m *WorkbookRangeFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBorders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBorders()))
        for i, v := range m.GetBorders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("borders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("columnWidth", m.GetColumnWidth())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("fill", m.GetFill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("font", m.GetFont())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("horizontalAlignment", m.GetHorizontalAlignment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("protection", m.GetProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("rowHeight", m.GetRowHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("verticalAlignment", m.GetVerticalAlignment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wrapText", m.GetWrapText())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBorders sets the borders property value. Collection of border objects that apply to the overall range selected Read-only.
func (m *WorkbookRangeFormat) SetBorders(value []WorkbookRangeBorderable)() {
    m.borders = value
}
// SetColumnWidth sets the columnWidth property value. The width of all columns within the range. If the column widths aren't uniform, null will be returned.
func (m *WorkbookRangeFormat) SetColumnWidth(value *float64)() {
    m.columnWidth = value
}
// SetFill sets the fill property value. Returns the fill object defined on the overall range. Read-only.
func (m *WorkbookRangeFormat) SetFill(value WorkbookRangeFillable)() {
    m.fill = value
}
// SetFont sets the font property value. Returns the font object defined on the overall range selected Read-only.
func (m *WorkbookRangeFormat) SetFont(value WorkbookRangeFontable)() {
    m.font = value
}
// SetHorizontalAlignment sets the horizontalAlignment property value. The horizontal alignment for the specified object. Possible values are: General, Left, Center, Right, Fill, Justify, CenterAcrossSelection, Distributed.
func (m *WorkbookRangeFormat) SetHorizontalAlignment(value *string)() {
    m.horizontalAlignment = value
}
// SetProtection sets the protection property value. Returns the format protection object for a range. Read-only.
func (m *WorkbookRangeFormat) SetProtection(value WorkbookFormatProtectionable)() {
    m.protection = value
}
// SetRowHeight sets the rowHeight property value. The height of all rows in the range. If the row heights aren't uniform null will be returned.
func (m *WorkbookRangeFormat) SetRowHeight(value *float64)() {
    m.rowHeight = value
}
// SetVerticalAlignment sets the verticalAlignment property value. The vertical alignment for the specified object. Possible values are: Top, Center, Bottom, Justify, Distributed.
func (m *WorkbookRangeFormat) SetVerticalAlignment(value *string)() {
    m.verticalAlignment = value
}
// SetWrapText sets the wrapText property value. Indicates whether Excel wraps the text in the object. A null value indicates that the entire range doesn't have a uniform wrap setting.
func (m *WorkbookRangeFormat) SetWrapText(value *bool)() {
    m.wrapText = value
}
type WorkbookRangeFormatable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBorders()([]WorkbookRangeBorderable)
    GetColumnWidth()(*float64)
    GetFill()(WorkbookRangeFillable)
    GetFont()(WorkbookRangeFontable)
    GetHorizontalAlignment()(*string)
    GetProtection()(WorkbookFormatProtectionable)
    GetRowHeight()(*float64)
    GetVerticalAlignment()(*string)
    GetWrapText()(*bool)
    SetBorders(value []WorkbookRangeBorderable)()
    SetColumnWidth(value *float64)()
    SetFill(value WorkbookRangeFillable)()
    SetFont(value WorkbookRangeFontable)()
    SetHorizontalAlignment(value *string)()
    SetProtection(value WorkbookFormatProtectionable)()
    SetRowHeight(value *float64)()
    SetVerticalAlignment(value *string)()
    SetWrapText(value *bool)()
}
