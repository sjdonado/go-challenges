package bankapi

import (
	"encoding/json"
	"example/go_challenges/bank/bankcore"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type JSONAccount struct {
	*bankcore.Account
}

func (c *JSONAccount) Statement() string {
	json, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}

	return string(json)
}

var accounts = map[float64]*JSONAccount{}

func seed() {
	accounts[1001] = &JSONAccount{
		Account: &bankcore.Account{
			Customer: bankcore.Customer{
				Name:    "John",
				Address: "Los Angeles, California",
				Phone:   "(213) 555 0147",
			},
			Number: 1001,
		},
	}
	accounts[1002] = &JSONAccount{
		Account: &bankcore.Account{
			Customer: bankcore.Customer{
				Name:    "Spooky",
				Address: "San Francisco, California",
				Phone:   "(213) 555 0148",
			},
			Number: 1002,
		},
	}
}

func Run() {
	seed()
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			fmt.Fprintf(w, account.Statement())
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	senderqs := req.URL.Query().Get("sender")
	recipientqs := req.URL.Query().Get("recipient")
	amountqs := req.URL.Query().Get("amount")

	if senderqs == "" || recipientqs == "" {
		fmt.Fprintf(w, "Account numbers are missing!")
		return
	}

	if sender, err := strconv.ParseFloat(senderqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid sender account number!")
	} else if recipient, err := strconv.ParseFloat(recipientqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid recipient account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		sender_account, ok := accounts[sender]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", sender_account)
			return
		}
		recipient_account, ok := accounts[recipient]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", recipient_account)
			return
		}
		err := sender_account.Transfer(amount, recipient_account.Account)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		} else {
			fmt.Fprintf(w, sender_account.Statement())
		}
	}
}
