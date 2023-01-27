package telegramClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"telegramSDK/types"
)

const (
	host              = "api.telegram.org"
	UpdateMethod      = "getUpdates"
	SendMessageMethod = "sendMessage"
)

type TelegramClient struct {
	client   http.Client
	host     string
	basePath string
}

func NewTelegramClient(token string) *TelegramClient {
	return &TelegramClient{
		client:   http.Client{},
		host:     host,
		basePath: "bot" + token,
	}
}

func (c *TelegramClient) Updates(offset int) ([]types.Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	body, err := c.makeRequest(UpdateMethod, q)
	if err != nil {
		return nil, fmt.Errorf("couldn't make request: %w", err)
	}
	var responseTelegram types.ResponseTelegram

	err = json.Unmarshal(body, &responseTelegram)
	if err != nil {
		return nil, fmt.Errorf("couldnt unmarshal: %w", err)
	}
	return responseTelegram.Result, nil
}

func (c *TelegramClient) makeRequest(method string, q url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path.Join(c.basePath, method),
	}

	body := q.Encode()
	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("couldnt read body: %w", err)
	}
	return data, nil

}

func (c *TelegramClient) SendMessage(chatId int, text string) error {
	q := url.Values{}
	switch text {
	case "/about":
		q.Set("text", types.AboutMsg)
	case "/links":
		q.Set("parse_mode", "HTML")
		q.Set("text", types.LinksMsg)
	case "/help", "/start":
		q.Set("text", types.HelpMsg)
		var keyboard = types.NewReplyKeyboardMarkup(
			types.NewKeyboardRow(types.KeyBoardButton{BtnText: "/help"}),
			types.NewKeyboardRow(types.KeyBoardButton{BtnText: "/start"}),
			types.NewKeyboardRow(types.KeyBoardButton{BtnText: "/links"}),
			types.NewKeyboardRow(types.KeyBoardButton{BtnText: "/about"}),
		)
		b, err := json.Marshal(keyboard)
		if err != nil {
			return fmt.Errorf("couldnt create keyboard: %w", err)
		}
		q.Set("reply_markup", string(b))
	default:
		q.Set("text", types.UnknownMsg)
	}
	q.Set("chat_id", strconv.Itoa(chatId))

	_, err := c.makeRequest(SendMessageMethod, q)
	if err != nil {
		return err
	}
	return nil
}
