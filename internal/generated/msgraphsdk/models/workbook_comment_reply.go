package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookCommentReply struct {
    Entity
    // The content of the reply.
    content *string
    // The content type for the reply.
    contentType *string
}
// NewWorkbookCommentReply instantiates a new WorkbookCommentReply and sets the default values.
func NewWorkbookCommentReply()(*WorkbookCommentReply) {
    m := &WorkbookCommentReply{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookCommentReplyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookCommentReplyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookCommentReply(), nil
}
// GetContent gets the content property value. The content of the reply.
// returns a *string when successful
func (m *WorkbookCommentReply) GetContent()(*string) {
    return m.content
}
// GetContentType gets the contentType property value. The content type for the reply.
// returns a *string when successful
func (m *WorkbookCommentReply) GetContentType()(*string) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookCommentReply) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WorkbookCommentReply) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content of the reply.
func (m *WorkbookCommentReply) SetContent(value *string)() {
    m.content = value
}
// SetContentType sets the contentType property value. The content type for the reply.
func (m *WorkbookCommentReply) SetContentType(value *string)() {
    m.contentType = value
}
type WorkbookCommentReplyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetContentType()(*string)
    SetContent(value *string)()
    SetContentType(value *string)()
}
