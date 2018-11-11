package executor

import (
	"gitlab.33.cn/chain33/chain33/types"
	pkt "gitlab.33.cn/chain33/plugin/plugin/dapp/pokerbull/types"
)

func (c *PokerBull) Exec_Start(payload *pkt.PBGameStart, tx *types.Transaction, index int) (*types.Receipt, error) {
	action := NewAction(c, tx, index)
	return action.GameStart(payload)
}

func (c *PokerBull) Exec_Continue(payload *pkt.PBGameContinue, tx *types.Transaction, index int) (*types.Receipt, error) {
	action := NewAction(c, tx, index)
	return action.GameContinue(payload)
}

func (c *PokerBull) Exec_Quit(payload *pkt.PBGameQuit, tx *types.Transaction, index int) (*types.Receipt, error) {
	action := NewAction(c, tx, index)
	return action.GameQuit(payload)
}