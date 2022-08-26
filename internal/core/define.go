package core

const (
	sharp = "#"
	flat  = "♭"
)

type Note string

const (
	NoteC Note = "C"
	NoteD Note = "D"
	NoteE Note = "E"
	NoteF Note = "F"
	NoteG Note = "G"
	NoteA Note = "A"
	NoteB Note = "B"
)

var noteIndex = map[Note]int8{
	NoteC: 0,
	NoteD: 1,
	NoteE: 2,
	NoteF: 3,
	NoteG: 4,
	NoteA: 5,
	NoteB: 6,
}

var noteInterval = map[Note]int8{
	NoteC: 0,
	NoteD: 2,
	NoteE: 4,
	NoteF: 5,
	NoteG: 7,
	NoteA: 9,
	NoteB: 11,
}

var intToCN = []string{"一", "二", "三", "四", "五", "六", "七", "八", "九"}
