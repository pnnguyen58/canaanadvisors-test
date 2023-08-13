package main

import (
	"context"
	"log"
	"os"
	"strings"

	"canaanadvisors-test/core/activities"
	"canaanadvisors-test/core/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	ctx := context.Background()
	//signalChan := make(chan os.Signal, 1)
	//signal.Notify(signalChan)
	//go func() {
	//	// Wait for termination signal
	//	<-signalChan
	//	// Trigger cancellation of the context
	//	cancel()
	//	// Wait for goroutine to finish
	//	fmt.Println("The workers terminated gracefully")
	//}()
	tempoHost := getEnv("TEMPO_HOST", "localhost:7233")
	tempoNS := getEnv("TEMPO_NAMESPACE", "canaanadvisors-test")
	tempoTasks := getEnv("TEMPO_TASK_QUEUE", "canaanadvisors-test-order,canaanadvisors-test-user" +
		",canaanadvisors-test-notification,canaanadvisors-test-management")
	cl, err := client.Dial(client.Options{
		HostPort: tempoHost,
		Namespace: tempoNS,
	})
	defer cl.Close()
	if err != nil {
		log.Fatalln(err)
	}
	taskQueues := strings.Split(tempoTasks, ",")
	register(cl, taskQueues)
	// Wait for a signal to shut down the server
	<-ctx.Done()
}

func register(cl client.Client, taskQueues []string) {
	for _, taskQueueName := range taskQueues {
		switch taskQueueName {
		case "canaanadvisors-test-order":
			w := worker.New(cl, taskQueueName, worker.Options{})
			w.RegisterWorkflow(workflows.CreateOrderWorkflow)
			w.RegisterActivity(activities.CreateOrder)
			w.RegisterActivity(activities.CreateOrderCompensation)
			// TODO: add more workflows and activities
			go func() {
				err := w.Run(worker.InterruptCh())
				if err != nil {
					log.Fatalln("Unable to start worker", err)
				}
			}()
		case "canaanadvisors-test-user":
			w := worker.New(cl, taskQueueName, worker.Options{})
			w.RegisterWorkflow(workflows.LoginWorkflow)
			w.RegisterWorkflow(workflows.LogoutWorkflow)
			w.RegisterActivity(activities.Login)
			w.RegisterActivity(activities.LoginCompensation)
			w.RegisterActivity(activities.Logout)
			w.RegisterActivity(activities.LogoutCompensation)
			// TODO: add more workflows and activities
			go func() {
				err := w.Run(worker.InterruptCh())
				if err != nil {
					log.Fatalln("Unable to start worker", err)
				}
			}()
		case "canaanadvisors-test-notification":
			w := worker.New(cl, taskQueueName, worker.Options{})
			w.RegisterWorkflow(workflows.SendNotificationWorkflow)
			w.RegisterActivity(activities.SendNotification)
			w.RegisterActivity(activities.SendNotificationCompensation)
			// TODO: add more workflows and activities
			go func() {
				err := w.Run(worker.InterruptCh())
				if err != nil {
					log.Fatalln("Unable to start worker", err)
				}
			}()
		case "canaanadvisors-test-management":
			w := worker.New(cl, taskQueueName, worker.Options{})
			w.RegisterWorkflow(workflows.GetMenuWorkflow)
			w.RegisterActivity(activities.GetMenu)
			// TODO: add more workflows and activities
			go func() {
				err := w.Run(worker.InterruptCh())
				if err != nil {
					log.Fatalln("Unable to start worker", err)
				}
			}()
		default:
			log.Println("app not defined")
		}
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}