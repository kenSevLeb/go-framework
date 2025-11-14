package set

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreadUnsafeSet_Add(t *testing.T) {
	set := ThreadSafe()
	//set := ThreadUnsafe()

	set.Add(1, 2, 3, 4)

	assert.True(t, set.Contain(1))
	set.Remove(2)
	assert.Equal(t, 3, set.Size())

	set.Clear()
	assert.Equal(t, 0, set.Size())
	assert.True(t, set.Empty())

	assert.False(t, set.Contain(5))
	set.Add(5)
	assert.True(t, set.Contain(5))
	fmt.Println(set.ToSlice())

	set.Add("hello", "world")
	fmt.Println(set.ToSlice())

	copySet := set.Duplicate()
	fmt.Println(copySet.ToSlice())
}
