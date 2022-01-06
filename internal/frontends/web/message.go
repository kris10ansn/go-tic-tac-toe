package web

const (
	MessageTypePresentBoard = "present-board"
	MessageTypeError        = "error"
	MessageTypeAddGame      = "add-game"
	MessageTypeAssignTic    = "assign-tic"
	MessageTypeEndGame      = "end-game"
	MessageTypeAwaitingMove = "await-move"
	MessageTypeClientMove   = "move"
)

type WebsocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
