package main

import (
	"fmt"
  "encoding/json"
	"log"
  "net/http"
	cmc "github.com/miguelmota/go-coinmarketcap"
  "github.com/gorilla/mux"

)


func coinCompare(w http.ResponseWriter, r *http.Request){

  params := mux.Vars(r)

  ticker1, err := cmc.Ticker(&cmc.TickerOptions{
    Symbol: params["ticker_symbol_1"],
	  Convert: "USD",
	})
  if err != nil {
		log.Fatal(err)
	}

  ticker2, err := cmc.Ticker(&cmc.TickerOptions{
    Symbol: params["ticker_symbol_2"],
	  Convert: "USD",
	})
  if err != nil {
		log.Fatal(err)
	}


  if (ticker1.Quotes["USD"].PercentChange1H > ticker2.Quotes["USD"].PercentChange1H){
		fmt.Println("Ticker ID: "+ ticker1.Symbol + " has more change.")
    json.NewEncoder(w).Encode(ticker1.Symbol)
	} else if (ticker1.Quotes["USD"].PercentChange1H < ticker2.Quotes["USD"].PercentChange1H){
		fmt.Println("Ticker ID: "+ ticker2.Symbol + " has more change.")
    json.NewEncoder(w).Encode(ticker2.Symbol)
	} else{
		// fmt.Println("They have the same percent change 1h")
    json.NewEncoder(w).Encode("They have the same percent change 1h")
  }

}

func main(){
  router := mux.NewRouter()
  fmt.Println("API start")
  router.HandleFunc("/ticker/{ticker_symbol_1}/{ticker_symbol_2}",coinCompare).Methods("GET")
  log.Fatal(http.ListenAndServe(":8080", router))
}
