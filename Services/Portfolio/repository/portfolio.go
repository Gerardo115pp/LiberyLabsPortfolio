package repository

import (
	"context"
	"libery_labs_portfolio/models"
)

type PortfolioRepository interface {
	InsertProject(ctx context.Context, portfolio models.Project) error
	GetProjects(ctx context.Context) ([]models.Project, error)
	GetProjectByID(ctx context.Context, id string) (models.Project, error)
	DeleteProject(ctx context.Context, id string) error
	UpdateProject(ctx context.Context, admin models.Project) error
	Close()
}

var admin_repo_implementation PortfolioRepository

func SetPortfolioRepositoryImplementation(implementation PortfolioRepository) {
	admin_repo_implementation = implementation
}

func InsertProject(ctx context.Context, project models.Project) error {
	return admin_repo_implementation.InsertProject(ctx, project)
}

func GetProjects(ctx context.Context) ([]models.Project, error) {
	return admin_repo_implementation.GetProjects(ctx)
}

func GetProjectByID(ctx context.Context, id string) (models.Project, error) {
	return admin_repo_implementation.GetProjectByID(ctx, id)
}

func DeleteProject(ctx context.Context, id string) error {
	return admin_repo_implementation.DeleteProject(ctx, id)
}

func UpdateProject(ctx context.Context, project models.Project) error {
	return admin_repo_implementation.UpdateProject(ctx, project)
}

func Close() {
	admin_repo_implementation.Close()
}
