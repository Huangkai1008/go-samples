package blogposts_test

import (
	"github.com/stretchr/testify/assert"
	blogposts "go-samples/learn_go_with_tests/readingfiles"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md": {Data: []byte("hi")},
		"hello world2md": {Data: []byte("hola")},
	}
	expected := len(fs)

	posts, err := blogposts.NewPostsFromFS(fs)

	assert.NoError(t, err)
	assert.Equal(t, len(posts), expected)
}
