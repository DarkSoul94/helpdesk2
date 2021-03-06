package usecase

import (
	"fmt"
	"time"

	"github.com/DarkSoul94/helpdesk2/cachettl"
	"github.com/DarkSoul94/helpdesk2/global_const/actions"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/comment_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/file_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/reg_fil_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
	"github.com/spf13/viper"
)

type TicketUsecase struct {
	repo      pkg_ticket.ITicketRepo
	catSecUC  cat_sec_manager.ICatSecUsecase
	regFilUC  reg_fil_manager.IRegFilUsecase
	fileUC    file_manager.IFileUsecase
	permUC    group_manager.IPermManager
	userUC    pkg_user.IUserUsecase
	suppUC    pkg_support.ISuppForTicket
	commentUC comment_manager.ICommentUsecase
	store     cachettl.ObjectStore
}

func NewTicketUsecase(
	tRepo pkg_ticket.ITicketRepo,
	catSecUC cat_sec_manager.ICatSecUsecase,
	regFilUC reg_fil_manager.IRegFilUsecase,
	fileUC file_manager.IFileUsecase,
	permUC group_manager.IPermManager,
	userUC pkg_user.IUserUsecase,
	suppUC pkg_support.ISuppForTicket,
	commentUC comment_manager.ICommentUsecase,
) *TicketUsecase {
	return &TicketUsecase{
		repo:      tRepo,
		catSecUC:  catSecUC,
		regFilUC:  regFilUC,
		fileUC:    fileUC,
		permUC:    permUC,
		userUC:    userUC,
		suppUC:    suppUC,
		commentUC: commentUC,
		store:     *cachettl.NewObjectStore(viper.GetDuration("app.ttl_cache.clear_period") * time.Second),
	}
}

func (u *TicketUsecase) CreateCategory(cat *internal_models.Category) (uint64, models.Err) {
	if len(cat.Name) == 0 {
		return 0, models.BadRequest("Имя не должно быть пустым")
	}
	return u.catSecUC.CreateCategory(cat)
}

func (u *TicketUsecase) UpdateCategory(cat *internal_models.Category) models.Err {
	if len(cat.Name) == 0 {
		return models.BadRequest("Имя не должно быть пустым")
	}

	return u.catSecUC.UpdateCategory(cat)
}

func (u *TicketUsecase) CreateCategorySection(sec *internal_models.CategorySection) (uint64, models.Err) {
	if sec.CategoryID == 0 {
		return 0, models.BadRequest("Не указана категория")
	}

	if len(sec.Name) == 0 {
		return 0, models.BadRequest("Имя не должно быть пустым")
	}

	if sec.NeedApproval && len(sec.ApprovalGroups) == 0 {
		return 0, models.BadRequest("Не выбрано ни одной согласовывающей группы")
	}

	return u.catSecUC.CreateCategorySection(sec)
}

func (u *TicketUsecase) UpdateCategorySection(sec *internal_models.CategorySection) models.Err {
	if sec.CategoryID == 0 {
		return models.BadRequest("Не указана категория")
	}

	if len(sec.Name) == 0 {
		return models.BadRequest("Имя не должно быть пустым")
	}

	if sec.NeedApproval && len(sec.ApprovalGroups) == 0 {
		return models.BadRequest("Не выбрано ни одной согласовывающей группы")
	}

	return u.catSecUC.UpdateCategorySection(sec)
}

func (u *TicketUsecase) GetCategorySection(forSearch bool) ([]*internal_models.SectionWithCategory, models.Err) {
	return u.catSecUC.GetCategorySection(forSearch)
}

func (u *TicketUsecase) GetCategorySectionList() ([]internal_models.CategorySectionList, models.Err) {
	return u.catSecUC.GetCategorySectionList()
}

func (u *TicketUsecase) CreateRegion(reg *internal_models.Region) (uint64, models.Err) {
	if len(reg.Name) == 0 {
		return 0, models.BadRequest("Имя региона не должно быть пустым")
	}

	return u.regFilUC.CreateRegion(reg)
}

func (u *TicketUsecase) UpdateRegion(reg *internal_models.Region) models.Err {
	if len(reg.Name) == 0 {
		return models.BadRequest("Имя региона не должно быть пустым")
	}

	return u.regFilUC.UpdateRegion(reg)
}

func (u *TicketUsecase) DeleteRegion(id uint64) models.Err {
	return u.regFilUC.DeleteRegion(id)
}

