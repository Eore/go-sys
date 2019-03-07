package pendaftaran

import "time"

type Antrian struct {
	UID          string `json:"uid" sql:"varchar(100)"`
	IDJadwal     string
	NomorAntrian int16
	WaktuAntri   time.Time
	WaktuSelesai time.Time
	Status       string
}

type JadwalDokter struct {
	UID         string
	IDDokter    string
	IDPoli      string
	Hari        string
	Mulai       time.Time
	Selesai     time.Time
	BatasPasien int8
}

type ListPoli struct {
	UID      string
	NamaPoli string
}
