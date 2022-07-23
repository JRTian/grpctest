package db

import (
	// "log"
	"git.neds.sh/matty/entain/sporting/proto/sporting"
	"strings"
	"time"

	"syreclabs.com/go/faker"
)

func (r *sportsRepo) seed() error {
	statement, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS sports (id INTEGER PRIMARY KEY, name TEXT, advertised_start_time DATETIME)`)
	if err == nil {
		_, err = statement.Exec()
	}

	for i := 1; i <= 100; i++ {
		statement, err = r.db.Prepare(`INSERT OR IGNORE INTO sports(id, name, advertised_start_time) VALUES (?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				faker.Team().Name(),
				faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2)).Format(time.RFC3339),
			)
		}
	}

	return err
}

func (r *sportsRepo) applyFilter(query string, filter *sporting.ListSportsRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)
    
	if filter == nil {
		return query, args
	}
	
	if len(filter.Names) > 0 {
		clauses = append(clauses, "name IN ("+strings.Repeat("?,", len(filter.Names)-1)+"?)")

		for _, name := range filter.Names {
			args = append(args, name)
		} 
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}
	// log.Println("query", query, args)
	return query, args
}