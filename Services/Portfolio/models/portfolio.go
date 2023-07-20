package models

import (
	"fmt"
)

type Project struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	TechStack   []string `json:"tech_stack"`
	Relevance   int      `json:"relevance"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	Github      string   `json:"github"`
}

func (p *Project) Create(project_data map[string]any) error {

	if project_data["id"] == nil || project_data["id"].(string) == "" {
		return fmt.Errorf("ID cannot be empty")
	}

	p.ID = project_data["id"].(string)

	if project_data["name"] == nil || project_data["name"].(string) == "" {
		return fmt.Errorf("Name cannot be empty")
	}

	p.Name = project_data["name"].(string)

	if project_data["tech_stack"] == nil {
		return fmt.Errorf("Tech stack cannot be empty")
	}

	parsed_tech_stack := project_data["tech_stack"].([]string)
	if len(parsed_tech_stack) == 0 {
		return fmt.Errorf("Tech stack cannot be empty")
	}

	p.TechStack = parsed_tech_stack

	if project_data["relevance"] == nil {
		return fmt.Errorf("Relevance cannot be empty")
	}

	p.Relevance = project_data["relevance"].(int)

	if project_data["description"] == nil {
		return fmt.Errorf("Description cannot be empty")
	}

	p.Description = project_data["description"].(string)

	if project_data["url"] == nil {
		return fmt.Errorf("URL cannot be empty")
	}

	p.Url = project_data["url"].(string)

	if project_data["github"] == nil {
		return fmt.Errorf("Github cannot be empty")
	}

	p.Github = project_data["github"].(string)

	return nil
}

func (p *Project) Update(new_data map[string]any) error {
	if new_data["name"] != nil && new_data["name"].(string) != "" {
		p.Name = new_data["name"].(string)
	}

	if new_data["tech_stack"] != nil {
		parsed_tech_stack := make([]string, 0)
		for _, tech := range new_data["tech_stack"].([]interface{}) {
			parsed_tech_stack = append(parsed_tech_stack, tech.(string))
		}

		if len(parsed_tech_stack) == 0 {
			return fmt.Errorf("Tech stack cannot be empty")
		}
		p.TechStack = parsed_tech_stack
	}

	if new_data["relevance"] != nil {
		p.Relevance = int(new_data["relevance"].(float64))
	}

	if new_data["description"] != nil {
		p.Description = new_data["description"].(string)
	}

	if new_data["url"] != nil {
		p.Url = new_data["url"].(string)
	}

	if new_data["github"] != nil {
		p.Github = new_data["github"].(string)
	}

	return nil
}
