package utils

import "database/sql"

func GetString(nullString sql.NullString) string {
	if nullString.Valid {
		return nullString.String
	}
	return ""
}
