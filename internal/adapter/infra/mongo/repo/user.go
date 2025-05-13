package repo

import (
	"context"
	"errors"
	"time"

	"github.com/babyplug/go-clean-arch/internal/core/domain"
	"github.com/babyplug/go-clean-arch/internal/core/port"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepoImpl struct {
	collection *mongo.Collection
}

func NewUserRepo(client *mongo.Client) port.UserRepository {
	col := client.Database("app").Collection("users")
	return &userRepoImpl{collection: col}
}

func (r *userRepoImpl) Create(user *domain.User) error {
	user.ID = primitive.NewObjectID().Hex()
	user.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}

func (r *userRepoImpl) GetByID(id string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return &user, err
}

func (r *userRepoImpl) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return &user, err
}

func (r *userRepoImpl) List() ([]*domain.User, error) {
	cur, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var users []*domain.User
	for cur.Next(context.TODO()) {
		var user domain.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *userRepoImpl) Update(user *domain.User) error {
	filter := bson.M{"id": user.ID}
	update := bson.M{"$set": bson.M{"name": user.Name, "email": user.Email}}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *userRepoImpl) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"id": id})
	return err
}

func (r *userRepoImpl) Count() (int, error) {
	count, err := r.collection.CountDocuments(context.TODO(), bson.M{})
	return int(count), err
}
