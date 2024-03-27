package timekit_test

import (
	"testing"
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/pkg/timekit"
	"github.com/stretchr/testify/assert"
)

func TestFreezeAndUnfreeze(t *testing.T) {
	t.Cleanup(func() {
		timekit.Unfreeze()
	})

	t1 := timekit.Now()

	time.Sleep(100 * time.Millisecond)
	timekit.Freeze()
	timeFroze := timekit.Now()
	assert.True(t, timeFroze.After(t1))

	time.Sleep(100 * time.Millisecond)
	assert.False(t, timekit.Now().After(timeFroze))

	timekit.Unfreeze()
	time.Sleep(100 * time.Millisecond)
	assert.True(t, timekit.Now().After(timeFroze))
}
