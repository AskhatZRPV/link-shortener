package mongodb

import (
	"context"
	"errors"
	"fmt"
	"link-shortener/internal/domain/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetOne(id string) (*entity.Link, error)
	GetOneByHash(hash string) (*entity.Link, error)
	GetOneByUrl(url string) (*entity.Link, error)
	Create(link *entity.Link) (*entity.Link, error)
	Delete(book *entity.Link) error
}

type mongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(collection *mongo.Collection) Repository {
	return &mongoRepository{collection: collection}
}

func (r *mongoRepository) GetOne(id string) (*entity.Link, error) {
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var out entity.Link
	err = r.collection.
		FindOne(context.Background(), bson.M{"_id": uid}).
		Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &entity.Link{}, ErrLinkNotFound
		}
		return &entity.Link{}, err
	}
	return &out, nil
}

func (r *mongoRepository) GetOneByHash(hash string) (*entity.Link, error) {
	var out entity.Link
	err := r.collection.
		FindOne(context.Background(), bson.M{"hash": hash}).
		Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &entity.Link{}, ErrLinkNotFound
		}
		return &entity.Link{}, err
	}
	return &out, nil
}

func (r *mongoRepository) GetOneByUrl(url string) (*entity.Link, error) {
	var out entity.Link
	err := r.collection.
		FindOne(context.Background(), bson.M{"url": url}).
		Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &entity.Link{}, ErrLinkNotFound
		}
		return &entity.Link{}, err
	}
	return &out, nil
}

func (r *mongoRepository) Create(link *entity.Link) (*entity.Link, error) {
	nCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := r.collection.InsertOne(nCtx, link)
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

func (r *mongoRepository) Delete(book *entity.Link) error {
	return nil
}
