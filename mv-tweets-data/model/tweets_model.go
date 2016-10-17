package model

type TWEETS struct {
	Id                        int64                `json:"id"`
	Id_str                    string               `json:"id_str"`
	Created_at                string               `json:"created_at"`
	Favorite_count            int                  `json:"favorite_count"`
	Favorited                 bool                 `json:"favorited"`
	Filter_level              string               `json:"filter_level"`
	In_reply_to_screen_name   string               `json:"in_reply_to_screen_name"`
	In_reply_to_status_id     int64                `json:"in_reply_to_status_id"`
	In_reply_to_status_id_str string               `json:"in_reply_to_status_id_str"`
	In_reply_to_user_id       int64                `json:"in_reply_to_user_id"`
	In_reply_to_user_id_str   string               `json:"in_reply_to_user_id_str"`
	Lang                      string               `json:"lang"`
	Possibly_sensitive        bool                 `json:"possibly_sensitive"`
	Quoted_status_id          int64                `json:"quoted_status_id"`
	Quoted_status_id_str      string               `json:"quoted_status_id_str"`
	Retweet_count             int                  `json:"retweet_count"`
	Retweeted                 bool                 `json:"retweeted"`
	Source                    string               `json:"source"`
	Text                      string               `json:"text"`
	Truncated                 bool                 `json:"truncated"`
	Withheld_in_countries     []string             `json:"withheld_in_countries"`
	Withheld_scope            string               `json:"withheld_scope"`
	Contributors              []CONTRIBUTOR        `json:"contributors"`
    Contributor               CONTRIBUTOR          `json:"contributor"`
	User                      USER                 `json:"user"`
	Geo                       COORDINATES          `json:"geo"`
	Places                    PLACES               `json:"place"`
	Scopes                    SCOPES               `json:"scopes"`
	Coordinates               COORDINATES          `json:"coordinates"`
	Current_user_retweet      CURRENT_USER_RETWEET `json:"current_user_retweet"`
	Entities                  ENTITIES             `json:"entities"`
	/* quoted_status TWEETS */
	/* retweeted_status TWEETS */
}

type CONTRIBUTOR struct {
	Id          int64
	Id_str      string
	Screen_name string
}

type COORDINATES struct {
	Coordinates []float64
	M_type      string
}

type CURRENT_USER_RETWEET struct {
	Id     int64
	Id_str string
}

type SCOPES struct {
	Followers bool
}