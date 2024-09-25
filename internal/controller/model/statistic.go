package transfer

import "github.com/google/uuid"

type StudentFormResultRequest struct {
	FormID    uuid.UUID
	StudentID uuid.UUID
}

type StudentFormResultResponse struct {
	MaxPossiblePoints uint
	ScoredPoints      uint
}
