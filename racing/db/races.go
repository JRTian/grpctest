package db

import (
	// "log"
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"sync"
	"time"

	"git.neds.sh/matty/entain/racing/proto/racing"
)

// RacesRepo provides repository access to races.
type RacesRepo interface {
	// Init will initialise our races repository.
	Init() error

	// List will return a list of races.
	List(filter *racing.ListRacesRequestFilter, orderBy *racing.ListRacesRequestOrderBy) ([]*racing.Race, error)
}

type racesRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewRacesRepo creates a new races repository.
func NewRacesRepo(db *sql.DB) RacesRepo {
	return &racesRepo{db: db}
}

// Init prepares the race repository dummy data.
func (r *racesRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy races.
		err = r.seed()
	})

	return err
}

func (r *racesRepo) List(filter *racing.ListRacesRequestFilter, orderBy *racing.ListRacesRequestOrderBy) ([]*racing.Race, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getRaceQueries()[racesList]

	query, args = r.applyFilter(query, filter, orderBy)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanRaces(rows)
}

func (r *racesRepo) applyFilter(query string, filter *racing.ListRacesRequestFilter, orderBy *racing.ListRacesRequestOrderBy) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)
    
	if filter == nil {
		return query, args
	}
	
	if len(filter.MeetingIds) > 0 {
		clauses = append(clauses, "meeting_id IN ("+strings.Repeat("?,", len(filter.MeetingIds)-1)+"?)")

		for _, meetingID := range filter.MeetingIds {
			args = append(args, meetingID)
		} 
	}

	// add by Gary Tian
	if filter.GetVisibleOnly(){
		clauses = append(clauses, "visible = ?")
		args = append(args, true /* visible = true*/)
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}
	
	// added order by 
	colunmNameForOrder := strings.ToLower(orderBy.GetOrderBy())
	order := strings.ToLower(orderBy.GetOrder())
	if isValidColunm(colunmNameForOrder) && isValidOrder(order){
		query += " order by " + colunmNameForOrder+ " " + order
	}
	
	// log.Println("query", query)
	return query, args
}

func isValidOrder(order string)bool{
    if order == "asc"{
		return true
	}

	if order == "desc"{
		return true
	}

	return false
}

func isValidColunm(colunmNameForOrder string)bool{
	if colunmNameForOrder == "id"{
		return true
	}

	if colunmNameForOrder == "meeting_id"{
		return true
	}

	if colunmNameForOrder == "name"{
		return true
	}

	if colunmNameForOrder == "number"{
		return true
	}

	if colunmNameForOrder == "visible"{
		return true
	}

	if colunmNameForOrder == "advertised_start_time"{
		return true
	}

	return false
}

func (m *racesRepo) scanRaces(
	rows *sql.Rows,
) ([]*racing.Race, error) {
	var races []*racing.Race

	for rows.Next() {
		var race racing.Race
		var advertisedStart time.Time

		if err := rows.Scan(&race.Id, &race.MeetingId, &race.Name, &race.Number, &race.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		race.AdvertisedStartTime = ts

		races = append(races, &race)
	}

	return races, nil
}

