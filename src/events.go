package main

import (
	"database/sql"
)

type Event struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func allEvents(db *sql.DB) ([]Event, error) {
	rows, err := db.Query("SELECT * FROM events ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		err = rows.Scan(&event.Id, &event.Name)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
