package repo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	driver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-hexagonal-architecture/internal/adapter/infra/mongo"
	"go-hexagonal-architecture/internal/adapter/infra/mongo/mock"
	"go-hexagonal-architecture/internal/adapter/infra/mongo/repo"
	"go-hexagonal-architecture/internal/core/domain"
)

// type

func TestUserRepoImpl_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name         string
		user         *domain.User
		dependency   func(ctrl *gomock.Controller) mongo.Client
		expectErr    bool
		expectErrMsg string
		expected     *domain.User
	}{
		{
			name: "when user is valid should create user",
			user: &domain.User{
				Name:     "Test",
				Email:    "test@google.com",
				Password: "password",
			},
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return(nil, nil)

				return client
			},
			expected: &domain.User{
				Name:     "Test",
				Email:    "test@google.com",
				Password: "password",
			},
		},
		{
			name: "when user is nil should return error",
			user: nil,
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)

				return client
			},
			expectErr:    true,
			expectErrMsg: repo.ErrNilValue.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := test.dependency(ctrl)
			r := repo.NewUserRepo(m)
			defer repo.ResetUserRepo()

			err := r.Create(context.Background(), test.user)

			if test.expectErr {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expectErrMsg)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, test.user)
				assert.NotEmpty(t, test.user.ID)
				assert.Equal(t, test.user.Name, test.expected.Name)
				assert.Equal(t, test.user.Email, test.expected.Email)
				assert.Equal(t, test.user.Password, test.expected.Password)
			}
		})
	}
}

func TestUserRepoImpl_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name         string
		id           string
		dependency   func(ctrl *gomock.Controller) mongo.Client
		expectErr    bool
		expectErrMsg string
		expected     *domain.User
	}{
		{
			name: "when user exists should return user",
			id:   "123",
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().FindOne(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) mongo.SingleResult {
						return driver.NewSingleResultFromDocument(&domain.User{
							ID:    "123",
							Name:  "Test User",
							Email: "test@google.com",
						}, nil, nil)
					})

				return client
			},
			expected: &domain.User{
				ID:    "123",
				Name:  "Test User",
				Email: "test@google.com",
			},
		},
		{
			name: "when user not found should return error",
			id:   "notfound",
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(driver.NewSingleResultFromDocument(nil, nil, nil))

				return client
			},
			expectErr:    true,
			expectErrMsg: "user not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := test.dependency(ctrl)
			r := repo.NewUserRepo(m)
			defer repo.ResetUserRepo()

			user, err := r.GetByID(context.Background(), test.id)

			if test.expectErr {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expectErrMsg)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, test.expected.ID, user.ID)
				assert.Equal(t, test.expected.Name, user.Name)
				assert.Equal(t, test.expected.Email, user.Email)
			}
		})
	}
}

func TestUserRepoImpl_GetByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name         string
		email        string
		dependency   func(ctrl *gomock.Controller) mongo.Client
		expectErr    bool
		expectErrMsg string
		expected     *domain.User
	}{
		{
			name:  "when user exists should return user",
			email: "test@google.com",
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().FindOne(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *driver.SingleResult {
						return driver.NewSingleResultFromDocument(&domain.User{
							ID:    "123",
							Name:  "Test User",
							Email: "test@google.com",
						}, nil, nil)
					})

				return client
			},
			expected: &domain.User{
				ID:    "123",
				Name:  "Test User",
				Email: "test@google.com",
			},
		},
		{
			name:  "when user not found should return error",
			email: "notfound@google.com",
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(driver.NewSingleResultFromDocument(nil, nil, nil))

				return client
			},
			expectErr:    true,
			expectErrMsg: "user not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := test.dependency(ctrl)
			r := repo.NewUserRepo(m)
			defer repo.ResetUserRepo()

			user, err := r.GetByEmail(context.Background(), test.email)

			if test.expectErr {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expectErrMsg)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, test.expected.ID, user.ID)
				assert.Equal(t, test.expected.Name, user.Name)
				assert.Equal(t, test.expected.Email, user.Email)
			}
		})
	}
}

