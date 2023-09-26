package druid

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestSupervisorService(t *testing.T) {
	// Set up druid containers using docker-compose.
	compose, err := tc.NewDockerCompose("testdata/docker-compose.yaml")
	assert.NoError(t, err, "NewDockerComposeAPI()")

	// Set up cleanup for druid containers.
	t.Cleanup(func() {
		assert.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	// Wait for druid contaners to start.
	assert.NoError(t, compose.Up(ctx, tc.Wait(true)), "compose.Up()")

	// Set up druid service and client.
	var druidOpts []ClientOption
	d, err := NewClient("http://localhost:8081", druidOpts...)
	assert.NoError(t, err, "error should be nil")
	var spec = NewIngestionSpec(SetType("kafka"),
		SetBrokers("telemetry-kafka.skaffold-telemetry-victorzaytsev.svc.cluster.local:9092"),
		SetTopic("test-topic"),
		SetDataSource("test-datasource"))
	assert.NoError(t, err, "error should be nil")
	assert.NotNil(t, spec, "specification should not be nil")

	// Waiting for druid coordinator service to start.
	err = compose.
		WaitForService("coordinator", wait.NewHTTPStrategy(processInformationPathPrefix).WithPort("8081/tcp").WithStartupTimeout(60*time.Second)).
		Up(ctx, tc.Wait(true))
	assert.NoError(t, err, "coordinator should be up with no error")

	// Test create supervisor -> get status -> terminate sequence.
	id, err := d.Supervisor().CreateOrUpdate(spec)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, id, spec.DataSchema.DataSource)
	status, err := d.Supervisor().GetStatus(spec.DataSchema.DataSource)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, "PENDING", status.Payload.State)
	id, err = d.Supervisor().Terminate(spec.DataSchema.DataSource)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, id, spec.DataSchema.DataSource)
}
