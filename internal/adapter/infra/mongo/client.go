//go:generate mockgen -source=client.go -destination=mock/client.go -package=mock
package mongo

import (
	"context"
	"go-hexagonal-architecture/internal/adapter/config"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type Client interface {
	Connect(ctx context.Context) error
	Database(name string, opts ...*options.DatabaseOptions) Database // <-- return Database interface
	Disconnect(ctx context.Context) error
	ListDatabaseNames(ctx context.Context, filter any, opts ...*options.ListDatabasesOptions) ([]string, error)
	ListDatabases(ctx context.Context, filter any, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error)
	NumberSessionsInProgress() int
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	StartSession(opts ...*options.SessionOptions) (mongo.Session, error)
	Timeout() *time.Duration
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
	UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error
	Watch(ctx context.Context, pipeline any, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
}

type Cursor interface {
	Next(context.Context) bool
	Decode(any) error
	Close(context.Context) error
}

type SingleResult interface {
	Decode(v interface{}) error
	DecodeBytes() (bson.Raw, error)
	Err() error
	Raw() (bson.Raw, error)
}

type Database interface {
	Aggregate(ctx context.Context, pipeline any, opts ...*options.AggregateOptions) (Cursor, error)
	Client() *mongo.Client
	Collection(name string, opts ...*options.CollectionOptions) Collection // <-- return Collection interface
	CreateCollection(ctx context.Context, name string, opts ...*options.CreateCollectionOptions) error
	CreateView(ctx context.Context, viewName string, viewOn string, pipeline any, opts ...*options.CreateViewOptions) error
	Drop(ctx context.Context) error
	ListCollectionNames(ctx context.Context, filter any, opts ...*options.ListCollectionsOptions) ([]string, error)
	ListCollectionSpecifications(ctx context.Context, filter any, opts ...*options.ListCollectionsOptions) ([]*mongo.CollectionSpecification, error)
	ListCollections(ctx context.Context, filter any, opts ...*options.ListCollectionsOptions) (Cursor, error)
	Name() string
	ReadConcern() *readconcern.ReadConcern
	ReadPreference() *readpref.ReadPref
	RunCommand(ctx context.Context, runCommand any, opts ...*options.RunCmdOptions) *mongo.SingleResult
	RunCommandCursor(ctx context.Context, runCommand any, opts ...*options.RunCmdOptions) (Cursor, error)
	Watch(ctx context.Context, pipeline any, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
	WriteConcern() *writeconcern.WriteConcern
}

type Collection interface {
	InsertOne(ctx context.Context, document any, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(ctx context.Context, documents []any, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	FindOne(ctx context.Context, filter any, opts ...*options.FindOneOptions) *mongo.SingleResult
	Find(ctx context.Context, filter any, opts ...*options.FindOptions) (Cursor, error)
	UpdateOne(ctx context.Context, filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter any, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter any, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	CountDocuments(ctx context.Context, filter any, opts ...*options.CountOptions) (int64, error)
	EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error)
	Aggregate(ctx context.Context, pipeline any, opts ...*options.AggregateOptions) (Cursor, error)
	BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
	Watch(ctx context.Context, pipeline any, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
	Drop(ctx context.Context) error
	Indexes() mongo.IndexView
}

var (
	_client *client
	_cOnce  = sync.Once{}
)

type client struct {
	*mongo.Client
}

// --- Wrapper for mongo.Client ---

// Ensure client implements Client interface
var _ Client = (*client)(nil)

func (c *client) Connect(ctx context.Context) error {
	return c.Client.Connect(ctx)
}

func (c *client) Database(name string, opts ...*options.DatabaseOptions) Database {
	return &database{c.Client.Database(name, opts...)}
}

func (c *client) Disconnect(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}

func (c *client) ListDatabaseNames(ctx context.Context, filter any, opts ...*options.ListDatabasesOptions) ([]string, error) {
	return c.Client.ListDatabaseNames(ctx, filter, opts...)
}

func (c *client) ListDatabases(ctx context.Context, filter any, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error) {
	return c.Client.ListDatabases(ctx, filter, opts...)
}

func (c *client) NumberSessionsInProgress() int {
	return c.Client.NumberSessionsInProgress()
}

func (c *client) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return c.Client.Ping(ctx, rp)
}

func (c *client) StartSession(opts ...*options.SessionOptions) (mongo.Session, error) {
	return c.Client.StartSession(opts...)
}

func (c *client) Timeout() *time.Duration {
	return c.Client.Timeout()
}

func (c *client) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return c.Client.UseSession(ctx, fn)
}

func (c *client) UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error {
	return c.Client.UseSessionWithOptions(ctx, opts, fn)
}

func (c *client) Watch(ctx context.Context, pipeline any, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return c.Client.Watch(ctx, pipeline, opts...)
}

// --- Wrapper for mongo.Cursor ---
type cursor struct {
	*mongo.Cursor
}

var _ Cursor = (*cursor)(nil)

func (c *cursor) Next(ctx context.Context) bool {
	return c.Cursor.Next(ctx)
}

func (c *cursor) Decode(val any) error {
	return c.Cursor.Decode(val)
}

func (c *cursor) Close(ctx context.Context) error {
	return c.Cursor.Close(ctx)
}

// --- Wrapper for mongo.Database ---
type database struct {
	*mongo.Database
}

var _ Database = (*database)(nil)

func (d *database) Aggregate(ctx context.Context, pipeline any, opts ...*options.AggregateOptions) (Cursor, error) {
	return d.Database.Aggregate(ctx, pipeline, opts...)
}

func (d *database) Client() *mongo.Client {
	return d.Database.Client()
}

func (d *database) Collection(name string, opts ...*options.CollectionOptions) Collection {
	return &collection{d.Database.Collection(name, opts...)}
}

func (d *database) CreateCollection(ctx context.Context, name string, opts ...*options.CreateCollectionOptions) error {
	return d.Database.CreateCollection(ctx, name, opts...)
}

func (d *database) CreateView(ctx context.Context, viewName string, viewOn string, pipeline any, opts ...*options.CreateViewOptions) error {
	return d.Database.CreateView(ctx, viewName, viewOn, pipeline, opts...)
}

func (d *database) Drop(ctx context.Context) error {
	return d.Database.Drop(ctx)
}

func (d *database) ListCollectionNames(ctx context.Context, filter any, opts ...*options.ListCollectionsOptions) ([]string, error) {
	return d.Database.ListCollectionNames(ctx, filter, opts...)
}

func (d *database) ListCollectionSpecifications(ctx context.Context, filter any, opts ...*options.ListCollectionsOptions) ([]*mongo.CollectionSpecification, error) {
	return d.Database.ListCollectionSpecifications(ctx, filter, opts...)
}

func (d *database) ListCollections(ctx context.Context, filter any, opts ...*options.ListCollectionsOptions) (Cursor, error) {
	return d.Database.ListCollections(ctx, filter, opts...)
}

func (d *database) Name() string {
	return d.Database.Name()
}

func (d *database) ReadConcern() *readconcern.ReadConcern {
	return d.Database.ReadConcern()
}

func (d *database) ReadPreference() *readpref.ReadPref {
	return d.Database.ReadPreference()
}

func (d *database) RunCommand(ctx context.Context, runCommand any, opts ...*options.RunCmdOptions) *mongo.SingleResult {
	return d.Database.RunCommand(ctx, runCommand, opts...)
}

func (d *database) RunCommandCursor(ctx context.Context, runCommand any, opts ...*options.RunCmdOptions) (Cursor, error) {
	return d.Database.RunCommandCursor(ctx, runCommand, opts...)
}

func (d *database) Watch(ctx context.Context, pipeline any, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return d.Database.Watch(ctx, pipeline, opts...)
}

func (d *database) WriteConcern() *writeconcern.WriteConcern {
	return d.Database.WriteConcern()
}

// --- Wrapper for mongo.Collection ---
type collection struct {
	*mongo.Collection
}

var _ Collection = (*collection)(nil)

func (c *collection) InsertOne(ctx context.Context, document any, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return c.Collection.InsertOne(ctx, document, opts...)
}

func (c *collection) InsertMany(ctx context.Context, documents []any, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return c.Collection.InsertMany(ctx, documents, opts...)
}

func (c *collection) FindOne(ctx context.Context, filter any, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return c.Collection.FindOne(ctx, filter, opts...)
}

func (c *collection) Find(ctx context.Context, filter any, opts ...*options.FindOptions) (Cursor, error) {
	return c.Collection.Find(ctx, filter, opts...)
}

func (c *collection) UpdateOne(ctx context.Context, filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return c.Collection.UpdateOne(ctx, filter, update, opts...)
}

func (c *collection) UpdateMany(ctx context.Context, filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return c.Collection.UpdateMany(ctx, filter, update, opts...)
}

func (c *collection) DeleteOne(ctx context.Context, filter any, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return c.Collection.DeleteOne(ctx, filter, opts...)
}

func (c *collection) DeleteMany(ctx context.Context, filter any, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return c.Collection.DeleteMany(ctx, filter, opts...)
}

func (c *collection) CountDocuments(ctx context.Context, filter any, opts ...*options.CountOptions) (int64, error) {
	return c.Collection.CountDocuments(ctx, filter, opts...)
}

func (c *collection) EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return c.Collection.EstimatedDocumentCount(ctx, opts...)
}

func (c *collection) Aggregate(ctx context.Context, pipeline any, opts ...*options.AggregateOptions) (Cursor, error) {
	return c.Collection.Aggregate(ctx, pipeline, opts...)
}

func (c *collection) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return c.Collection.BulkWrite(ctx, models, opts...)
}

func (c *collection) Watch(ctx context.Context, pipeline any, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return c.Collection.Watch(ctx, pipeline, opts...)
}

func (c *collection) Drop(ctx context.Context) error {
	return c.Collection.Drop(ctx)
}

func (c *collection) Indexes() mongo.IndexView {
	return c.Collection.Indexes()
}

func New(ctx context.Context, cfg *config.Config) (Client, error) {
	var err error
	_cOnce.Do(func() {
		var c *mongo.Client
		c, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
		if err != nil {
			log.Printf("MongoDB connect error: %v", err)
			return
		}

		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		if err = c.Ping(ctx, readpref.Primary()); err != nil {
			log.Printf("MongoDB ping error: %v", err)
			return
		}

		_client = &client{c}
	})

	return _client, nil
}

func Reset() {
	_cOnce = sync.Once{}
}
