package util_test

import (
	"testing"

	"github.com/whutchinson98/secretsmanager-cli/internal/util"
)

func TestBuildJsonStringFromEnv(t *testing.T) {
	t.Run("Creates a correctly formatted JSON string from an env file", func(t *testing.T) {
		envFileData := []string{"FOO=baz", "BAR=boo", "BAZ=foo"}

		res, _ := util.BuildJsonStringFromEnv(envFileData)

		want := "{\"FOO\":\"baz\",\"BAR\":\"boo\",\"BAZ\":\"foo\"}"
		if res != want {
			t.Errorf("expected \"%s\" got \"%s\"\n", want, res)
		}
	})
	t.Run("Errors if an env file has a key with no value", func(t *testing.T) {
		envFileData := []string{"FOO=baz", "BAR", "BAZ=foo"}

		_, err := util.BuildJsonStringFromEnv(envFileData)

		if err == nil {
			t.Errorf("expected to error")
		}

		expected := "entry in env file does not have a value BAR"
		if err.Error() != expected {
			t.Errorf("incorrect error message got \"%s\" wanted \"%s\"\n", err.Error(), expected)
		}
	})
}
