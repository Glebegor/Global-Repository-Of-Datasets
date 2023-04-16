package grod

type Dataset struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Id          int    `json:"id" db:"id"`
}
type UpdateDataset struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
type DatasetItem struct {
	Datainfo string `json:"datainfo" binding:"required"`
	Solution string `json:"solution" binding:"required"`
	Id       int    `json:"id" db:"id"`
}
type UpdateDatasetItem struct {
	Datainfo string `json:"datainfo"`
	Solution string `json:"solution"`
}
