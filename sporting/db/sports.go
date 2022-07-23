package db

import (
	// "strconv"
	// "log"
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	// "strings"
	"sync"
	"time"
	"git.neds.sh/matty/entain/sporting/proto/sporting"
)

// SportsRepo provides repository access to races.
type SportsRepo interface {
	// Init will initialise our races repository.
	Init() error

	// List will return a list of races.
	List(filter *sporting.ListSportsRequestFilter) ([]*sporting.Sport, error)
 
}

type sportsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewSportsRepocreates a new races repository.
func NewSportsRepo(db *sql.DB) SportsRepo {
	return &sportsRepo{db: db}
}

// Init prepares the race repository dummy data.
func (r *sportsRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy races.
		err = r.seed()
	})

	return err
}

func (r *sportsRepo) List(filter *sporting.ListSportsRequestFilter) ([]*sporting.Sport, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getSportsQueries()[sportsList]

	query, args = r.applyFilter(query, filter)
    // log.Println(query, args)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanSports(rows)
	
}

func (m *sportsRepo) scanSports(
	rows *sql.Rows,
) ([]*sporting.Sport, error) {
	var sports []*sporting.Sport

	for rows.Next() {
		var sport sporting.Sport
		var advertisedStart time.Time

		if err := rows.Scan(&sport.Id, &sport.Name, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		sport.AdvertisedStartTime = ts

		sports = append(sports, &sport)
	}

	return sports, nil
}

