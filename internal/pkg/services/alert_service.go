package services

import (
	"my-app/internal/pkg/models"
	"my-app/internal/pkg/repository"
)

func CreateTest(test models.Test) (string, error) {
	return repository.InsertTest(test)
}

func DeleteTestById(testId string) error {
	return repository.DeleteTestById(testId)
}

func GetTestIdsByName(name string) ([]string, error) {
	return repository.GetTestIdsByName(name)
}
