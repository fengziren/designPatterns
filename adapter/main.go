package main

import "github.com/fengziren/designPatterns/adapter/player1"

func main() {
	audioPlayer := new(player1.AudioPlayer)

	audioPlayer.Play("mp3", "beyond the horizon.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "far far away.vlc")
	audioPlayer.Play("avi", "mind me.avi")
}
