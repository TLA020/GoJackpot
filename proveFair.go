package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"log"
	rnd "math/rand"
	"strconv"
	"sync"
)

// ErrInvalidNonce is returned when doesn't create a valid random number
var ErrInvalidNonce = errors.New("invalid nonce")

const alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// Proof represents the current state of a dice proof for a single client
type Proof struct {
	ClientSeed        []byte
	ServerSeed        []byte
	BlindedServerSeed []byte
	Nonce             int64
	Lock              sync.Mutex
}

func (p *Proof) LogState() {
	wnr, _ := p.Calculate()
	log.Printf("[Proof] Random number: %f", wnr)
	log.Printf("[Proof] Client Seed (public) : %s", string(p.ClientSeed))
	log.Printf("[Proof] Server Seed (secret): %s", string(p.ServerSeed))
	log.Printf("[Proof] Server Seed (public): %s", hex.EncodeToString(p.BlindedServerSeed))
	log.Printf("[Proof] Nonce: %d", p.Nonce)
	log.Print(p.Nonce)
}

// NewProof creates a new proof from the given seeds. A clientSeed is required
// If the serverSeed is nil we create a random one
func NewProof(clientSeed []byte, serverSeed []byte, nonce int64) (*Proof, error) {
	// Validate the clientSeed
	if clientSeed == nil || len(clientSeed) == 0 {
		clientSeed = []byte("germany")
	}

	// Generate a random server seed if one isn't provided
	if serverSeed == nil || len(serverSeed) == 0 {
		serverSeed = newSeed(64)
	}

	// Hash the serverSeed to show the client
	blindedSeed := sha256.Sum256(serverSeed)

	return &Proof{
		Nonce:             nonce,
		ClientSeed:        clientSeed,
		ServerSeed:        serverSeed,
		BlindedServerSeed: blindedSeed[:],
	}, nil
}

// Roll calculates the number for the current nonce, then increments the nonce
// Doing it in this order ensures that the first once we use is 0
func (p *Proof) Roll() (float64, error) {
	p.Lock.Lock()
	defer func() {
		p.Lock.Unlock()
		p.LogState()
	}()

	// Calculate the current number from the current state
	roll, err := p.Calculate()
	if err != nil {
		log.Print(err)
		return roll, err
	}

	// Increment the nonce for next time
	p.Nonce++
	return roll, nil
}

// Calculate calculates the current value from the current state of the proof
// It does not advance the state in anyway; i.e. simply calling Calculate
// multiple times will always result in the same value unless the Nonce changes
func (p *Proof) Calculate() (float64, error) {
	// Calculate the HMAC for the current nonce
	ourHMAC := string(p.CalculateHMAC())

	// Find the first 5 character segment that converts to decimal < the max
	var randNum uint64
	var err error
	for i := 0; i < len(ourHMAC)-5; i++ {
		// Get the index for this segment and ensure it doesn't overrun the slice
		idx := i * 5
		if len(ourHMAC) < (idx + 5) {
			break
		}

		// Get 5 characters and convert them to decimal
		randNum, err = strconv.ParseUint(ourHMAC[idx:idx+5], 16, 0)
		if err != nil {
			return 0, err
		}

		// Continue unless our number was greater than our max
		if randNum <= 999999 {
			break
		}
	}

	// If even the last segment was invalid we must give up
	if randNum > 999999 {
		return 0, ErrInvalidNonce
	}

	// Normalize the number to [0,100]
	return float64(randNum%10000) / 100, nil
}

// CalculateHMAC calculates the hmac of "client seed-nonce" as a hex string
func (p *Proof) CalculateHMAC() []byte {
	h := hmac.New(sha512.New, p.ServerSeed)
	h.Write(append(append(p.ClientSeed, '-'), []byte(strconv.FormatInt(p.Nonce, 10))...))

	ourHMAC := make([]byte, 128)
	hex.Encode(ourHMAC, h.Sum(nil))
	return ourHMAC
}

// Verify takes a state and checks that the supplied number was fairly generated
func Verify(clientSeed []byte, serverSeed []byte, nonce int64, randNum float64) (bool, error) {
	proof, _ := NewProof(clientSeed, serverSeed, nonce)
	//proof.Nonce = nonce

	roll, err := proof.Calculate()
	if err != nil {
		return false, err
	}

	log.Printf("Verify returns: %f Expected: %f", roll, randNum)
	return roll == randNum, nil
}

func newSeed(size int) []byte {
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rnd.Intn(len(alpha))]
	}
	return buf
}
