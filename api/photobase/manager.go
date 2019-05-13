package photobase

import (
	"errors"

	"github.com/JulzDiverse/eirinidotcf/api"
)

type InMemoryPhotobase struct {
	photos map[string]api.Photo
}

func NewInMemoryPhotobase() InMemoryPhotobase {
	return InMemoryPhotobase{
		photos: map[string]api.Photo{},
	}
}

func (p *InMemoryPhotobase) List() ([]api.Photo, error) {
	photoList := []api.Photo{}
	for _, photo := range p.photos {
		photoList = append(photoList, photo)
	}
	return photoList, nil
}

func (p *InMemoryPhotobase) Add(photo api.Photo) error {
	p.photos[photo.Hero] = photo
	return nil
}

func (p *InMemoryPhotobase) Get(id string) (api.Photo, error) {
	photo, ok := p.photos[id]
	if !ok {
		return api.Photo{}, errors.New("photo doesn't exist")
	}

	return photo, nil
}
