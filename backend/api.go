package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/sajari/regression"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var serverhost string // localhost:8000
var nodes []Microservice
var data string

type Microservice struct {
	// Node struct
	remotehost string
	data       string
	input      BodyInput
	bussy      bool
}

type BodyInput struct {
	CO      float64
	H2S     float64
	NO2     float64
	O3      float64
	PM10    float64
	PM2     float64
	SO2     float64
	Ruido   float64
	UV      float64
	Humedad float64
	Presion float64
}

type Temperature struct {
	CO      float64
	H2S     float64
	NO2     float64
	O3      float64
	PM10    float64
	PM2     float64
	SO2     float64
	Ruido   float64
	UV      float64
	Humedad float64
	Presion float64
}

func generateCsvs() {
	resp, err := http.Get("https://raw.githubusercontent.com/MarcoMnrq/tf-concurrent-programming/main/backend/temperature.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	// we create a new csv reader specifying
	// the number of columns it has
	temperatureData := csv.NewReader(resp.Body)
	temperatureData.FieldsPerRecord = 20
	// we read all the records
	records, err := temperatureData.ReadAll()
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
		"training.csv": trainingSet,
		"testing.csv":  testingSet,
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

func train() {
	// we open the csv file from the disk
	f, err := os.Open("training.csv")
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

func test() {
	// we open the csv file from the disk
	f, err := os.Open("testing.csv")
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

func getPrediction(w http.ResponseWriter, r *http.Request) {
	location := Temperature{}
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error")
	}
	err = json.Unmarshal(jsn, &location)
	if err != nil {
		log.Fatal("Error")
	}
	generateCsvs()
	train()
	test()
	//predict(location.CO, location.H2S, location.NO2, location.O3, location.PM10, location.PM2, location.SO2, location.Ruido, location.UV, location.Humedad, location.Presion)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("content-type", "application/json")

	json.NewEncoder(w).Encode(predict(location.CO, location.H2S, location.NO2, location.O3, location.PM10, location.PM2, location.SO2, location.Ruido, location.UV, location.Humedad, location.Presion))
	log.Printf("Received: %v\n", location)
	fmt.Println("[Servidor] Petición hacia el endpoint recibida")
}

func main() {
	// Run API server
	fmt.Print("[Servidor] Puerto TCP para el API Server: ")
	serverhost = fmt.Sprintf("localhost:%s", strings.TrimSpace(strconv.Itoa(8000)))

	// Specify server nodes
	fmt.Println("[Servidor] Ingrese detalles de los 3 nodos...")
	for i := 0; i < 3; i++ {
		fmt.Print("[Nodo ", i+1, "]")
		fmt.Print(" Puerto TCP de este nodo: ")
		port := strings.TrimSpace(strconv.Itoa(8000 + i))
		remotehost := fmt.Sprintf("localhost:%s", port)
		nodes = append(nodes, Microservice{remotehost: remotehost, data: "", bussy: false})
	}
	fmt.Println("[Servidor] Iniciando API Server...")
	//fmt.Println(nodes)
	// Start server and expose 4000
	go initServer()
	http.HandleFunc("/getPrediction", getPrediction)
	http.ListenAndServe(":4000", nil)
}

func initServer() {
	// Initialize server on specified port
	ln, _ := net.Listen("tcp", serverhost) // localhost:8000
	fmt.Println("[Servidor] Escuchando TCP en:", serverhost)
	fmt.Println("[Servidor] ¡Listo! Endpoint: http://localhost:4000/getPrediction")

	defer ln.Close()
	for {
		con, _ := ln.Accept()
		go handleRequest(con)
	}
}

func handleRequest(con net.Conn) {
	// Handle TCP requests
	defer con.Close()
	for {
		msg, _ := bufio.NewReader(con).ReadString('\n')
		msg = strings.TrimSpace(msg)
		data = "Hello world"
		communicateToNodes()
	}
}

func communicateToNodes() {
	for _, node := range nodes {
		if node.bussy == false {
			node.bussy = true
			conn, _ := net.Dial("tcp", node.remotehost)
			defer conn.Close()
			encoder := json.NewEncoder(conn)
			encoder.Encode(333)
			node.bussy = false
		}
	}
}
