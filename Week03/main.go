package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {
	g := new(errgroup.Group)
	sigs := make(chan os.Signal,1)
	serverErr := make(chan error,10)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_,err := fmt.Fprint(w, "Hello world!")
		if err != nil {
			serverErr <- err
		}
	})
	server := &http.Server{
		Addr:         ":8087",
		Handler:      mux,
	}
	g.Go(func() error {
		signal.Notify(sigs, syscall.SIGINT|syscall.SIGTERM|syscall.SIGKILL)
		<-sigs
		return server.Shutdown(context.TODO())
	})

	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case tmp := <-sigs:
			fmt.Println(tmp)
			return errors.New("信号")
		case err:=<-serverErr:
			return err
		}

	})

	err := g.Wait()
	if err != nil {
		fmt.Println(err)
		ctx,cancel := context.WithTimeout(context.Background(),1*time.Second)
		_ = server.Shutdown(ctx)
		cancel()
	}
}




