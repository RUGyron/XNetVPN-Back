package generics

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
	"reflect"
	"sort"
	"sync"
)

type Identifiable[K comparable, V any] interface {
	GetKey() K
	GetValue() V
}

// RunConcurrently run function(InArray) -> OutArray for huge arrays.
// Run concurrently by <batchSize> items per batch (Max <semaphoreSize> goroutines)
func RunConcurrently[IN any, OUT any](
	fn func([]IN) (*[]OUT, error),
	array []IN,
	batchSize int,
	semaphoreSize int,
	logger *log.Logger,
	logLabel string,
	logProcess string,
) []OUT {
	wg := sync.WaitGroup{}
	semaphore := make(chan struct{}, semaphoreSize)

	batchesCount := (len(array) + batchSize - 1) / batchSize
	results := make(chan *[]OUT, batchesCount)
	counter := 0

	for i := 0; i < len(array); i += batchSize {
		batchEnd := i + batchSize
		if batchEnd > len(array) {
			batchEnd = len(array)
		}
		batch := array[i:batchEnd]

		wg.Add(1)
		semaphore <- struct{}{}
		go func() {
			defer func() {
				wg.Done()
				<-semaphore
				counter++
				if logger != nil {
					if logLabel == "" {
						logger.Printf("(%-4d / %-4d) x%-3d %-15s %d %%", counter, batchesCount, batchSize, logProcess, counter*100/batchesCount)
					} else {
						logger.Printf("(%-4d / %-4d) x%-3d %-40s %-15s %d %%", counter, batchesCount, batchSize, logLabel, logProcess, counter*100/batchesCount)
					}
				}
			}()
			result, err := fn(batch)
			if err != nil {
				return
			}
			results <- result
		}()
	}
	wg.Wait()
	close(results)

	var finalResponse []OUT
	for result := range results {
		if result != nil {
			finalResponse = append(finalResponse, *result...)
		}
	}
	return finalResponse
}

// BindStructWithResponse Template for fast usage: read, serialize, unmarshal <http.Response> -> <struct>
func BindStructWithResponse[OUT interface{}](response http.Response, model *OUT, verbose bool) error {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if verbose {
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, body, "", "  ")
		if err != nil {
			fmt.Println(string(body))
		} else {
			fmt.Println(prettyJSON.String())
		}
	}

	err = json.Unmarshal(body, &model)
	if err != nil {
		return err
	}

	return nil
}

// BindStructArrayWithCursor Template for fast usage: <mongo.Cursor> -> <[]struct>
func BindStructArrayWithCursor[OUT interface{}](ctx context.Context, cursor *mongo.Cursor, models *[]OUT) error {
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
		}
	}(cursor, ctx)
	var newModels []OUT
	for cursor.Next(ctx) {
		var model OUT
		err := cursor.Decode(&model)
		if err != nil {
			return err
		}
		newModels = append(newModels, model)
	}
	*models = AlwaysArray(newModels)

	return nil
}

func AlwaysArray[T any](array []T) []T {
	if len(array) == 0 {
		return []T{}
	}
	return array
}

func ArrayToMap[K comparable, V any](array []Identifiable[K, V]) map[K]V {
	result := make(map[K]V)
	for _, obj := range array {
		result[obj.GetKey()] = obj.GetValue()
	}
	return result
}

func Contains[T comparable](array []T, elem T) bool {
	for _, item := range array {
		if item == elem {
			return true
		}
	}
	return false
}

// ArraysEqual Order does not depend
func ArraysEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	sortedA := append([]T{}, a...)
	sortedB := append([]T{}, b...)
	sort.Slice(sortedA, func(i, j int) bool {
		return fmt.Sprint(sortedA[i]) < fmt.Sprint(sortedA[j])
	})
	sort.Slice(sortedB, func(i, j int) bool {
		return fmt.Sprint(sortedB[i]) < fmt.Sprint(sortedB[j])
	})

	return reflect.DeepEqual(sortedA, sortedB)
}
