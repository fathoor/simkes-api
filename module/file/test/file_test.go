package test

import (
	"bytes"
	_ "embed"
	"github.com/fathoor/simkes-api/core/config"
	"github.com/fathoor/simkes-api/core/provider"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ProvideTestApp() *fiber.App {
	var (
		cfg = config.ProvideConfig()
		app = config.ProvideApp(cfg)
	)

	provider.ProvideFile(app)

	return app
}

var app = ProvideTestApp()

//go:embed resource/default.png
var fileTest []byte

func TestFile_Upload(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	file, err := writer.CreateFormFile("file", "default.png")
	assert.Nil(t, err)
	_, err = file.Write(fileTest)
	assert.Nil(t, err)

	err = writer.WriteField("type", "image")
	assert.Nil(t, err)

	err = writer.Close()
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/file", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}

func TestFile_Download(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/api/v1/file/image/4d32b23d-0927-457d-ab12-25bd632224b1.png/download", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}

func TestFile_Get(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/api/v1/file/image/4d32b23d-0927-457d-ab12-25bd632224b1.png", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}

func TestFile_Delete(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/api/v1/file/image/4d32b23d-0927-457d-ab12-25bd632224b1.png", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusNoContent, response.StatusCode)
}
