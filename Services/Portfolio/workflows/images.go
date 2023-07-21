package workflows

import (
	"fmt"
	app_config "libery_labs_portfolio/Config"
	"libery_labs_portfolio/helpers"
	"os"
	"strings"
)

func GetProjectImages(project_id string) (map[string][]string, error) {
	project_files, err := getProjectFiles(project_id)
	if err != nil {
		return map[string][]string{}, err
	}

	var project_images map[string][]string = make(map[string][]string)
	project_images["mobile"] = make([]string, 0)
	project_images["desktop"] = make([]string, 0)

	for _, file := range project_files {
		if strings.HasPrefix(file, "m-") {
			project_images["mobile"] = append(project_images["mobile"], file)
		} else if strings.HasPrefix(file, "d-") {
			project_images["desktop"] = append(project_images["desktop"], file)
		}
	}

	return project_images, nil
}

func GetProjectImage(project_id string, image_name string) (*os.File, error) {
	var image_path string = fmt.Sprintf("%s/%s/%s", app_config.PROJECTS_DATA_PATH, project_id, image_name)
	if !helpers.FileExists(image_path) {
		return nil, fmt.Errorf("Image not found")
	}

	return os.Open(image_path)
}
