package service

import (
	"github.com/rodrigoberes/rest-api/data/request"
	"github.com/rodrigoberes/rest-api/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}