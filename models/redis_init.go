package models

func RedisInit() RedisService {
	// configNames := make(map[string]RedisService)

	return RedisService{
		Addr:        "127.0.0.1:6379",
		Password:    "",
		MaxIdle:     10,
		IdleTimeout: 3600,
	}

	// configNames["default"] = redisService

	// InitRedis(configNames)

}
