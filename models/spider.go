package models

// "github.com/duanduan2288/spidermodels/models"
// "fmt"

var RedisModels = &RedisModel{}

type RedisModel struct{}

func (s *RedisModel) SetContent(key, content string) error {
	conn := NewRedis(RedisInit()).Get()
	defer conn.Close()

	if _, err := conn.Do("SET", key, content); err != nil {
		return err
	}

	return nil
}
