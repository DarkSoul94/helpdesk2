package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/spf13/viper"
)

func (u *TicketUsecase) DistributeTicket(ctx context.Context) {
	fmt.Println("distributor started")
loop:

	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			u.distributor()
		}

		time.Sleep(viper.GetDuration("app.distribute.delay") * time.Second)
	}
}

func (u *TicketUsecase) distributor() {
	u.removeOldActivity()

	tickets, err := u.repo.GetTicketListForDistribute()
	if err != nil || len(tickets) == 0 {
		return
	}
	fmt.Println(len(tickets))

	for _, ticket := range tickets {
		support := u.suppUC.GetSupportForDistribution(ticket.Support.ID)
		fmt.Println(support)
		if support == nil {
			break
		}
		if ticket.Support.ID == support.ID {
			err := u.suppUC.UpdateSupportActivity(support.ID, ticket.ID)
			if err != nil {
				return
			}
		} else {
			err := u.suppUC.AddSupportActivity(support, ticket.ID)
			if err != nil {
				return
			}
		}

		ticket.Support = &models.User{ID: support.ID}
		ticket.Status.Set(internal_models.KeyTSInWork)
		err := u.UpdateTicket(ticket, &models.User{ID: 1}, false)
		if err != nil {
			return
		}
	}
}

func (u *TicketUsecase) removeOldActivity() {

}
