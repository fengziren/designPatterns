package player2

type Mp4Player struct{}

func (m *Mp4Player) PlayMp4(fileName string) {
	println("Playing mp4 file. Name: " + fileName)
}

func (m *Mp4Player) PlayVlc(fileName string) {}