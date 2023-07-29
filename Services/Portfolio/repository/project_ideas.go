package repository

type ProjectIdeasRepository interface {
	AddIdea(idea string) error
	GetRandomIdea() string
	GetIdeas() []string
	GetNIdeas(n int) []string
}

var project_ideas_repo_implementation ProjectIdeasRepository

func SetProjectIdeasRepositoryImplementation(implementation ProjectIdeasRepository) {
	project_ideas_repo_implementation = implementation
}

func AddIdea(idea string) error {
	return project_ideas_repo_implementation.AddIdea(idea)
}

func GetRandomIdea() string {
	return project_ideas_repo_implementation.GetRandomIdea()
}

func GetIdeas() []string {
	return project_ideas_repo_implementation.GetIdeas()
}

func GetNIdeas(n int) []string {
	return project_ideas_repo_implementation.GetNIdeas(n)
}
