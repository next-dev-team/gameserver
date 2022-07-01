package msg

import (
	"github.com/name5566/leaf/network/json"
)

type Hello struct {
	Name string
}

type GameState struct {
	Status int // 0 stop, 1 start
}
type GameStateResult struct {
	Msg string
}

type GuessGame struct {
	Text string
}
type GuessGameResult struct {
	Text   string
	Status bool
	Remark string
}

// create json processor
var Processor = json.NewProcessor()

func init() {
	// register message hello
	Processor.Register(&Hello{})
	// register game state message
	Processor.Register(&GameState{})
	// register game state result message
	Processor.Register(&GameStateResult{})
	// register guess game
	Processor.Register(&GuessGame{})
	// register guess game result
	Processor.Register(&GuessGameResult{})
}
