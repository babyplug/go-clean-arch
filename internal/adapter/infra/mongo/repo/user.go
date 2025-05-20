package repo

import (
	"context"
	"errors"
	"sync"
	"time"

	"go-hexagonal-architecture/internal/adapter/infra/mongo"
	"go-hexagonal-architecture/internal/core/domain"
	"go-hexagonal-architecture/internal/core/port"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"

	driver "go.mongodb.org/mongo-driver/mongo"
)

var (
	_userRepo     port.UserRepository
	_userRepoOnce sync.Once
)

type userRepoImpl struct {
	collection mongo.Collection
}

func NewUserRepo(client mongo.Client) port.UserRepository {
	_userRepoOnce.Do(func() {
		col := client.Database("app").Collection("users")
		_userRepo = &userRepoImpl{collection: col}
	})

	return _userRepo
}

func ResetUserRepo() {
	_userRepoOnce = sync.Once{}
}

func (r *userRepoImpl) Create(ctx context.Context, user *domain.User) error {
	if user == nil {
		return ErrNilValue
	}
	user.ID = primitive.NewObjectID().Hex()
	user.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepoImpl) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	if errors.Is(err, driver.ErrNilDocument) {
		return nil, errors.New("user not found")
	}
	return &user, err
}

func (r *userRepoImpl) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if errors.Is(err, driver.ErrNilDocument) {
		return nil, errors.New("user not found")
	}
	return &user, err
}

func (r *userRepoImpl) List(ctx context.Context, page, size int64) ([]*domain.User, error) {
	skip := CalculateSkip(page, size)
	cur, err := r.collection.Find(ctx, bson.M{}, &options.FindOptions{Limit: &size, Skip: &skip})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users []*domain.User
	for cur.Next(ctx) {
		var user domain.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *userRepoImpl) Update(ctx context.Context, user *domain.User) error {
	filter := bson.M{"id": user.ID}
	update := bson.M{"$set": bson.M{"name": user.Name, "email": user.Email}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *userRepoImpl) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (r *userRepoImpl) Count(ctx context.Context) (int, error) {
	count, err := r.collection.CountDocuments(ctx, bson.M{})
	return int(count), err
}
