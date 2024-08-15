package main

import (
	"database/sql"
)

type Country struct {
	Id   int
	Name string
}

func allCountries(db *sql.DB) ([]Country, error) {
	rows, err := db.Query("SELECT * FROM countries")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var countries []Country

	for rows.Next() {
		var cty Country

		err := rows.Scan(&cty.Id, &cty.Name)
		if err != nil {
			return nil, err
		}

		countries = append(countries, cty)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return countries, nil
}
