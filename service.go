package main

import (
	"context"
)

type Service interface {
	Start(ctx context.Context, phoneNumber string) (string, error)
	Stop(ctx context.Context, userId int) (int, error)
}
