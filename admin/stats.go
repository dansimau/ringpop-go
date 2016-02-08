package admin

type Stats struct {
	Membership struct {
		Checksum float64
	}
	Ring struct {
		Checksum float64
	}
	Uptime  float64
	Version string
}
