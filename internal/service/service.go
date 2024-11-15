package service

import (
    "sync"
    "retail-pulse/internal/models"
)

type JobService struct {
    jobs     map[int64]*models.Job
    jobsLock sync.RWMutex
}

func NewJobService() *JobService {
    return &JobService{
        jobs: make(map[int64]*models.Job),
    }
}

func (s *JobService) CreateJob(job *models.Job) {
    s.jobsLock.Lock()
    defer s.jobsLock.Unlock()
    s.jobs[job.ID] = job
}

func (s *JobService) GetJob(id int64) (*models.Job, bool) {
    s.jobsLock.RLock()
    defer s.jobsLock.RUnlock()
    job, exists := s.jobs[id]
    return job, exists
}