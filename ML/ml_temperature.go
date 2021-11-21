package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func generateCsvs(){
	f, err := os.Open("datasets/temperature.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	// we create a new csv reader specifying
	// the number of columns it has
	salesData := csv.NewReader(f)
	salesData.FieldsPerRecord = 20
	// we read all the records
	records, err := salesData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// save the header
	header := records[0]
	// we have to shuffle the dataset before splitting
	// to avoid having ordered data
	// if the data is ordered, the data in the train set
	// and the one in the test set, can have different
	// behavior
	shuffled := make([][]string, len(records)-1)
	perm := rand.Perm(len(records) - 1)
	for i, v := range perm {
		shuffled[v] = records[i+1]
	}
	// split the training set
	trainingIdx := (len(shuffled)) * 4 / 5
	trainingSet := shuffled[1 : trainingIdx+1]
	// split the testing set
	testingSet := shuffled[trainingIdx+1:]
	// we write the splitted sets in separate files
	sets := map[string][][]string{
		"datasets/training.csv": trainingSet,
		"datasets/testing.csv":  testingSet,
	}
	for fn, dataset := range sets {
		f, err := os.Create(fn)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		out := csv.NewWriter(f)
		if err := out.Write(header); err != nil {
			log.Fatal(err)
		}
		if err := out.WriteAll(dataset); err != nil {
			log.Fatal(err)
		}
		out.Flush()
	}
}

func train(){
	// we open the csv file from the disk
	f, err := os.Open("datasets/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// we create a new csv reader specifying
	// the number of columns it has
	temperatureData := csv.NewReader(f)
	temperatureData.FieldsPerRecord = 20

	// we read all the records
	records, err := temperatureData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// In this case we are going to try and model our house price (y)
	// by the grade feature.
	var r regression.Regression
	r.SetObserved("Temperatura")
	r.SetVar(0, "CO")
	r.SetVar(1, "H2S")
	r.SetVar(2, "NO2")
	r.SetVar(3, "O3")
	r.SetVar(4, "PM10")
	r.SetVar(5, "PM2")
	r.SetVar(6, "SO2")
	r.SetVar(7, "Ruido")
	r.SetVar(8, "UV")
	r.SetVar(9, "Humedad")
	r.SetVar(10, "Presion")

	// Loop of records in the CSV, adding the training data to the regressionvalue.
	for i, record := range records {
		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the house price, "y".
		Temperatura, err := strconv.ParseFloat(records[i][19], 64)
		if err != nil {
			log.Fatal(err)
		}

		CO, err := strconv.ParseFloat(record[6], 64)
		if err != nil {
			log.Fatal(err)
		}

		H2S, err := strconv.ParseFloat(records[i][7], 64)
		if err != nil {
			log.Fatal(err)
		}

		NO2, err := strconv.ParseFloat(records[i][8], 64)
		if err != nil {
			log.Fatal(err)
		}

		O3, err := strconv.ParseFloat(records[i][9], 64)
		if err != nil {
			log.Fatal(err)
		}

		PM10, err := strconv.ParseFloat(records[i][10], 64)
		if err != nil {
			log.Fatal(err)
		}

		PM2, err := strconv.ParseFloat(records[i][11], 64)
		if err != nil {
			log.Fatal(err)
		}

		SO2, err := strconv.ParseFloat(records[i][12], 64)
		if err != nil {
			log.Fatal(err)
		}

		Ruido, err := strconv.ParseFloat(records[i][13], 64)
		if err != nil {
			log.Fatal(err)
		}

		UV, err := strconv.ParseFloat(records[i][14], 64)
		if err != nil {
			log.Fatal(err)
		}

		Humedad, err := strconv.ParseFloat(records[i][15], 64)
		if err != nil {
			log.Fatal(err)
		}

		Presion, err := strconv.ParseFloat(records[i][18], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Add these points to the regression value.
		r.Train(regression.DataPoint(Temperatura, []float64{CO, H2S, NO2, O3, PM10, PM2, SO2, Ruido, UV, Humedad, Presion}))
	}

	// Train/fit the regression model.
	r.Run()
	// Output the trained model parameters.
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)
}

func test(){
	// we open the csv file from the disk
	f, err := os.Open("datasets/testing.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// we create a new csv reader specifying
	// the number of columns it has
	salesData := csv.NewReader(f)
	salesData.FieldsPerRecord = 20
	// we read all the records
	records, err := salesData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// by slicing the records we skip the header
	records = records[1:]
	// Loop over the test data predicting y
	observed := make([]float64, len(records))
	predicted := make([]float64, len(records))
	var sumObserved float64
	for i, record := range records {
		// Parse the house price, "y".
		Temperatura, err := strconv.ParseFloat(records[i][19], 64)
		if err != nil {
			log.Fatal(err)
		}
		observed[i] = Temperatura
		sumObserved += Temperatura

		// Parse the grade value.
		CO, err := strconv.ParseFloat(records[i][6], 64)
		if err != nil {
			log.Fatal(err)
		}

		H2S, err := strconv.ParseFloat(records[i][7], 64)
		if err != nil {
			log.Fatal(err)
		}

		NO2, err := strconv.ParseFloat(records[i][8], 64)
		if err != nil {
			log.Fatal(err)
		}

		O3, err := strconv.ParseFloat(records[i][9], 64)
		if err != nil {
			log.Fatal(err)
		}

		PM10, err := strconv.ParseFloat(records[i][10], 64)
		if err != nil {
			log.Fatal(err)
		}

		PM2, err := strconv.ParseFloat(records[i][11], 64)
		if err != nil {
			log.Fatal(err)
		}

		SO2, err := strconv.ParseFloat(records[i][12], 64)
		if err != nil {
			log.Fatal(err)
		}

		Ruido, err := strconv.ParseFloat(record[13], 64)
		if err != nil {
			log.Fatal(err)
		}

		UV, err := strconv.ParseFloat(records[i][14], 64)
		if err != nil {
			log.Fatal(err)
		}

		Humedad, err := strconv.ParseFloat(records[i][15], 64)
		if err != nil {
			log.Fatal(err)
		}

		Presion, err := strconv.ParseFloat(records[i][18], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict y with our trained model.
		predicted[i] = predict(CO, H2S, NO2, O3, PM10, PM2, SO2, Ruido, UV, Humedad, Presion)
	}
	mean := sumObserved / float64(len(observed))
	var observedCoefficient, predictedCoefficient float64
	for i := 0; i < len(observed); i++ {
		observedCoefficient += math.Pow((observed[i] - mean), 2)
		predictedCoefficient += math.Pow((predicted[i] - mean), 2)
	}
	rsquared := predictedCoefficient / observedCoefficient
	// Output the R-squared to standard out.
	fmt.Printf("R-squared = %0.2f\n\n", rsquared)
}

func predict(CO float64, H2S float64, NO2 float64,
	O3 float64, PM10 float64, PM2 float64, SO2 float64,
	Ruido float64, UV float64, Humedad float64, Presion float64) float64 {
	return 19.3196 + CO*0.0008 + H2S*-0.1078 + NO2*-0.0013 + O3*0.0590 + PM10*0.0001 + PM2*0.0022 + SO2*1.3090 + Ruido*0.0056 + UV*-0.1535 + Humedad*-0.1362 + Presion*0.0000
}

func main() {
	generateCsvs()
	train()
	test()
}