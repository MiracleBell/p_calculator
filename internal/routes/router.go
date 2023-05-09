package routes

import (
	. "./features"
	. "./milestones"
	. "./projects"
	. "./rates"
	. "./teams"
	. "./users"
	"net/http"
)

func MainRouter() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/users", RegisterNewUser)

	http.HandleFunc("/projects", AddNewProject)
	http.HandleFunc("/projects/{projectId}", GetProjectList)

	http.HandleFunc("/projects/{projectId}/features", FeaturesRouter)
	http.HandleFunc("/projects/{projectId}/features/{featureId}", PutFeature)

	http.HandleFunc("/projects/{projectId}/milestones", MilestoneRouter)
	http.HandleFunc("/projects/{projectId}/milestones/{milestoneId}", MilestoneRouterWithId)

	http.HandleFunc("/projects/{projectId}/rates", RatesRouter)

	http.HandleFunc("/projects/{projectId}/team-members", TeamRouter)

	// запускаем сервер
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