func TestUserRepoImpl_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name         string
		user         *domain.User
		dependency   func(ctrl *gomock.Controller) mongo.Client
		expectErr    bool
		expectErrMsg string
	}{
		{
			name: "when update is successful",
			user: &domain.User{ID: "123", Name: "Updated", Email: "updated@google.com"},
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)

				return client
			},
		},
		{
			name: "when update fails",
			user: &domain.User{ID: "123", Name: "Updated", Email: "updated@google.com"},
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

				return client
			},
			expectErr:    true,
			expectErrMsg: assert.AnError.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := test.dependency(ctrl)
			r := repo.NewUserRepo(m)
			defer repo.ResetUserRepo()

			err := r.Update(context.Background(), test.user)

			if test.expectErr {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expectErrMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserRepoImpl_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name         string
		id           string
		dependency   func(ctrl *gomock.Controller) mongo.Client
		expectErr    bool
		expectErrMsg string
	}{
		{
			name: "when delete is successful",
			id:   "123",
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().DeleteOne(gomock.Any(), gomock.Any()).Return(nil, nil)

				return client
			},
		},
		{
			name: "when delete fails",
			id:   "123",
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().DeleteOne(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

				return client
			},
			expectErr:    true,
			expectErrMsg: assert.AnError.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := test.dependency(ctrl)
			r := repo.NewUserRepo(m)
			defer repo.ResetUserRepo()

			err := r.Delete(context.Background(), test.id)

			if test.expectErr {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expectErrMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserRepoImpl_Count(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name         string
		dependency   func(ctrl *gomock.Controller) mongo.Client
		expectErr    bool
		expectErrMsg string
		expected     int
	}{
		{
			name: "when count is successful",
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().CountDocuments(gomock.Any(), gomock.Any()).Return(int64(5), nil)

				return client
			},
			expected: 5,
		},
		{
			name: "when count fails",
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().CountDocuments(gomock.Any(), gomock.Any()).Return(int64(0), assert.AnError)

				return client
			},
			expectErr:    true,
			expectErrMsg: assert.AnError.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := test.dependency(ctrl)
			r := repo.NewUserRepo(m)
			defer repo.ResetUserRepo()

			count, err := r.Count(context.Background())

			if test.expectErr {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expectErrMsg)
				assert.Equal(t, 0, count)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, count)
			}
		})
	}
}

func TestUserRepoImpl_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name         string
		page, size   int64
		dependency   func(ctrl *gomock.Controller) mongo.Client
		expectErr    bool
		expectErrMsg string
		expected     []*domain.User
	}{
		{
			name: "when users exist should return user list",
			page: 0, size: 2,
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)
				cur := mock.NewMockCursor(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(cur, nil)
				cur.EXPECT().Next(gomock.Any()).Return(true).Times(2)
				cur.EXPECT().Next(gomock.Any()).Return(false)
				cur.EXPECT().Decode(gomock.Any()).DoAndReturn(
					func(v interface{}) error {
						u := v.(*domain.User)
						u.ID = "1"
						u.Name = "User1"
						u.Email = "user1@google.com"
						return nil
					},
				).Times(1)
				cur.EXPECT().Decode(gomock.Any()).DoAndReturn(
					func(v interface{}) error {
						u := v.(*domain.User)
						u.ID = "2"
						u.Name = "User2"
						u.Email = "user2@google.com"
						return nil
					},
				).Times(1)
				cur.EXPECT().Close(gomock.Any()).Return(nil)

				return client
			},
			expected: []*domain.User{
				{ID: "1", Name: "User1", Email: "user1@google.com"},
				{ID: "2", Name: "User2", Email: "user2@google.com"},
			},
		},
		{
			name: "when find returns error",
			page: 0, size: 2,
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)
				col.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

				return client
			},
			expectErr:    true,
			expectErrMsg: assert.AnError.Error(),
		},
		{
			name: "when decode returns error",
			page: 0, size: 2,
			dependency: func(ctrl *gomock.Controller) mongo.Client {
				client := mock.NewMockClient(ctrl)
				db := mock.NewMockDatabase(ctrl)
				col := mock.NewMockCollection(ctrl)

				cur := mock.NewMockCursor(ctrl)
				cur.EXPECT().Next(gomock.Any()).Return(true)
				cur.EXPECT().Close(gomock.Any()).Return(nil)
				cur.EXPECT().Decode(gomock.Any()).Return(assert.AnError)

				client.EXPECT().Database(gomock.Any()).Return(db)
				db.EXPECT().Collection(gomock.Any()).Return(col)

				col.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(cur, nil)

				return client
			},
			expectErr:    true,
			expectErrMsg: assert.AnError.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := test.dependency(ctrl)
			r := repo.NewUserRepo(m)
			defer repo.ResetUserRepo()

			users, err := r.List(context.Background(), test.page, test.size)

			if test.expectErr {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expectErrMsg)
				assert.Nil(t, users)
			} else {
				assert.NoError(t, err)
				assert.Len(t, users, len(test.expected))
				for i, u := range users {
					assert.Equal(t, test.expected[i].ID, u.ID)
					assert.Equal(t, test.expected[i].Name, u.Name)
					assert.Equal(t, test.expected[i].Email, u.Email)
				}
			}
		})
	}
}
