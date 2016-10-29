package handle

import (
	"fmt"
	"github.com/cgao/mv-tweets-data/model"
	"strconv"
	"strings"
)

func analyze_places(places *model.PLACES) int64 {
	var attributes_street_address string
	var attributes_twitter string
	if &places.Attributes != nil {
		attributes_street_address = places.Attributes.Street_address
		attributes_twitter = places.Attributes.Twitter
	}

	var bounding_box_coordinates_slice []string
	var bounding_box_coordinates string
	var bounding_box_m_type string
	if &places.Bounding_box != nil {
		for i := range places.Bounding_box.Coordinates {
			for j := range places.Bounding_box.Coordinates[i] {
				for k := range places.Bounding_box.Coordinates[i][j] {
					bounding_box_coordinates_slice = append(bounding_box_coordinates_slice, strconv.FormatFloat(places.Bounding_box.Coordinates[i][j][k], 'f', -1, 64))
				}
			}
		}
		bounding_box_coordinates = strings.Join(bounding_box_coordinates_slice, ",")
		bounding_box_m_type = places.Bounding_box.M_type
	}

	stmt, err := tx.Prepare("insert into places(country, country_code," +
		"full_name, id, name, place_type, url, attributes_street_address, attributes_twitter, " +
		"bounding_box_coordinates, bounding_box_m_type) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("analyze_places tx.Prepare error =", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(places.Country, places.Country_code,
		places.Full_name, places.Id, places.Name, places.Place_type, places.Url, attributes_street_address, attributes_twitter,
		bounding_box_coordinates, bounding_box_m_type)
	if err != nil {
		fmt.Println("analyze_places stmt.Exec error =", err)
	}

	var lastid int64
	if err := tx.QueryRow("select last_insert_id() as lastid").Scan(&lastid); err != nil {
		fmt.Println("analyze_places tx.QueryRow err =", err)
	}
	return lastid
}
