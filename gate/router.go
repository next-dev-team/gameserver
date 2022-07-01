package gate

import (
	"gameserver/game"
	"gameserver/msg"
)

func init() {
	// Route Hello to game
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
	// Route to game state
	msg.Processor.SetRouter(&msg.GameState{}, game.ChanRPC)
	// Route guess game
	msg.Processor.SetRouter(&msg.GuessGame{}, game.ChanRPC)
}
