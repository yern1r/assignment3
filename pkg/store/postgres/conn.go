import (
	"context"
	"fmt"
)

package repository

import (
"context"
"fmt"

"github.com/jackc/pgx/v5/pgxpool"
)

func ConnPostgres(host, user string, port int, password, dbname string) (*pgxpool.Pool, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s ",
		user, password, host, port, dbname)

	pool, err := pgxpool.New(context.Background(), psqlInfo)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
