package workspace

type Workspace struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Logo        string `json:"logo"`

	CreatedBy uint64
	UpdateBy  uint64
}
