package fileprocessing

//Delim is the FS-safe delimeter that should replace any other delimeter before sending filepath to any peer.
//
//Its use will reduce number of possible delimeters to check in path and it does not need any escape in strings.
//So conversion becomes very simple by replacing 'X' to Delim to 'Y'
var Delim = ">"

//RootPointer is the text representation of root path in FS-safe way
var RootPointer = "!ROOT_DIR!"
