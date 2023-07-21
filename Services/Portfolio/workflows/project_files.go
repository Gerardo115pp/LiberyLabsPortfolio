package workflows

import (
	"fmt"
	"io/ioutil"
	app_config "libery_labs_portfolio/Config"
	"strings"
)

func getProjectFiles(project_id string) ([]string, error) {
	files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", app_config.PROJECTS_DATA_PATH, project_id))
	if err != nil {
		return []string{}, err
	}

	var project_files []string = make([]string, 0)
	for _, file := range files {
		if !file.IsDir() || file.Name() != "project.json" || !strings.HasPrefix(file.Name(), ".") {
			project_files = append(project_files, file.Name())
		}
	}

	return project_files, nil
}
