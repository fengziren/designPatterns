package player1

import "fmt"

type AudioPlayer struct {
	mediaAdapter
}

func (a *AudioPlayer) Play(audioType string, fileName string) {
	if audioType == "mp3" {
		fmt.Println("Playing mp3 file. Name: " + fileName)
	} else if audioType == "vlc" || audioType == "mp4" {
		a.GetMediaAdapter(audioType)
		a.mediaAdapter.Play(audioType, fileName)
	} else {
		fmt.Println("Invalid media. " + audioType + " format not supported")
	}

}
