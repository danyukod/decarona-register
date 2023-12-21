package persistence

import (
	"context"
	"fmt"
	document2 "github.com/danyukod/decarona-register/internal/domain/document"
	"github.com/danyukod/decarona-register/internal/domain/user"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"testing"
)

var mongoClient *mongo.Client

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not ping docker: %s", err)
	}

	resource, err := pool.Run("mongo", "latest", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		mongoClient, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://localhost:%s", resource.GetPort("27017/tcp")),
			),
		)
		if err != nil {
			return err
		}
		return mongoClient.Ping(context.TODO(), nil)
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestUserPersistence_Save(t *testing.T) {
	t.Run("should save user", func(t *testing.T) {
		ctx := context.Background()
		persistence := NewUserPersistence(mongoClient)
		userMock, err := getUserMock()
		assert.Nil(t, err)
		assert.NotNil(t, userMock)
		save, err := persistence.Save(ctx, userMock)
		assert.Nil(t, err)
		assert.NotNil(t, save)
	})

}

func getUserMock() (user.IUser, error) {
	document, err := document2.FromText(document2.CPF, "39357160876")
	if err != nil {
		return nil, err
	}
	return user.NewUser(
		"Danilo",
		"dyk@email.com",
		"M",
		"123456",
		append(make([]document2.IDocument, 0), document),
	)
}
