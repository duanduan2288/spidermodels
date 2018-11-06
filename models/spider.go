package models

// import (
// 	"github.com/duanduan2288/spidermodels/models"
// )

var RedisModel = &Redis{}

type Redis struct{}

func (s *Redis) SetContent(key, content string) error {
	conn := models.Redis().Get()
	defer conn.Close()

	if _, err = conn.Do("SET", key, content); err != nil {
		return models.Error(err)
	}

	return nil
}
