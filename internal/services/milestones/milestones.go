package milestones

import (
	db "../../database/milestones"
	. "../../models/milestones"
	"time"
)

func AddMilestone(milestone *Milestone) error {
	milestone.CreatedAt = time.Now()
	milestone.LastUpdatedAt = time.Now()
	err := db.AddNewMilestone(milestone)
	return err
}

func ChangeMilestone(milestone *Milestone) error {
	milestone.LastUpdatedAt = time.Now()
	err := db.ChangeMilestone(milestone)
	return err
}

func GetMilestones(projectId uint64) ([]Milestone, error) {
	return db.GetProjectMilestones(projectId)
}

func DeleteMilestone(milestoneId uint64) error {
	return db.DeleteMilestoneById(milestoneId)
}
