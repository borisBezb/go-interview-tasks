package main

import (
	"context"
	"fmt"

	"github.com/borisBezb/go-interview-tasks/internal/search"
)

/**
 * Мы выполняем поиск вариантов размещения по 3-ем разным поставщикам
 * Опрашивая каждого поставщика мы формируем общий результат поиска, мы не знаем сколько времени
 * занимает поиск у каждого поставщика, поэтому каждый раз время поиска будет отличаться
 *
 * Вопросы:
 * - В чем недочет текущего поиска?
 * - Как нам оптимизировать время поиска?
 *
 * Задачи:
 * - Модифицируйте код программы main, чтобы сократить общее время поиска
 * - Затем модифицируйте код, чтобы поиск укладывался в 5 секунд, а результаты поставщиков, которые не успели ответить во время
 * мы отсекаем
 *
 */
func main() {
	providers := []search.Provider{
		search.NewProvider("a"),
		search.NewProvider("b"),
		search.NewProvider("c"),
	}
	ctx := context.Background()
	var results []search.Offer

	for _, p := range providers {
		providerResults, err := p.Search(ctx)
		if err != nil {
			panic(err)
		}

		results = append(results, providerResults...)
	}

	fmt.Println(results)
}
