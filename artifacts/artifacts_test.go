package artifacts

import (
	"errors"
	"testing"

	"generics/genmock"

	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	expectedArtifact := Artifact{Name: "dootdoot", S3Key: "its/in/here/i/promise.tar.gz"}
	database := &fakeDatabase{}
	database.F = func(id string) (Artifact, error) {
		return expectedArtifact, nil
	}

	result, err := Query("0118999", database)

	assert.Nil(t, err)

	assert.Equal(t, expectedArtifact, result)

	// using contains
	assert.Contains(t, database.Calls, genmock.Mock1ArgCall[string]{Arg0: "0118999"})

	// or if it's comparable we can do this
	assert.True(t, genmock.WasCalledWith1Arg(database.Calls, "0118999"))
}

func TestDatabaseErr(t *testing.T) {
	expectedErr := errors.New("oh no")
	database := &fakeDatabase{}
	database.F = func(id string) (Artifact, error) {
		return Artifact{}, expectedErr
	}

	_, err := Query("0118999", database)

	assert.Equal(t, expectedErr, err)
}

type fakeDatabase struct {
	genmock.Mock1Arg2Results[string, Artifact, error]
}

func (d *fakeDatabase) Find(id string) (Artifact, error) {
	return d.Call(id)
}
