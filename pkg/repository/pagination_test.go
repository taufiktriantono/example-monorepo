package repository_test

import (
	"encoding/base64"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/taufiktriantono/api-first-monorepo/pkg/repository"
)

func TestEncodeDecodeCursor(t *testing.T) {
	cursor := repository.Cursor{
		CreatedAt: time.Now().Format(time.RFC3339Nano),
		ID:        "template-123",
	}

	encoded, err := repository.EncodeCursor(cursor)
	assert.NoError(t, err)
	assert.NotEmpty(t, encoded)

	decoded, err := repository.DecodeCursor(encoded)
	assert.NoError(t, err)
	assert.Equal(t, cursor.ID, decoded.ID)
	assert.Equal(t, cursor.CreatedAt, decoded.CreatedAt)
}

func TestDecodeCursor_InvalidBase64(t *testing.T) {
	_, err := repository.DecodeCursor("!!!invalid-base64$$$")
	assert.Error(t, err)
}

func TestBuildCursorPageInfo(t *testing.T) {
	type Dummy struct {
		ID        string
		CreatedAt time.Time
	}

	time1 := time.Now()
	time2 := time1.Add(time.Second)
	time3 := time1.Add(2 * time.Second)

	dummyData := []*Dummy{
		{ID: "1", CreatedAt: time1},
		{ID: "2", CreatedAt: time2},
		{ID: "3", CreatedAt: time3},
	}

	extractor := func(d *Dummy) string {
		encoded, _ := repository.EncodeCursor(repository.Cursor{
			ID:        d.ID,
			CreatedAt: d.CreatedAt.Format(time.RFC3339Nano),
		})
		return encoded
	}

	pageInfo := repository.BuildCursorPageInfo(dummyData, 2, extractor)

	assert.True(t, pageInfo.HasMore)
	decoded, err := base64.StdEncoding.DecodeString(pageInfo.NextCursor)
	assert.NoError(t, err)
	assert.True(t, strings.Contains(string(decoded), "3"))
}

func TestBuildCursorPageInfo_NoMore(t *testing.T) {
	type Dummy struct {
		ID        string
		CreatedAt time.Time
	}

	dummyData := []*Dummy{
		{ID: "1", CreatedAt: time.Now()},
		{ID: "2", CreatedAt: time.Now()},
	}

	extractor := func(d *Dummy) string {
		encoded, _ := repository.EncodeCursor(repository.Cursor{
			ID:        d.ID,
			CreatedAt: d.CreatedAt.Format(time.RFC3339Nano),
		})
		return encoded
	}

	pageInfo := repository.BuildCursorPageInfo(dummyData, 2, extractor)

	assert.False(t, pageInfo.HasMore)
	assert.Empty(t, pageInfo.NextCursor)
}
