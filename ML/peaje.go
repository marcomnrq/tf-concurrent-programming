package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/sajari/regression"
)

func main() {
	// we open the csv file from the disk
	f, err := os.Open("datasets/peajev4.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// we create a new csv reader specifying
	// the number of columns it has
	salesData := csv.NewReader(f)
	salesData.FieldsPerRecord = 6

	// we read all the records
	records, err := salesData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// In this case we are going to try and model our house price (y)
	// by the grade feature.
	var r regression.Regression
	r.SetObserved("total_veh_exon")
	r.SetVar(0, "anio")
	r.SetVar(1, "mes")
	r.SetVar(2, "sentido")
	r.SetVar(3, "total_veh_pagan")
	r.SetVar(4, "total_veh_eva")


	// Loop of records in the CSV, adding the training data to the regressionvalue.
	for i, record := range records {
		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the house price, "y".
		total_veh_exon, err := strconv.ParseFloat(records[i][4], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the house price, "y".
		anio, err := strconv.ParseFloat(records[i][0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the grade value.
		mes, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		sentido, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		total_veh_pagan, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		total_veh_eva, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Add these points to the regression value.
		r.Train(regression.DataPoint(total_veh_exon, []float64{anio, mes, sentido, total_veh_pagan, total_veh_eva}))
	}

	// Train/fit the regression model.
	r.Run()
	// Output the trained model parameters.
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)
}