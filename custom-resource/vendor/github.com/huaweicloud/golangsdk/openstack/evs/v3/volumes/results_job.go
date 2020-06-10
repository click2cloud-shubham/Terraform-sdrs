package volumes

import (
	"fmt"
	"strings"

	"github.com/huaweicloud/golangsdk"
)

type JobResponse struct {
	JobID string `json:"job_id"`
}

type JobStatus struct {
	Status     string      `json:"status"`
	Entities   JobEntity   `json:"entities"`
	JobID      string      `json:"job_id"`
	JobType    string      `json:"job_type"`
	BeginTime  string      `json:"begin_time"`
	EndTime    string      `json:"end_time"`
	ErrorCode  string      `json:"error_code"`
	FailReason string      `json:"fail_reason"`
	Message    string      `json:"message"`
	Code       string      `json:"code"`
	SubJobs    []JobStatus `json:"sub_jobs"`
}

type JobEntity struct {
	VolumeID string `json:"volume_id"`
}

type JobResult struct {
	golangsdk.Result
}

func (r JobResult) ExtractJobResponse() (*JobResponse, error) {
	job := new(JobResponse)
	err := r.ExtractInto(job)
	return job, err
}

func (r JobResult) ExtractJobStatus() (*JobStatus, error) {
	job := new(JobStatus)
	err := r.ExtractInto(job)
	return job, err
}

func WaitForJobSuccess(client *golangsdk.ServiceClient, secs int, jobID string) error {

	jobClient := *client
	jobClient.Endpoint = strings.Replace(jobClient.Endpoint, "v3", "v1", 1)
	jobClient.ResourceBase = jobClient.Endpoint
	return golangsdk.WaitFor(secs, func() (bool, error) {
		job := new(JobStatus)
		_, err := jobClient.Get(jobClient.ServiceURL("jobs", jobID), &job, nil)
		if err != nil {
			return false, err
		}

		if job.Status == "SUCCESS" {
			return true, nil
		}
		if job.Status == "FAIL" {
			err = fmt.Errorf("Job failed with code %s: %s.\n", job.ErrorCode, job.FailReason)
			return false, err
		}

		return false, nil
	})
}

func GetJobEntity(client *golangsdk.ServiceClient, jobId string, label string) (interface{}, error) {

	if label != "volume_id" {
		return nil, fmt.Errorf("Unsupported label %s in GetJobEntity.", label)
	}

	jobClient := *client
	jobClient.Endpoint = strings.Replace(jobClient.Endpoint, "v3", "v1", 1)
	jobClient.ResourceBase = jobClient.Endpoint
	job := new(JobStatus)
	_, err := jobClient.Get(jobClient.ServiceURL("jobs", jobId), &job, nil)
	if err != nil {
		return nil, err
	}

	if job.Status == "SUCCESS" {
		if e := job.Entities.VolumeID; e != "" {
			return e, nil
		}
	}

	return nil, fmt.Errorf("Unexpected conversion error in GetJobEntity.")
}
