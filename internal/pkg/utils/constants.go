package utils

import "business-rule-manager/internal/pkg/models"

var AvailableCategories = []models.Category{
	{Key: "sec"},
}

var AvailableEvents = map[string][]models.Event{
	"sec": {
		{Key: "13f"},
	},
}
