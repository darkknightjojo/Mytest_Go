package gokit

import (
	"context"
	"testing"
)

func TestHashService(t *testing.T) {
	service := NewService()
	ctx := context.Background()
	hash, err := service.Hash(ctx, "JiaoShouKun")
	if err != nil {
		t.Errorf("Hash：%s", err)
	}
	ok, err := service.Validate(ctx, "JiaoShouKun", hash)
	if err != nil {
		t.Errorf("Validate: %s", err)
	}
	if !ok {
		t.Errorf("Hahs is not same with password's")
	}

}
