package main

import (
	u "goprac/utils"
	"log"
	"math/rand"
	"sync"
	"time"
)

const gameDuration = 60

const (
	Idle       = 0
	InProgress = 1
	Ended      = 2
	WinnerPicked = 3
)

type Game struct {
	ID        int64     `json:"id"`
	NewTime   time.Time `json:"new_time,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
	Duration  int       `json:"duration"`
	UserBets  []UserBet `json:"userBets"`
	State     int        `json:"state"`
	BetsMutex *sync.Mutex
	StateMutex *sync.Mutex
}

type GameManager struct {
	mutex       sync.Mutex
	pastGames   map[int64]Game
	currentGame Game
	events      chan interface{}
}

type Player struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func NewPlayer(uid int, email string) *Player {
	return &Player{
		Id:    uid,
		Email: email,
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
	Bets   []*Bet  `json:"bets"`
	Player *Player `json:"player"`
}

func NewUserBet(bet *Bet, player *Player) *UserBet {
	return &UserBet{
		Bets:   []*Bet{bet},
		Player: player,
	}
}

func NewGameManager() *GameManager {
	return &GameManager{
		mutex:     sync.Mutex{},
		pastGames: make(map[int64]Game),
		events:    make(chan interface{}),
	}
}

func (gm *GameManager) Events() chan interface{} {
	return gm.events
}

func (gm *GameManager) NewGame() {
	gm.mutex.Lock()
	now := time.Now()
	gameID := now.UnixNano()

	newGame := Game{
		ID:        gameID,
		NewTime:   now,
		Duration:  gameDuration,
		BetsMutex: &sync.Mutex{},
		StateMutex: &sync.Mutex{},
		UserBets:  make([]UserBet, 0),
		State: 		Idle,
	}

	gm.currentGame = newGame

	gm.mutex.Unlock()

	gm.events <- GameEvent{
		Type: "new-game",
		Game: newGame,
	}

	log.Println("[GAME] New game started")
	log.Println("[GAME] Waiting for bets from at least 2 ppl..")
}

func (gm *GameManager) GetCurrentGame() *Game {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	return &gm.currentGame
}

func (gm *GameManager) StartGame() {
	gm.mutex.Lock()
	gm.currentGame.StartTime = time.Now()

	gm.currentGame.SetState(InProgress)

	gm.events <- GameEvent{
		Type: "start-game",
		Game: gm.currentGame,
	}

	gm.mutex.Unlock()

	log.Println("[GAME] Game Started")

	defer func() {
		for d := range u.Countdown(u.NewTicker(time.Second), 5*time.Second) {
			gm.events <- CountDownEvent{
				TimeLeft: d.Seconds(),
			}
		}
		gm.EndGame()
	}()
}

func (gm *GameManager) EndGame() {
	gm.mutex.Lock()

	gm.currentGame.EndTime = time.Now()
	gm.pastGames[gm.currentGame.ID] = gm.currentGame

	gm.currentGame.SetState(Ended)

	gm.events <- GameEvent{
		Type: "end-game",
		Game: gm.currentGame,
	}

	gm.mutex.Unlock()

	log.Print("[GAME] Has ended, no more bets!")
	_ = gm.currentGame.GetWinner()

	defer func() {
		log.Println("[GAME] starting new game in 15 seconds...")
		time.Sleep(time.Second * 15)
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
	return totalPrice
}

func (g *Game) GetWinner() *int {

	log.Print("[GAME] picking a winner...")
	totalPricePerUser := make(map[int]float64)

	g.BetsMutex.Lock()
	defer g.BetsMutex.Unlock()

	totalPrice := g.GetTotalPrice()

	for _, userBet := range g.UserBets {
		for _, bet := range userBet.Bets {
			totalPricePerUser[userBet.Player.Id] = totalPricePerUser[userBet.Player.Id] + bet.Amount
		}
	}

	log.Printf("[GAME] Total price: %.2f", totalPrice)

	// Fill pool
	pool := make([]int, 100)
	startPoint := 0
	for userID, p := range totalPricePerUser {
		share := (p / totalPrice) * 100
		for i := startPoint; i <= startPoint + int(share); i++ {
			pool[i-1] = userID
		}
		startPoint = int(share)
	}

	log.Printf("[GAME] Pool length: %d, Pool: %v", len(pool), pool)
	// Pick random number from pool
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomInt := r.Intn(100)
	winningUserId := pool[randomInt]

	log.Printf("[GAME] Winner userId: %v", winningUserId)
	log.Printf("====================================")

	for _, userBet := range g.UserBets {
		if userBet.Player.Id == winningUserId {
			g.SetState(WinnerPicked)
			gameManager.events <- GameEvent{
				Type:   "winner-picked",
				Game:   *g,
				Player: userBet.Player,
				Amount: totalPricePerUser[winningUserId],
			}
		}
	}

	return &pool[randomInt]
}
