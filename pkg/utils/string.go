package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Converter struct{}

func (c *Converter) ToString(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case bool:
		return strconv.FormatBool(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case []uint8:
		return string(v)
	default:
		b, err := json.Marshal(v)
		if err != nil {
			log.Println("Error on ConvertString: ", err)
			return ""
		}
		return string(b)
	}
}

func (c *Converter) ToInt(v interface{}) int {
	switch v := v.(type) {
	case string:
		res, _ := strconv.Atoi(strings.TrimSpace(v))
		return res
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	case []byte:
		res, _ := strconv.Atoi(string(v))
		return res
	default:
		return 0
	}
}

func (c *Converter) StructToString(s interface{}) string {
	b, err := json.Marshal(s)
	if err != nil {
		log.Println("Error on StructToString: ", err)
		return ""
	}
	return string(b)
}

func (c *Converter) ToInt64(v interface{}) int64 {
	switch v := v.(type) {
	case string:
		res, _ := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		return res
	case int:
		return int64(v)
	case int64:
		return v
	case float64:
		return int64(v)
	case []byte:
		res, _ := strconv.ParseInt(string(v), 10, 64)
		return res
	default:
		return int64(0)
	}
}

type TimeHelper struct{}

func (t *TimeHelper) GetLocalTime() time.Time {
	return time.Now().Local()
}

type UuidHelper struct{}

func (u *UuidHelper) GenerateUUID() uuid.UUID {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Println(fmt.Errorf("failed to generate UUID: %w", err))
		return uuid.Nil
	}
	return id
}

func (u *UuidHelper) ParseUuid(v string) uuid.UUID {
	return uuid.MustParse(v)
}

type Security struct{}

func (s *Security) GenerateToken(email string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (s *Security) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}

func (s *Security) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
