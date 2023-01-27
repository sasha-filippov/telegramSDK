package types

type ResponseTelegram struct {
	Ok          bool     `json:"ok"`
	Description string   `json:"description,omitempty"`
	Result      []Update `json:"result"`
}

type Update struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message"`
}

type Message struct {
	MessageID   int                  `json:"message_id"`
	Text        string               `json:"text"`
	Chat        Chat                 `json:"chat"`
	ReplyMarkup *ReplyKeyboardMarkup `json:"reply_markup"`
}

type Chat struct {
	ChatID int `json:"id"`
}

type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyBoardButton `json:"keyboard"`
	OneTime  bool               `json:"one_time_keyboard"`
	Resize   bool               `json:"resize_keyboard"`
}

type KeyBoardButton struct {
	BtnText string `json:"text"`
}

func NewReplyKeyboardMarkup(rows ...[]KeyBoardButton) ReplyKeyboardMarkup {
	var keyboard [][]KeyBoardButton
	keyboard = append(keyboard, rows...)
	return ReplyKeyboardMarkup{
		Keyboard: keyboard,
		Resize:   true,
		OneTime:  true,
	}
}

func NewKeyboardRow(buttons ...KeyBoardButton) []KeyBoardButton {
	var row []KeyBoardButton
	row = append(row, buttons...)
	return row
}
