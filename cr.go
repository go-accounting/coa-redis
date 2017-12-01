package coaredis

import (
	"strings"

	"github.com/go-accounting/coa"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type redisStore struct {
	c      redis.UniversalClient
	prefix *string
}

func NewStore(master string, addrs []string, prefix *string) (coa.KeyValueStore, error) {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		MasterName: master,
		Addrs:      addrs,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "ping failed")
	}
	return redisStore{client, prefix}, nil
}

func (rs redisStore) Get(key []byte) ([]byte, error) {
	cmd := rs.c.Get(prefix(*rs.prefix) + string(key))
	if cmd.Err() == redis.Nil {
		return nil, nil
	}
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	b, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (rs redisStore) Put(key []byte, value []byte) error {
	return rs.c.Set(prefix(*rs.prefix)+string(key), value, 0).Err()
}

func prefix(s string) string {
	if strings.HasSuffix(s, "/") {
		return s
	}
	return s + "/"
}
