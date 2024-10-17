package entities

type Tasks struct {
	ID        int64  `json:"id"`
	TaskName  string `json:"task_name"`
	TaskDesc  string `json:"task_desc"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}
