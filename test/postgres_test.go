package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/backend-school/tracing/internal/models"
	"github.com/backend-school/tracing/internal/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func Test_postgres(t *testing.T) {
	ctx := context.Background()

	psql, err := storage.NewPostgres(ctx)
	assert.NoError(t, err)

	t.Run("insert tokens", func(t *testing.T) {
		query := `INSERT INTO token (id, name, network_id, currency_id, is_active) VALUES (@ID, @name, @netID, @currID, @isActive)`
		args := pgx.NamedArgs{
			"ID":       uuid.NewString(),
			"name":     "token0",
			"netID":    0,
			"currID":   0,
			"isActive": true,
		}
		_, err := psql.Exec(ctx, query, args)
		assert.NoError(t, err)
	})

	t.Run("list all tokens", func(t *testing.T) {
		sql := "select * from token;"
		data, err := storage.RunSelect[models.Token](ctx, psql, sql)

		fmt.Println(data)
		assert.NoError(t, err)
	})
}
