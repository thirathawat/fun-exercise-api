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
}
