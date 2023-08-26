package vespa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {

	t.Run("configが設定されていない場合、エラーを返す、", func(t *testing.T) {
		_, err := NewTypedClient(nil)
		assert.Equal(t, "conf is required", err.Error())
	})

	t.Run("configのURlが設定されていない場合、エラーを返す", func(t *testing.T) {
		conf := Config{
			Url: "",
		}
		_, err := NewTypedClient(&conf)
		assert.Equal(t, "url is not determined", err.Error())
	})
}
