package website

type portfolio struct {
	ID     string
	Albums []album
}

type album struct {
	ID          string
	PortfolioID string
	CoverID     string
	Photos      []photo
}

type photo struct {
	ID    string
	Album album
}
