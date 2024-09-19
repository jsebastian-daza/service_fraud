package models

// Stats represents statistics related to a specific country.
type Stats struct {
	Country  string
	Distance string
	Invokes  int
}

// StatsRequest is used to capture parameters for requesting statistics.
type StatsRequest struct {
	Country string
	Lat     float64
	Lon     float64
}
