package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/lithammer/shortuuid/v3"
)

type journeyService struct{}

func NewService() Service { return &journeyService{} }

// Start daily quiz (and create a new user, if necessary)
func (w *journeyService) Start(_ context.Context, phoneNumber string) (string, error) {
	newUserID := shortuuid.New()
	return newUserID, nil
}

// Stop daily quiz
func (w *journeyService) Stop(_ context.Context, userId int) (int, error) {
	logger.Log("Checking the Service health...")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
