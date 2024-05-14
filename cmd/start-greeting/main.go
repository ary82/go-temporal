package main

import (
	"context"
	"log"
	"time"

	"github.com/ary82/go-temporal/app"
	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: "GREETING_TASK_QUEUE",

		// Every minute
		CronSchedule: "* * * * *",

		// Expire after 3 minutes
		WorkflowExecutionTimeout: 3 * time.Minute,
	}

	// Start the Workflow
	name := "World"
	_, err = c.ExecuteWorkflow(context.Background(), options, app.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
}
