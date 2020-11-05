package ro

//QueryRequest set of query paramters
type QueryRequest struct {
	CityID        string
	ProvinceID    string
	SubdistrictID string
	Origin        string
	Destination   string
	Weight        int
	Courier       string
}