func (u *TicketUsecase) CreateFilial(fil *internal_models.Filial) (uint64, models.Err) {
	if fil.RegionID == 0 {
		return 0, models.BadRequest("У данного филиала не указан регион")
	}

	if len(fil.Name) == 0 {
		return 0, models.BadRequest("Имя не должно быть пустым")
	}

	return u.regFilUC.CreateFilial(fil)
}

func (u *TicketUsecase) UpdateFilial(fil *internal_models.Filial) models.Err {
	if fil.RegionID == 0 {
		return models.BadRequest("У данного филиала не указан регион")
	}

	if len(fil.Name) == 0 {
		return models.BadRequest("Имя не должно быть пустым")
	}

	return u.regFilUC.UpdateFilial(fil)
}

func (u *TicketUsecase) DeleteFilial(id uint64) models.Err {
	return u.regFilUC.DeleteFilial(id)
}

func (u *TicketUsecase) GetRegionsWithFilials() ([]*internal_models.RegionWithFilials, models.Err) {
	return u.regFilUC.GetRegionsWithFilials()
}

func (u *TicketUsecase) GetTicketStatuses(group_id uint64, all bool) ([]*internal_models.TicketStatus, models.Err) {
	list, err := u.repo.GetTicketStatuses()
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	if !all {
		list = u.ticketStatusFilter(list, group_id)
	}

	return list, nil
}

func (u *TicketUsecase) ticketStatusFilter(list []*internal_models.TicketStatus, group_id uint64) []*internal_models.TicketStatus {
	filteredList := make([]*internal_models.TicketStatus, 0)
	if u.permUC.CheckPermission(group_id, actions.AdminTA) {
		return list
	}

	if u.permUC.CheckPermission(group_id, actions.TicketTA_Work) {
		for _, stat := range list {
			if stat.ID == internal_models.TicketStatusMap[internal_models.KeyTSResolve].ID ||
				stat.ID == internal_models.TicketStatusMap[internal_models.KeyTSWait].ID {
				continue
			} else {
				filteredList = append(filteredList, stat)
			}
		}
	} else if u.permUC.CheckPermission(group_id, actions.TicketTA_Resolve) {
		for _, stat := range list {
			if stat.ID == internal_models.TicketStatusMap[internal_models.KeyTSRevision].ID ||
				stat.ID == internal_models.TicketStatusMap[internal_models.KeyTSRejected].ID {
				filteredList = append(filteredList, stat)
			}
		}
	}

	return filteredList
}

func (u *TicketUsecase) CreateTicketStatusHistory(ticketID, changedUserID uint64, newStatus *internal_models.TicketStatus) models.Err {
	currentTime := time.Now().Truncate(time.Second)

	existHistory, err := u.repo.GetLastTicketStatusHistory(ticketID)
	if err == nil {
		existHistory.Duration = time.Duration(currentTime.Sub(existHistory.SelectTime).Seconds())
		err = u.repo.UpdateTicketStatusHistory(existHistory)
		if err != nil {
			return models.InternalError(err.Error())
		}
	}

	newHistory := new(internal_models.TicketStatusHistory)
	newHistory.New(ticketID, changedUserID, newStatus, currentTime)
	err = u.repo.CreateTicketStatusHistory(newHistory)
	if err != nil {
		return models.InternalError(err.Error())
	}

	return nil
}

func (u *TicketUsecase) GetAllTicketStatusHistory(ticketID uint64, user *models.User) ([]*internal_models.TicketStatusHistory, models.Err) {
	historyList, err := u.repo.GetAllTicketStatusHistory(ticketID)
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	seeAdditionalInfo := u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_SeeAdditionalInfo)
	ticket, _ := u.repo.GetTicket(ticketID)

	for _, history := range historyList {
		changedUser, _ := u.userUC.GetUserByID(history.ChangedUser.ID)

		if !seeAdditionalInfo && history.ChangedUser.ID != ticket.Author.ID && history.ChangedUser.ID != user.ID {
			history.ChangedUser.Name = changedUser.Group.Name
		} else {
			history.ChangedUser.Name = changedUser.Name
		}
	}

	return historyList, nil
}

