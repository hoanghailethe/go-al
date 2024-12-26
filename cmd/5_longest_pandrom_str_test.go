package cmd_test

import (
	"HaiRepoGh/go-al/cmd"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsPanlindrome(t *testing.T) {
	t.Run("is Pand", func(t *testing.T) {
		assert.True(t, true, cmd.IsPanlindrome("abba"))
	})

	t.Run("is not Pand", func(t *testing.T) {
		assert.False(t, false, cmd.IsPanlindrome("ab"))
	})

}
