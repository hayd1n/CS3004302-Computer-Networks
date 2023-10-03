package iperfer

func calcBandwidth(bytes uint64, seconds int) float32 {
	bits := bytes * 8
	mbps := float32(bits) / float32(1000000*seconds)
	return mbps
}
