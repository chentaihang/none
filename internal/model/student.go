package model

// Student 表结构体
type Student struct {
	StudentID int     `json:"student_id" gorm:"primaryKey"`
	Name      string  `json:"name"`
	Gender    string  `json:"gender"`
	Major     string  `json:"major"`
	Class     string  `json:"class"`
	Contact   *string `json:"contact,omitempty"`
	Email     *string `json:"email,omitempty"`
	AdvisorID *int    `json:"advisor_id,omitempty"`
	UserID    int     `json:"user_id"`
}
