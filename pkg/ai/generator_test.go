package ai_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"goat/pkg/ai"
	"io/ioutil"
	"strings"
	"testing"
)

func TestSimpleGenerator(t *testing.T) {
	testcases := []struct {
		FilePath, ExpectedPrefix string
	}{
		{FilePath: "testdata/simple-1.txt", ExpectedPrefix: "Ala ma kota."},
		{FilePath: "testdata/simple-2.txt", ExpectedPrefix: "Ala ma"},
	}

	for _, tc := range testcases {
		buf, err := ioutil.ReadFile(tc.FilePath)
		if assert.NoError(t, err) {
			g := ai.NewGenerator(bytes.NewReader(buf))
			assert.True(t, strings.HasPrefix(g.RandomTitle(), tc.ExpectedPrefix))
		}
	}
}

func TestGeneratorBruteForce(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/simple-2.txt")
	require.NoError(t, err)

	g := ai.NewGenerator(bytes.NewReader(buf))

	for i := 0; i < 100; i++ {
		title := g.RandomTitle()
		assert.True(t, strings.HasPrefix(title, "Ala ma "))
		assert.True(t, strings.HasSuffix(title, "psa.") || strings.HasSuffix(title, "kota."))
	}
}
