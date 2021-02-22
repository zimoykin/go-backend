package config

type Configuration struct {
	JwtSecret string `env:"JWT_SECRET,required"`
	RedisAddr string `env:"REDIS_ADDR" envDefault:":6379"`
	RedisPwd  string `env:"REDIS_PWD"`
}

func newConfig() {

}
