package cycapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func Error(ctx context.Context, message string, err error) {
	values := getValues(ctx)
	if values == nil {
		values = make(map[string]string)
	}

	values["ERROR"] = message

	values["CAUSE"] = fmt.Sprintf("%v", err)
	manageLog(values, err)
}

// Add adds a new value to the existing context to be tracked in logs and returns a new context
func Add(ctx context.Context, key, value string) context.Context {
	keys := copyLogValues(getKeys(ctx))

	keys[key] = ""

	ctx = context.WithValue(ctx, "keys", keys)
	ctx = context.WithValue(ctx, key, value)

	return ctx
}

func manageLog(values map[string]string, err error) {
	values["timestamp"] = time.Now().Format(time.RFC3339)

	errLog, err := json.Marshal(values)
	if err != nil {
		log.Println("Could not create JSON from values")
		return
	}

	fmt.Println(string(errLog))
}

func copyLogValues(old map[string]string) map[string]string {
	newMap := make(map[string]string)

	for k, v := range old {
		newMap[k] = v
	}

	return newMap
}

func getValues(ctx context.Context) map[string]string {
	values := make(map[string]string)
	keys := getKeys(ctx)

	for k := range keys {
		if v, ok := ctx.Value(k).(string); ok {
			values[string(k)] = v
		}
	}

	return values
}

func getKeys(ctx context.Context) map[string]string {
	keys := make(map[string]string)

	if k, ok := ctx.Value("keys").(map[string]string); ok {
		keys = k
	}

	return keys
}