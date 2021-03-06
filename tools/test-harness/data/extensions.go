package data

import "database/sql"

// Extension is a data structure that holds information
// about extensions found in the database.
type Extension struct {
	DefaultVersion   string
	InstalledVersion string
	Name             string
}

/// AllExtensions method returns all extensions available.
func (db *DB) AllExtensions() ([]Extension, error) {
	statement := "SELECT name, default_version, installed_version " +
		" FROM pg_available_extensions"

	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	extensions := []Extension{}
	for rows.Next() {
		var name, defaultVersion, installedVersion sql.NullString
		if err := rows.Scan(&name, &defaultVersion, &installedVersion); err != nil {
			return nil, err
		}
		x := Extension{
			Name:             name.String,
			DefaultVersion:   defaultVersion.String,
			InstalledVersion: installedVersion.String,
		}
		extensions = append(extensions, x)
	}
	return extensions, nil
}

// InstalledExtensions method returns all installed extensions in the
// database currently being queried.
func (db *DB) InstalledExtensions() ([]Extension, error) {
	statement := "SELECT name, default_version, installed_version " +
		" FROM pg_available_extensions WHERE installed_version IS NOT NULL"

	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	extensions := []Extension{}
	for rows.Next() {
		var name, defaultVersion, installedVersion sql.NullString
		if err := rows.Scan(&name, &defaultVersion, &installedVersion); err != nil {
			return nil, err
		}
		x := Extension{
			Name:             name.String,
			DefaultVersion:   defaultVersion.String,
			InstalledVersion: installedVersion.String,
		}
		extensions = append(extensions, x)
	}
	return extensions, nil
}
