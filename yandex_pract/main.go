package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unicode"
)

func main() {
	ints := []int{1, 2, 3, 4, 5}
	in := make(chan int)
	pred := func(i int) bool {
		return i%2 == 0
	}
	go func() {
		for p := range Filter(in, pred) {
			fmt.Println(p)
		}
	}()

	go func() {
		for _, i := range ints {
			in <- i
		}
	}()

	time.Sleep(1 * time.Second)
}

//Описание:
//Проверьте, что строка — корректный slug:
//- состоит только из строчных латинских букв a–z, цифр 0–9 и дефисов -;
//- не начинается и не заканчивается дефисом;
//- не содержит подряд идущих дефисов "--";
//- длина от 1 до 64 символов.
//
//Вход: строка s.
//Выход: true, если строка корректна; иначе false.

func IsValidSlug(s string) bool {
	runes := []rune(s)
	l := len(runes)

	if l < 1 || l > 64 {
		return false
	}

	if runes[0] == '-' || runes[l-1] == '-' {
		return false
	}

	prev := rune(0)
	for i, r := range runes {
		if !unicode.IsNumber(r) && !(unicode.Is(unicode.Latin, r) && unicode.IsLower(r)) && r != '-' {
			return false
		}

		if i != 0 && prev == '-' && r == '-' {
			return false
		}

		prev = r
	}

	return true
}

//Описание:
//Замените любой непустой блок пробельных символов (пробел, табуляция, перевод строки) одним обычным пробелом. Уберите ведущие и хвостовые пробелы.
//
//Вход: строка s.
//Выход: нормализованная строка.

func NormalizeSpaces(s string) string {
	builder := strings.Builder{}
	trimmed := strings.TrimSpace(s)

	prev := rune(32)
	for _, r := range trimmed {
		if unicode.IsSpace(r) {
			if !unicode.IsSpace(prev) {
				builder.WriteRune(' ')
			}
		} else {
			builder.WriteRune(r)
		}

		prev = r
	}

	return builder.String()
}

//Описание:
//Примените функцию fn к каждому элементу in параллельно с пулом из workers горутин. Результаты должны быть в том же порядке, что и во входном слайсе.
//
//Вход:
//in — входные данные
//fn — функция обработки
//workers — количество рабочих горутин (≥1)
//Выход:
//Срез результатов длины len(in).

type MapFunc[T any, R any] func(T) R

func ParallelMap[T any, R any](in []T, fn MapFunc[T, R], workers int) []R {
	l := len(in)
	result := make([]R, l)
	indxs := make(chan int)

	if workers < 1 {
		workers = 1
	}

	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range indxs {
				result[idx] = fn(in[idx])
			}
		}()
	}

	go func() {
		for i := 0; i < l; i++ {
			indxs <- i
		}
		close(indxs)
	}()

	wg.Wait()

	return result
}

//Описание:
//Запустите все jobs параллельно в горутинах. Дождитесь завершения всех. Верните первую ненулевую ошибку. Если ошибок не было — вернуть nil.
//
//Вход: срез функций jobs.
//Выход: error.

type Job func() error

func RunAll(jobs []Job) error {
	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		res error
	)

	for _, job := range jobs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if job != nil {
				if err := job(); err != nil {
					mu.Lock()
					defer mu.Unlock()
					if res == nil {
						res = err
					}
				}
			}
		}()
	}

	wg.Wait()
	return res
}

//Описание:
//Прочитайте элементы из канала in и отправьте в выходной канал только те, для которых pred возвращает true. При закрытии входа закройте выход.
//
//Вход: канал in, предикат pred.
//Выход: канал отфильтрованных элементов.

type Pred[T any] func(T) bool

func Filter[T any](in <-chan T, pred Pred[T]) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for v := range in {
			if pred(v) {
				out <- v
			}
		}
	}()

	return out
}

//Описание:
//Скопируйте все элементы из входного канала в два выходных канала. Каждый элемент должен появиться в обоих выходах в том же порядке. При закрытии входа закройте оба выхода.
//
//Вход: канал in.
//Выход: каналы out1, out2.

func Tee[T any](in <-chan T) (chan<- T, chan<- T) {
	out1 := make(chan T)
	out2 := make(chan T)

	go func() {
		defer close(out1)
		defer close(out2)
		for v := range in {
			out1 <- v
			out2 <- v
		}
	}()

	return out1, out2
} // TODO: НЕ РАБОТАЕТ

//Описание:
//Для списка urls сделайте HTTP GET с таймаутом и верните коды ответов в том же порядке, что и вход.
//
//Требования:
//Ограничить одновременные запросы числом concurrency.
//У каждого запроса свой таймаут timeout.
//Результаты должны быть в порядке входных URL.
//Если запрос не удался — вернуть Code=0 и заполнить Err.
//
//Вход:
//список URL, concurrency≥1, timeout>0.
//
//Выход:
//срез []StatusResult длины len(urls).

type StatusResult struct {
	URL  string
	Code int   // 0, если не удалось получить ответ
	Err  error // nil, если ответ получен
}

func FetchStatuses(
	urls []string,
	concurrency int, // общий лимит одновременных запросов
	timeout time.Duration,
) []StatusResult {
	return nil
} // TODO: РЕШИТЬ СВОИМИ СИЛАМИ