func (u *TicketUsecase) CreateTicket(ticket *internal_models.Ticket) (uint64, models.Err) {
	var (
		hash, ticketHash string
		err              models.Err
	)

	if len(ticket.Text) == 0 {
		return 0, models.BadRequest("Текст запроса не должен быть пустым")
	}

	if ticket.CatSect, err = u.catSecUC.GetSectionWithCategoryByID(ticket.CatSect.ID); err != nil {
		return 0, err
	}

	ticketHash = ticket.HashСalculation()
	if err := u.store.Get(ticket.Author.Email, &hash); err == nil && hash == ticketHash {
		return 0, nil
	}

	if ticket.CatSect.NeedApproval {
		ticket.Status.Set(internal_models.KeyTSResolve)
		ticket.NeedResolve = true
	} else {
		ticket.Status.Set(internal_models.KeyTSWait)
		ticket.NeedResolve = false
	}

	ticket.Date = time.Now().Truncate(time.Second)

	if mask, err := ticket.IpMask(); err != nil {
		ticket.Filial = "not found"
	} else {
		if fil, reg, err := u.regFilUC.GetFilialByIp(mask); err != nil {
			ticket.Filial = "not found"
		} else {
			ticket.Filial = fmt.Sprintf("%s / %s", fil.Name, reg.Name)
		}
	}

	id, er := u.repo.CreateTicket(ticket)
	if er != nil {
		return 0, models.InternalError(er.Error())
	}

	err = u.CreateTicketStatusHistory(id, ticket.Author.ID, ticket.Status)
	if err != nil {
		return 0, err
	}

	if len(ticket.Files) > 0 {
		err = u.fileUC.CreateFiles(ticket.Files, id)
		if err != nil {
			return 0, err
		}
	}

	u.store.Add(ticket.Author.Email, ticketHash, viper.GetInt64("app.ttl_cache.life_time"))

	return id, nil
}

func (u *TicketUsecase) GenerateTicket(ticketGenerate internal_models.TicketGenerate, author *models.User) models.Err {
	if len(ticketGenerate.Text) == 0 {
		return models.BadRequest("Тест запроса не должен быть пустым")
	}

	if _, err := u.catSecUC.GetCategorySectionByID(ticketGenerate.SectionID); err != nil {
		return models.BadRequest("Указанного раздела категории не существует")
	}

	go func() {
		for _, user := range ticketGenerate.Users {
			ticket := &internal_models.Ticket{
				Date:    time.Now().Truncate(time.Second),
				CatSect: &internal_models.SectionWithCategory{ID: ticketGenerate.SectionID},
				Text:    ticketGenerate.Text,
				Status:  &internal_models.TicketStatus{ID: internal_models.TSCompletedID},
				Author:  author,
				Support: &models.User{ID: user.UserID},
			}

			for i := 0; i < user.Count; i++ {
				id, err := u.repo.CreateTicket(ticket)
				if err != nil {
					continue
				}

				u.CreateTicketStatusHistory(id, author.ID, ticket.Status)
			}
		}
	}()

	return nil
}

func (u *TicketUsecase) UpdateTicket(ticket *internal_models.Ticket, user *models.User, checkControls bool) models.Err {
	if checkControls {
		existTicket, _ := u.repo.GetTicket(ticket.ID)
		existTicket.CatSect, _ = u.catSecUC.GetSectionWithCategoryByID(existTicket.CatSect.ID)
		err := u.prepareBeforeUpdate(ticket, existTicket, user)
		if err != nil {
			return err
		}
	}

	if err := u.repo.UpdateTicket(ticket); err != nil {
		return models.InternalError(err.Error())
	}

	if err := u.CreateTicketStatusHistory(ticket.ID, user.ID, ticket.Status); err != nil {
		return models.InternalError(err.Error())
	}

	if err := u.fileUC.CreateFiles(ticket.Files, ticket.ID); err != nil {
		return err
	}

	return nil
}

func (u *TicketUsecase) GetTicketList(user *models.User, limit, offset int) ([]*internal_models.Ticket, []string, map[uint]uint, models.Err) {
	var (
		list     []*internal_models.Ticket
		priority map[uint]uint
		err      error
	)

	if u.permUC.CheckPermission(user.Group.ID, actions.AdminTA) {
		list, err = u.repo.GetTicketListForAdmin(limit, offset)
		priority = nil
	} else if u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Work) {
		list, err = u.repo.GetTicketListForSupport(user.ID, limit, offset)
		priority = u.repo.GetTicketStatusesSortPriority(true)
	} else {
		list, err = u.repo.GetTicketListForUser(user.ID, limit, offset)
		priority = u.repo.GetTicketStatusesSortPriority(false)
	}
	if err != nil {
		return nil, nil, nil, models.InternalError(err.Error())
	}

	if err := u.loadAddInfoForTicket(list, user); err != nil {
		return nil, nil, nil, err
	}

	return list, u.makeTagList(user.Group.ID), priority, nil
}

