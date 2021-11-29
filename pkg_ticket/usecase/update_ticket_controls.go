package usecase

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/global_const/actions"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/spf13/viper"
)

func (u *TicketUsecase) prepareBeforeUpdate(ticket, existTicket *internal_models.Ticket, user *models.User) models.Err {
	if u.permUC.CheckPermission(user.Group.ID, actions.AdminTA) {
		return nil
	}

	err := u.generalChecks(existTicket, user)
	if err != nil {
		return err
	}

	if ticket.Status.ID != existTicket.Status.ID {
		err := u.checkChangeStatus(ticket, existTicket, user)
		if err != nil {
			return err
		}

		err = u.changeStatus(ticket, existTicket, user)
	}

	if ticket.Support.ID != 0 && ticket.Support.ID != existTicket.Support.ID {
		err := u.checkChangeSupport(user)
		if err != nil {
			return err
		}
	}

	if ticket.CatSect.ID != 0 && ticket.CatSect.ID != existTicket.CatSect.ID {
		err := u.checkChangeSection(ticket.Status.ID, user)
		if err != nil {
			return err
		}
	}

	if len(ticket.ServiceComment) > 0 && ticket.ServiceComment != existTicket.ServiceComment {
		err := u.checkChangeServiceComment(user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *TicketUsecase) generalChecks(existTicket *internal_models.Ticket, user *models.User) models.Err {
	type hDate struct {
		year  int
		month time.Month
		day   int
	}
	var nowDate, completeDate hDate

	switch existTicket.Status.ID {
	case internal_models.TSCompletedID, internal_models.TSRejectedID:
		if !u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Work) {
			if existTicket.Author.ID == user.ID {
				history, err := u.repo.GetLastTicketStatusHistory(existTicket.ID)
				if err != nil {
					return models.InternalError(err.Error())
				}

				completeDate.year, completeDate.month, completeDate.day = history.SelectTime.Local().Date()
				nowDate.year, nowDate.month, nowDate.day = time.Now().Local().Date()

				if nowDate != completeDate {
					return errTicketOld
				}
			}
		}
	case internal_models.TSRevisionID:
		if existTicket.Author.ID != user.ID {
			return errCannotUpdateTicket
		}
	default:
		return nil
	}

	return nil
}

func (u *TicketUsecase) checkChangeStatus(ticket, existTicket *internal_models.Ticket, user *models.User) models.Err {
	workOnTicket := u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Work)
	resolveTicket := u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Resolve)

	switch ticket.Status.ID {
	case internal_models.TSWaitID:
		if ticket.Author.ID != user.ID {
			return errCannotUpdateTicket
		}

		if existTicket.Status.ID == internal_models.TSInWorkID ||
			existTicket.Status.ID == internal_models.TSImplementationID ||
			ticket.Status.ID == internal_models.TSPostponedID {
			return errCannotSelectStatus
		}

		if existTicket.CatSect.NeedApproval && existTicket.ResolvedUser.ID == 0 {
			ticket.Status.ID = internal_models.TSWaitForResolveID
		}

	case internal_models.TSInWorkID:
		if !workOnTicket || existTicket.Status.ID == internal_models.TSImplementationID {
			return errCannotSelectStatus
		}

	case internal_models.TSImplementationID:
		if ticket.Support == nil || ticket.Support.ID == 0 {
			return errTicketWithoutSupport
		}

		if !workOnTicket {
			return errCannotSelectStatus
		}

		count := u.repo.GetTicketsCount(ticket.Support.ID, ticket.Status.ID)
		if count >= viper.GetInt("app.limitation.implementation") {
			return errLimitImplementation
		}

	case internal_models.TSPostponedID:
		if ticket.Support == nil || ticket.Support.ID == 0 {
			return errTicketWithoutSupport
		}

		if !workOnTicket {
			return errCannotSelectStatus
		}

		count := u.repo.GetTicketsCount(ticket.Support.ID, ticket.Status.ID)
		if count >= viper.GetInt("app.limitation.postponed") {
			return errLimitImplementation
		}

	case internal_models.TSRevisionID:
		if !(workOnTicket || resolveTicket) {
			return errCannotSelectStatus
		}

	case internal_models.TSRejectedID:
		if !(workOnTicket || resolveTicket) || existTicket.Author.ID == user.ID {
			return errCannotSelectStatus
		}
	case internal_models.TSCompletedID:
		if !workOnTicket || existTicket.Author.ID == user.ID {
			return errCannotSelectStatus
		}
	}

	return nil
}

func (u *TicketUsecase) changeStatus(ticket, existTicket *internal_models.Ticket, user *models.User) models.Err {
	if existTicket.Support.ID == 0 && ticket.Support.ID == 0 {
		if u.permUC.CheckPermission(user.Group.ID, actions.AdminTA) || u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Work) {
			ticket.Support = user
		}
	}

	switch ticket.Status.ID {
	case internal_models.TSInWorkID:
		err := u.suppUC.UpdateSupportActivity(ticket.Support.ID, ticket.ID)
		if err != nil {
			return err
		}

	default:
		u.suppUC.RemoveSupportActivity(ticket.ID)
	}

	return nil
}

func (u *TicketUsecase) checkChangeSupport(user *models.User) models.Err {
	if !u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Work) {
		return errCannotUpdateTicket
	}

	return nil
}

func (u *TicketUsecase) changeSupport(ticket *internal_models.Ticket) models.Err {
	if ticket.Status.ID == internal_models.TSInWorkID {
		err := u.suppUC.UpdateSupportActivity(ticket.Support.ID, ticket.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *TicketUsecase) checkChangeSection(statusID uint64, user *models.User) models.Err {
	if statusID == internal_models.TSRevisionID || u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Work) {
		return errCannotUpdateTicket
	}

	return nil
}

func (u *TicketUsecase) changeSection(ticket *internal_models.Ticket) models.Err {
	section, err := u.catSecUC.GetSectionWithCategoryByID(ticket.CatSect.ID)
	if err != nil {
		return err
	}

	if section.NeedApproval {
		ticket.Status.ID = internal_models.TSWaitForResolveID
	}

	return nil
}

func (u *TicketUsecase) checkChangeServiceComment(user *models.User) models.Err {
	if !u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Work) {
		return errCannotUpdateTicket
	}

	return nil
}