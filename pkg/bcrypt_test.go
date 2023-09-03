package pkg

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var password = "abcd1234"

func TestHashPassword(t *testing.T) {
	hashPassword, err := HashPassword(password)
	assert.NoError(t, err, "error occured")
	assert.NotEqual(t, password, hashPassword, "password is not hashing")
}

func TestVerifyPassword(t *testing.T) {
	hashPassword, _ := HashPassword(password)
	t.Run("Verify success", func(t *testing.T) {
		err := VerifyPassword(hashPassword, password)
		assert.Nil(t, err, "password incorrect")
	})
	t.Run("Verify fail", func(t *testing.T) {
		err := VerifyPassword(hashPassword, "abcd")
		log.Println(err)
		assert.NotNil(t, err, "password correct")
	})
}
