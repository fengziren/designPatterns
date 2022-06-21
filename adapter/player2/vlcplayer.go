package player2

type VlcPlayer struct{}

func (v *VlcPlayer) PlayVlc(fileName string) {
	println("Playing vlc file. Name: " + fileName)
}

func (v *VlcPlayer) PlayMp4(fileName string) {}
