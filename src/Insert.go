package src

import (
	"context"

	"github.com/uptrace/bun"
)

func InsertData(db *bun.DB, ctx context.Context,h string, v map[string]interface{}) {
	_, err := db.NewInsert().Model(&v).TableExpr(h).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
