package database

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	app_config "libery_labs_portfolio/Config"
	"libery_labs_portfolio/helpers"
	"libery_labs_portfolio/models"
	"os"
)

type PortfolioRepository struct {
	projects_directory string
	projects           map[string]models.Project
}

func NewPortfolioRepository() (PortfolioRepository, error) {
	projects_directory := app_config.PROJECTS_DATA_PATH
	var projects map[string]models.Project = make(map[string]models.Project)

	new_repo := new(PortfolioRepository)

	new_repo.projects_directory = projects_directory
	projects, err := new_repo.loadProjects()
	if err != nil {
		return PortfolioRepository{}, err
	}

	new_repo.projects = projects

	return *new_repo, nil
}

func (portfolio_repo PortfolioRepository) InsertProject(ctx context.Context, new_project models.Project) error {
	if new_project.ID == "" {
		return fmt.Errorf("Name cannot be empty")
	}

	project_directory := fmt.Sprintf("%s/%s", app_config.PROJECTS_DATA_PATH, new_project.ID)

	if helpers.FileExists(project_directory) {
		return fmt.Errorf("Project already exists")
	}

	err := os.Mkdir(project_directory, os.ModePerm) // 0777
	if err != nil {
		return err
	}

	json_data, err := json.Marshal(new_project)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/project.json", project_directory), json_data, 0777)
	if err != nil {
		return err
	}

	portfolio_repo.projects[new_project.ID] = new_project

	return nil
}

func (portfolio_repo PortfolioRepository) GetProjects(ctx context.Context) ([]models.Project, error) {
	var projects []models.Project = make([]models.Project, 0)

	for _, project := range portfolio_repo.projects {
		projects = append(projects, project)
	}

	return projects, nil
}

func (portfolio_repo PortfolioRepository) loadProjects() (map[string]models.Project, error) {
	var projects map[string]models.Project = make(map[string]models.Project)

	files, err := ioutil.ReadDir(portfolio_repo.projects_directory)
	if err != nil {
		return projects, err
	}

	for _, file := range files {
		if file.IsDir() {
			project_json, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/project.json", portfolio_repo.projects_directory, file.Name()))
			if err != nil {
				return projects, err
			}

			var project models.Project
			err = json.Unmarshal(project_json, &project)
			if err != nil {
				return projects, err
			}

			projects[project.ID] = project
		}
	}

	return projects, nil
}

func (portfolio_repo PortfolioRepository) GetProjectByID(ctx context.Context, id string) (models.Project, error) {
	project, _ := portfolio_repo.projects[id]
	return project, nil
}

func (portfolio_repo PortfolioRepository) DeleteProject(ctx context.Context, id string) error {
	project_directory := fmt.Sprintf("%s/%s", app_config.PROJECTS_DATA_PATH, id)

	if !helpers.FileExists(project_directory) {
		return fmt.Errorf("Project does not exist")
	}

	err := os.RemoveAll(project_directory)
	if err != nil {
		return err
	}

	delete(portfolio_repo.projects, id)

	return nil
}

func (portfolio_repo PortfolioRepository) UpdateProject(ctx context.Context, project models.Project) error {
	new_data, err := json.Marshal(project)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/project.json", app_config.PROJECTS_DATA_PATH, project.ID), new_data, 0777)
	if err != nil {
		return err
	}

	portfolio_repo.projects[project.ID] = project

	return nil
}

func (portfolio_repo PortfolioRepository) Close() {
	// Nothing to close
}
