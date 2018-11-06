package models

func RedisInit() {
	configNames := make(map[string]RedisService)

	redisService := RedisService{
		Addr:        "127.0.0.1",
		Password:    "",
		MaxIdle:     10,
		IdleTimeout: 3600,
	}

	configNames["default"] = redisService

	InitRedis(configNames)
}
