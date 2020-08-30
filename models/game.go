package models

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var GlobalGameManger *GameManager
var gameDuration = 60

type Game struct {
	ID         int64               `json:"id"`
	NewTime    time.Time           `json:"new_time,omitempty"`
	StartTime  time.Time           `json:"start_time,omitempty"`
	EndTime    time.Time           `json:"end_time,omitempty"`
	Duration   int                 `json:"duration"`
	Bets      map[uint]UserBet	    `json:"bets"`
	itemsMutex *sync.Mutex `json:"lik"`
}

type UserBet struct {
	Gambler  *Gambler
	Amount float64
}

type NewGameEvent struct {
	Game Game
}

type StartGameEvent struct {
	Game Game
}

type EndGameEvent struct {
	Game Game
}

type NewBetEvent struct {
	Game  Game
	Bet	UserBet
}

type GameManager struct {
	mutex       sync.Mutex
	pastGames   map[int64]Game
	currentGame Game
	events      chan interface{}
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
		Bets:      make(map[uint]UserBet),
	}
	gm.currentGame = newGame

	gm.mutex.Unlock()

	// Fire event
	gm.events <- NewGameEvent{
		Game: newGame,
	}

	log.Println("New game")
	log.Println("Waiting for bets..")

	time.Sleep(time.Second * time.Duration(gameDuration))
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

	log.Println("Bets are placed, starting game..")

	defer func() {
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

	winner := gm.currentGame.GetWinner()
	log.Print(winner)

	//log.Println("Game is ended, winner is...")
	//log.Println("Starting next game..")
	// Create new game
	gm.NewGame()
}

func (g *Game) PlaceBet(gambler *Gambler, bet UserBet) {
	log.Printf("Placing ($ %f)bet for: %d ", bet.Amount, gambler.Account.ID)
	g.itemsMutex.Lock()
	defer g.itemsMutex.Unlock()

	// lookup current user bet if exist.
	userBet, found := g.Bets[gambler.Account.ID]
	if !found {
		g.Bets[gambler.Account.ID] = UserBet{}
	}

	userBet.Gambler = gambler
	userBet.Amount = userBet.Amount + bet.Amount
	g.Bets[gambler.Account.ID] = userBet

	GlobalGameManger.events <- NewBetEvent{
		Game:  *g,
		Bet: userBet ,
	}

	log.Printf("Total price of pot is now: %f", g.GetTotalPrice())
}

func (g Game) GetTotalPrice() (totalPrice float64) {
	for _, bet := range g.Bets {
			totalPrice = totalPrice + bet.Amount
	}
	return totalPrice
}

func (g *Game) GetWinner() *Gambler	{

	totalPricePerUser := make(map[uint]float64)
	var totalPrice float64

	g.itemsMutex.Lock()
	defer g.itemsMutex.Unlock()


	for _, bet := range g.Bets {
		totalPrice = totalPrice + bet.Amount
		totalPricePerUser[bet.Gambler.Account.ID] = totalPricePerUser[bet.Gambler.Account.ID] + bet.Amount
	}

	log.Printf("Total price: %f", totalPrice)
	log.Print(totalPricePerUser)

	// Fill pool
	pool := make([]uint, 100)
	for userID, p := range totalPricePerUser {
		share := (p / totalPrice) * 100
		for i := 1; i <= int(share); i++ {
			pool[i - 1] = userID
		}
	}

	log.Printf("Pool length: %d, Pool: %v", len(pool), pool)
	// Pick random number from pool
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomInt := r.Intn(100)

	log.Printf("Winner: %v", pool[randomInt])


	return &Gambler{}
}
