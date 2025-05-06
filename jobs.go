package aijobs

func JobSetup(config Config) *JobConfig {

	return &JobConfig{
		DB: config.DB,
	}
}

func (config *JobConfig) CreateJob(jobDetails TblJobs) error {

	err := CreateJob(jobDetails, config.DB)
	if err != nil {
		return err
	}

	return nil
}

func (config *JobConfig) DeleteJobById(jobId int) error {

	var deletingDetails JobDeleteDetails
	deletingDetails.JobId = jobId

	err := DeleteJob(deletingDetails, config.DB)
	if err != nil {
		return err
	}

	return nil
}

func (config *JobConfig) DeleteJobByUuid(jobUUID string) error {

	var deletingDetails JobDeleteDetails
	deletingDetails.JobUUID = jobUUID

	err := DeleteJob(deletingDetails, config.DB)
	if err != nil {
		return err
	}

	return nil
}
