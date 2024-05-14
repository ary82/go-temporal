package app

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	var err error

	for i := 0; i < 10; i++ {
		err = workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
	}

	return result, err
}

func ComposeGreeting(ctx context.Context, name string) (string, error) {
	n := rand.Intn(5)
	if n >= 3 {
		return "", fmt.Errorf("example error")
	}

	greeting := fmt.Sprintf("Hello %s!", name)
	return greeting, nil
}
