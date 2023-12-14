package druid

import "strings"

const (
	tasksEndpoint             = "druid/indexer/v1/tasks"
	taskStatusEndpoint        = "druid/indexer/v1/task/:taskId/status"
	tasksCompleteEndpoint     = "druid/indexer/v1/completeTasks"
	tasksRunningEndpoint      = "druid/indexer/v1/runningTasks"
	tasksWaitingEndpoint      = "druid/indexer/v1/waitingTasks"
	tasksPendingEndpoint      = "druid/indexer/v1/pendingTasks"
	taskPayloadEndpoint       = "druid/indexer/v1/task/:taskId"
	taskSegmentsEndpoint      = "druid/indexer/v1/task/:taskId/segments"
	taskReportEndpoint        = "druid/indexer/v1/task/:taskId/reports"
	taskSubmitEndpoint        = "druid/indexer/v1/task"
	taskShutdownEndpoint      = "druid/indexer/v1/task/:taskId/shutdown"
	tasksShutdownAllEndpoint  = "druid/indexer/v1/datasources/:datasource/shutdownAllTasks"
	tasksStatusesEndpoint     = "druid/indexer/v1/taskStatus"
	taskDeletePendingSegments = "druid/indexer/v1/pendingSegments/:datasource"
)

// TasksService is a service that runs requests to druid tasks API.
type TasksService struct {
	client *Client
}

// SubmitTaskResponse is a response object of Druid Task API Submit task method.
type SubmitTaskResponse struct {
	Task string `json:"task"`
}

// ShutdownTaskResponse is a response object of Druid SupervisorService's Terminate method.
type ShutdownTaskResponse struct {
	Task string `json:"task"`
}

// SubmitTask submits an ingestion specification to druid tasks API with a pre-configured druid client.
// https://druid.apache.org/docs/latest/api-reference/tasks-api/#submit-a-task
func (s *TasksService) SubmitTask(spec *TaskIngestionSpec) (string, error) {
	r, err := s.client.NewRequest("POST", taskSubmitEndpoint, spec)
	if err != nil {
		return "", err
	}
	var result SubmitTaskResponse
	_, err = s.client.Do(r, &result)
	if err != nil {
		return "", err
	}
	return result.Task, nil
}

// GetStatus calls druid tasks service's Get status API.
// https://druid.apache.org/docs/latest/api-reference/tasks-api/#get-task-status
func (s *TasksService) GetStatus(taskId string) (TaskStatusResponse, error) {
	r, err := s.client.NewRequest("GET", applyTaskId(taskStatusEndpoint, taskId), nil)
	var result TaskStatusResponse
	if err != nil {
		return result, err
	}
	_, err = s.client.Do(r, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Shutdown calls druid task service's shutdown task API.
// https://druid.apache.org/docs/latest/api-reference/tasks-api/#shut-down-a-task
func (s *TasksService) Shutdown(taskId string) (string, error) {
	r, err := s.client.NewRequest("POST", applyTaskId(taskShutdownEndpoint, taskId), "")
	var result ShutdownTaskResponse
	if err != nil {
		return "", err
	}
	_, err = s.client.Do(r, &result)
	if err != nil {
		return result.Task, err
	}
	return result.Task, nil
}

func applyTaskId(input string, taskId string) string {
	return strings.Replace(input, ":taskId", taskId, 1)
}
