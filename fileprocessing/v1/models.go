package v1

import "time"

//Delim is the FS-safe delimeter that should replace any other delimeter before sending filepath to any peer.
//
//Its use will reduce number of possible delimeters to check in path and it does not need any escape in strings.
//So conversion becomes very simple by replacing 'X' to Delim to 'Y'
var Delim = ">"

//RootPointer is the text representation of root path in FS-safe way
var RootPointer = "!ROOT_DIR!"

//File represents file data into DB
type File struct {
	ID          uint `gorm:"primaryKey"`
	Hash        string
	Name        string `gorm:"uniqueIndex:file"`
	Path        string `gorm:"uniqueIndex:file"`
	Owner       uint
	Size        int64
	Ext         string
	FSUpdatedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsRenamed   bool
}

//Folder represents folder data into DB
type Folder struct {
	ID          uint `gorm:"primaryKey"`
	Hash        string
	Name        string `gorm:"uniqueIndex:folder"`
	Path        string `gorm:"uniqueIndex:folder"`
	Owner       uint
	Size        int64
	Items       int
	FSUpdatedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsRenamed   bool
}
