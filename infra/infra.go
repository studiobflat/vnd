package infra

import (
	"errors"

	"github.com/thienhaole92/vnd/firebase"
	"github.com/thienhaole92/vnd/mongo"
	"github.com/thienhaole92/vnd/postgre"
	"github.com/thienhaole92/vnd/redis"
)

type Infra struct {
	redis    *redis.Redis
	mongo    *mongo.Mongo
	postgre  *postgre.Postgre
	firebase *firebase.Firebase
}

func (i *Infra) Redis() (*redis.Redis, error) {
	if i.redis == nil {
		return nil, errors.New("redis client is not set")
	}

	return i.redis, nil
}

func (i *Infra) SetRedis(r *redis.Redis) {
	i.redis = r
}

func (i *Infra) Mongo() (*mongo.Mongo, error) {
	if i.mongo == nil {
		return nil, errors.New("mongo client is not set")
	}

	return i.mongo, nil
}

func (i *Infra) SetMongo(m *mongo.Mongo) {
	i.mongo = m
}

func (i *Infra) Firebase() (*firebase.Firebase, error) {
	if i.firebase == nil {
		return nil, errors.New("firebase client is not set")
	}

	return i.firebase, nil
}

func (i *Infra) SetFirebase(f *firebase.Firebase) {
	i.firebase = f
}

func (i *Infra) Postgre() (*postgre.Postgre, error) {
	if i.postgre == nil {
		return nil, errors.New("postgres client is not set")
	}

	return i.postgre, nil
}

func (i *Infra) SetPostgre(p *postgre.Postgre) {
	i.postgre = p
}
