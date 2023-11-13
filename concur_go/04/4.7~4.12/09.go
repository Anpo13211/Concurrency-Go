package main

import (
	"context"
	"fmt"
)

func main() {
	ProcessRequest("Johnson", "abc123")
}

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToekn
)

func UserID(c context.Context) string {
	return c.Value(ctxUserID).(string)
}

func AuthToken(c context.Context) string {
	return c.Value(ctxAuthToekn).(string)
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToekn, authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"handling response for %v (auth: %v)\n",
		UserID(ctx),
		AuthToken(ctx),
	)
}
