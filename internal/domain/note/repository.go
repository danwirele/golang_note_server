package note

import (
	"fmt"
	"notes_server/internal/algorithm"
	"notes_server/internal/cryp"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	database *gorm.DB
}

type Repository interface {
	CreateNote(data string) (*Note, error)
	ReadNote(code, uid string) (*Note, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{database: db}
}

// CreateNote implements Repository.
func (r *repository) CreateNote(data string) (*Note, error) {
	//create unique secret key
	UID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	//Count all models in table
	var RowsCount int64
	r.database.Model(&Note{}).Count(&RowsCount)

	//geting short link for GET request
	code := algorithm.GetShortLink(int(RowsCount) + 1)

	//Encoding secret key for data encrypt
	key := cryp.Encode(UID.String())

	//encrypt data
	var slice []byte = key[:]
	HashData, err := cryp.EncryptMessage(slice, data)
	if err != nil {
		return nil, err
	}

	//create entity in db
	note := &Note{Uid: UID.String(), Data: HashData, Code: code}
	res := r.database.Create(note)
	if err := res.Error; err != nil {
		return nil, err
	}

	return note, nil
}

// ReadNote implements Repository.
func (r *repository) ReadNote(code, uid string) (*Note, error) {
	//decode code
	index := algorithm.GetIdFromLink(code)
	var Result Note
	//find row with this id in bd
	r.database.Table("notes").Select("*").Where("id = ?", index).First(&Result)
	key := cryp.Encode(uid)
	var slice []byte = key[:]
	data, err := cryp.DecryptMessage(slice, Result.Data)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	Result.Data = data
	return &Result, nil
}
