package testbehavior_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/testenvs"
)

func TestMain(m *testing.M) {
	exitVal := execute(m)
	os.Exit(exitVal)
}

func execute(m *testing.M) int {
	defer cleanup()
	setup()
	exitVal := m.Run()
	return exitVal
}

func setup() {
	err := initialize()
	if err != nil {
		cleanup()
		os.Exit(1)
	}
}

func initialize() error {
	enableAcceptance := os.Getenv(fmt.Sprintf("%v", testenvs.EnableAcceptance))
	if enableAcceptance == "" {
		return fmt.Errorf("acceptance tests cannot be run; set %s env to run them", testenvs.EnableAcceptance)
	}

	requireGeneratedRandomValue := os.Getenv(fmt.Sprintf("%v", testenvs.RequireGeneratedRandomValue))
	if requireGeneratedRandomValue == "" {
		return fmt.Errorf("generated random value is required for this test run; set %s env", testenvs.RequireGeneratedRandomValue)
	}

	return nil
}

func cleanup() {
}
