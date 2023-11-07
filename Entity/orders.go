package entity

type Orders struct {
	Uuid       string `json:"uuid"`
	Contents   []Food `json:"contents"`
	Completion []bool `json:"completion"`
}
