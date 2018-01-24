package repository

import (
	"fmt"
	. "../model"
)

func selectContent( )(Contents, error) {

	statement := fmt.Sprintf("SELECT id, name, age FROM users LIMIT %d OFFSET %d", 1, 1)
	rows, err := DBCon.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	contents := Contents{}

	for rows.Next() {
		var c Content
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		contents = append(contents, c)
	}

	return contents, nil

}