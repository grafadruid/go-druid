package druid

import (
	"errors"
	"strings"
)

const (
	supervisorEndpoint             = "druid/indexer/v1/supervisor"
	supervisorSpecEndpoint         = "druid/indexer/v1/supervisor/:supervisorId"
	supervisorAllActiveEndpoint    = "druid/indexer/v1/supervisor?full"
	supervisorStatusEndpoint       = "druid/indexer/v1/supervisor/:supervisorId/status"
	supervisorHistoryAllEndpoint   = "druid/indexer/v1/supervisor/history"
	supervisorHistoryEndpoint      = "druid/indexer/v1/supervisor/:supervisorId/history"
	supervisorSuspendEndpoint      = "druid/indexer/v1/supervisor/:supervisorId/suspend"
	supervisorSuspendAllEndpoint   = "druid/indexer/v1/supervisor/suspendAll"
	supervisorResumeEndpoint       = "druid/indexer/v1/supervisor/:supervisorId/resume"
	supervisorResumeAllEndpoint    = "druid/indexer/v1/supervisor/resumeAll"
	supervisorResetEndpoint        = "druid/indexer/v1/supervisor/:supervisorId/reset"
	supervisorTerminateEndpoint    = "druid/indexer/v1/supervisor/:supervisorId/terminate"
	supervisorTerminateAllEndpoint = "druid/indexer/v1/supervisor/terminateAll"
	supervisorShutdownEndpoint     = "druid/indexer/v1/supervisor/:supervisorId/shutdown"
)

// SupervisorService is a service that submits ingestion tasks to druid supervisor API.
type SupervisorService struct {
	client *Client
}

// CreateSupervisorResponse is a response object of Druid SupervisorService's SubmitTask method.
type CreateSupervisorResponse struct {
	SupervisorID string `json:"id"`
}

// TerminateSupervisorResponse is a response object of Druid SupervisorService's Terminate method.
type TerminateSupervisorResponse struct {
	SupervisorID string `json:"id"`
}

// CreateOrUpdate submits an ingestion specification to druid Supervisor API with a pre-configured druid client.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#create-or-update-a-supervisor
func (s *SupervisorService) CreateOrUpdate(spec *InputIngestionSpec) (string, error) {
	r, err := s.client.NewRequest("POST", supervisorEndpoint, spec)
	if err != nil {
		return "", err
	}
	var result CreateSupervisorResponse
	_, err = s.client.Do(r, &result)
	if err != nil {
		return "", err
	}
	return result.SupervisorID, nil
}

func applySupervisorId(input string, supervisorId string) string {
	return strings.Replace(input, ":supervisorId", supervisorId, 1)
}

// GetActiveIDs returns array of active supervisor IDs.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#get-an-array-of-active-supervisor-ids
func (s *SupervisorService) GetActiveIDs() ([]string, error) {
	return nil, errors.New("method GetActiveIDs not implemented")
}

// GetAllActiveSupervisors returns array of active supervisor objects.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#get-an-array-of-active-supervisor-objects
func (s *SupervisorService) GetAllActiveSupervisors() ([]SupervisorStateWithSpec, error) {
	return nil, errors.New("method GetAllActiveSpecs not implemented")
}

// GetAllActiveStates returns an array of supervisor states.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#get-an-array-of-supervisor-states
func (s *SupervisorService) GetActiveStates() ([]SupervisorState, error) {
	return nil, errors.New("method GetAllActiveStates not implemented")
}

// GetSpec calls druid Supervisor Status API.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#get-supervisor-specification
func (s *SupervisorService) GetSpec(supervisorId string) (OutputIngestionSpec, error) {
	r, err := s.client.NewRequest("GET", applySupervisorId(supervisorSpecEndpoint, supervisorId), nil)
	var result OutputIngestionSpec
	if err != nil {
		return result, err
	}
	_, err = s.client.Do(r, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetStatus calls druid Supervisor service's Get status API.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#get-supervisor-status
func (s *SupervisorService) GetStatus(supervisorId string) (SupervisorStatus, error) {
	r, err := s.client.NewRequest("GET", applySupervisorId(supervisorStatusEndpoint, supervisorId), nil)
	var result SupervisorStatus
	if err != nil {
		return result, err
	}
	_, err = s.client.Do(r, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetAuditHistory calls druid Supervisor service's AuditHistory API.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#get-audit-history-for-a-specific-supervisor
func (s *SupervisorService) GetAuditHistory(string) (SupervisorAuditHistory, error) {
	var res SupervisorAuditHistory
	return res, errors.New("method GetAuditHistory not implemented")
}

// GetAuditHistoryAll calls druid Supervisor service's AuditHistoryAll API.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#get-audit-history-for-all-supervisors
func (s *SupervisorService) GetAuditHistoryAll() (map[string]SupervisorAuditHistory, error) {
	return nil, errors.New("method GetAuditHistoryAll not implemented")
}

// Suspend calls druid Supervisor service's Suspend API.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#suspend-a-running-supervisor
func (s *SupervisorService) Suspend(string) (OutputIngestionSpec, error) {
	var res OutputIngestionSpec
	return res, errors.New("method Suspend not implemented")
}

// SuspendAll calls druid Supervisor service's SuspendAll API.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#suspend-all-supervisors
func (s *SupervisorService) SuspendAll() (string, error) {
	return "", errors.New("method SuspendAll not implemented")
}

// Resume calls druid Supervisor service's Resume API.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#resume-a-supervisor
func (s *SupervisorService) Resume(string) (OutputIngestionSpec, error) {
	var res OutputIngestionSpec
	return res, errors.New("method Resume not implemented")
}

// ResumeAll calls druid Supervisor service's ResumeAll API.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#resume-all-supervisors
func (s *SupervisorService) ResumeAll() (string, error) {
	return "", errors.New("method ResumeAll not implemented")
}

// Reset calls druid Supervisor service's Reset API.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#reset-a-supervisor
func (s *SupervisorService) Reset(string) (string, error) {
	return "", errors.New("method Reset not implemented")
}

// Terminate calls druid Supervisor service's Terminate API.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#terminate-a-supervisor
func (s *SupervisorService) Terminate(supervisorId string) (string, error) {
	r, err := s.client.NewRequest("POST", applySupervisorId(supervisorTerminateEndpoint, supervisorId), "")
	if err != nil {
		return "", err
	}
	var result TerminateSupervisorResponse
	_, err = s.client.Do(r, &result)
	if err != nil {
		return "", err
	}
	return result.SupervisorID, nil
}

// TerminateAll calls druid Supervisor service's TerminateAll API.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#terminate-all-supervisors
func (s *SupervisorService) TerminateAll() (string, error) {
	return "", errors.New("method TerminateAll not implemented")
}

// Shutdown calls druid Supervisor service's Shutdown API.
// https://druid.apache.org/docs/latest/api-reference/supervisor-api/#shut-down-a-supervisor
func (s *SupervisorService) Shutdown(string) (string, error) {
	return "", errors.New("method Shutdown not implemented")
}
