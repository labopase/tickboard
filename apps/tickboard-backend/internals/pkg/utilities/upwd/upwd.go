package upwd

import "golang.org/x/crypto/bcrypt"

type Cost int

const (
	// the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
	DefaultCost Cost = Cost(bcrypt.DefaultCost)
	// the minimum allowable cost as passed in to GenerateFromPassword
	MinCost Cost = Cost(bcrypt.MinCost)
	// the maximum allowable cost as passed in to GenerateFromPassword
	MaxCost Cost = Cost(bcrypt.MaxCost)
)

type Config struct {
	Cost Cost
}

// HashWithConfig hashes a string with the given cost.
func HashWithConfig(str string, cfg Config) (string, error) {
	if cfg.Cost < MinCost {
		cfg.Cost = DefaultCost
	}
	if cfg.Cost > MaxCost {
		cfg.Cost = MaxCost
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), int(cfg.Cost))
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// Hash hashes a string with the default cost.
func Hash(str string) (string, error) {
	return HashWithConfig(str, Config{
		Cost: DefaultCost,
	})
}

// Verify checks if a string matches a hashed string.
func Verify(str string, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str))
}