func (u *TicketUsecase) loadAddInfoForTicket(ticketList []*internal_models.Ticket, user *models.User) models.Err {
	var err models.Err
	for _, ticket := range ticketList {
		ticket.CatSect, err = u.catSecUC.GetSectionWithCategoryByID(ticket.CatSect.ID)
		if err != nil {
			return err
		}

		if ticket.Author.ID != 0 {
			ticket.Author, err = u.userUC.GetUserByID(ticket.Author.ID)
			if err != nil {
				return err
			}
		}

		if ticket.Support.ID != 0 && u.permUC.CheckPermission(user.Group.ID, actions.AdminTA) {
			ticket.Support, err = u.userUC.GetUserByID(ticket.Support.ID)
			if err != nil {
				return err
			}
		} else {
			ticket.Support = nil
		}

		if ticket.ResolvedUser.ID != 0 {
			ticket.ResolvedUser, err = u.userUC.GetUserByID(ticket.ResolvedUser.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (u *TicketUsecase) makeTagList(groupID uint64) []string {
	tags := []string{
		"ticket_id",
		"ticket_date",
		"category",
		"section",
		"ticket_text",
		"status",
		"filial",
		"ticket_author",
	}

	if u.permUC.CheckPermission(groupID, actions.AdminTA) {
		tags = append(tags, "support")
	}
	if !u.permUC.CheckPermission(groupID, actions.TicketTA_Resolve) {
		tags = append(tags, "grade")
	}

	return tags
}

func (u *TicketUsecase) GetFilteredTicketsList(filter map[string]interface{}, user *models.User) ([]*internal_models.Ticket, []string, models.Err) {
	fullSearch := u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_FullSearch)

	for key, value := range filter {
		v := value
		switch v.(type) {
		case string:
			if len(value.(string)) == 0 {
				delete(filter, key)
			}
		case float64:
			if value.(float64) == 0 {
				delete(filter, key)
			}
		case []interface{}:
			if len(value.([]interface{})) == 0 {
				delete(filter, key)
			}
		default:
			delete(filter, key)
		}
	}

	if len(filter) == 0 {
		return nil, nil, models.BadRequest("Не выбрано ни одного параметра для поиска")
	}

	if !fullSearch {
		if u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Resolve) {
			delete(filter, pkg_ticket.SupportID)
			delete(filter, pkg_ticket.AuthorID)
			filter[pkg_ticket.AuthorAndResolve] = user.ID

			if filter[pkg_ticket.TicketID] != nil {
				filter[pkg_ticket.TicketIDBACK] = filter[pkg_ticket.TicketID]
				delete(filter, pkg_ticket.TicketID)
			}
		} else {
			delete(filter, pkg_ticket.SupportID)
			filter[pkg_ticket.AuthorID] = user.ID
			if filter[pkg_ticket.TicketID] != nil {
				filter[pkg_ticket.TicketIDREG] = filter[pkg_ticket.TicketID]
				delete(filter, pkg_ticket.TicketID)
			}
		}
	} else {
		if filter[pkg_ticket.TicketID] != nil {
			filter[pkg_ticket.TicketIDAll] = filter[pkg_ticket.TicketID]
			delete(filter, pkg_ticket.TicketID)
		}
	}

	query, args := u.repo.FilterDispatcher(filter)

	ticketsList, err := u.repo.GetFilteredTicketsList(query, args, fullSearch)
	if err != nil {
		return nil, nil, models.InternalError(err.Error())
	}

	if err := u.loadAddInfoForTicket(ticketsList, user); err != nil {
		return nil, nil, err
	}

	return ticketsList, u.makeTagList(user.Group.ID), nil
}

func (u *TicketUsecase) GetTicket(ticketID uint64, user *models.User) (*internal_models.Ticket, models.Err) {
	seeAdditionalInfo := u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_SeeAdditionalInfo)

	ticket, err := u.repo.GetTicket(ticketID)
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	ticket.CatSect, err = u.catSecUC.GetSectionWithCategoryByID(ticket.CatSect.ID)
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	ticket.Author, err = u.userUC.GetUserByID(ticket.Author.ID)
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	ticket.Files, err = u.fileUC.GetTicketFiles(ticketID)
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	if comments, err := u.commentUC.GetTicketComments(ticket.ID); err != nil {
		return nil, err
	} else {
		for _, comment := range comments {
			author, _ := u.userUC.GetUserByID(comment.Author.ID)

			if !seeAdditionalInfo && comment.Author.ID != ticket.Author.ID && comment.Author.ID != user.ID {
				comment.Author.Name = author.Group.Name
			} else {
				comment.Author.Name = author.Name
			}
			ticket.Comments = append(ticket.Comments, comment)
		}
	}

	if seeAdditionalInfo {
		if ticket.Support.ID != 0 {
			ticket.Support, err = u.userUC.GetUserByID(ticket.Support.ID)
			if err != nil {
				return nil, models.InternalError(err.Error())
			}
		}

		if ticket.ResolvedUser.ID != 0 {
			ticket.ResolvedUser, err = u.userUC.GetUserByID(ticket.ResolvedUser.ID)
			if err != nil {
				return nil, models.InternalError(err.Error())
			}
		}

		if u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Resolve) &&
			!(u.permUC.CheckPermission(user.Group.ID, actions.AdminTA) ||
				u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Work)) {
			ticket.ServiceComment = ""
		}
	}

	return ticket, nil
}

