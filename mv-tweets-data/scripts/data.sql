drop database if exists json_data_tweets;
create database json_data_tweets;
use json_data_tweets;

drop table if exists tweets;
CREATE TABLE `tweets` (
  `tweets_index` bigint NOT NULL AUTO_INCREMENT,
  `id` bigint DEFAULT NULL,
  `id_str` varchar(20) DEFAULT NULL,
  `created_at` varchar(50) DEFAULT NULL,
  `favorite_count` bigint DEFAULT NULL,
  `favorited` tinyint(1) DEFAULT NULL,
  `filter_level` varchar(20) DEFAULT NULL,
  `in_reply_to_screen_name` varchar(255) DEFAULT NULL,
  `in_reply_to_status_id` bigint DEFAULT NULL,
  `in_reply_to_status_id_str` varchar(20) DEFAULT NULL,
  `in_reply_to_user_id` bigint DEFAULT NULL,
  `in_reply_to_user_id_str` varchar(20) DEFAULT NULL,
  `lang` varchar(20) DEFAULT NULL,
  `possibly_sensitive` tinyint(1) DEFAULT NULL,
  `quoted_status_id` bigint DEFAULT NULL,
  `quoted_status_id_str` varchar(20) DEFAULT NULL,
  `retweet_count` bigint DEFAULT NULL,
  `retweeted` tinyint(1) DEFAULT NULL,
  `source` varchar(255) DEFAULT NULL,
  `text` varchar(255) DEFAULT NULL,
  `truncated` tinyint(1) DEFAULT NULL,
  `withheld_in_countries` varchar(255) DEFAULT NULL,
  `withheld_scope` varchar(255) DEFAULT NULL,
  `contributors_id` varchar(255) DEFAULT NULL,
  `contributors_id_str` varchar(255) DEFAULT NULL,
  `contributors_screen_name` varchar(2048) DEFAULT NULL,
  `contributor_id` bigint DEFAULT NULL,
  `contributor_id_str` varchar(20) DEFAULT NULL,
  `contributor_screen_name` varchar(255) DEFAULT NULL,
  `geo_coordinates` varchar(50) DEFAULT NULL,
  `geo_m_type` varchar(20) DEFAULT NULL,
  `coordinates_followers` tinyint(1) DEFAULT NULL,
  `coordinates_coordinates` varchar(50) DEFAULT NULL,
  `coordinates_m_type` varchar(20) DEFAULT NULL,
  `current_user_retweet_id` bigint DEFAULT NULL,
  `current_user_retweet_id_str` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`tweets_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

drop table if exists user;
CREATE TABLE `user` (
  `user_index` bigint NOT NULL AUTO_INCREMENT,
  `id` bigint DEFAULT NULL,
  `id_str` varchar(20) DEFAULT NULL,
  `contributors_enabled` tinyint(1) DEFAULT NULL,
  `created_at` varchar(50) DEFAULT NULL,
  `default_profile` tinyint(1) DEFAULT NULL,
  `default_profile_image` tinyint(1) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `favourites_count` bigint DEFAULT NULL,
  `follow_request_sent` tinyint(1) DEFAULT NULL,
  `following` tinyint(1) DEFAULT NULL,
  `followers_count` bigint DEFAULT NULL,
  `friends_count` bigint DEFAULT NULL,
  `geo_enabled` tinyint(1) DEFAULT NULL,
  `is_translator` tinyint(1) DEFAULT NULL,
  `lang` varchar(20) DEFAULT NULL,
  `listed_count` bigint DEFAULT NULL,
  `location` varchar(50) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `notifications` tinyint(1) DEFAULT NULL,
  `profile_background_color` varchar(20) DEFAULT NULL,
  `profile_background_image_url` varchar(1024) DEFAULT NULL,
  `profile_background_image_url_https` varchar(1024) DEFAULT NULL,
  `profile_background_tile` tinyint(1) DEFAULT NULL,
  `profile_banner_url` varchar(1024) DEFAULT NULL,
  `profile_image_url` varchar(1024) DEFAULT NULL,
  `profile_image_url_https` varchar(1024) DEFAULT NULL,
  `profile_link_color` varchar(20) DEFAULT NULL,
  `profile_sidebar_border_color` varchar(20) DEFAULT NULL,
  `profile_sidebar_fill_color` varchar(20) DEFAULT NULL,
  `profile_text_color` varchar(20) DEFAULT NULL,
  `profile_use_background_image` tinyint(1) DEFAULT NULL,
  `protected` tinyint(1) DEFAULT NULL,
  `screen_name` varchar(255) DEFAULT NULL,
  `show_all_inline_media` tinyint(1) DEFAULT NULL,
  `statuses_count` bigint DEFAULT NULL,
  `time_zone` varchar(50) DEFAULT NULL,
  `url` varchar(1024) DEFAULT NULL,
  `utc_offset` bigint DEFAULT NULL,
  `verified` tinyint(1) DEFAULT NULL,
  `m_is_translator` tinyint(1) DEFAULT NULL,
  `withheld_in_countries` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`user_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

drop table if exists places;
CREATE TABLE `places` (
  `places_index` bigint NOT NULL AUTO_INCREMENT,
  `country` varchar(50) DEFAULT NULL,
  `country_code` varchar(20) DEFAULT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `id` varchar(20) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `place_type` varchar(20) DEFAULT NULL,
  `url` varchar(1024) DEFAULT NULL,
  `attributes_street_address` varchar(255) DEFAULT NULL,
  `attributes_twitter` varchar(50) DEFAULT NULL,
  `bounding_box_coordinates` varchar(255) DEFAULT NULL,
  `bounding_box_m_type` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`places_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

drop table if exists entities;
CREATE TABLE `entities` (
  `entities_index` bigint NOT NULL AUTO_INCREMENT,
  `hashtags_index` varchar(255) DEFAULT NULL,
  `media_index` varchar(255) DEFAULT NULL,
  `urls_index` varchar(255) DEFAULT NULL,
  `user_mentions_index` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`entities_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

drop table if exists hashtags;
CREATE TABLE `hashtags` (
  `hashtags_index` bigint NOT NULL AUTO_INCREMENT,
  `indices` varchar(255) DEFAULT NULL,
  `hashtags_text` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`hashtags_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

drop table if exists media;
CREATE TABLE `media` (
  `media_index` bigint NOT NULL AUTO_INCREMENT,
  `display_url` varchar(1024) DEFAULT NULL,
  `expanded_url` varchar(1024) DEFAULT NULL,
  `id` bigint DEFAULT NULL,
  `id_str` varchar(20) DEFAULT NULL,
  `indices` varchar(255) DEFAULT NULL,
  `media_url` varchar(255) DEFAULT NULL,
  `media_url_https` varchar(255) DEFAULT NULL,
  `source_status_id` bigint DEFAULT NULL,
  `source_status_id_str` varchar(20) DEFAULT NULL,
  `m_type` varchar(20) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `sizes_index` bigint DEFAULT NULL,
  PRIMARY KEY (`media_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

drop table if exists urls;
CREATE TABLE `urls` (
  `urls_index` bigint NOT NULL AUTO_INCREMENT,
  `display_url` varchar(255) DEFAULT NULL,
  `expanded_url` varchar(255) DEFAULT NULL,
  `indices` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`urls_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

drop table if exists user_mentions;
CREATE TABLE `user_mentions` (
  `user_mentions_index` bigint NOT NULL AUTO_INCREMENT,
  `id` bigint DEFAULT NULL,
  `id_str` varchar(20) DEFAULT NULL,
  `indices` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `screen_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`user_mentions_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

drop table if exists sizes;
CREATE TABLE `sizes` (
  `sizes_index` int NOT NULL AUTO_INCREMENT,
  `thumb_h` int DEFAULT NULL,
  `thumb_w` int DEFAULT NULL,
  `thumb_resize` varchar(50) DEFAULT NULL,
  `large_h` int DEFAULT NULL,
  `large_w` int DEFAULT NULL,
  `large_resize` varchar(50) DEFAULT NULL,
  `medium_h` int DEFAULT NULL,
  `medium_w` int DEFAULT NULL,
  `medium_resize` varchar(50) DEFAULT NULL,
  `small_h` int DEFAULT NULL,
  `small_w` int DEFAULT NULL,
  `small_resize` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`sizes_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

drop table if exists rel_tweets_user;
CREATE TABLE `rel_tweets_user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `tweets_index` bigint DEFAULT NULL,
  `user_index` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ref_idx_tweets_index` (`tweets_index`),
  KEY `ref_idx_user_index` (`user_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

drop table if exists rel_tweets_places;
CREATE TABLE `rel_tweets_places` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `tweets_index` bigint DEFAULT NULL,
  `places_index` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ref_idx_tweets_index` (`tweets_index`),
  KEY `ref_idx_places_index` (`places_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

drop table if exists rel_tweets_entities;
CREATE TABLE `rel_tweets_entities` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `tweets_index` bigint DEFAULT NULL,
  `entities_index` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ref_idx_tweets_index` (`tweets_index`),
  KEY `ref_idx_entities_index` (`entities_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

drop table if exists rel_user_entities;
CREATE TABLE `rel_user_entities` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_index` bigint DEFAULT NULL,
  `entities_index` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ref_idx_user_index` (`user_index`),
  KEY `ref_idx_entities_index` (`entities_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;