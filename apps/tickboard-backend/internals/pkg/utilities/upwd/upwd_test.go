package upwd_test

import (
	"testing"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/utilities/upwd"
	"golang.org/x/crypto/bcrypt"
)

func TestHash(t *testing.T) {
	t.Parallel()

	password := "super_secret_password"

	hashed, err := upwd.Hash(password)
	if err != nil {
		t.Fatalf("Hash() failed: %v", err)
	}

	if hashed == "" {
		t.Error("Hash() returned empty string")
	}

	if hashed == password {
		t.Error("Hash() returned the plaintext password")
	}

	// Verify the hash is valid
	err = upwd.Verify(password, hashed)
	if err != nil {
		t.Errorf("Verify() failed for valid password: %v", err)
	}
}

func TestVerify(t *testing.T) {
	t.Parallel()

	password := "my_secure_password"
	hashed, err := upwd.Hash(password)
	if err != nil {
		t.Fatalf("Failed to generate test hash: %v", err)
	}

	tests := []struct {
		name    string
		pwd     string
		hash    string
		wantErr bool
	}{
		{
			name:    "valid password",
			pwd:     password,
			hash:    hashed,
			wantErr: false,
		},
		{
			name:    "invalid password",
			pwd:     "wrong_password",
			hash:    hashed,
			wantErr: true,
		},
		{
			name:    "invalid hash format",
			pwd:     password,
			hash:    "invalid_hash_string",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := upwd.Verify(tt.pwd, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashWithConfig(t *testing.T) {
	t.Parallel()

	password := "test_password"

	tests := []struct {
		name         string
		cfg          upwd.Config
		expectedCost int
	}{
		{
			name: "explicit default cost",
			cfg: upwd.Config{
				Cost: upwd.DefaultCost,
			},
			expectedCost: int(upwd.DefaultCost),
		},
		{
			name: "min cost",
			cfg: upwd.Config{
				Cost: upwd.MinCost,
			},
			expectedCost: int(upwd.MinCost),
		},
		{
			name: "below min cost gets upgraded to default",
			cfg: upwd.Config{
				Cost: upwd.MinCost - 1,
			},
			expectedCost: int(upwd.DefaultCost),
		},
		{
			name: "zero cost gets upgraded to default",
			cfg: upwd.Config{
				Cost: 0,
			},
			expectedCost: int(upwd.DefaultCost),
		},
		// Exclude tests with `MaxCost` or values > `MaxCost`.
		// Hashing with a bcrypt cost of 31 requires an extreme amount of CPU time
		// and will cause the test suite to hang practically indefinitely.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			hashed, err := upwd.HashWithConfig(password, tt.cfg)
			if err != nil {
				t.Fatalf("HashWithConfig() error = %v", err)
			}

			if hashed == "" {
				t.Error("HashWithConfig() returned empty string")
			}

			// Verify the generated hash actually uses the expected cost using the bcrypt library
			actualCost, err := bcrypt.Cost([]byte(hashed))
			if err != nil {
				t.Fatalf("Failed to extract cost from hash: %v", err)
			}

			if actualCost != tt.expectedCost {
				t.Errorf("HashWithConfig() resulted in cost %d, want %d", actualCost, tt.expectedCost)
			}
		})
	}
}
