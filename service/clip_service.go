package service

import (
	"github.com/rodrigoberes/rest-api/data/request"
	"github.com/rodrigoberes/rest-api/data/response"
)

type ClipsService interface {
	Create(clips request.CreateClipsRequest)
	Update(clips request.UpdateClipsRequest)
	Delete(clipsId int)
	FindById(clipsId int) response.ClipsResponse
	FindAll() []response.ClipsResponse
}