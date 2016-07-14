package event

import (
	"bytes"
)

type AuthEvent struct {
	EventHeader
	User   string
	Token  string
	Index  uint32
	Reauth bool
}

func (ev *AuthEvent) Encode(buffer *bytes.Buffer) {
	EncodeStringValue(buffer, ev.User)
	EncodeStringValue(buffer, ev.Token)
	EncodeUInt32Value(buffer, ev.Index)
	EncodeBoolValue(buffer, ev.Reauth)
}
func (ev *AuthEvent) Decode(buffer *bytes.Buffer) (err error) {
	ev.User, err = DecodeStringValue(buffer)
	if nil == err {
		ev.Token, err = DecodeStringValue(buffer)
		if nil == err {
			ev.Index, err = DecodeUInt32Value(buffer)
			if nil == err {
				ev.Reauth, err = DecodeBoolValue(buffer)
			}
		}
	}
	return
}
