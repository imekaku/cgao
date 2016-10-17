package model

type PLACES struct {
	Country      string
	Country_code string
	Full_name    string
	Id           string
	Name         string
	Place_type   string
	Url          string
	Attributes   ATTRIBUTES
	Bounding_box BOUNDING_BOX
}

type ATTRIBUTES struct {
	Street_address string
	Twitter        string
	/* Id          string */
	// 不清楚这个是什么意思
}

type BOUNDING_BOX struct {
	Coordinates [][][]float64
	M_type      string
}
