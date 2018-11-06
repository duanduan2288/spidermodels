package spider

import (
	"models"
)

var RedisModel = &Redis{}

type Redis struct{}

func (s *Redis) SetContent(key, content string) error {
	conn := models.Redis().Get()
	defer conn.Close()

}
