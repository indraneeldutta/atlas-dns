package models

type DroneRequest struct {
	ID  string `json:"id"`
	X   string `json:"x"`
	Y   string `json:"y"`
	Z   string `json:"z"`
	Vel string `json:"vel"`
}

type DroneDetails struct {
	ID       string `bson:"_id" json:"id"`
	SectorID int64  `bson:"sector_id" json:"sector_id"`
}

type DroneResponse struct {
	Loc float64 `json:"loc"`
}
