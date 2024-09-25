package controller

import transfer "github.com/utf2/utf-statistic-service/internal/controller/model"

type StatisticControllerAPI interface {
	StudentFormResult(transfer.StudentFormResultRequest) transfer.StudentFormResultResponse
}
