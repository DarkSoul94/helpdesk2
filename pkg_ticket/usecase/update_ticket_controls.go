package usecase

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/global_const/actions"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/spf13/viper"
)

const (
	fieldSection string = "section"
	fieldStatus  string = "status"
	fieldSupport string = "support"
	fieldComment string = "comment"
	fieldFiles   string = "files"
)

func (u *TicketUsecase) prepareBeforeUpdate(ticket, existTicket *internal_models.Ticket, user *models.User) models.Err {
	fields, err := u.checkChanges(ticket, existTicket, user)
	if err != nil {
		return err
	}

	if !u.permUC.CheckPermission(user.Group.ID, actions.AdminTA) {
		err := u.generalChecks(existTicket, user)
		if err != nil {
			return err
		}

		err = u.checkRules(fields, ticket, existTicket, user)
		if err != nil {
			return err
		}
	}

	err = u.prepareTicket(fields, ticket, existTicket, user)
	if err != nil {
		return err
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

func (u *TicketUsecase) checkChanges(ticket, existTicket *internal_models.Ticket, user *models.User) ([]string, models.Err) {
	fields := make([]string, 0)

	if ticket.Status.ID != existTicket.Status.ID {
		fields = append(fields, fieldStatus)
	}

	if ticket.Support.ID != 0 && ticket.Support.ID != existTicket.Support.ID {
		fields = append(fields, fieldSupport)
	}

	if ticket.CatSect.ID != 0 && ticket.CatSect.ID != existTicket.CatSect.ID {
		fields = append(fields, fieldSection)
	}

	if len(ticket.ServiceComment) > 0 && ticket.ServiceComment != existTicket.ServiceComment {
		fields = append(fields, fieldComment)
	}

	if len(ticket.Files) > 0 {
		fields = append(fields, fieldFiles)
	}

	return fields, nil
}

func (u *TicketUsecase) checkRules(fields []string, ticket, existTicket *internal_models.Ticket, user *models.User) models.Err {
	for _, key := range fields {
		switch key {
		case fieldStatus:
			err := u.checkChangeStatus(ticket, existTicket, user)
			if err != nil {
				return err
			}

		case fieldSupport:
			err := u.checkChangeSupport(user)
			if err != nil {
				return err
			}

		case fieldSection:
			err := u.checkChangeSection(ticket.Status.ID, user)
			if err != nil {
				return err
			}

		case fieldComment:
			err := u.checkChangeServiceComment(user)
			if err != nil {
				return err
			}

		case fieldFiles:
			err := u.checkChangeFiles(existTicket, user)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (u *TicketUsecase) prepareTicket(fields []string, ticket, existTicket *internal_models.Ticket, user *models.User) models.Err {
	var err models.Err
	ticket.WasReturned = existTicket.WasReturned
	for _, key := range fields {
		switch key {
		case fieldStatus:
			err = u.changeStatus(ticket, existTicket, user)
		case fieldSupport:
			err = u.changeSupport(ticket)
		case fieldSection:
			err = u.changeSection(ticket)
		case fieldComment:
			continue
		case fieldFiles:
			continue
		}
	}

	return err
}

func (u *TicketUsecase) checkChangeStatus(ticket, existTicket *internal_models.Ticket, user *models.User) models.Err {
	workOnTicket := u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Work)
	resolveTicket := u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Resolve)

	switch ticket.Status.ID {
	case internal_models.TSWaitID:
		if existTicket.Author.ID != user.ID {
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
	case internal_models.TSWaitID:
		if existTicket.Status.ID == internal_models.TSCompletedID || existTicket.Status.ID == internal_models.TSRejectedID {
			ticket.WasReturned = true
		}

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

	ticket.NeedResolve = section.NeedApproval

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

func (u *TicketUsecase) checkChangeFiles(existTicket *internal_models.Ticket, user *models.User) models.Err {
	switch existTicket.Status.ID {
	case internal_models.TSInWorkID, internal_models.TSImplementationID:
		if !u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Work) {
			return errCannotUpdateTicket
		}

	case internal_models.TSRevisionID:
		if existTicket.Author.ID != user.ID {
			return errCannotUpdateTicket
		}

	case internal_models.TSWaitForResolveID:
		if !u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Resolve) {
			return errCannotUpdateTicket
		}
	}

	return nil
}
