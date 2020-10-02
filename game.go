package main

import (
	m "goprac/models"
	u "goprac/utils"
	"log"
	"math"
	"sync"
	"time"
)

const gameDuration = 60

var proof *Proof

const (
	Idle         = 0
	InProgress   = 1
	Ended        = 2
	WinnerPicked = 3
)

type Game struct {
	ID        int64     `json:"id"`
	NewTime   time.Time `json:"new_time,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
	Duration  int       `json:"duration"`
	UserBets  []UserBet `json:"userBets"`
	State     int       `json:"state"`

	BetsMutex  *sync.Mutex
	StateMutex *sync.Mutex
}

type GameManager struct {
	mutex       sync.Mutex
	pastGames   map[int64]Game
	currentGame *Game
	events      chan m.Event
}

func (gm *GameManager) GetCurrentGame() *Game {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	return gm.currentGame
}

func NewGameManager() *GameManager {
	return &GameManager{
		mutex:     sync.Mutex{},
		pastGames: make(map[int64]Game),
		events:    make(chan m.Event),
	}
}

func (gm *GameManager) Events() chan m.Event {
	return gm.events
}

type Player struct {
	Id    int    `json:"id"`
	Username string `json:"email"`
	Avatar string `json:"avatar"`
}

func NewPlayer(uid int, username string, avatar string) *Player {
	return &Player{
		Id:    uid,
		Username: username,
		Avatar: avatar,
	}
}

type Bet struct {
	Amount  float64 `json:"amount"`
	Created time.Time
}

func NewBet(amount float64) *Bet {
	return &Bet{
		Amount:  amount,
		Created: time.Now(),
	}
}

type UserBet struct {
	Bets        []*Bet  `json:"bets"`
	Player      *Player `json:"player"`
	Share       float64 `json:"share"`
	StartTicket int     `json:"startTicket,omitempty"`
	EndTicket   int     `json:"endTicket,omitempty"`
}

func NewUserBet(bet *Bet, player *Player) *UserBet {
	return &UserBet{
		Bets:   []*Bet{bet},
		Player: player,
	}
}

func (ub UserBet) GetTotalBet() (total float64) {
	for _, bet := range ub.Bets {
		total = total + bet.Amount
	}
	return
}

func (gm *GameManager) NewGame() {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	now := time.Now()
	gameID := now.UnixNano()

	newGame := &Game{
		ID:         gameID,
		NewTime:    now,
		Duration:   gameDuration,
		BetsMutex:  &sync.Mutex{},
		StateMutex: &sync.Mutex{},
		UserBets:   make([]UserBet, 0),
		State:      Idle,
	}

	proof, _ = NewProof(nil, nil, gameID)
	randomNr, _ := proof.Calculate()
	log.Printf("[PROOF] Random number used to pick winner: %.2f", randomNr)

	gm.currentGame = newGame
	gm.events <- GameEvent{
		Type: "new-game",
		Game: *newGame,
	}

	log.Println("[GAME] New game started")
	log.Println("[GAME] Waiting for bets from at least 2 ppl..")
}

func (gm *GameManager) StartGame() {
	gm.mutex.Lock()
	defer func() {
		for d := range u.Countdown(u.NewTicker(time.Second), 30*time.Second) {
			gm.events <- CountDownEvent{
				TimeLeft: d.Seconds(),
			}
		}
		gm.EndGame()
	}()

	gm.currentGame.StartTime = time.Now()
	gm.currentGame.SetState(InProgress)
	gm.events <- GameEvent{
		Type: "start-game",
		Game: *gm.currentGame,
	}

	gm.mutex.Unlock()
}

func (gm *GameManager) EndGame() {
	gm.mutex.Lock()

	gm.currentGame.EndTime = time.Now()
	gm.pastGames[gm.currentGame.ID] = *gm.currentGame

	gm.currentGame.SetState(Ended)

	gm.events <- GameEvent{
		Type: "end-game",
		Game: *gm.currentGame,
	}

	gm.mutex.Unlock()
	log.Print("[GAME] Has ended, no more bets!")
	gm.currentGame.GetWinner()

	defer func() {
		log.Println("[GAME] starting new game in 5 seconds...")
		time.Sleep(time.Second * 20)
		gm.NewGame()
	}()
}

func (g *Game) SetState(state int) {
	g.StateMutex.Lock()
	defer g.StateMutex.Unlock()
	g.State = state
}

func (g *Game) PlaceBet(player *Player, amount float64) {
	log.Printf("[GAME] NEW BET:($%.2f) FROM => Id: %d ", amount, player.Id)
	g.BetsMutex.Lock()

	if g.State != InProgress && g.State != Idle {
		// betting not allowed in this state
		g.BetsMutex.Unlock()
		return
	}

	bet := NewBet(amount)
	// lookup current user bet if exist.
	found := false
	for i, userBet := range g.UserBets {
		if userBet.Player.Id == player.Id {
			g.UserBets[i].Bets = append(userBet.Bets, bet)
			found = true
		}
	}

	if !found {
		userBet := NewUserBet(bet, player)
		g.UserBets = append(g.UserBets, *userBet)
	}

	g.BetsMutex.Unlock()

	gameManager.events <- GameEvent{
		Type:   "bet-placed",
		Game:   *gameManager.GetCurrentGame(),
		Player: player,
		Amount: amount,
	}

	go g.CalculateShares()

	log.Printf("[GAME] TOTAL BETS:($%.2f) ", g.GetTotalPrice())

	if g.StartTime.IsZero() && len(g.UserBets) >= 2 {
		log.Print("[GAME] Enough players starting game...")
		go gameManager.StartGame()
	}
}

func (g Game) GetTotalPrice() (totalPrice float64) {
	for _, userBet := range g.UserBets {
		for _, bet := range userBet.Bets {
			totalPrice = totalPrice + bet.Amount
		}
	}
	return
}

func (g *Game) CalculateShares() {
	g.BetsMutex.Lock()
	defer func() {
		gameManager.events <- GameEvent{
			Type: "shares-updated",
			Game: *g,
		}
		g.BetsMutex.Unlock()
	}()

	// using cents to increase accuracy of 'user-tickets'.
	total := math.Round(g.GetTotalPrice())
	totalCents := total * 100
	startTicket := 0

	log.Printf("[CALC-SHARES] Total pot: $%f", total)

	for i := range g.UserBets {
		ub := &g.UserBets[i]

		betInCents := int(ub.GetTotalBet()) * 100
		ub.StartTicket = startTicket
		ub.EndTicket = startTicket + betInCents
		ub.Share = (100 / totalCents) * float64(betInCents)

		startTicket += betInCents + 1
		//	log.Printf("[CALC-SHARES] User: %d | StartTicket: %d | EndTicket: %d | Share: %f |", ub.Player.Id, ub.StartTicket, ub.EndTicket, ub.Share)
	}
}

func (g *Game) GetWinner() {
	log.Print("[GAME] picking a winner...")
	g.BetsMutex.Lock()
	defer g.BetsMutex.Unlock()

	totalTickets := math.Round(g.GetTotalPrice()) * 100.0

	randomNr, _ := proof.Calculate()
	winningTicket := int(randomNr / 100.0 * totalTickets)
	winningPercentage := (100.0 / totalTickets) * (randomNr / 100.0 * totalTickets)

	log.Printf("[GAME] winning ticket: %d", winningTicket)

	for _, userBet := range g.UserBets {
		if userBet.StartTicket <= winningTicket && userBet.EndTicket >= winningTicket {
			g.SetState(WinnerPicked)
			gameManager.events <- WinnerPickedEvent{
				Player:     userBet.Player,
				Ticket:     winningTicket,
				Percentage: winningPercentage,
				Amount:     g.GetTotalPrice() - userBet.GetTotalBet(),
			}
			log.Printf("[GAME]:::The winning userID: %d:::...", userBet.Player.Id)
		}
	}
}
