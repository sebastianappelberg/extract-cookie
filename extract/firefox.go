package extract

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func firefoxCookies() ([]Cookie, error) {
	paths, err := Firefox.GetPaths()
	if err != nil {
		return nil, fmt.Errorf("failed to find cookie db path: %w", err)
	}
	return readFirefoxCookies(paths.CookieDbPath)
}

func readFirefoxCookies(dbPath string) ([]Cookie, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	rows, err := db.Query(`
        SELECT host, name, value, path, expiry, isSecure, isHttpOnly
        FROM moz_cookies
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to run query: %w", err)
	}
	defer rows.Close()

	var cookies []Cookie

	for rows.Next() {
		var c Cookie
		var isSecure, isHttpOnly int
		if err := rows.Scan(&c.Domain, &c.Name, &c.Value, &c.Path, &c.Expires, &isSecure, &isHttpOnly); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		c.Secure = isSecure != 0
		c.HttpOnly = isHttpOnly != 0
		cookies = append(cookies, c)
	}

	return cookies, nil
}
