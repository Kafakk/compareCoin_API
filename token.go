package main

import (
  "fmt"
  "github.com/stellar/go/clients/horizon"
  "log"
  "github.com/gorilla/mux"
  "encoding/json"
  "net/http"
  "github.com/stellar/go/build"

)

func sendToken(w http.ResponseWriter, r *http.Request){

  params := mux.Vars(r)
  if _, err := horizon.DefaultTestNetClient.LoadAccount(params["destination_account"]); err != nil {
        panic(err)
    }

    // passphrase := network.TestNetworkPassphrase

    tx, err := build.Transaction(
        build.TestNetwork,
        build.SourceAccount{params["source_account"]},
        build.AutoSequence{horizon.DefaultTestNetClient},
        build.Payment(
            build.Destination{params["destination_account"]},
            build.NativeAmount{params["amount"]},
        ),
    )

    if err != nil {
        panic(err)
    }

    txe, err := tx.Sign(params["source_account"])
    if err != nil {
        panic(err)
    }

    txeB64, err := txe.Base64()
    if err != nil {
        panic(err)
    }

    // And finally, send it off to Stellar!
    resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
    if err != nil {
        panic(err)
    }

    fmt.Println("Successful Transaction:")
    fmt.Println("Send ", params["amount"], "lumens to ", params["destination_account"])
    fmt.Println("Ledger:", resp.Ledger)
    fmt.Println("Hash:", resp.Hash)
    json.NewEncoder(w).Encode("Successful Transaction")

}

func main(){

  // Account 1
  // SourceAccount1 := "GA2E74U3ABUX7VRASW2QAZDX7BP4MNF7RY7JKRVCZD7DBP7ZTOPX2PCX"
  // SeedAccount1 := "SAOCTAK6BPGRXYLLQGYU2PZI4GU5JJATEI6K7T6D5QNJJZET5SIOYQNL"
  // Account2
  // SourceAccount2 := "GAR4RAQYNXYH3YFYXEHTYBOF3DGV6C5B77FKE775VVGSK7DITDAFJIZX"
  // SeedAccount2 := "SDG46YHEAL7RJBLZ4QXU3AELGJXMVYKGXEBVEKCHAOBWS45OGDQPNSHT"

  router := mux.NewRouter()
  fmt.Println("API start")
  router.HandleFunc("/token/{source_account}/{destination_account}/{amount}",sendToken).Methods("GET")
  log.Fatal(http.ListenAndServe(":8080", router))

}
