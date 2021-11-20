package main

import (
	"context"
	"fmt"
	"github.com/Ksloveyuan/channelx"
	"github.com/pkg/errors"
	"time"
)

func main() {
	eventBus()
}

func eventBus() {
	logger := channelx.NewConsoleLogger()
	eventBus := channelx.NewEventBus(logger, 4, 4, 2, time.Second, 5*time.Second)
	eventBus.Start()

	handler := NewExampleHandler(logger)
	eventBus.Subscribe(ExampleEventID, handler)
	eventBus.Publish(NewExampleEvent())

	eventBus.Stop()
}

func parallelRunner() {
	//worker := func(ctx context.Context, input interface{}) (interface{}, error) {
	//	num := input.(int)
	//	return num+1, nil
	//}
	//
	//inputs := []interface{}{1,2,3,4,5}
	//outputs, err := channelx.RunInParallel(context.Background(), inputs, worker, 4)
	//fmt.Println(outputs, err)

	var pr = channelx.NewParallelRunner(5)
	var outputs, err = pr.Run(context.TODO(), []interface{}{1, 2, 3, 4, 5}, func(ctx context.Context, i interface{}) (interface{}, error) {
		fmt.Println(i)
		return i.(int) + 1, nil
	})

	fmt.Println(outputs, err)
}

func promise() {
	promise := channelx.NewPromise(func() (res interface{}, err error) {
		// do work asynchronously here
		return 1, errors.New("234")
	}).Then(func(input interface{}) (interface{}, error) {
		fmt.Println("=============t", input)
		return 2, errors.New("123")
	}, func(err error) interface{} {
		fmt.Println("=============r", err)
		// here is the error handler, which aslo runs asynchronously
		return err
	})

	// await: wait until it completes.
	res, err := promise.Done()
	fmt.Println(res, err)
}

func actor() {
	actor := channelx.NewActor(channelx.SetActorBuffer(2))
	defer actor.Close()

	var ps = make([]*channelx.Promise, 0)

	for i := 0; i < 5; i++ {
		// do some work asynchroniously.
		call := actor.Do(func() (interface{}, error) {
			time.Sleep(5 * time.Second)
			return 0, nil
		})
		fmt.Println("====a")
		ps = append(ps, call)
	}

	// can to some other synchroniouse work here
	// ......

	// wait for the call completes.
	for _, call := range ps {
		res, err := call.Done()
		fmt.Println(res, err)
	}
}

func aggregatorf() {
	var cnt = 0

	batchProcess := func(items []interface{}) error {
		time.Sleep(2 * time.Second)
		fmt.Println(items)
		cnt += len(items)
		return nil
	}

	aggregator := channelx.NewAggregator(batchProcess, func(option channelx.AggregatorOption) channelx.AggregatorOption {
		//option.ChannelBufferSize = 10
		return option
	})

	aggregator.Start()

	var rs = map[bool]int{
		true:  0,
		false: 0,
	}

	for i := 0; i < 100; i++ {
		aggregator.Enqueue(i)
		//fmt.Println("===== te", aggregator.TryEnqueue(i))
		//b:= aggregator.TryEnqueue(i)
		//rs[b]++
	}

	//time.Sleep(10*time.Second)

	aggregator.SafeStop()

	fmt.Println(rs, cnt)
}

// ====================================

const ExampleEventID channelx.EventID = 1

type ExampleEvent struct {
	id channelx.EventID
}

func (evt ExampleEvent) ID() channelx.EventID {
	return evt.id
}

type ExampleHandler struct {
	logger channelx.Logger
}

func NewExampleEvent() ExampleEvent {
	return ExampleEvent{id: ExampleEventID}
}

func NewExampleHandler(logger channelx.Logger) *ExampleHandler {
	return &ExampleHandler{
		logger: logger,
	}
}

func (h ExampleHandler) Logger() channelx.Logger {
	return h.logger
}

func (h ExampleHandler) CanAutoRetry(err error) bool {
	return false
}

func (h ExampleHandler) OnEvent(ctx context.Context, event channelx.Event) error {
	if event.ID() != ExampleEventID {
		return fmt.Errorf("subscribe wrong event(%d)", event.ID())
	}

	_, ok := event.(ExampleEvent)
	if !ok {
		return fmt.Errorf("failed to convert received event to ExampleEvent")
	}

	h.Logger().Infof("event handled")

	return nil
}
