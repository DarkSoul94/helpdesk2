package internal_models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
)

type Ticket struct {
	ID             uint64
	Date           time.Time
	CatSect        *SectionWithCategory
	Text           string
	Status         *TicketStatus
	Filial         string
	IP             string
	Author         *models.User
	Support        *models.User
	ResolvedUser   *models.User
	ServiceComment string
	//Files          []*File
	//Comments       []*CommentHistory
	Grade uint
}

//HashСalculation расчитывает хэш над текстом тикета, email-ом автора и разделом категории.
func (t *Ticket) HashСalculation() string {
	hashTarget := fmt.Sprintf("%s;%s;%s", t.Text, t.Author.Email, t.CatSect.Name)
	hasher := md5.New()
	hasher.Write([]byte(hashTarget))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (t *Ticket) IpMask() (string, error) {
	octets := strings.Split(t.IP, ".")
	if len(octets) > 1 {
		return octets[0] + "." + octets[1] + "." + octets[2], nil
	}

	return "", errors.New("Wrong ip")
}
