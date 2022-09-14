package util_test

import (
	"testing"

	"github.com/whutchinson98/secretsmanager-cli/internal/util"
)

func TestBuildEnvStringFromSecret(t *testing.T) {
	t.Run("Creates a correctly formatted env file string from a secretMap", func(t *testing.T) {
		secretMap := make(map[string]interface{})
		secretMap["FOO"] = "baz"
		secretMap["BAR"] = "boo"
		secretMap["BAZ"] = "foo"

		res := util.BuildEnvStringFromSecret(secretMap)

		want := "FOO=baz\nBAR=boo\nBAZ=foo\n"
		if res != want {
			t.Errorf("expected \"%s\" got \"%s\"\n", want, res)
		}
	})
}
