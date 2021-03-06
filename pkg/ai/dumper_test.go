package ai_test

import (
	"bytes"
	"errors"
	"goat/pkg/ai"
	"goat/pkg/hn"
	"goat/pkg/hn/mock"
	"io/ioutil"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type errWriter struct {
}

func (e errWriter) Write(p []byte) (int, error) {
	return 0, errors.New("boom")
}

func TestDumper(t *testing.T) {

	t.Run("writer error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockClient(ctrl)

		client.EXPECT().Top().Return([]int{1})
		client.EXPECT().Get(1).Return(hn.Item{Title: "A"})

		err := ai.NewDumper(client).Dump(&errWriter{})
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockClient(ctrl)

		client.EXPECT().Top().Return([]int{1, 2})
		client.EXPECT().Get(1).Return(hn.Item{Title: "A"})
		client.EXPECT().Get(2).Return(hn.Item{Title: "B"})

		var b bytes.Buffer
		err := ai.NewDumper(client).Dump(&b)
		require.NoError(t, err)
		assert.Equal(t, "A\nB\n", b.String())
	})
}

func TestDumperIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	f, err := ioutil.TempFile(os.TempDir(), "goat-dump-")
	require.NoError(t, err)
	defer f.Close()

	err = ai.NewDumper(hn.NewClient()).Dump(f)
	require.NoError(t, err)

	spew.Dump(f.Name())
}
