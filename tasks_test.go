package druid

import (
	"bytes"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

// testDO represents entry with payload.
type testDO struct {
	Timestamp time.Time      `db:"ts"`
	Id        uuid.UUID      `db:"id"`
	Payload   types.JSONText `db:"payload"`
}

var testObjects = []testDO{
	{
		Id:        uuid.New(),
		Timestamp: time.Now(),
		Payload:   types.JSONText("{\"test\": \"json\"}"),
	},
	{
		Id:        uuid.New(),
		Timestamp: time.Now().Add(time.Hour),
		Payload:   types.JSONText("{\"test\": \"json2\"}"),
	},
}

// triggerIngestionTask initiates inline ingestion task with druid client.
func triggerIngestionTask[T any](d *Client, dataSourceName string, entries []T) (string, error) {
	csvEntriesBuff := &bytes.Buffer{}

	err := gocsv.MarshalWithoutHeaders(entries, csvEntriesBuff)
	if err != nil {
		return "", err
	}

	var spec = NewTaskIngestionSpec(
		SetTaskType("index_parallel"),
		SetTaskTimestampColumn("ts"),
		SetTaskDataSource(dataSourceName),
		SetTaskTuningConfig("index_parallel", 25000, 5000000),
		SetTaskIOConfigType("index_parallel"),
		SetTaskInputFormat("csv", "false", []string{"ts", "id", "payload"}),
		SetTaskInlineInputData(csvEntriesBuff.String()),
	)
	taskID, err := d.Tasks().SubmitTask(spec)
	return taskID, err
}

// awaitTaskCompletion waits for the task to complete. Function timeouts with an error after awaitTimeout nanoseconds.
func awaitTaskCompletion(client *Client, taskID string, awaitTimeout time.Duration, tickerDuration time.Duration) error {
	ticker := time.NewTicker(tickerDuration)
	defer ticker.Stop()
	afterTimeout := time.After(awaitTimeout)
	for {
		select {
		case <-ticker.C:
			res, err := client.Tasks().GetStatus(taskID)
			if err != nil {
				return err
			}

			if res.Status.Status == "RUNNING" {
				continue
			}
			return nil
		case <-afterTimeout:
			return errors.New("AwaitTaskRunning timeout")
		}
	}
}

// awaitTaskStatus waits for the druid task status for the maximum of awaitTimeout duration, querying druid task API.
func awaitTaskStatus(client *Client, taskID string, status string, awaitTimeout time.Duration, tickerDuration time.Duration) error {
	ticker := time.NewTicker(tickerDuration)
	defer ticker.Stop()
	afterTimeout := time.After(awaitTimeout)
	for {
		select {
		case <-ticker.C:
			res, err := client.Tasks().GetStatus(taskID)
			if err != nil {
				return err
			}

			if res.Status.Status == status {
				return nil
			}
		case <-afterTimeout:
			return errors.New("AwaitTaskRunning timeout")
		}
	}
}

// runInlineIngestionTask initiates inline ingestion task with druid client and runs until it is complete.
func runInlineIngestionTask[T any](client *Client, dataSourceName string, entries []T, recordsCount int) error {
	taskID, err := triggerIngestionTask(client, dataSourceName, entries)
	if err != nil {
		return err
	}

	err = awaitTaskCompletion(client, taskID, 180*time.Second, 500*time.Millisecond)
	if err != nil {
		return err
	}

	err = client.metadata(WithMetadataQueryTicker(500*time.Millisecond), WithMetadataQueryTimeout(180*time.Second)).awaitDataSourceAvailable(dataSourceName)
	if err != nil {
		return err
	}

	err = client.metadata().awaitRecordsCount(dataSourceName, recordsCount)
	if err != nil {
		return err
	}

	return nil
}

func TestTaskService(t *testing.T) {
	// Set up druid containers using docker-compose.
	compose, err := tc.NewDockerCompose("testdata/docker-compose.yaml")
	require.NoError(t, err, "NewDockerComposeAPI()")

	// Set up cleanup for druid containers.
	t.Cleanup(func() {
		require.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveVolumes(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up druid service and client.
	d, err := NewClient("http://localhost:8888")
	require.NoError(t, err, "error should be nil")

	// Waiting for druid services to start.
	err = compose.
		WaitForService("coordinator", wait.NewHTTPStrategy(processInformationPathPrefix).WithPort("8081/tcp").WithStartupTimeout(180*time.Second)).
		WaitForService("router", wait.NewHTTPStrategy(processInformationPathPrefix).WithPort("8888/tcp").WithStartupTimeout(180*time.Second)).
		WaitForService("broker", wait.NewHTTPStrategy(processInformationPathPrefix).WithPort("8082/tcp").WithStartupTimeout(180*time.Second)).
		WaitForService("middlemanager", wait.NewHTTPStrategy(processInformationPathPrefix).WithPort("8091/tcp").WithStartupTimeout(180*time.Second)).
		Up(ctx, tc.Wait(true))
	require.NoError(t, err, "druid services should be up with no error")

	// Test create ingestion task -> get status -> complete sequence.
	runInlineIngestionTask(d, "test-submit-task-datasource", testObjects, 2)
	require.NoError(t, err, "error should be nil")
}

func TestTerminateTask(t *testing.T) {
	// Set up druid containers using docker-compose.
	compose, err := tc.NewDockerCompose("testdata/docker-compose.yaml")
	require.NoError(t, err, "NewDockerComposeAPI()")

	// Set up cleanup for druid containers.
	t.Cleanup(func() {
		require.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveVolumes(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up druid service and client.
	d, err := NewClient("http://localhost:8888")
	require.NoError(t, err, "error should be nil")

	// Waiting for druid services to start.
	err = compose.
		WaitForService("coordinator", wait.NewHTTPStrategy(processInformationPathPrefix).WithPort("8081/tcp").WithStartupTimeout(180*time.Second)).
		WaitForService("router", wait.NewHTTPStrategy(processInformationPathPrefix).WithPort("8888/tcp").WithStartupTimeout(180*time.Second)).
		WaitForService("broker", wait.NewHTTPStrategy(processInformationPathPrefix).WithPort("8082/tcp").WithStartupTimeout(180*time.Second)).
		WaitForService("middlemanager", wait.NewHTTPStrategy(processInformationPathPrefix).WithPort("8091/tcp").WithStartupTimeout(180*time.Second)).
		Up(ctx, tc.Wait(true))
	require.NoError(t, err, "druid services should be up with no error")

	// Test create ingestion task -> get status -> terminate sequence.
	taskID, err := triggerIngestionTask(d, "test-terminate-task-datasource", testObjects)
	require.NoError(t, err, "error should be nil")

	err = awaitTaskStatus(d, taskID, "RUNNING", 180*time.Second, 200*time.Millisecond)
	require.NoError(t, err, "error should be nil")

	shutdownTaskID, err := d.Tasks().Shutdown(taskID)
	require.NoError(t, err, "error should be nil")
	require.Equal(t, shutdownTaskID, taskID)
}
