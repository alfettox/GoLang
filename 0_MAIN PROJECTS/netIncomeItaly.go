/*
Author: Giovanni De Franceschi
*/

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"encoding/json"
)

type ExchangeRateResponse struct {
	ConversionRate float64 `json:"conversion_rate"`
}

func readAPIKey(filename string) (string, error) {
	apiKey, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(apiKey), nil
}

func fetchExchangeRate(apiKey, base, target string) (float64, error) {
	baseURL := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/%s/%s", apiKey, base, target)
	response, err := http.Get(baseURL)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	var exchangeRateResponse ExchangeRateResponse
	if err := json.Unmarshal(body, &exchangeRateResponse); err != nil {
		return 0, err
	}

	return exchangeRateResponse.ConversionRate, nil
}

func main() {
	var choice int
	fmt.Println("Choose an option:")
	fmt.Println("1. Enter yearly gross salary in EUR")
	fmt.Println("2. Enter yearly gross salary in CAD for Quebec")
	fmt.Print("Choice: ")
	fmt.Scan(&choice)

	if choice != 1 && choice != 2 {
		fmt.Println("Invalid choice.")
		return
	}

	var grossIncome float64
	if choice == 1 {
		fmt.Println("Enter the yearly gross income in EUR: ")
		fmt.Scan(&grossIncome)
	} else {
		fmt.Println("Enter the yearly gross income in CAD for Quebec: ")
		fmt.Scan(&grossIncome)
	}

	var monthlyPayments int
	fmt.Println("Enter the number of monthly payments (12 or 13): ")
	fmt.Scan(&monthlyPayments)

	if monthlyPayments != 12 && monthlyPayments != 13 {
		fmt.Println("Invalid number of monthly payments.")
		return
	}

	apiKey, err := readAPIKey("API_SECRET")
	if err != nil {
		fmt.Println("Error reading API key:", err)
		return
	}

	socialSecurityContributions := grossIncome * 0.23
	incomeTax := math.Min(grossIncome*0.43, 15000)
	personalExemptionFloat := float64(12912)

	netIncomeYearly := grossIncome - socialSecurityContributions - incomeTax + personalExemptionFloat
	netIncomeMonthly := netIncomeYearly / float64(monthlyPayments)

	fmt.Printf("Yearly gross income: %.2f ", grossIncome)
	if choice == 1 {
		fmt.Println("EUR")
	} else {
		fmt.Println("CAD (Quebec)")
	}

	fmt.Printf("Yearly net income: %.2f ", netIncomeYearly)
	if choice == 1 {
		fmt.Println("EUR")
	} else {
		fmt.Println("CAD")
	}

	fmt.Printf("Monthly net income (%d payments): %.2f ", monthlyPayments, netIncomeMonthly)
	if choice == 1 {
		fmt.Println("EUR")
	} else {
		fmt.Println("CAD")
	}

	fmt.Printf("Bi-weekly net income (%d payments): %.2f ", monthlyPayments*2, netIncomeMonthly/2)
	if choice == 1 {
		fmt.Println("EUR")
	} else {
		fmt.Println("CAD")
	}

	if choice == 2 {
		exchangeRateCADtoEUR, err := fetchExchangeRate(apiKey, "CAD", "EUR")
		if err != nil {
			fmt.Println("Error fetching exchange rate:", err)
			return
		}

		equivalentNetIncomeInEUR := netIncomeMonthly / exchangeRateCADtoEUR
		equivalentNetIncomeYearlyItaly := equivalentNetIncomeInEUR * 12
		equivalentGrossIncomeItaly := equivalentNetIncomeYearlyItaly + socialSecurityContributions + incomeTax - personalExemptionFloat
		fmt.Printf("Equivalent gross income in Italy (EUR): %.2f EUR\n", equivalentGrossIncomeItaly)
	} else if choice == 1 {
		exchangeRateEURtoCAD, err := fetchExchangeRate(apiKey, "EUR", "CAD")
		if err != nil {
			fmt.Println("Error fetching exchange rate:", err)
			return
		}

		equivalentNetIncomeInCAD := netIncomeMonthly * exchangeRateEURtoCAD
		equivalentNetIncomeYearlyQuebec := equivalentNetIncomeInCAD * 12
		equivalentGrossIncomeQuebec := equivalentNetIncomeYearlyQuebec + socialSecurityContributions + incomeTax - personalExemptionFloat
		fmt.Printf("Equivalent gross income in Quebec (CAD): %.2f CAD\n", equivalentGrossIncomeQuebec)
	}
}
