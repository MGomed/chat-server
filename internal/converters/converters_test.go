package converters

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	chat_api "github.com/MGomed/chat_server/pkg/chat_api"
)

func TestToChatInfoFromAPI(t *testing.T) {
	t.Run("Sunny case", func(t *testing.T) {
		var info = &chat_api.ChatInfo{
			Name: "Chat",
			Members: []*chat_api.ChatMember{
				{
					Name:  "Alex",
					Email: "Alex@mail.com",
				},
			},
		}

		chat := ToChatInfoFromAPI(info)
		assert.Equal(t, info.Name, chat.Name)
		assert.Equal(t, len(info.Members), len(chat.Members))
	})

	t.Run("Rainy case", func(t *testing.T) {
		chat := ToChatInfoFromAPI(nil)
		assert.Equal(t, chat == nil, true)
	})
}

func TestToMessageInfoFromAPI(t *testing.T) {
	t.Run("Sunny case", func(t *testing.T) {
		var info = &chat_api.MessageInfo{
			From:      "Alex",
			Text:      "Hello!",
			Timestamp: timestamppb.Now(),
		}

		msg := ToMessageInfoFromAPI(info)
		assert.Equal(t, info.From, msg.From)
		assert.Equal(t, info.Text, msg.Text)
		assert.Equal(t, info.Timestamp.AsTime(), msg.Timestamp)
	})

	t.Run("Rainy case", func(t *testing.T) {
		msg := ToMessageInfoFromAPI(nil)
		assert.Equal(t, msg == nil, true)
	})
}
