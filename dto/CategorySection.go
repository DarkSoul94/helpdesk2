package dto

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

type OutCategorySection struct {
	ID             uint64   `json:"section_id,omitempty"`
	CategoryID     uint64   `json:"category_id"`
	Name           string   `json:"section_name"`
	Significant    bool     `json:"significant"`
	Old            bool     `json:"old_category"`
	NeedApproval   bool     `json:"need_approval"`
	Template       string   `json:"template"`
	ApprovalGroups []uint64 `json:"approval_groups"`
}

type OutSectionWithCategory struct {
	ID           uint64      `json:"section_id"`
	Name         string      `json:"section_name"`
	Significant  bool        `json:"significant"`
	Old          bool        `json:"old_category"`
	NeedApproval bool        `json:"need_approval"`
	Template     string      `json:"template"`
	Cat          InpCategory `json:"category"`
}

func ToOutCategorySection(sec *internal_models.CategorySection) OutCategorySection {
	return OutCategorySection{
		ID:             sec.ID,
		CategoryID:     sec.CategoryID,
		Name:           sec.Name,
		Significant:    sec.Significant,
		Old:            sec.Old,
		NeedApproval:   sec.NeedApproval,
		Template:       sec.Template,
		ApprovalGroups: sec.ApprovalGroups,
	}
}

func ToModelCategorySection(inpSec OutCategorySection) *internal_models.CategorySection {
	return &internal_models.CategorySection{
		ID:             inpSec.ID,
		CategoryID:     inpSec.CategoryID,
		Name:           inpSec.Name,
		Significant:    inpSec.Significant,
		Old:            inpSec.Old,
		NeedApproval:   inpSec.NeedApproval,
		Template:       inpSec.Template,
		ApprovalGroups: inpSec.ApprovalGroups,
	}
}

func ToOutSectionWithCategory(sec *internal_models.SectionWithCategory) OutSectionWithCategory {
	return OutSectionWithCategory{
		ID:           sec.ID,
		Name:         sec.Name,
		Significant:  sec.Significant,
		Old:          sec.Old,
		NeedApproval: sec.NeedApproval,
		Template:     sec.Template,
		Cat:          ToOutCategory(sec.Cat),
	}
}
