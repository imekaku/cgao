package model

type USER struct {
	Id                                 int64    `json:"id"`
	Id_str                             string   `json:"id_str"`
	Contributors_enabled               bool     `json:"contributors_enabled"`
	Created_at                         string   `json:"created_at"`
	Default_profile                    bool     `json:"default_profile"`
	Default_profile_image              bool     `json:"default_profile_image"`
	Description                        string   `json:"description"`
	Favourites_count                   int      `json:"favourites_count"`
	Follow_request_sent                bool     `json:"follow_request_sent"`
	Following                          bool     `json:"following"`
	Followers_count                    int      `json:"followers_count"`
	Friends_count                      int      `json:"friends_count"`
	Geo_enabled                        bool     `json:"geo_enabled"`
	Is_translator                      bool     `json:"is_translator"`
	Lang                               string   `json:"lang"`
	Listed_count                       int      `json:"listed_count"`
	Location                           string   `json:"location"`
	Name                               string   `json:"name"`
	Notifications                      bool     `json:"notifications"`
	Profile_background_color           string   `json:"profile_background_color"`
	Profile_background_image_url       string   `json:"profile_background_image_url"`
	Profile_background_image_url_https string   `json:"profile_background_image_url_https"`
	Profile_background_tile            bool     `json:"profile_background_tile"`
	Profile_banner_url                 string   `json:"profile_banner_url"`
	Profile_image_url                  string   `json:"profile_image_url"`
	Profile_image_url_https            string   `json:"profile_image_url_https"`
	Profile_link_color                 string   `json:"profile_link_color"`
	Profile_sidebar_border_color       string   `json:"profile_sidebar_border_color"`
	Profile_sidebar_fill_color         string   `json:"profile_sidebar_fill_color"`
	Profile_text_color                 string   `json:"profile_text_color"`
	Profile_use_background_image       bool     `json:"profile_use_background_image"`
	Protected                          bool     `json:"protected"`
	Screen_name                        string   `json:"screen_name"`
	Show_all_inline_media              bool     `json:"show_all_inline_media"`
	Statuses_count                     int      `json:"statuses_count"`
	Time_zone                          string   `json:"time_zone"`
	Url                                string   `json:"url"`
	Utc_offset                         int      `json:"utc_offset"`
	Verified                           bool     `json:"verified"`
	M_is_translator                    bool     `json:"is_translator"`
	Withheld_in_countries              []string `json:"withheld_in_countries"`
	Entities                           ENTITIES `json:"entities"`
	/* status Tweets */
}
