package learning

import (
	"context"
	"fmt"
	"log"

	"time"
)

func MainCtx() {
	start := time.Now()
	ctx := context.Background()
	userId := 10
	val, err := fetchUserData(ctx, userId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result: ", val)
	fmt.Println("took: ", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	respch := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took too long")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 150)
	return 666, nil
}
