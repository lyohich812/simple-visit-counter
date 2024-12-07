package main

import (
	"log/slog"
	"math/rand"
	"sync"
	"sync/atomic"
)

func main() {
	var visitCount uint64 = 0
	var wg sync.WaitGroup

	dumbVisitorsCount := rand.Intn(10) + 1

	slog.Info("Dumb visitors coming", "count", dumbVisitorsCount)

	for i := 0; i < dumbVisitorsCount; i++ {
		wg.Add(1)
		go incVisitCount(&visitCount, &wg)
	}
	wg.Wait()
}

func incVisitCount(visitCount *uint64, wg *sync.WaitGroup) {
	atomic.AddUint64(visitCount, 1)
	slog.Info("Counter incremented", "value", atomic.LoadUint64(visitCount))
	wg.Done()
}

// Задача 1: Подсчёт количества посетителей на веб-сайте
// Описание задачи:
// Вы разрабатываете систему для отслеживания количества уникальных пользователей,
// посещающих веб-сайт. Каждое посещение сайта увеличивает счётчик на единицу.
// Нужно обеспечить корректную работу программы, когда одновременно несколько пользователей
// (горутин) могут увеличивать этот счётчик.

// Требования:
// Реализуйте программу, которая с помощью горутин увеличивает счётчик посещений.
// Используйте пакет sync/atomic для безопасного увеличения счётчика.
// Подсчитайте итоговое количество посещений, которое должно быть равно количеству запущенных
// горутин.
