package player1

import "github.com/fengziren/designPatterns/adapter/player2"

type mediaAdapter struct {
	advancedMediaPlayer player2.AdvancedMediaPlayer
}

func (m *mediaAdapter) GetMediaAdapter(audioType string) {
	if audioType == "vlc" {
		m.advancedMediaPlayer = new(player2.VlcPlayer)
	} else if audioType == "mp4" {
		m.advancedMediaPlayer = new(player2.Mp4Player)
	}
}

func (m *mediaAdapter) Play(audioType string, fileName string) {
	if audioType == "vlc" {
		m.advancedMediaPlayer.PlayVlc(fileName)
	} else if audioType == "mp4" {
		m.advancedMediaPlayer.PlayMp4(fileName)
	}
}