func (u *TicketUsecase) StealTicket(ticketID uint64, newSupport *models.User) models.Err {
	existTicket, err := u.repo.GetTicket(ticketID)
	if err != nil {
		return models.BadRequest("Запроса с таким id не существует")
	}

	if existTicket.Support.ID == newSupport.ID {
		return models.BadRequest("Запрос уже закреплен за вами")
	}

	switch existTicket.Status.ID {
	case internal_models.TSCompletedID,
		internal_models.TSRejectedID:
		return models.InternalError("Невозможно забрать запрос")

	case internal_models.TSWaitID, internal_models.TSWaitForResolveID:
		err := u.repo.StealTicket(ticketID, newSupport.ID, true)
		if err != nil {
			return models.InternalError(err.Error())
		}

		err = u.suppUC.UpdateSupportActivity(newSupport.ID, ticketID)
		if err != nil {
			return models.InternalError(err.Error())
		}

	default:
		err := u.repo.StealTicket(ticketID, newSupport.ID, false)
		if err != nil {
			return models.InternalError(err.Error())
		}

		err = u.suppUC.UpdateSupportActivity(newSupport.ID, ticketID)
		if err != nil {
			return models.InternalError(err.Error())
		}
	}
	return nil
}

func (u *TicketUsecase) TicketGrade(ticketID uint64, grade uint, user *models.User) models.Err {
	ticket, err := u.repo.GetTicket(ticketID)
	if err != nil {
		return models.BadRequest("Запроса с таким ид не найдено")
	}

	if ticket.Grade != 0 {
		return models.BadRequest("Запрос уже оценен")
	}

	if ticket.Author.ID != user.ID {
		return models.Forbidden("Запрос может оценить только автор")
	}

	if ticket.Status.ID != internal_models.TSCompletedID {
		return models.BadRequest("Оценивать можно только выполненые запросы")
	}

	err = u.repo.TicketGrade(ticketID, user.ID, grade)
	if err != nil {
		return models.InternalError(err.Error())
	}

	return nil
}

func (u *TicketUsecase) CheckNeedApprovalTicketExist(groupID uint64) bool {
	var (
		exist bool
		err   error
	)

	if u.permUC.CheckPermission(groupID, actions.AdminTA) || u.permUC.CheckPermission(groupID, actions.TicketTA_Work) {
		exist, err = u.repo.CheckNeedApprovalTicketExist(groupID, false)

	} else if u.permUC.CheckPermission(groupID, actions.TicketTA_Resolve) {
		exist, err = u.repo.CheckNeedApprovalTicketExist(groupID, true)
	}
	if err != nil {
		return false
	}

	return exist
}

func (u *TicketUsecase) GetApprovalTicketList(groupID uint64, limit, offset int) ([]*internal_models.Ticket, []string, models.Err) {
	var (
		list []*internal_models.Ticket
		err  error
	)
	if u.permUC.CheckPermission(groupID, actions.AdminTA) || u.permUC.CheckPermission(groupID, actions.TicketTA_Work) {
		list, err = u.repo.GetTicketListForApproval(groupID, limit, offset, false)

	} else if u.permUC.CheckPermission(groupID, actions.TicketTA_Resolve) {
		list, err = u.repo.GetTicketListForApproval(groupID, limit, offset, true)
	}

	if err != nil {
		return nil, nil, models.InternalError(err.Error())
	}

	for _, ticket := range list {
		ticket.CatSect, err = u.catSecUC.GetSectionWithCategoryByID(ticket.CatSect.ID)
		if err != nil {
			return nil, nil, models.InternalError(err.Error())
		}

		if ticket.Author != nil {
			ticket.Author, err = u.userUC.GetUserByID(ticket.Author.ID)
			if err != nil {
				return nil, nil, models.InternalError(err.Error())
			}
		}

		ticket.Support = nil
	}

	return list, u.makeTagList(groupID), nil
}

