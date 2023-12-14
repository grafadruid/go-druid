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
		assert.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveVolumes(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	// Wait for druid contaners to start.
	assert.NoError(t, compose.Up(ctx, tc.Wait(true)), "compose.Up()")

	// Set up druid service and client.
	var druidOpts []ClientOption
	d, err := NewClient("http://localhost:8888", druidOpts...)
	assert.NoError(t, err, "error should be nil")
	var spec = NewIngestionSpec(SetType("kafka"),
		SetBrokers("kafka:9092"),
		SetTopic("test-topic"),
		SetDataSource("test-datasource"),
		SetDimensions(DimensionSet{{"dim1"}, {"dim2"}}))
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
	assert.False(t, status.Payload.Suspended)

	// suspend and check status
	suspendedSpec, err := d.Supervisor().Suspend(spec.DataSchema.DataSource)
	assert.True(t, suspendedSpec.Suspended)
	assert.NoError(t, err, "error should be nil")

	status, err = d.Supervisor().GetStatus(spec.DataSchema.DataSource)
	assert.NoError(t, err, "error should be nil")
	assert.True(t, status.Payload.Suspended)

	// resume and check status
	_, err = d.Supervisor().Resume(spec.DataSchema.DataSource)
	assert.NoError(t, err, "error should be nil")

	status, err = d.Supervisor().GetStatus(spec.DataSchema.DataSource)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, "PENDING", status.Payload.State)
	assert.False(t, status.Payload.Suspended)

	// terminate
	id, err = d.Supervisor().Terminate(spec.DataSchema.DataSource)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, id, spec.DataSchema.DataSource)
}
