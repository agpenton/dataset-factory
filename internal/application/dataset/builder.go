package dataset

import "encoding/json"

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type record struct {
	Messages []message `json:"messages"`
}

func BuildRecord(user, assistant string) string {
	r := record{
		Messages: []message{
			{
				Role:    "user",
				Content: user,
			},
			{
				Role:    "assistant",
				Content: assistant,
			},
		},
	}

	data, _ := json.Marshal(r)

	return string(data)
}
