package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var user_id string = "1"
var email string = "hauzan41200@gmail.com"
var role string = "user"

func TestVerifyToken(t *testing.T) {
	jwtt := NewToken(user_id, email, role)
	token, _ := jwtt.Generate()

	t.Run("token success verify", func(t *testing.T) {
		_, err := VerifyToken(token)
		assert.Nil(t, err, "verify token fail")
	})

	t.Run("token fail", func(t *testing.T) {
		_, err := VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
		assert.NotNil(t, err, "verify success")
	})

	t.Run("check data token", func(t *testing.T) {
		decode, _ := VerifyToken(token)
		assert.Equal(t, email, decode.Email, "email not same")
	})
}

func TestGenerate(t *testing.T) {
	jwtt := NewToken(user_id, email, role)
	_, err := jwtt.Generate()
	assert.NoError(t, err, "Error has occured when generated token")
}
