package pkg

import (
	"ninja1cak/coffeshop-be/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

var status = [10]int{
	200, 201, 400,
	401, 403, 404,
	500, 501, 304,
	100,
}

var desc = [10]string{
	"OK", "Created", "Bad Request",
	"Unauthorized", "Forbidden", "Not Found",
	"Internal Server Error", "Bad Gateway", "Not Modified",
	"",
}

func TestGetStatus(t *testing.T) {
	for i := 0; i < len(status); i++ {
		description := getStatus(status[i])
		assert.Equal(t, desc[i], description, "desc isn't match")
	}
}

func TestNewResponse(t *testing.T) {
	var dummyData = config.Result{
		Data:    "data",
		Message: "Data found",
	}

	t.Run("check message", func(t *testing.T) {
		res := NewResponse(201, &dummyData)
		assert.NotNil(t, res.Message, "input data success")
		assert.Equal(t, "Data found", res.Description)
	})

	var dummyData1 = config.Result{
		Data: "data",
		Meta: "meta",
	}

	t.Run("check data input", func(t *testing.T) {
		res := NewResponse(201, &dummyData1)
		assert.NotNil(t, res.Data, "empty value or nil")
		assert.Equal(t, "data", res.Data, "difference value")

	})

	t.Run("check when error", func(t *testing.T) {
		res := NewResponse(400, &dummyData1)
		assert.NotNil(t, res.Description, "empty value or nil")
		assert.Equal(t, "data", res.Description, "difference value")

	})

	t.Run("check meta", func(t *testing.T) {
		res := NewResponse(201, &dummyData1)
		assert.NotNil(t, res.Meta, "empty value or nil")
		assert.Equal(t, "meta", res.Meta, "difference value")

	})
}
