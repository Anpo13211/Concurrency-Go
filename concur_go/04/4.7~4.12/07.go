package main

import (
	"context"
	"fmt"
)

func main() {
	ProcessRequest("Johnson", "abc13211")
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), "userID", userID)
	ctx = context.WithValue(ctx, "authToken", authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"handling response for %v (%v)\n",
		ctx.Value("userID"),
		ctx.Value("authToken"),
	)
}
