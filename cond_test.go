package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder_Cond(t *testing.T) {
	assert := assert.New(t)
	{
		cond, err := ToCondition(CondTopic{"all"})
		assert.Equal("'all' in topics", cond)
		assert.NoError(err)
	}

	{
		cond, err := ToCondition(CondTopic{"test1"}.And(CondTopic{"test2"}))
		assert.Equal("'test1' in topics && 'test2' in topics", cond)
		assert.NoError(err)
	}

	{
		cond, err := ToCondition(CondTopic{"test1"}.Or(CondTopic{"test2"}))
		assert.Equal("'test1' in topics || 'test2' in topics", cond)
		assert.NoError(err)
	}

	{
		cond, err := ToCondition(CondTopic{"test1"}.And(CondTopic{"test2"}.And(CondTopic{"test3"})))
		assert.Equal("'test1' in topics && 'test2' in topics && 'test3' in topics", cond)
		assert.NoError(err)
	}

	{
		cond, err := ToCondition(CondTopic{"test1"}.And(CondTopic{"test2"}.Or(CondTopic{"test3"})))
		assert.Equal("'test1' in topics && ('test2' in topics || 'test3' in topics)", cond)
		assert.NoError(err)
	}
}
