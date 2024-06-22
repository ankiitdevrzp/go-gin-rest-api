package service

import "github.com/ankiitdev/golang/go-gin-api/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	if len(service.videos) == 0 {
		var emptyVideo entity.Video
		// emptyVideo.Title = "Empty Title"
		// emptyVideo.Description = "Empty Description"
		// emptyVideo.URL = "Empty URL"
		return append(service.videos, emptyVideo)
	}
	return service.videos
}
