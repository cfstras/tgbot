package tgbot

type Bot struct {
	key string
}

func New(apiKey string) (*Bot, error) {
	if len(apiKey) < 3 {
		return nil, ErrorInvalidArgs
	}
	b := &Bot{apiKey}
	err := b.connect()
	return b, err
}

func (b *Bot) connect() error {
	return ErrorNotImplemented
}

func (b *Bot) Name() string {
	return ""
}
