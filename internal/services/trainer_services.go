package services

import (
	"github.com/faveroferreira/maromba-center/internal/models"
	"gorm.io/gorm"
)

type TrainerServices struct {
	DB *gorm.DB
}

func (ts TrainerServices) CreateTrainer(trainer *models.Trainer) error {
	result := ts.DB.Create(trainer)

	return result.Error
}

func (ts TrainerServices) GetTrainerById(trainer *models.Trainer, id string) error {
	result := ts.DB.First(trainer, id)

	return result.Error
}

func (ts TrainerServices) GetAllTrainers(trainers *[]models.Trainer) error {
	result := ts.DB.Find(trainers)

	return result.Error
}

func (ts TrainerServices) UpdateTrainer(body *models.Trainer, trainer *models.Trainer, id string) error {

	if err := ts.GetTrainerById(trainer, id); err != nil {
		return err
	}

	trainer.Name = body.Name
	trainer.Price = body.Price
	trainer.HatesGO = body.HatesGO

	ts.DB.Save(trainer)

	return nil
}

func (ts TrainerServices) DeleteTrainer(id string) error {
	var trainer models.Trainer

	if err := ts.GetTrainerById(&trainer, id); err != nil {
		return err
	}

	ts.DB.Delete(&trainer)
	return nil
}
