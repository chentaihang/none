package model

type Teacher struct {
	TeacherID  int     `json:"teacher_id" gorm:"primaryKey"`
	Name       string  `json:"name"`
	Gender     string  `json:"gender"`
	Department string  `json:"department"`
	Email      *string `json:"email,omitempty"`
	UserID     int     `json:"user_id"`
}
