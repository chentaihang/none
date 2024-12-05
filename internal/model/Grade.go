package model

import "time"

type Grade struct {
	GradeID     int       `json:"grade_id"`
	ProjectID   int       `json:"project_id"`
	StudentID   int       `json:"student_id"`
	TeacherID   int       `json:"teacher_id"`
	Score       float64   `json:"score"`
	GradingDate time.Time `json:"grading_date"`
	Comments    string    `json:"comments,omitempty"`
}
