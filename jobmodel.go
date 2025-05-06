package aijobs

import (
	"time"

	"gorm.io/gorm"
)

type TblJobs struct {
	Id                int       `gorm:"primaryKey;type:serial"`
	Uuid              string    `gorm:"type:character varying"`
	CompanyId         int       `gorm:"type:integer"`
	RecruiterId       int       `gorm:"type:integer"`
	JobTitle          string    `gorm:"type:character varying"`
	JobDescription    string    `gorm:"type:character varying"`
	JobSlug           string    `gorm:"type:character varying"`
	MinExperience     string    `gorm:"type:character varying"`
	MaxExperience     string    `gorm:"type:character varying"`
	MinSalary         string    `gorm:"type:character varying"`
	MaxSalary         string    `gorm:"type:character varying"`
	Industry          string    `gorm:"type:character varying"`
	WorkMode          string    `gorm:"type:character varying"`
	Location          string    `gorm:"type:character varying"`
	SkillsRequired    string    `gorm:"type:character varying"`
	EducationRequired string    `gorm:"type:character varying"`
	EmploymentType    string    `gorm:"type:character varying"`
	PerksAndBenefits  string    `gorm:"type:character varying"`
	LastDateToApply   time.Time `gorm:"type:timestamp without time zone"`
	OpeningDate       time.Time `gorm:"type:timestamp without time zone"`
	AdditionalNotes   string    `gorm:"type:character varying"`
	CreatedOn         time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy         int       `gorm:"type:integer"`
	ModifiedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"DEFAULT:NULL;type:integer"`
	IsDeleted         int       `gorm:"DEFAULT:0;type:integer"`
	DeletedBy         int       `gorm:"DEFAULT:NULL;type:integer"`
	DeletedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsActive          int       `gorm:"DEFAULT:1;type:integer"`
}

type JobDeleteDetails struct {
	JobId     int
	JobUUID   string
	IsDeleted int
	DeletedOn time.Time
	DeletedBy int
}

func CreateJob(jobData TblJobs, DB *gorm.DB) error {

	query := DB.Debug().Table("tbl_jobs").Create(&jobData)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func DeleteJob(deletingDetails JobDeleteDetails, DB *gorm.DB) error {
	query := DB.Debug().Table("tbl_jobs")

	if deletingDetails.JobId != 0 {
		query = query.Where("id = ?", deletingDetails.JobId)
	} else if deletingDetails.JobUUID != "" {
		query = query.Where("uuid = ?", deletingDetails.JobUUID)
	}

	query.Updates(map[string]interface{}{"is_deleted": deletingDetails.IsDeleted, "deleted_by": deletingDetails.DeletedBy, "deleted_on": deletingDetails.DeletedOn})
	if query.Error != nil {
		return query.Error
	}

	return nil
}
