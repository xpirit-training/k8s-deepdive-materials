package database

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	guuid "github.com/google/uuid"
	"github.com/pathogende/docustore/pkg/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var (
	// ErrDocumentNotFound is returned when a document is not found in the database.
	ErrDocumentNotFound = fmt.Errorf("document not found")
)

const collectionName = "docs"

// DB is a database instance.
type DB struct {
	client  *mongo.Client
	col     *mongo.Collection
	options MongoOptions
}

type MongoOptions struct {
	uri     string
	dbName  string
	colName string
}

type DBChecker interface {
	Ping(ctx context.Context) error
}

// Implement the DBChecker interface for *DB
func (db *DB) Ping(ctx context.Context) error {
	return db.client.Ping(ctx, nil)
}

func getMongoOptions() MongoOptions {
	return MongoOptions{
		uri:     viper.GetString(config.MONGODB_URL),
		dbName:  viper.GetString(config.MONGODB_DATABASE),
		colName: viper.GetString(config.MONGODB_COLLECTION),
	}
}

// NewDB creates a new database instance.
func NewDB() (*DB, error) {

	options := getMongoOptions()
	client, err := initMongoClient()
	if err != nil {
		return nil, err
	}

	col := client.Database(options.dbName).Collection(options.colName)
	return &DB{
		client:  &client,
		col:     col,
		options: options,
	}, nil
}

func initMongoClient() (mongo.Client, error) {
	url := viper.GetString(config.MONGODB_URL)

	credential := options.Credential{
		AuthMechanism: viper.GetString(config.MONGODB_AUTH_MECHANISM),
		Username:      viper.GetString(config.MONGODB_USER),
		Password:      viper.GetString(config.MONGODB_USER_PASSWORD),
		AuthSource:    viper.GetString(config.MONGODB_DATABASE),
	}

	var clientOpts = &options.ClientOptions{}
	// if config TLS_ENABLE is set to false this will be null, no TLS will be used (client-side)
	if viper.GetBool(config.ENABLE_TLS) {
		tlsConfig, err := getCustomTLSConfig(viper.GetString(config.TLS_CA_BUNDLE_PATH))
		if err != nil {
			return mongo.Client{}, err
		}
		clientOpts = options.Client().ApplyURI(url).SetAuth(credential).SetTLSConfig(tlsConfig)
	} else if viper.GetString(config.MONGODB_USER_PASSWORD) == "" {
		clientOpts = options.Client().ApplyURI(url)
	} else {
		clientOpts = options.Client().ApplyURI(url).SetAuth(credential)
	}

	clientOpts.SetRetryWrites(false)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return mongo.Client{}, err
	}

	defer ctx.Done()

	return *client, nil
}

func IsMongoDBConnected(db DBChecker) bool {
	// Create a timeout context for the ping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping the MongoDB server
	err := db.Ping(ctx)
	return err == nil
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {

	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("Failed parsing pem file")
	}

	return tlsConfig, nil
}

// CreateDocument inserts a new document into the database.
func (db *DB) CreateDocument(document Document) error {
	// If id is not set create one.
	// Alternatively mongo can manage the id wich would be preffered in most cases
	// I wanted to try it out how it could work without being managed by mongo
	if document.ID == "" {
		document.ID = guuid.New().String()
	}
	_, err := db.col.InsertOne(context.TODO(), document)
	if err != nil {
		return err
	}
	return nil
}

// GetDocument retrieves a document from the database by its ID.
func (db *DB) GetDocument(id string) (Document, error) {
	var document Document
	err := db.col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&document)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Document{}, ErrDocumentNotFound
		}
		return Document{}, err
	}
	return document, nil
}

// UpdateOrCreateDocument updates a document in the database.
func (db *DB) UpdateOrCreateDocument(id string, document Document) error {
	res, err := db.col.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": document})
	if res.MatchedCount == 0 {
		return db.CreateDocument(document)
	}
	if err != nil {
		return err
	}

	return nil
}

// DeleteDocument deletes a document from the database.
func (db *DB) DeleteDocument(id string) error {
	res, err := db.col.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return ErrDocumentNotFound
	}
	return nil
}

// ListDocumentIDs retrieves all document IDs from the database.
func (db *DB) ListDocumentIDs() ([]string, error) {
	// Set up a context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use the Find method with a projection to retrieve only the _id field.
	opts := options.Find().SetProjection(bson.M{"_id": 1})
	cursor, err := db.col.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and collect the _id values.
	var ids []string
	for cursor.Next(ctx) {
		var result struct {
			ID string `bson:"_id"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		ids = append(ids, result.ID)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}
