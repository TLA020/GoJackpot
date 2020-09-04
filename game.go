package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var gameDuration = 60

type Game struct {
	ID         int64       `json:"id"`
	NewTime    time.Time   `json:"new_time,omitempty"`
	StartTime  time.Time   `json:"start_time,omitempty"`
	EndTime    time.Time   `json:"end_time,omitempty"`
	Duration   int         `json:"duration"`
	Bets       []UserBet   `json:"bets"`
	itemsMutex *sync.Mutex
}

type GameManager struct {
	mutex       sync.Mutex
	pastGames   map[int64]Game
	currentGame Game
	events      chan interface{}
}

type Player struct {
	UserId int `json:"userId"`
	Email string `json:"email"`
}

func NewPlayer(uid int, email string) *Player {
	return &Player{
		UserId: uid,
		Email: email,
	}
}


type Bet struct {
	Amount float64 `json:"amount"`
	Created time.Time
}

func NewBet(amount float64) *Bet {
	return &Bet{
		Amount: amount,
		Created: time.Now(),
	}
}

type UserBet struct {
	Bets []*Bet `json:"bets"`
	Player *Player `json:"player"`
}

func NewUserBet(bet *Bet, player *Player) *UserBet {
	return &UserBet{
		Bets: []*Bet {bet},
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
		ID:         gameID,
		NewTime:    now,
		Duration:   gameDuration,
		itemsMutex: &sync.Mutex{},
		Bets:       make([]UserBet, 0),
	}

	gm.currentGame = newGame

	gm.mutex.Unlock()

	//Fire event
	gm.events <- NewGameEvent{
		Game: newGame,
	}

	log.Println("[GAME] New game started")
	log.Println("[GAME] Waiting for bets from at least 2 ppl..")

	//time.Sleep(time.Second * time.Duration(gameDuration))
}

func (gm *GameManager) GetCurrentGame() *Game {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	return &gm.currentGame
}

func (gm *GameManager) StartGame() {
	gm.mutex.Lock()
	gm.currentGame.StartTime = time.Now()

	gm.events <- StartGameEvent{
		Game: gm.currentGame,
	}

	gm.mutex.Unlock()

	log.Println("[GAME] Game Started")

	defer func() {
		log.Println("[GAME] 30 seconds left...")
		time.Sleep(time.Second * 30)
		gm.EndGame()
	}()
}

func (gm *GameManager) EndGame() {
	gm.mutex.Lock()

	gm.currentGame.EndTime = time.Now()
	gm.pastGames[gm.currentGame.ID] = gm.currentGame

	gm.events <- EndGameEvent{
		Game: gm.currentGame,
	}

	gm.mutex.Unlock()

	log.Print("[GAME] Has ended, no more bets!")
	_ = gm.currentGame.GetWinner()
	gm.NewGame()
}

func (g *Game) PlaceBet(player *Player, amount float64) {
	log.Printf("[GAME] NEW BET:($%.2f) FROM => UserId: %d ", amount, player.UserId)
	g.itemsMutex.Lock()

	bet := NewBet(amount)
	// lookup current user bet if exist.
	found := false
	for i, userBet := range g.Bets {
		if userBet.Player.UserId == player.UserId {
			g.Bets[i].Bets = append(userBet.Bets, bet)
			found = true
		}
	}

	if !found {
		userBet := NewUserBet(bet, player)
		g.Bets = append(g.Bets, *userBet)
	}

	gameManager.events <- NewBetEvent{
		Game: *g,
		Bet:  *bet,
	}

	log.Printf("[GAME] TOTAL BETS:($%.2f) ", g.GetTotalPrice())
	g.itemsMutex.Unlock()

	if g.StartTime.IsZero() && len(g.Bets) >= 2 {
		log.Print("[GAME] Enough players starting game...")
		gameManager.StartGame()
	}
}

func (g Game) GetTotalPrice() (totalPrice float64) {
	for _, userBet := range g.Bets {
		for _, bet := range userBet.Bets {
			totalPrice = totalPrice + bet.Amount
		}
	}
	return totalPrice
}

func (g *Game) GetWinner() *int {

	log.Print("[GAME] picking a winner...")
	totalPricePerUser := make(map[int]float64)

	g.itemsMutex.Lock()
	defer g.itemsMutex.Unlock()

	totalPrice := g.GetTotalPrice()

	for _, userBet := range g.Bets {
		for _, bet := range userBet.Bets {
			totalPricePerUser[userBet.Player.UserId] = totalPricePerUser[userBet.Player.UserId] + bet.Amount
		}
	}

	log.Printf("[GAME] Total price: %.2f", totalPrice)
	//log.Print(totalPricePerUser)

	// Fill pool
	pool := make([]int, 100)
	for userID, p := range totalPricePerUser {
		share := (p / totalPrice) * 100
		for i := 1; i <= int(share); i++ {
			pool[i-1] = userID
		}
	}

	//log.Printf("[GAME] Pool length: %d, Pool: %v", len(pool), pool)
	// Pick random number from pool
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomInt := r.Intn(100)

	log.Printf("[GAME] Winner userId: %v", pool[randomInt])
	log.Printf("====================================")

	return &pool[randomInt]
}
