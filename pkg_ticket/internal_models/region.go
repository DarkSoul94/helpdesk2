package internal_models

type Region struct {
	ID   uint64
	Name string
}

type RegionWithFilials struct {
	Region  *Region
	Filials []*Filial
}
