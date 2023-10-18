package usermongo

import (
	"context"
	"errors"
	"fmt"
	"link-shortener/internal/domain/user/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(user *entity.User) (*entity.User, error)
	GetOne(id string) (*entity.User, error)
	GetOneByUsername(username string) (*entity.User, error)
	Delete(id string) error
}

type mongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(collection *mongo.Collection) Repository {
	return &mongoRepository{collection: collection}
}

// Create...
func (r *mongoRepository) Create(user *entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.GetOneByUsername(user.Username)
	if err == nil {
		return nil, ErrUserExists
	}

	result, err := r.collection.InsertOne(ctx, user)
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

// GetOne...
func (r *mongoRepository) GetOne(id string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var out entity.User

	if err = r.collection.
		FindOne(ctx, bson.M{"_id": oid}).
		Decode(&out); err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			return &entity.User{}, ErrUserNotFound
		}

		return &entity.User{}, err
	}

	return &out, nil
}

// GetOneByUsername...
func (r *mongoRepository) GetOneByUsername(username string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var out entity.User

	if err := r.collection.
		FindOne(ctx, bson.M{"username": username}).
		Decode(&out); err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			return &entity.User{}, ErrUserNotFound
		}

		return &entity.User{}, err
	}

	return &out, nil
}

// Delete...
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
