package internal

import (
	"gameserver/msg"
	"reflect"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	// register the handler of `Hello` message to `game` module handleHello
	handler(&msg.Hello{}, handleHello)
	handler(&msg.GameState{}, handleGameState)
	handler(&msg.GuessGame{}, handleGuessGame)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)

	// content message
	log.Debug("hello %v", m.Name)
	a.WriteMsg(&msg.Hello{
		Name: "hi, thank",
	})
}
func handleGameState(args []interface{}) {
	m := args[0].(*msg.GameState)
	a := args[1].(gate.Agent)

	// content message
	log.Debug("game state %v", m.Status)

	if m.Status == 1 {
		stopRollingDice()
		a.WriteMsg(&msg.GameStateResult{
			Msg: "Game stopped",
		})
	}
	if m.Status == 2 {
		rollDice()
		a.WriteMsg(&msg.GameStateResult{
			Msg: "Game started",
		})
	}

}
func handleGuessGame(args []interface{}) {
	m := args[0].(*msg.GuessGame)
	a := args[1].(gate.Agent)

	// content message
	log.Debug("game state %v", m.Text)

	status, text, remark := guessDice(m.Text)
	a.WriteMsg(&msg.GuessGameResult{Status: status, Text: text, Remark: remark})
}
