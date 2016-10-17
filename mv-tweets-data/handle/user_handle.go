package handle

import (
	"fmt"
	"github.com/cgao/mv-tweets-data/model"
	"strings"
)

func analyze_user(user *model.USER) int64 {
	var entities_index int64 = -1
	if &user.Entities != nil {
		entities_index = analyze_entities(&user.Entities)
	}

	var withheld_in_countries string
	withheld_in_countries = strings.Join(user.Withheld_in_countries, ",")

	stmt, err := tx.Prepare("insert into user(id, id_str, contributors_enabled, created_at, default_profile, default_profile_image, " +
		"description, favourites_count, follow_request_sent, following, followers_count, friends_count, " +
		"geo_enabled, is_translator, lang, listed_count, location, name, notifications, profile_background_color, " +
		"profile_background_image_url, profile_background_image_url_https, profile_background_tile, " +
		"profile_banner_url, profile_image_url, profile_image_url_https, profile_link_color, profile_sidebar_border_color, " +
		"profile_sidebar_fill_color, profile_text_color, profile_use_background_image, protected, screen_name, " +
		"show_all_inline_media, statuses_count, time_zone, url, utc_offset, verified, m_is_translator, withheld_in_countries) " +
		"values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		fmt.Println("analyze_user tx.Prepare insert err =", err)
	}
	_, err = stmt.Exec(user.Id, user.Id_str, user.Contributors_enabled, user.Created_at, user.Default_profile, user.Default_profile_image, user.
		Description, user.Favourites_count, user.Follow_request_sent, user.Following, user.Followers_count, user.Friends_count, user.
		Geo_enabled, user.Is_translator, user.Lang, user.Listed_count, user.Location, user.Name, user.Notifications, user.Profile_background_color, user.
		Profile_background_image_url, user.Profile_background_image_url_https, user.Profile_background_tile, user.
		Profile_banner_url, user.Profile_image_url, user.Profile_image_url_https, user.Profile_link_color, user.Profile_sidebar_border_color, user.
		Profile_sidebar_fill_color, user.Profile_text_color, user.Profile_use_background_image, user.Protected, user.Screen_name, user.
		Show_all_inline_media, user.Statuses_count, user.Time_zone, user.Url, user.Utc_offset, user.Verified, user.M_is_translator, withheld_in_countries)
	if err != nil {
		fmt.Println("analyze_user stmt.Exec err =", err)
	}

	var lastid int64
	if err := tx.QueryRow("select last_insert_id() as lastid").Scan(&lastid); err != nil {
		fmt.Println("analyze_user last_insert_id err =", err)
	}

	if entities_index > 0 {
		insert_rel_user_entities(lastid, entities_index)
	}

	return lastid
}

func insert_rel_user_entities(user_index int64, entities_index int64) {
	stmt, err := tx.Prepare("insert into rel_user_entities(user_index, entities_index) values(?, ?)")
	if err != nil {
		fmt.Println("insert_rel_user_entities tx.Prepare err =", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user_index, entities_index)
	if err != nil {
		fmt.Println("insert_rel_user_entities stmt.Exec err =", err)
	}
}
