package usecase

import (
	"context"
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/spf13/viper"
)

func (u *TicketUsecase) DistributeTicket(ctx context.Context) {
loop:

	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			go u.removeOldActivity()
			u.distributor()
		}

		time.Sleep(viper.GetDuration("app.distribute.delay") * time.Second)
	}
}

func (u *TicketUsecase) distributor() {

	tickets, err := u.repo.GetTicketListForDistribute()
	if err != nil || len(tickets) == 0 {
		return
	}

	for _, ticket := range tickets {
		support := u.suppUC.GetSupportForDistribution(ticket.Support.ID)
		if support == nil || support.ID == 0 {
			break
		}
		if ticket.Support.ID == support.ID {
			err := u.suppUC.UpdateSupportActivity(support.ID, ticket.ID)
			if err != nil {
				continue
			}
		} else {
			err := u.suppUC.AddSupportActivity(support, ticket.ID)
			if err != nil {
				continue
			}
		}

		ticket.Support = &models.User{ID: support.ID}
		ticket.Status.Set(internal_models.KeyTSInWork)
		err := u.UpdateTicket(ticket, &models.User{ID: 1}, false)
		if err != nil {
			continue
		}
	}
}

func (u *TicketUsecase) removeOldActivity() {
	tickets, err := u.repo.GetTicketListForReturnToDistribute()
	if err != nil || len(tickets) == 0 {
		return
	}

	for _, ticket := range tickets {
		ticket.Status.Set(internal_models.KeyTSWait)
		ticket.Support = nil

		if err := u.UpdateTicket(ticket, &models.User{ID: 1}, false); err != nil {
			continue
		}

		if err := u.suppUC.RemoveSupportActivity(ticket.ID); err != nil {
			continue
		}
	}
}
