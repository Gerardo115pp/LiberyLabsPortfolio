package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	app_config "libery_labs_portfolio/Config"
	"libery_labs_portfolio/helpers"
	"math/rand"
	"time"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

type ProjectIdeasRepository struct {
	ProjectIdeas []string `json:"project_ideas"`
}

func NewProjectIdeasRepository() (*ProjectIdeasRepository, error) {
	new_repo := new(ProjectIdeasRepository)

	project_ideas_file := fmt.Sprintf("%s/project_ideas.json", app_config.OPERATION_DATA_PATH)
	if !helpers.FileExists(project_ideas_file) {
		return nil, fmt.Errorf("Project ideas file does not exist on '%s'", app_config.OPERATION_DATA_PATH)
	}

	file_data, err := ioutil.ReadFile(project_ideas_file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file_data, &new_repo)

	echo.Echo(echo.GreenFG, fmt.Sprintf("Loaded %d project ideas", len(new_repo.ProjectIdeas)))

	return new_repo, err
}

func (project_ideas_repo *ProjectIdeasRepository) AddIdea(idea string) error {
	project_ideas_repo.ProjectIdeas = append(project_ideas_repo.ProjectIdeas, idea)

	err := project_ideas_repo.writeIdeas()
	if err != nil {
		return err
	}

	return nil
}

func (project_ideas_repo ProjectIdeasRepository) GetRandomIdea() string {
	rand.Seed(time.Now().UnixNano())

	random_index := rand.Intn(len(project_ideas_repo.ProjectIdeas))

	return project_ideas_repo.ProjectIdeas[random_index]
}

func (project_ideas_repo ProjectIdeasRepository) GetIdeas() []string {
	return project_ideas_repo.ProjectIdeas
}

func (project_ideas_repo ProjectIdeasRepository) GetNIdeas(n int) []string {
	if n > len(project_ideas_repo.ProjectIdeas) {
		n = len(project_ideas_repo.ProjectIdeas)
	}

	return project_ideas_repo.ProjectIdeas[:n]
}

func (project_ideas_repo *ProjectIdeasRepository) writeIdeas() (err error) {
	project_ideas_file := fmt.Sprintf("%s/project_ideas.json", app_config.OPERATION_DATA_PATH)

	file_data, err := json.Marshal(project_ideas_repo)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(project_ideas_file, file_data, 0644)
	if err != nil {
		return err
	}

	return nil
}
