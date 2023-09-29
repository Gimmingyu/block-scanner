package entity

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	gorm.Model `gorm:"embedded"`
	UUID       string `gorm:"unique;not null"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
}

func (u User) Table() string {
	return "users"
}

func (u *User) ComparePassword(password string) (err error) {
	// Parsing the Argon2 hash string
	parts := strings.Split(u.Password, "$")
	if len(parts) != 6 {
		return errors.New("invalid hash format")
	}

	var version int
	_, err = fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		return err
	}

	if version != argon2.Version {
		return fmt.Errorf("incompatible argon2 version (stored: %d, current: %d)", version, argon2.Version)
	}

	var memory, iterations uint32
	var parallelism uint8
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil {
		return err
	}

	salt, err := base64.StdEncoding.DecodeString(parts[4])
	if err != nil {
		return err
	}

	storedHash, err := base64.StdEncoding.DecodeString(parts[5])
	if err != nil {
		return err
	}

	// Generate the hash to compare
	comparisonHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, uint32(len(storedHash)))

	// Perform constant-time comparison
	if !constantTimeCompare(storedHash, comparisonHash) {
		return errors.New("passwords do not match")
	}

	return nil
}

func constantTimeCompare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	var result byte
	for i := 0; i < len(a); i++ {
		result |= a[i] ^ b[i]
	}

	return result == 0
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	salt := make([]byte, 16)
	if _, err = rand.Read(salt); err != nil {
		return
	}

	u.UUID = uuid.New().String()
	hash := argon2.IDKey([]byte(u.Password), salt, 1, 64*1024, 4, 32)
	encodedHash := base64.StdEncoding.EncodeToString(hash)
	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	u.Password = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, 64*1024, 1, 4, encodedSalt, encodedHash)
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Password == "" {
		return
	}

	salt := make([]byte, 16)
	if _, err = rand.Read(salt); err != nil {
		return
	}

	hash := argon2.IDKey([]byte(u.Password), salt, 1, 64*1024, 4, 32)
	encodedHash := base64.StdEncoding.EncodeToString(hash)
	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	u.Password = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, 64*1024, 1, 4, encodedSalt, encodedHash)
	return
}
