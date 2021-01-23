package service

import (
	"go-gin-gorm/database"
	"go-gin-gorm/model"
)

func SavePost(post *model.Post) error {
	return database.SavePost(post.Title, post.Content)
}

func GetPostById(id int64) (*model.Post, error) {
	post, err := database.GetPost(id)
	if err != nil {
		return nil, err
	}
	return &model.Post{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}, nil
}
