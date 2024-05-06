package domain

import "time"

type Message struct {
	ID      string    // Identificador único da mensagem
	From    string    // ID do usuário que envia a mensagem
	To      string    // ID do usuário receptor
	Content string    // Conteúdo da mensagem
	SentAt  time.Time // Timestamp de quando a mensagem foi enviada
}
