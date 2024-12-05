package model

import "time"

type FailedSubmission struct {
	SubmissionID   int       `json:"submission_id"`
	ProjectID      int       `json:"project_id"`
	SubmissionDate time.Time `json:"submission_date"`
	Reason         string    `json:"reason"`
}
