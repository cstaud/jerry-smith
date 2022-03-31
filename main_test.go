package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_send(t *testing.T) {
	messenger := messengerMock{}

	tests := []struct {
		name     string
		text     string
		receiver user
		mocks    func()
		wantErr  error
	}{
		{
			name:     "error on too long text",
			text:     "a much toooooooooooooooooooo long text",
			receiver: user{Name: "dummy"},
			mocks:    func() {},
			wantErr:  fmt.Errorf("message too long"),
		},
		{
			name:     "error on deliver",
			text:     "some text",
			receiver: user{Name: "dummy"},
			mocks: func() {
				messenger.ReturnErr = fmt.Errorf("could not deliver")
			},
			wantErr: fmt.Errorf("could not deliver"),
		},
		{
			name:     "success",
			text:     "some text",
			receiver: user{Name: "dummy"},
			mocks: func() {
				messenger.ReturnErr = nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks()
			err := send(messenger, tt.receiver, tt.text)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

type messengerMock struct {
	ReturnErr error
}

func (m messengerMock) deliver(receiver, text string) error {
	return m.ReturnErr
}
