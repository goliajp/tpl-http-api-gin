package core

import (
	"context"
	"fmt"
	"github.com/goliajp/http-api-gin/data/kvm"
	"github.com/goliajp/http-api-gin/env"
	"github.com/goliajp/http-api-gin/utils/idx"
	"github.com/goliajp/http-api-gin/utils/tx"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"time"
)

type Session kvm.Session

func (s *Session) MustJSON() string {
	str, err := jsoniter.MarshalToString(s)
	if err != nil {
		panic(err)
	}
	return str
}

func ParseFromJSON(str string) (*Session, error) {
	s := Session{}
	if err := jsoniter.UnmarshalFromString(str, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func (s *Session) Save(ctx context.Context, kv *redis.Client) error {
	if err := kv.Set(ctx, s.Token, s.MustJSON(), s.Expires).Err(); err != nil {
		return fmt.Errorf("failed to save session %s: %v", s.Token, err)
	}
	return nil
}

func (s *Session) IsExpired() bool {
	return s.RefreshedAt.Add(s.Expires).Before(tx.Now())
}

func (s *Session) Refresh(ctx context.Context, kv *redis.Client) error {
	s.RefreshedAt = tx.NowP()
	if err := s.Save(ctx, kv); err != nil {
		return fmt.Errorf("refresh session %s failed: %v", s.Token, err)
	}
	return nil
}

func CreateSession(ctx context.Context, kv *redis.Client, userId int, payload *string) (*Session, error) {
	s := Session{
		Token:       generateToken(),
		UserId:      userId,
		Payload:     payload,
		RefreshedAt: tx.NowP(),
		Expires:     time.Hour * time.Duration(env.ExpireHours),
	}
	if err := s.Save(ctx, kv); err != nil {
		return nil, fmt.Errorf("create session failed: %v", err)
	}
	return &s, nil
}

func DeleteSession(ctx context.Context, kv *redis.Client, token string) error {
	if err := kv.Del(ctx, token).Err(); err != nil {
		return fmt.Errorf("failed to delete session %s: %v", token, err)
	}
	return nil
}

func GetSessionByToken(ctx context.Context, kv *redis.Client, token string) (*Session, error) {
	str, err := kv.Get(ctx, token).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get session %s: %v", token, err)
	}
	return ParseFromJSON(str)
}

func generateToken() string {
	return idx.Uuid()
}
