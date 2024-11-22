package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"

	"github.com/borisBezb/go-interview-tasks/internal/dictionary"
	"github.com/borisBezb/go-interview-tasks/internal/repo"
)

/**
 * У нас есть базовый репозиторий, который умеет получать список всех записей из указанной таблицы
 * Например, hotelRepo.GetList возвращает нам список всех отелей
 *
 * Модифицируйте метод GetList таким образом, чтобы мы имели возможность указывать условия выборки, такие как:
 * WHERE city_id = ?, ORDER BY name DESC, LIMIT 20, OFFSET 40
 *
 * При этом все эти опции должны работать так:
 * hotelRepo.GetList(ctx, Where("city_id", 25), OrderBy("name", "desc"), Limit(20))
 *
 */
func main() {
	dialect := goqu.Dialect("postgres")

	pgsql, err := sql.Open("postgres", os.Getenv("DB_STRING"))
	if err != nil {
		panic(err.Error())
	}

	db := dialect.DB(pgsql)

	hotelRepo := repo.NewRepo[dictionary.Hotel](db, "dictionary.hotels")

	ctx := context.Background()

	hotels, err := hotelRepo.GetList(ctx)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(hotels)
}
