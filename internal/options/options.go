package options

type Opts struct {
	Repeat     bool   `short:"r" long:"repeat" description:"Start repeating all words"`
	Count      int    `short:"c" long:"count" default:"0" description:"Start repeating selected number of words"`
	List       bool   `short:"l" long:"list" description:"Print all words in storage"`
	AddWord    string `short:"a" long:"add" description:"Add a word to storage"`
	DeleteWord string `short:"d" long:"delete" description:"Delete a word from storage"`
}
