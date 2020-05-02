package cli

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	// RedisPasswordFlag is the ENV var for the Redis password
	RedisPasswordFlag string = "redis-password"
	// RedisHostFlag is the ENV var for the Redis hostname
	RedisHostFlag string = "redis-host"
	// RedisPortFlag is the ENV var for the Redis port
	RedisPortFlag string = "redis-port"
	// RedisDBNameFlag is the ENV var for the Redis database name, which
	// is represented by a positive integer. Using multiple databases in
	// the same Redis instance allows separating concerns.
	RedisDBNameFlag string = "redis-db-name"
	// RedisConnectTimeoutFlag specifies how long to wait to establish a
	// connection to the Redis instance
	RedisConnectTimeoutFlag string = "redis-connect-timeout-in-seconds"
	// RedisEnabledFlag specifies whether or not we attempt to connect
	// to Redis. For example, apps that use mTLS don't need Redis.
	RedisEnabledFlag string = "redis-enabled"
)

// InitRedisFlags initializes RedisFlags command line flags
func InitRedisFlags(flag *pflag.FlagSet) {
	flag.String(RedisPasswordFlag, "", "Redis password")
	flag.String(RedisHostFlag, "localhost", "Redis hostname")
	flag.Int(RedisPortFlag, 6379, "Redis port")
	flag.Int(RedisDBNameFlag, 0, "Redis database")
	flag.Int(RedisConnectTimeoutFlag, 2, "Redis connect timeout in seconds")
	flag.Bool(RedisEnabledFlag, true, "Whether or not Redis is enabled")
}

// CheckRedis validates Redis command line flags
func CheckRedis(v *viper.Viper) error {
	enabled := v.GetBool(RedisEnabledFlag)
	if !enabled {
		return nil
	}

	if err := ValidatePort(v, RedisPortFlag); err != nil {
		return err
	}

	environment := v.GetString(EnvironmentFlag)
	if environment == EnvironmentProd {
		if err := ValidateRedisHost(v, RedisHostFlag); err != nil {
			return err
		}
	}

	if err := ValidateRedisConnectTimeout(v, RedisConnectTimeoutFlag); err != nil {
		return err
	}

	return nil
}

// InitRedis initializes a Redis pool from command line flags.
// v is the viper Configuration.
// logger is the application logger.
func InitRedis(v *viper.Viper, logger Logger) (*redis.Pool, error) {
	enabled := v.GetBool(RedisEnabledFlag)
	if !enabled {
		return nil, nil
	}

	redisPassword := v.GetString(RedisPasswordFlag)
	redisHost := v.GetString(RedisHostFlag)
	redisPort := strconv.Itoa(v.GetInt(RedisPortFlag))
	redisDBName := v.GetInt(RedisDBNameFlag)
	redisConnectTimeout := v.GetInt(RedisConnectTimeoutFlag)
	timeoutDuration := time.Duration(redisConnectTimeout) * time.Second

	redisURITemplate := "%s:%s"
	redisURL := fmt.Sprintf(redisURITemplate, redisHost, redisPort)

	if err := logRedisConnection(redisURL, logger); err != nil {
		return nil, err
	}

	if err := testRedisConnection(redisURL, redisPassword, redisDBName, timeoutDuration, logger); err != nil {
		return nil, err
	}

	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			connection, err := redis.Dial(
				"tcp",
				redisURL,
				redis.DialDatabase(redisDBName),
				redis.DialPassword(redisPassword),
			)
			if err != nil {
				return nil, err
			}
			return connection, err
		},
	}

	return pool, nil
}

// ValidateRedisHost validates the REDIS_HOST environment variable in prod
func ValidateRedisHost(v *viper.Viper, flagname string) error {
	r := regexp.MustCompile(`([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}`)

	if host := v.GetString(flagname); r.MatchString(host) == false {
		return errors.Errorf("%s can only contain letters, numbers, periods, and dashes", flagname)
	}

	return nil
}

// ValidateRedisConnectTimeout validates a Redis connect timeout passed in from
// the command line
func ValidateRedisConnectTimeout(v *viper.Viper, flagname string) error {
	timeout := v.GetInt(flagname)

	if timeout < 1 || timeout > 5 {
		return errors.Errorf("%s should be between 1 and 5 seconds", flagname)
	}

	return nil
}

func testRedisConnection(redisURL string, redisPassword string, redisDBName int, timeout time.Duration, logger Logger) error {
	connection, err := redis.Dial(
		"tcp",
		redisURL,
		redis.DialDatabase(redisDBName),
		redis.DialPassword(redisPassword),
		redis.DialConnectTimeout(timeout),
	)
	logger.Info("Testing Redis connection...")
	errorString := fmt.Sprintf("Failed to connect to Redis host %s after %s", redisURL, timeout)
	var finalErrorString string
	if redisPassword == "" {
		finalErrorString = errorString + ". No password provided."
	}
	if err != nil {
		logger.Error(finalErrorString, zap.Error(err))
		return err
	}
	logger.Info("...Redis connection successful!")

	logger.Info("Starting Redis ping...")
	_, pingError := redis.String(connection.Do("PING"))
	if err != nil {
		logger.Error("failed to ping Redis", zap.Error(pingError))
	}
	logger.Info("...Redis ping successful!")

	return nil
}

func logRedisConnection(redisURL string, logger Logger) error {
	logger.Info("Connecting to Redis", zap.String("url", redisURL))

	return nil
}
