package telego

import (
	"github.com/golang/mock/gomock"

	mockAPI "github.com/mymmrac/telego/api/mock"
)

type mockedBot struct {
	MockAPICaller          *mockAPI.MockCaller
	MockRequestConstructor *mockAPI.MockRequestConstructor
	Bot                    *Bot
}

func newMockedBot(ctrl *gomock.Controller) mockedBot {
	mb := mockedBot{
		MockAPICaller:          mockAPI.NewMockCaller(ctrl),
		MockRequestConstructor: mockAPI.NewMockRequestConstructor(ctrl),
	}

	bot, _ := NewBot(token,
		CustomAPICaller(mb.MockAPICaller),
		CustomRequestConstructor(mb.MockRequestConstructor),
		DefaultLogger(false, false))

	mb.Bot = bot

	return mb
}

type testNamedReade struct{}

func (t testNamedReade) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func (t testNamedReade) Name() string {
	return "test"
}

var testInputFile = InputFile{
	File: testNamedReade{},
}