func (u *TicketUsecase) ResolveTicket(ticketID uint64, user *models.User) models.Err {
	if !u.permUC.CheckPermission(user.Group.ID, actions.TicketTA_Resolve) {
		return models.Forbidden("У вас нет прав согласовывать запрос")
	}

	ticket, err := u.repo.GetTicket(ticketID)
	if err != nil {
		return models.InternalError(err.Error())
	}

	if !ticket.NeedResolve {
		return models.BadRequest("Запрос не нуждается в согласовании")
	}

	if !u.catSecUC.CheckExistInResolveGroupList(ticket.CatSect.ID, user.Group.ID) {
		return models.Forbidden("У вас нет прав согласовывать запрос данной категории")
	}

	if ticket.ResolvedUser != nil && ticket.ResolvedUser.ID != 0 {
		return models.BadRequest("Запрос уже согласован")
	}

	ticket.ResolvedUser = user
	ticket.NeedResolve = false
	if ticket.Status.ID == internal_models.TSWaitForResolveID {
		ticket.Status.ID = internal_models.TSWaitID
	}

	return u.UpdateTicket(ticket, user, false)
}

func (u *TicketUsecase) CreateComment(comment *internal_models.Comment) (uint64, models.Err) {
	if err := u.createCommentDispatcher(comment); err != nil {
		return 0, err
	}

	id, err := u.commentUC.CreateComment(comment)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *TicketUsecase) createCommentDispatcher(comment *internal_models.Comment) models.Err {
	type hDate struct {
		year  int
		month time.Month
		day   int
	}
	var nowDate, completeDate hDate

	if u.permUC.CheckPermission(comment.Author.Group.ID, actions.AdminTA) {
		return nil
	}

	ticket, _ := u.repo.GetTicket(comment.TicketId)

	if u.permUC.CheckPermission(comment.Author.Group.ID, actions.TicketTA_Resolve) {
		switch ticket.Status.ID {
		case internal_models.TSWaitForResolveID:
			return nil
		default:
			return errCannotUpdateTicket
		}
	}

	if comment.Author.ID == ticket.Author.ID {
		switch ticket.Status.ID {
		case internal_models.TSCompletedID, internal_models.TSRejectedID:
			history, _ := u.repo.GetLastTicketStatusHistory(comment.TicketId)

			completeDate.year, completeDate.month, completeDate.day = history.SelectTime.Local().Date()
			nowDate.year, nowDate.month, nowDate.day = time.Now().Local().Date()
			if nowDate != completeDate {
				return models.InternalError("Данный запрос нельзя комментировать")
			}
			return nil

		case internal_models.TSRevisionID:
			return nil

		default:
			return errCannotUpdateTicket
		}
	}

	if comment.Author.ID == ticket.Support.ID {
		switch ticket.Status.ID {
		case internal_models.TSInWorkID,
			internal_models.TSImplementationID,
			internal_models.TSPostponedID:
			return nil
		default:
			return errCannotUpdateTicket
		}
	}

	return errCannotUpdateTicket
}

func (u *TicketUsecase) GetFile(fileID uint64) (*internal_models.File, models.Err) {
	return u.fileUC.GetFile(fileID)
}

func (u *TicketUsecase) AutoCreateTicket(text, email, ip string, priority bool) (uint64, models.Err) {
	var (
		author *models.User = new(models.User)
		err    models.Err
	)

	if len(email) > 0 {
		author, err = u.userUC.GetUserByEmail(email)
		if err != nil {
			author = &models.User{
				Email: email,
				Name:  email,
			}

			author.ID, err = u.userUC.CreateUser(author)
			if err != nil {
				return 0, models.InternalError("Не удалось создать пользователя")
			}
		}
	} else {
		return 0, models.BadRequest("Email пользователя пустой")
	}

	sectionID, err := u.catSecUC.GetServiceSectionID(priority)
	if err != nil {
		return 0, err
	}

	ticket := &internal_models.Ticket{
		Text:    text,
		CatSect: &internal_models.SectionWithCategory{ID: sectionID},
		Status:  &internal_models.TicketStatus{},
		IP:      ip,
		Author:  author,
	}

	return u.CreateTicket(ticket)
}
