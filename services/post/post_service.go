package services

import (
	"errors"
	"strconv"

	"github.com/dakasakti/postingan/entities"
	pm "github.com/dakasakti/postingan/repositories/post"
)

type postService struct {
	Pm pm.PostModel
}

func NewPostService(pm pm.PostModel) *postService {
	return &postService{Pm: pm}
}

func (ps *postService) Register(user_id uint, data entities.PostRequest) error {
	dataPost := entities.Post{
		UserID:      user_id,
		Description: data.Description,
		Name:        data.Name,
		Deadline:    data.Deadline,
	}

	err := ps.Pm.Insert(dataPost)
	if err != nil {
		if err.Error() == "Error 1452: Cannot add or update a child row: a foreign key constraint fails (`postingan`.`posts`, CONSTRAINT `fk_posts_post_type` FOREIGN KEY (`post_type_id`) REFERENCES `post_types` (`id`))" {
			return errors.New("post_type_id not found")
		} else {
			return err
		}
	}

	return nil
}

func (ps *postService) GetAll() ([]entities.Post, error) {
	data, err := ps.Pm.Gets()
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	return data, nil
}

func (ps *postService) CheckParamId(id string) (uint, error) {
	idConv, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(idConv), nil
}

func (ps *postService) GetById(id uint) (entities.Post, error) {
	data, err := ps.Pm.Get(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (ps *postService) CheckUser(id uint, user_id uint) (entities.Post, error) {
	data, err := ps.Pm.Get(id)
	if data.UserID != user_id {
		return data, errors.New("user not authorized")
	}

	if err != nil {
		return data, err
	}

	return data, nil
}

func (ps *postService) UpdateById(id uint, data entities.PostUpdateRequest) error {
	dataPost := entities.Post{
		Description: data.Description,
		Name:        data.Name,
		Deadline:    data.Deadline,
	}

	err := ps.Pm.Update(id, dataPost)
	if err != nil {
		return err
	}

	return nil
}

func (ps *postService) UpdateMarkById(id uint) error {
	dataPost := entities.Post{
		Status: "Done",
	}

	err := ps.Pm.Update(id, dataPost)
	if err != nil {
		return err
	}

	return nil
}

func (ps *postService) DeleteById(id uint) error {
	err := ps.Pm.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
