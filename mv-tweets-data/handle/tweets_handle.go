package handle

import (
	"fmt"
	"github.com/cgao/mv-tweets-data/model"
	"regexp"
	"strconv"
	"strings"
)

func analyze_tweets(twi *model.TWEETS) {
	var tweets_index = insert_tweets(twi)

	if &twi.Entities != nil {
		entities_index := analyze_entities(&twi.Entities)
		insert_rel_tweets_entities(tweets_index, entities_index)
	}

	if &twi.Places != nil {
		places_index := analyze_places(&twi.Places)
		insert_rel_tweets_places(tweets_index, places_index)
	}

	if &twi.User != nil {
		user_index := analyze_user(&twi.User)
		insert_rel_tweets_user(tweets_index, user_index)
	}
}

func insert_tweets(twi *model.TWEETS) int64 {
	var withheld_in_countries string
	if twi.Withheld_in_countries != nil {
		withheld_in_countries = strings.Join(twi.Withheld_in_countries, ",")
	}

	var contributors_id string = ""
	var contributors_id_str string = ""
	var contributors_screen_name string = ""
	if &twi.Contributors != nil {
		for i := range twi.Contributors {
			contributors_id += strconv.FormatInt(twi.Contributors[i].Id, 10) + ","
			contributors_id_str += twi.Contributors[i].Id_str + ","
			contributors_screen_name += twi.Contributors[i].Screen_name + ","
		}
		contributors_id = strings.TrimRight(contributors_id, ",")
		contributors_id_str = strings.TrimRight(contributors_id_str, ",")
		contributors_screen_name = strings.TrimRight(contributors_screen_name, ",")
	}

	var contributor_id int64
	var contributor_id_str string
	var contributor_screen_name string
	if &twi.Contributor != nil {
		contributor_id = twi.Contributor.Id
		contributor_id_str = twi.Contributor.Id_str
		contributor_screen_name = twi.Contributor.Screen_name
	}

	var geo_coordinates string
	var geo_coordinates_slice []string
	var geo_m_type string
	if &twi.Geo != nil {
		geo_m_type = twi.Geo.M_type
		for i := range twi.Geo.Coordinates {
			num := twi.Geo.Coordinates[i]
			text := strconv.FormatFloat(num, 'f', -1, 64)
			geo_coordinates_slice = append(geo_coordinates_slice, text)
		}

		geo_coordinates = strings.Join(geo_coordinates_slice, ",")
	}

	var coordinates_followers bool
	if &twi.Scopes != nil {
		coordinates_followers = twi.Scopes.Followers
	}

	var current_user_retweet_id int64
	var current_user_retweet_id_str string
	if &twi.Current_user_retweet != nil {
		current_user_retweet_id = twi.Current_user_retweet.Id
		current_user_retweet_id_str = twi.Current_user_retweet.Id_str
	}

	stmt, err := tx.Prepare("insert into tweets(id, id_str, created_at, favorite_count, favorited, filter_level, in_reply_to_screen_name, in_reply_to_status_id, " +
		"in_reply_to_status_id_str, in_reply_to_user_id, in_reply_to_user_id_str, lang, possibly_sensitive, quoted_status_id, quoted_status_id_str, " +
		"retweet_count, retweeted, source, text, truncated, withheld_in_countries, withheld_scope, contributors_id, contributors_id_str, " +
		"contributors_screen_name, contributor_id, contributor_id_str, contributor_screen_name, geo_coordinates, geo_m_type, coordinates_followers, " +
		"coordinates_coordinates, coordinates_m_type, current_user_retweet_id, current_user_retweet_id_str) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, " +
		"?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		fmt.Println("analyze_tweets Insert table Error =", err)
	}

	// for all emoji, url = http://apps.timwhitlock.info/emoji/tables/unicode
	// I am too lazy to want to collect all the situation, so "lazy regexp".
	// r := regexp.MustCompile(`[\x{1F601}-\x{1F64F}]|[\x{2702}-\x{27B0}]|[\x{1F680}-\x{1F6C0}]|[\x{24C2}-\x{1F251}]|[\x{0080}-\x{00FF}]|[\x{1F600}-\x{1F636}]|[\x{1F681}-\x{1F6C5}]|[\x{1F30D}-\x{1F567}]`)
	// r := regexp.MustCompile(`[\x{1F000}-\x{1FFFF}]|[\x{0000}-\x{3FFF}]`)
	r := regexp.MustCompile(`[\x{1F000}-\x{1FFFF}]|[\x{2000}-\x{2FFF}]|[\x{3000}-\x{3FFF}]`)
	twi.Text = r.ReplaceAllString(twi.Text, "")

	_, err = stmt.Exec(twi.Id, twi.Id_str, twi.Created_at, twi.Favorite_count, twi.Favorited, twi.Filter_level, twi.In_reply_to_screen_name,
		twi.In_reply_to_status_id, twi.In_reply_to_status_id_str, twi.In_reply_to_user_id, twi.In_reply_to_user_id_str, twi.Lang, twi.Possibly_sensitive,
		twi.Quoted_status_id, twi.Quoted_status_id_str, twi.Retweet_count, twi.Retweeted, twi.Source, twi.Text, twi.Truncated,
		withheld_in_countries, twi.Withheld_scope, contributors_id, contributors_id_str, contributors_screen_name, contributor_id, contributor_id_str,
		contributor_screen_name, geo_coordinates, geo_m_type, coordinates_followers, geo_coordinates, geo_m_type,
		current_user_retweet_id, current_user_retweet_id_str)
	if err != nil {
		fmt.Println("analyze_tweets Exec err = ", err)
	}

	var lastid int64
	if err := tx.QueryRow("select last_insert_id() as lastid").Scan(&lastid); err != nil {
		fmt.Println("insert_tweets tx.QueryRow err =", err)
	}
	return lastid
}

func insert_rel_tweets_user(tweets_index int64, user_index int64) {
	stmt, err := tx.Prepare("insert into rel_tweets_user(tweets_index, user_index) values(?, ?)")
	if err != nil {
		fmt.Println("insert_rel_tweets_user tx.Prepare err =", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(tweets_index, user_index)
	if err != nil {
		fmt.Println("insert_rel_tweets_user stmt.Exec err =", err)
	}
}

func insert_rel_tweets_places(tweets_index int64, places_index int64) {
	stmt, err := tx.Prepare("insert into rel_tweets_places(tweets_index, places_index) values(?, ?)")
	if err != nil {
		fmt.Println("insert_rel_tweets_places tx.Prepare err =", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(tweets_index, places_index)
	if err != nil {
		fmt.Println("insert_rel_tweets_places stmt.Exec err =", err)
	}
}

func insert_rel_tweets_entities(tweets_index int64, entities_index int64) {
	stmt, err := tx.Prepare("insert into rel_tweets_entities(tweets_index, entities_index) values(?, ?)")
	if err != nil {
		fmt.Println("insert_rel_tweets_entities tx.Prepare err =", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(tweets_index, entities_index)
	if err != nil {
		fmt.Println("insert_rel_tweets_entities stmt.Exec err =", err)
	}
}
