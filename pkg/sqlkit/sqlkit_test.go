package sqlkit_test

import (
	"testing"

	"github.com/KKGo-Software-engineering/fun-exercise-api/pkg/sqlkit"
	"github.com/stretchr/testify/assert"
)

func TestQueryBuilder(t *testing.T) {
	t.Run("Select", func(t *testing.T) {
		builder := sqlkit.NewQueryBuilder().
			Select("name", "age").
			From("users").
			Where("name", "=", "John").
			Where("age", ">", 18)

		query, args := builder.Build()

		assert.Equal(t, "SELECT name, age FROM users WHERE name = $1 AND age > $2", query)
		assert.Equal(t, []any{"John", 18}, args)
	})

	t.Run("Insert", func(t *testing.T) {
		builder := sqlkit.NewQueryBuilder().
			Insert().
			Table("users").
			Set("name", "John").
			Set("age", 25).
			Returning("id")

		query, args := builder.Build()

		assert.Equal(t, "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id", query)
		assert.Equal(t, []any{"John", 25}, args)
	})
}
