package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/backend-school/tracing/internal/models"
	"github.com/backend-school/tracing/internal/storage"
	"github.com/stretchr/testify/assert"
)

func Test_postgres(t *testing.T) {
	ctx := context.Background()

	psql, err := storage.NewPostgres(ctx)
	assert.NoError(t, err)

	t.Run("list all tokens", func(t *testing.T) {
		sql := "select * from token;"
		data, err := storage.RunSelect[models.Token](ctx, psql, sql)

		fmt.Println(data)
		assert.NoError(t, err)
	})
}
