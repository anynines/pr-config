package backend

type RedisBackend struct {
}

func NewRedisBackend() *RedisBackend {
	return &RedisBackend{}
}

func (r *RedisBackend) Write(key, value string) error {
	return nil
}
