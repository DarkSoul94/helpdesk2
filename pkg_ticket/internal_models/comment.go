package internal_models

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
)

type Comment struct {
	ID       uint64
	TicketId uint64
	Date     time.Time
	Author   *models.User
	Text     string
}
