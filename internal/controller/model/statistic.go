package transfer

import "github.com/google/uuid"

type StudentFormResultRequest struct {
	FormID    uuid.UUID
	StudentID uuid.UUID
}

// Should be cached in Redis.
type StudentFormResultResponse struct {
	MaxPossiblePoints uint
	ScoredPoints      uint
}

type FormGeneralResultRequest struct {
	FormID uuid.UUID
}

type FormGeneralResultResponse struct {
	GroupResults      []GroupResultDTO
	MaxPossiblePoints uint
}

type GroupResultDTO struct {
	Group          GroupDTO
	StudentResults []StudentResultDTO
}

type GroupDTO struct {
	ID                 uuid.UUID
	SpecializationCode string
	GroupNumber        string
}

type StudentResultDTO struct {
	Student         StudentDTO
	ScoredPoints    uint
	QuestionResults []QuestionResultDTO
}

type StudentDTO struct {
	ID               uuid.UUID
	FirstName        string
	LastName         string
	MiddleName       string
	EducationalEmail string
	Username         string
}

type QuestionResultDTO struct {
	QuestionID   uuid.UUID
	ChosenAnswer AnswerDTO
}

type AnswerDTO struct {
	ID      uuid.UUID
	Text    string
	Correct bool
}
