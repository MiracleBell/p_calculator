package projects

import (
	db "../../database/projects"
	. "../../models/positions"
	. "../../models/projects"
	. "../../models/rates"
	positionService "../positions"
	rateService "../rates"
	"time"
)

func AddProject(project *Project) error {
	project.CreatedAt = time.Now()
	project.LastUpdatedAt = time.Now()
	err := db.AddNewProject(project)
	if err != nil {
		return err
	}

	positions, err := positionService.GetPositions()
	if err != nil {
		return err
	}

	for i := 0; i < len(positions); i++ {
		position := Position{Name: positions[i].Name, DefaultRate: positions[i].DefaultRate}
		var rate = Rate{0, position, position.DefaultRate, *project}
		rateService.AddRate(&rate)
	}
	return err
}

func GetProjects() ([]Project, error) {
	return db.GetProjects()
}
