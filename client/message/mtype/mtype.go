package mtype

import "fmt"

type MessageType uint8

const (
	Reserve MessageType = iota
	Calc
	Poll
	Error
)

func (mType MessageType) String() string {
	switch mType {
	case Reserve:
		return "reserve"
	case Calc:
		return "calc"
	case Poll:
		return "poll"
	case Error:
		return "error"
	}

	return "undefined"
}

func (mType MessageType) Encode() (code uint8, err error) {
	if mType == Error {
		err = fmt.Errorf("Invalid request type: %s", mType)
		return
	}

	code = uint8(mType)
	return
}

func Decode(code uint8) (mType MessageType, err error) {
	if code > uint8(Error) {
		err = fmt.Errorf("Unknown message type code: %d", code)
		return
	}

	mType = MessageType(code)
	return
}
