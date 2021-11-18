package internal_models

type CategorySection struct {
	ID             uint64
	CategoryID     uint64
	Name           string
	Significant    bool
	Old            bool
	NeedApproval   bool
	Template       string
	ApprovalGroups []uint64
}

type SectionWithCategory struct {
	ID           uint64
	Name         string
	Significant  bool
	Old          bool
	NeedApproval bool
	Template     string
	Cat          *Category
}

type CategorySectionList struct {
	Category *Category
	Sections []*CategorySection
}
