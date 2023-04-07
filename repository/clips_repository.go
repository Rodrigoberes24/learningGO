package repository

import "github.com/rodrigoberes/rest-api/model"

type ClipsRepository interface {
	Save(clips model.Clips)
	Update(clips model.Clips)
	Delete(clipsId int)
	FindById(clipsId int) (tags model.Clips, err error)
	FindAll() []model.Clips
}