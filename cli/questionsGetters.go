package cli

import (
	"github.com/AlecAivazis/survey/v2"
	"templater/config"
	"templater/models"
)

func getPropertyName() *survey.Input {
	return &survey.Input{Message: "Choose property name:"}
}

func getPropertyValueType() *survey.Select {
	return &survey.Select{Message: "Choose property valueType:", Options: models.GetValueTypesStrings()}
}

func getName() *survey.Input {
	return &survey.Input{Message: "Choose entity name:"}
}

func getProjectName() *survey.Select {
	return &survey.Select{Message: "Choose Project:", Options: config.GetProjectsNames()}
}

func getServerBuildOperations() *survey.MultiSelect {
	return &survey.MultiSelect{Message: "Choose Server operations:", Options: models.GetServerBuildOperations()}
}

func getMode() *survey.Select {
	return &survey.Select{Message: "Choose mode:", Options: models.GetModels()}
}

func getWorks(works models.Works) *survey.MultiSelect {
	return &survey.MultiSelect{Message: "Choose works:", Options: works.GetNames()}
}
