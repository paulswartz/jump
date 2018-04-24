package cmd

import (
	"os"
	p "path"
	"strings"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/cli"
	s "github.com/gsamokovarov/jump/scoring"
)

func Test_hintCmd_short(t *testing.T) {
	p1 := p.Join(td, "web-console")
	p2 := p.Join(td, "/client/website")

	conf := &testConfig{
		Entries: s.Entries{
			&s.Entry{p2, &s.Score{Weight: 90, Age: s.Now}},
			&s.Entry{p1, &s.Score{Weight: 100, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, hintCmd(cli.Args{"webcons"}, conf))
	})

	lines := strings.Fields(output)
	assert.Len(t, 1, lines)

	assert.Equal(t, p1, lines[0])
}
