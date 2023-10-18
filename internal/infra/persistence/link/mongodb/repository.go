package linkmongo

import (
	"context"
	"errors"
	"fmt"
	"link-shortener/internal/domain/link/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(link *entity.Link) (*entity.Link, error)
	GetOne(id string) (*entity.Link, error)
	GetOneByHash(hash string) (*entity.Link, error)
	Delete(id string) error
}

type mongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(collection *mongo.Collection) Repository {
	return &mongoRepository{collection: collection}
}

func (r *mongoRepository) GetOne(id string) (*entity.Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var out entity.Link

	if err = r.collection.
		FindOne(ctx, bson.M{"_id": oid}).
		Decode(&out); err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			return &entity.Link{}, ErrLinkNotFound
		}

		return &entity.Link{}, err
	}

	return &out, nil
}

func (r *mongoRepository) GetOneByHash(hash string) (*entity.Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var out entity.Link

	if err := r.collection.
		FindOne(ctx, bson.M{"hash": hash}).
		Decode(&out); err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			return &entity.Link{}, ErrLinkNotFound
		}

		return &entity.Link{}, err
	}

	return &out, nil
}

func (r *mongoRepository) Create(link *entity.Link) (*entity.Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, link)
	if err != nil {
		return nil, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		fmt.Println(oid.Hex())
		ent, err := r.GetOne(oid.Hex())
		if err != nil {
			return nil, err
		}
		return ent, nil
	}

	return nil, err
}

func (r *mongoRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return err
	}

	return nil
}
