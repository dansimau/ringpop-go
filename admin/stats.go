package admin

type Stats struct {
	data map[string]interface{}
}

// RingChecksum gets the node ring checksum using the ringpop admin endpoint.
func (s *Stats) RingChecksum() float64 {
	ringInfo := s.data["ring"].(map[string]interface{})
	checksum := ringInfo["checksum"].(float64)
	return checksum
}
