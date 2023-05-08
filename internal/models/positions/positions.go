package positions

type Position struct {
	Name        string
	DefaultRate float64
}

const (
	REGULAR_DEVELOPER string = "REGULAR_DEVELOPER"
	SENIOR_DEVELOPER         = "SENIOR_DEVELOPER"
	QA_ENGINEER              = "QA_ENGINEER"
	PROJECT_MANAGER          = "PROJECT_MANAGER"
	DEVOPS_ENGINEER          = "DEVOPS_ENGINEER"
	ARCHITECT                = "ARCHITECT"
)
