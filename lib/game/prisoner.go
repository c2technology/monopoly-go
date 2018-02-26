package game

type Prisoner interface {
	InJail() bool
	SentenceRemaining() int
	PostBail()
}
