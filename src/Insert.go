package src

import (
	"context"

	"github.com/uptrace/bun"
)

func InsertData(db *bun.DB, ctx context.Context, v map[string]interface{}) {
	_, err := db.NewInsert().Model(&v).TableExpr("transfer_log").Exec(ctx)
	if err != nil {
		panic(err)
	}
}
