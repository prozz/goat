package api_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"goat/pkg/ai/mock"
	"goat/pkg/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiGenerator(t *testing.T) {

	t.Run("200", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		generator := mock.NewMockGenerator(ctrl)
		generator.EXPECT().RandomTitle().Return("foo!")

		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/fakenews", nil)
		require.NoError(t, err)

		api.NewApp(generator).ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "foo!", w.Body.String())
	})

	t.Run("500", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		generator := mock.NewMockGenerator(ctrl)
		generator.EXPECT().RandomTitle().Return("")

		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/fakenews", nil)
		require.NoError(t, err)

		api.NewApp(generator).ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
		assert.Empty(t, w.Body.String())
	})

}
