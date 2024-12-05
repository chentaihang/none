package model

import "time"

type Process struct {
	ProcessID      int       `json:"process_id"`
	ProjectID      int       `json:"project_id"`
	ProgressDate   time.Time `json:"progress_date"`
	ProgressDesc   string    `json:"progress_desc"`
	ProgressStatus string    `json:"progress_status"`
}
