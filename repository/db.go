package repository

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	// "structs"
	"Test_task/models"
	"log"

	_ "github.com/lib/pq"
	"errors"
	"strconv"
	// "fmt"
)

var database *sql.DB

func ConnectDatabase(logger *slog.Logger) {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Failed connect database")
		panic(err)
	}
	logger.Debug("Connect database")
	database = db
	// defer db.Close()
}

func AddNewSong(song models.Song) error {
	row := findSong(song)
	if row.Song != "" {
		log.Printf("Song already exists")
		return errors.New("song already exists")
	}
	_, err := database.Exec("INSERT INTO songs (song, group_name, release_date, text, link) VALUES ($1, $2, $3, $4, $5)", song.Song, song.GroupName, song.ReleaseDate, song.Text, song.Link)
	if err != nil {
		log.Println(err)
		return err
	}
	// defer rows.Close()
	return nil
}

func findSong(song models.Song) models.Song {
	var foundSong models.Song
	err := database.QueryRow("SELECT * FROM songs WHERE song = $1 AND group_name = $2", song.Song, song.GroupName).Scan(&foundSong.Id, &foundSong.Song, &foundSong.GroupName, &foundSong.ReleaseDate, &foundSong.Text, &foundSong.Link)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Song{}
		}
		log.Println("Error while retrieving song")
		return models.Song{}
	}
	return foundSong
}

func GetAllSongs(filter models.Song, lastId string, limit int) ([]*models.Song, error) { //переделать если будет время
	query := "SELECT * FROM songs WHERE"
	var args []interface{}
	argCount := 1
	conditionsAdded := false //флаг для проверки существования условий

	if filter.Song != "" {
		if conditionsAdded {
			query += " AND"
		}
		query += " song = $" + strconv.Itoa(argCount)
		args = append(args, filter.Song)
		argCount++
		conditionsAdded = true
	}
	if filter.GroupName != "" {
		if conditionsAdded {
			query += " AND"
		}
		query += " group_name = $" + strconv.Itoa(argCount)
		args = append(args, filter.GroupName)
		argCount++
		conditionsAdded = true
	}
	if filter.ReleaseDate != "" {
		if conditionsAdded {
			query += " AND"
		}
		query += " release_date = $" + strconv.Itoa(argCount)
		args = append(args, filter.ReleaseDate)
		argCount++
		conditionsAdded = true
	}
	if filter.Text != "" {
		if conditionsAdded {
			query += " AND"
		}
		query += " text = $" + strconv.Itoa(argCount)
		args = append(args, filter.Text)
		argCount++
		conditionsAdded = true
	}
	if filter.Link != "" {
		if conditionsAdded {
			query += " AND"
		}
		query += " link = $" + strconv.Itoa(argCount)
		args = append(args, filter.Link)
		argCount++
		conditionsAdded = true
	}

	if lastId != "" {
		if conditionsAdded {
			query += " AND"
		}
		query += " id > $" + strconv.Itoa(argCount)
		args = append(args, lastId)
		argCount++
		conditionsAdded = true
	}

	if limit > 0 {
		query += " ORDER BY id ASC LIMIT $" + strconv.Itoa(argCount)
		args = append(args, limit)
	}

	if !conditionsAdded {
		query = "SELECT * FROM songs ORDER BY id ASC LIMIT $" + strconv.Itoa(argCount)
		args = append(args, limit)
	}

	rows, err := database.Query(query, args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	songs := make([]*models.Song, 0)
	for rows.Next() {
		song := new(models.Song)
		err := rows.Scan(&song.Id, &song.Song, &song.GroupName, &song.ReleaseDate, &song.Text, &song.Link)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		songs = append(songs, song)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return songs, nil
}

func DeleteSong(id string) error {
	result, err := database.Exec("DELETE FROM songs WHERE id = $1", id)
	if err != nil {
		log.Println("Error deleting song")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected")
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("song not found")
	}
	return nil
}

func UpdateSong(id string, song models.Song) error {
	query := `
		UPDATE songs 
		SET 
			song = COALESCE(NULLIF($1, ''), song),
			group_name = COALESCE(NULLIF($2, ''), group_name),
			release_date = COALESCE(NULLIF($3, ''), release_date),
			text = COALESCE(NULLIF($4, ''), text),
			link = COALESCE(NULLIF($5, ''), link)
		WHERE id = $6
	`
	_, err := database.Exec(query, song.Song, song.GroupName, song.ReleaseDate, song.Text, song.Link, id)
	if err != nil {
		log.Println("Error updating song")
		return err
	}
	return nil
}

func GetSongText(id string) (string, error) {
	query := `
		SELECT text 
		FROM songs 
		WHERE id = $1
	`
	var text string
	err := database.QueryRow(query, id).Scan(&text)
	if err != nil {
		log.Println("Error retrieving song text for song")
		return "", err
	}
	return text, nil
}

