package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func PlaceOrderWithContext (ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		log.Printf("Xu ly don hang thanh cong\n")
		return nil
	case <-ctx.Done():
		log.Printf("Huy xu ly don hang vi li do: %s\n", ctx.Err())
		return ctx.Err()
	}
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5 * time.Second)
	defer cancel()

	err := PlaceOrderWithContext(ctx)
	if err != nil {
		http.Error(w, "xu ly don hang that bai", http.StatusRequestTimeout)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Xu ly don hang thanh cong"))
}

func main() {
	http.HandleFunc("/order", orderHandler)
	log.Println("Server dang chay tai http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
