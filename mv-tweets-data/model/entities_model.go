package model

type ENTITIES struct {
	Hashtags      []HASHTAGS
	Media         []MEDIA
	Urls          []URLS
	User_mentions []USER_MENTIONS
}

type HASHTAGS struct {
	Indices []int
	Text    string
}

type MEDIA struct {
	Display_url          string
	Expanded_url         string
	Id                   int64
	Id_str               string
	Indices              []int
	Media_url            string
	Media_url_https      string
	Source_status_id     int64
	Source_status_id_str string
	M_type               string
	Url                  string
	Sizes                SIZES
}

type URLS struct {
	Display_url  string
	Expanded_url string
	Indices      []int
	Url          string
}

type USER_MENTIONS struct {
	Id          int64
	Id_str      string
	Indices     []int
	Name        string
	Screen_name string
}

type SIZES struct {
	Thumb  SIZE
	Large  SIZE
	Medium SIZE
	Small  SIZE
}

type SIZE struct {
	H      int
	Resize string
	W      int
}
