package songdb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// SongDB look up and store songs in the database.
type SongDB struct {
	dbFilename string
}

type songBlock struct {
	UID    string
	URI    string
	Artist string    `json:",omitempty"`
	Album  string    `json:",omitempty"`
	Song   string    `json:",omitempty"`
	Added  time.Time `json:",omitempty"`
}

// NewSongDB - create new SongDB instance with provided filename.
func NewSongDB(filename string) *SongDB {
	return &SongDB{dbFilename: filename}
}

// Lookup find a song by its uid.
func (db *SongDB) Lookup(uid string) (string, error) {
	contents, err := ioutil.ReadFile(db.dbFilename)

	if err != nil {
		return "", err
	}

	var songs map[string]songBlock

	if err := json.Unmarshal(contents, &songs); err != nil {
		return "", err
	}

	if song, ok := songs[uid]; ok {
		return song.URI, nil
	}

	return "", errors.New("LookupFailed")
}

// Save a uri to a uid.
func (db *SongDB) Save(uid string, uri string) error {
	log.Print("doing store for: " + uid + " :: " + uri)
	contents, err := ioutil.ReadFile(db.dbFilename)

	if err != nil {
		return err
	}

	var songs map[string]songBlock

	if err := json.Unmarshal(contents, &songs); err != nil {
		return err
	}
	os.Rename(db.dbFilename, db.dbFilename+".old")

	songs[uid] = songBlock{UID: uid, URI: uri, Added: time.Now()}

	output, err := json.Marshal(&songs)

	if err != nil {
		log.Print("Error Writing file...")
	}
	ioutil.WriteFile(db.dbFilename, output, 0644)

	return nil

}
