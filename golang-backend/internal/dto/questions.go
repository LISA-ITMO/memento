package dto

type UPDQuestionIn struct {
	QuestionID int
	UserID     int
	Text       string
	Answer     string
}

type DelQuestionIn struct {
	QuestionID int
	UserID     int
}
