package handler

import (
	"asura/database"
	"context"
	"fmt"
	"strconv"
	"time"
)

func SetCooldown(ctx context.Context, userId string, command Command) {
	k := fmt.Sprintf("%s_%s", command.Name, userId)

	database.Cache.SetNX(ctx, k, time.Now().Unix(), time.Duration(command.Cooldown)*time.Second)
}

func GetCooldown(ctx context.Context, userId string, command Command) (time.Time, bool) {
	k := fmt.Sprintf("%s_%s", command.Name, userId)
	v := database.Cache.Get(ctx, k)

	if resp, _ := v.Result(); resp != "" {
		date, _ := strconv.ParseInt(resp, 10, 64)

		return time.Unix(date, 0), true
	} else {
		return time.Now(), false
	}
}
