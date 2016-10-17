package handle

import (
	"fmt"
	"github.com/cgao/mv-tweets-data/model"
	"strconv"
	"strings"
)

func analyze_entities(entities *model.ENTITIES) int64 {
	var hashtags_index_slice []string
	var hashtags_index string
	if &entities.Hashtags != nil {
		for i := range entities.Hashtags {
			hashtags_index_slice = append(hashtags_index_slice, strconv.FormatInt(insert_hashtags(&entities.Hashtags[i]), 10))
		}
		hashtags_index = strings.Join(hashtags_index_slice, ",")
	}

	var entities_media_index_slice []string
	var entities_media_index string
	if len(entities.Media) != 0 {
		for i := range entities.Media {
			entities_media_index_slice = append(entities_media_index_slice, strconv.FormatInt(insert_media(&entities.Media[i]), 10))
		}
		entities_media_index = strings.Join(entities_media_index_slice, ",")
	}

	var urls_index_slice []string
	var urls_index string
	if &entities.Urls != nil {
		for i := range entities.Urls {
			urls_index_slice = append(urls_index_slice, strconv.FormatInt(insert_urls(&entities.Urls[i]), 10))
		}
		urls_index = strings.Join(urls_index_slice, ",")
	}

	var user_mentions_index_slice []string
	var user_mentions_index string
	if &entities.User_mentions != nil {
		for i := range entities.User_mentions {
			user_mentions_index_slice = append(user_mentions_index_slice, strconv.FormatInt(insert_user_mentions(&entities.User_mentions[i]), 10))
		}
		user_mentions_index = strings.Join(user_mentions_index_slice, ",")
	}

	stmt, _ := tx.Prepare("insert into entities(hashtags_index, media_index, urls_index, user_mentions_index) values(?, ?, ?, ?)")
	defer stmt.Close()
	stmt.Exec(hashtags_index, entities_media_index, urls_index, user_mentions_index)

	var lastid int64
	if err := tx.QueryRow("select last_insert_id() as lastid").Scan(&lastid); err != nil {
		fmt.Println("insert_hashtags err =", err)
	}
	return lastid
}

func insert_hashtags(hashtags *model.HASHTAGS) int64 {
	var indices string
	var indices_slice []string
	for i := range hashtags.Indices {
		num := hashtags.Indices[i]
		text := strconv.Itoa(num)
		indices_slice = append(indices_slice, text)
	}
	indices = strings.Join(indices_slice, ",")
	stmt, _ := tx.Prepare("insert into hashtags(indices, hashtags_text) values(?, ?)")
	defer stmt.Close()
	stmt.Exec(indices, hashtags.Text)

	var lastid int64
	if err := tx.QueryRow("select last_insert_id() as lastid").Scan(&lastid); err != nil {
		fmt.Println("insert_hashtags err =", err)
	}
	return lastid
}

func insert_media(media *model.MEDIA) int64 {
	stmt, _ := tx.Prepare("insert into media(display_url, expanded_url," +
		"id, id_str, indices, media_url, media_url_https, source_status_id, source_status_id_str, " +
		"m_type, url, sizes_index) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()

	var sizes_index int64
	if &media.Sizes != nil {
		sizes_index = insert_media_sizes(&media.Sizes)
	}

	stmt.Exec(media.Display_url, media.Expanded_url, media.Id, media.Id_str, media.Indices, media.Media_url,
		media.Media_url_https, media.Source_status_id, media.Source_status_id_str, media.M_type, media.Url, sizes_index)

	var lastid int64
	if err := tx.QueryRow("select last_insert_id() as lastid").Scan(&lastid); err != nil {
		fmt.Println("insert_hashtags err =", err)
	}
	return lastid
}

func insert_media_sizes(sizes *model.SIZES) int64 {
	var thumb_h int
	var thumb_w int
	var thumb_resize string = ""
	if &sizes.Thumb != nil {
		thumb_h = sizes.Thumb.H
		thumb_w = sizes.Thumb.W
		thumb_resize = sizes.Thumb.Resize
	}

	var large_h int
	var large_w int
	var large_resize string = ""
	if &sizes.Large != nil {
		large_h = sizes.Large.H
		large_w = sizes.Large.W
		large_resize = sizes.Large.Resize
	}

	var medium_h int
	var medium_w int
	var medium_resize string = ""
	if &sizes.Medium != nil {
		medium_h = sizes.Medium.H
		medium_w = sizes.Medium.W
		medium_resize = sizes.Medium.Resize
	}

	var small_h int
	var small_w int
	var small_resize string = ""
	if &sizes.Small != nil {
		small_h = sizes.Small.H
		small_w = sizes.Small.W
		small_resize = sizes.Small.Resize
	}

	stmt, _ := tx.Prepare("insert into sizes(thumb_h, thumb_w, thumb_resize, large_h, large_w, large_resize, " +
		"medium_h, medium_w, medium_resize, small_h, small_w, small_resize) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()
	stmt.Exec(thumb_h, thumb_w, thumb_resize, large_h, large_w, large_resize, medium_h, medium_w, medium_resize, small_h, small_w, small_resize)

	var lastid int64
	if err := tx.QueryRow("select last_insert_id() as lastid").Scan(&lastid); err != nil {
		fmt.Println("insert_hashtags err =", err)
	}
	return lastid
}

func insert_urls(urls *model.URLS) int64 {
	var indices string
	var indices_slice []string
	for i := range urls.Indices {
		num := urls.Indices[i]
		text := strconv.Itoa(num)
		indices_slice = append(indices_slice, text)
	}
	indices = strings.Join(indices_slice, ",")
	stmt, _ := tx.Prepare("insert into urls(display_url, expanded_url, indices, url) values(?, ?, ?, ?)")
	defer stmt.Close()
	stmt.Exec(urls.Display_url, urls.Expanded_url, indices, urls.Url)

	var lastid int64
	if err := tx.QueryRow("select last_insert_id() as lastid").Scan(&lastid); err != nil {
		fmt.Println("insert_urls err =", err)
	}
	return lastid
}

func insert_user_mentions(user_mentions *model.USER_MENTIONS) int64 {
	var indices string
	var indices_slice []string
	for i := range user_mentions.Indices {
		num := user_mentions.Indices[i]
		text := strconv.Itoa(num)
		indices_slice = append(indices_slice, text)
	}
	indices = strings.Join(indices_slice, ",")
	stmt, _ := tx.Prepare("insert into user_mentions(id, id_str, indices, name, screen_name) values(?, ?, ?, ?, ?)")
	defer stmt.Close()
	stmt.Exec(user_mentions.Id, user_mentions.Id_str, indices, user_mentions.Name, user_mentions.Screen_name)

	var lastid int64
	if err := tx.QueryRow("select last_insert_id() as lastid").Scan(&lastid); err != nil {
		fmt.Println("insert_uer_mentions err =", err)
	}
	return lastid
}
