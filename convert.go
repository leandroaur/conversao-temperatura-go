package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(nil)

	http.HandleFunc("/", temperatureHandler)
	http.ListenAndServe(":8080", nil)
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		form := `
			<html>
			  <head>
			    <style>
			      h1 {
			        font-size: 36px;
			        text-align: center;
			        color: blue;
			      }
			    </style>
			  </head>
			  <body>
			    <h1>Temperature Converter</h1>
			    <form action="/" method="post">
			      <input type="text" name="temperature">
			      <select name="unit">
			        <option value="C">Celsius</option>
			        <option value="F">Fahrenheit</option>
			      </select>
			      <button type="submit">Convert</button>
			    </form>
			  </body>
			</html>
		`
		fmt.Fprintf(w, form)
		return
	}

	temperature, err := strconv.ParseFloat(r.FormValue("temperature"), 64)
	if err != nil {
		http.Error(w, "Invalid temperature", http.StatusBadRequest)
		return
	}

	unit := r.FormValue("unit")
	var result float64
	switch unit {
	case "C":
		result = (temperature*9/5 + 32)
	case "F":
		result = (temperature - 32) * 5 / 9
	default:
		http.Error(w, "Invalid unit", http.StatusBadRequest)
		return
	}

	err = saveToDB(temperature, unit, result)
	if err != nil {
		http.Error(w, "Failed to save to database", http.StatusInternalServerError)
		return
	}

	displayResult(w, temperature, unit, result)
}

func displayResult(w http.ResponseWriter, temperature float64, unit string, result float64) {
	var resultUnit string
	if unit == "C" {
		resultUnit = "Fahrenheit"
	} else {
		resultUnit = "Celsius"
	}
	resultPage := fmt.Sprintf(`
		<html>
		  <head>
		    <style>
		      h1 {
		        font-size: 36px;
		        text-align: center;
		        color: blue;
		      }
		    </style>
		  </head>
		  <body>
		    <h1>Result</h1>
		    <p>%f %s is equivalent to %f %s</p>
		  </body>
		</html>
	`, temperature, unit, result, resultUnit)
	fmt.Fprintf(w, resultPage)
}

func saveToDB(temperature float64, fromUnit string, result float64) error {
	collection := client.Database("test").Collection("temperatures")
	_, err := collection.InsertOne(nil, map[string]interface{}{
		"temperature": temperature,
		"from_unit":   fromUnit,
		"to_unit":     "F", /* assume always converted to Fahrenheit */
		"result":      result,
	})
	return err
}

