package appconfig

import "testing"

var data = `
APP_PORT = 3001

# dbs
# postgres
POSTGRES_DATABASE_NAME=nest-prototype
# for windows
POSTGRES_DATABASE_HOST=localhost
POSTGRES_DATABASE_PORT=5433
POSTGRES_DATABASE_USERNAME=user
POSTGRES_DATABASE_PASSWORD=password
`
var APP_PORT__Value = 3001
var POSTGRES_DATABASE_NAME__Value = "golang-prototype"

var POSTGRES_DATABASE_HOST__Value = "localhost"
var POSTGRES_DATABASE_PORT__Value = 5433
var POSTGRES_DATABASE_USERNAME__Value = "user"
var POSTGRES_DATABASE_PASSWORD__Value = "password"

func TestSomeTestEnvFile(t *testing.T) {
	env := ReadEnvFile("./.env.mock.dev")

	if env.APP_PORT != APP_PORT__Value {
		t.Errorf("Should be %d but got %d	", APP_PORT__Value, env.APP_PORT)
	}

	if env.POSTGRES_DATABASE_NAME != POSTGRES_DATABASE_NAME__Value {
		t.Errorf("Should be %s but got %s", POSTGRES_DATABASE_NAME__Value, env.POSTGRES_DATABASE_NAME)
	}

	if env.POSTGRES_DATABASE_HOST != POSTGRES_DATABASE_HOST__Value {
		t.Errorf("Should be %s but got %s", POSTGRES_DATABASE_HOST__Value, env.POSTGRES_DATABASE_HOST)
	}

	if env.POSTGRES_DATABASE_PORT != POSTGRES_DATABASE_PORT__Value {
		t.Errorf("Should be %d but got %d", POSTGRES_DATABASE_PORT__Value, env.POSTGRES_DATABASE_PORT)
	}

	if env.POSTGRES_DATABASE_USERNAME != POSTGRES_DATABASE_USERNAME__Value {
		t.Errorf("Should be %s but got %s", POSTGRES_DATABASE_USERNAME__Value, env.POSTGRES_DATABASE_USERNAME)
	}

	if env.POSTGRES_DATABASE_PASSWORD != POSTGRES_DATABASE_PASSWORD__Value {
		t.Errorf("Should be %s but got %s", POSTGRES_DATABASE_PASSWORD__Value, env.POSTGRES_DATABASE_PASSWORD)
	}

}
