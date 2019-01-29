package main

import "testing"

func TestGetRightLink(t *testing.T) {
	var link = []struct {
		TestLink    string
		CorrectLink string
	}{
		{"Rag'n'Bone Man - Human (Official Video) https://www.youtube.com/watch?v=L3wKzyIN1yk", "https://www.youtube.com/watch?v=L3wKzyIN1yk"},
		{"Bon Jovi - It's My Life https://youtu.be/vx2u5uUu3DE", "https://youtu.be/vx2u5uUu3DE"},
		{"https://www.youtube.com/watch?v=0I647GU3Jsc", "https://www.youtube.com/watch?v=0I647GU3Jsc"},
	}

	for _, tt := range link {
		checkLink := getRightLink(tt.TestLink)
		if tt.CorrectLink != checkLink {
			t.Errorf("Не правильная передача файла (%s) != (%s)", tt.CorrectLink, tt.TestLink)
		}
	}
}
