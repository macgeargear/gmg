package mongodb

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	m "github.com/macgeargear/gmg/models"
	repo "github.com/macgeargear/gmg/repositories"
)

type userMongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newUserMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, e := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if e != nil {
		return nil, e
	}
	if e = client.Ping(ctx, readpref.Primary()); e != nil {
		return nil, e
	}
	return client, e
}

func NewUserRepository(mongoURL, mongoDB string, mongoTimeout int) (repo.UserRepo, error) {
	repo := &userMongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}

	client, err := newUserMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	repo.client = client
	return repo, nil
}

func (r *userMongoRepository) GetAll() ([]m.User, error) {
	res := []m.User{}
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	c := r.client.Database(r.database).Collection(new(m.User).TableName())
	queryResult, err := c.Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	if err := queryResult.Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *userMongoRepository) GetBy(filter map[string]interface{}) (*m.User, error) {
	res := new(m.User)
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	convertID(filter)
	c := r.client.Database(r.database).Collection(res.TableName())

	if err := c.FindOne(ctx, filter).Decode(res); err != nil {
		if err == mongo.ErrNoDocuments {
			return res, errors.New(err.Error())
		}
	}
	return res, nil

}

func (r *userMongoRepository) Create(data m.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	c := r.client.Database(r.database).Collection(data.TableName())

	if _, err := c.InsertOne(ctx, data); err != nil {
		return err
	}
	return nil
}

func (r *userMongoRepository) Update(user *m.User) error {
	return nil
}

func (r *userMongoRepository) Delete(userId string) error {
	return nil
}

func convertID(data map[string]interface{}) {
	if _, ok := data["ID"]; ok {
		data["_id"] = data["ID"]
		delete(data, "ID")
	}
}
