package model

import "time"

// Project 表结构体
type Project struct {
	ProjectID      int        `json:"project_id"`
	Title          string     `json:"title"`
	Description    string     `json:"description,omitempty"`
	TeacherID      int        `json:"teacher_id"`
	StudentID      *int       `json:"student_id,omitempty"`
	Status         string     `json:"status"`
	StartDate      time.Time  `json:"start_date"`
	EndDate        *time.Time `json:"end_date,omitempty"`
	ProgressDate   *time.Time `json:"progress_date,omitempty"`
	ProgressDesc   string     `json:"progress_desc,omitempty"`
	ProgressStatus *string    `json:"progress_status,omitempty"`
	Type           string     `json:"type"`
	Major          string     `json:"major"`
}
