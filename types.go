package tgbot

import (
	"encoding/json"
	"fmt"
)

type Integer int32

type TGResponse struct {
	Ok          bool
	Description string `json:",omitempty"`
	Result      json.RawMessage
}

type TGID struct {
	Id Integer `json:"id"`
}

type TGUser struct {
	TGID
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:",omitempty"`
}

type TGGroupChat struct {
	TGID
	Title string
}

type TGUserGroupChat struct {
	TGUser
	TGGroupChat
}

type TGUpdate struct {
	UpdateId Integer   `json:"update_id"`
	Message  TGMessage `json:"message,omitempty"`
}

type TGMessage struct {
	MessageId      Integer         `json:"message_id"`       // Unique message identifier
	From           TGUser          `json:"from"`             // Sender
	Date           Integer         `json:"date"`             // Date the message was sent in Unix time
	Chat           TGUserGroupChat `json:"chat"`             // Conversation the message belongs to — user in case of a private message, GroupChat in case of a group
	ForwardFrom    *TGUser         `json:"forward_from"`     // Optional. For forwarded messages, sender of the original message
	ForwardDate    *Integer        `json:"forward_date"`     // Optional. For forwarded messages, date the original message was sent in Unix time
	ReplyToMessage *TGMessage      `json:"reply_to_message"` // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	Text           *string         `json:"text"`             // Optional. For text messages, the actual UTF-8 text of the message
	//Audio               *Audio          `json:"asdf"` // Optional. Message is an audio file, information about the file
	//Document            *Document       `json:"asdf"` // Optional. Message is a general file, information about the file
	//Photo               []PhotoSize     `json:"asdf"` // Optional. Message is a photo, available sizes of the photo
	//Sticker             *Sticker        `json:"asdf"` // Optional. Message is a sticker, information about the sticker
	//Video               *Video          `json:"asdf"` // Optional. Message is a video, information about the video
	//Contact             *Contact        `json:"asdf"` // Optional. Message is a shared contact, information about the contact
	//Location            *Location       `json:"asdf"` // Optional. Message is a shared location, information about the location
	NewChatParticipant  *TGUser `json:"new_chat_participant"`  // Optional. A new member was added to the group, information about them (this member may be bot itself)
	LeftChatParticipant *TGUser `json:"left_chat_participant"` // Optional. A member was removed from the group, information about them (this member may be bot itself)
	NewChatTitle        *string `json:"new_chat_title"`        // Optional. A group title was changed to this value
	//NewChatPhoto        []PhotoSize    `json:"new_chat_photo"` // Optional. A group photo was change to this value
	DeleteChatPhoto  bool `json:"delete_chat_photo"`  // Optional. Informs that the group photo was deleted
	GroupChatCreated bool `json:"group_chat_created"` // Optional. Informs that the group has been created
}

func (t TGMessage) String() string {
	str := t.From.String()
	if t.Text != nil {
		str += ": " + *t.Text
	}
	return str
}

func (u TGUser) String() string {
	return fmt.Sprintf("%s %s (%s)", u.FirstName, u.LastName, u.Username)
}
