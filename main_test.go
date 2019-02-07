package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVideoId(t *testing.T) {
	var link = []struct {
		TestLink string
		VideoId  string
	}{
		{"Rag'n'Bone Man - Human (Official Video) https://www.youtube.com/watch?v=L3wKzyIN1yk", "L3wKzyIN1yk"},
		{"https://www.youtube.com/watch?v=L3wKzyIN1y", "fail"},
		{"Bon Jovi - It's My Life https://youtu.be/vx2u5uUu3DE", "vx2u5uUu3DE"},
		{"https://www.youtube.com/watch?v=", "fail"},
		{"https://www.youtube.com/watch?v=0I647GU3Jsc", "0I647GU3Jsc"},
		{"https://www.youtube.com/watch?v0I647GU3Jsc", "fail"},
		{"https://youtu.be/vx2u5uUuDE", "fail"},
	}

	for _, tt := range link {
		checkLink, _ := getVideoID(tt.TestLink)
		assert.Equal(t, checkLink, tt.VideoId, "they should be equal")

	}
}

func TestYoutubeDown(t *testing.T) {
	testCases := []string{
		"https://www.youtube.com/watch?v=hT_nvWreIhg",
	}
	for _, url := range testCases {
		videoId, _ := getVideoID(url)
		youtubeDown(url)
		if _, err := os.Stat(videoId + ".mp3"); os.IsNotExist(err) {
			t.Error("Файла не существует")
		}

	}
}
