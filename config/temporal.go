package config

import (
	"context"
	"time"
)

type TempoConfig struct {
	HostPort string `json:"hostPort"`
	Namespace *Namespace `json:"namespace"`
	Workflows map[string]*Workflow `json:"workflows"`
}

type Namespace struct {
	Name string `json:"namespace"`
	WorkflowExecutionRetentionPeriod time.Duration `json:"workflowExecutionRetentionPeriod"` // seconds
}

type Workflow struct {
	TaskQueueName string `json:"taskQueueName"`
	SearchAttributes map[string]interface{} `json:"searchAttributes"`
	ExecutionTimeout time.Duration `json:"executionTimeout"` // seconds
	RunTimeout time.Duration `json:"runTimeout"` // seconds
	TaskTimeout time.Duration `json:"taskTimeout"` // seconds
	ScheduleToCloseTimeout time.Duration `json:"scheduleToCloseTimeout"` // seconds
	StartToCloseTimeout time.Duration `json:"startToCloseTimeout"` // seconds
	HeartbeatTimeout time.Duration `json:"heartbeatTimeout"` // seconds
	WaitForCancellation bool `json:"waitForCancellation"` // seconds
}

func LoadTempoConfig(ctx context.Context) *TempoConfig {
	return mockTempoConfig()
}

func mockTempoConfig() (tc *TempoConfig) {
	tc = &TempoConfig{
		HostPort: C.Server.TempoHost,
		Namespace: &Namespace{
			Name: "canaanadvisors-test",
			WorkflowExecutionRetentionPeriod: 1720*time.Hour,
		},
		Workflows: map[string]*Workflow{},
	}
	tc.Workflows["canaanadvisors-test-order"] = &Workflow{
		TaskQueueName: "canaanadvisors-test-order",
		ExecutionTimeout: 300*time.Second,
		RunTimeout: 300*time.Second,
		TaskTimeout: 300*time.Second,
	}
	tc.Workflows["canaanadvisors-test-auth"] = &Workflow{
		TaskQueueName: "canaanadvisors-test-auth",
		ExecutionTimeout: 300*time.Second,
		RunTimeout: 300*time.Second,
		TaskTimeout: 300*time.Second,
	}
	return
}