package task1

import (
	"testing"
)

func TestOpenSqlConnectionsPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Паника не сработала с неправильными параметрами")
		}
	}()
	UserEnv = "panic"
	PasswordEnv = "panic"
	HostEnv = "panic"
	DBEnv = "panic"
	OpenSqlConnections()
}
