package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/dakasakti/todolist-web/entities"
	pm "github.com/dakasakti/todolist-web/repositories/post"
)

type postService struct {
	Pm pm.PostModel
}

func NewPostService(pm pm.PostModel) *postService {
	return &postService{Pm: pm}
}

func (ps *postService) Register(user_id uint, data entities.PostRequest) error {
	parseTime, err := time.Parse("2006-01-02", data.Deadline)
	if err != nil {
		return err
	}

	dataPost := entities.Post{
		UserID:      user_id,
		Description: data.Description,
		Name:        data.Name,
		Deadline:    parseTime.Format("02-01-2006"),
	}

	err = ps.Pm.Insert(dataPost)
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

	parseTime, _ := time.Parse("02-01-2006", data.Deadline)
	data.Deadline = parseTime.Format("2006-01-02")
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
	parseTime, err := time.Parse("2006-01-02", data.Deadline)
	if err != nil {
		return err
	}

	dataPost := entities.Post{
		Description: data.Description,
		Name:        data.Name,
		Deadline:    parseTime.Format("02-01-2006"),
	}

	err = ps.Pm.Update(id, dataPost)
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
