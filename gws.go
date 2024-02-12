package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"embed"
	"io"
	//"html/template"
)
func main() {
	http.HandleFunc("/json", jsonTest)
	http.HandleFunc("/", index)
	http.HandleFunc("/embedGWS", embedGWS)
	http.HandleFunc("/html", html)
	http.HandleFunc("/testAPI", testAPI)
	http.HandleFunc("/delete", deleteStubbed)
	http.HandleFunc("/create", createStubbed)
	http.HandleFunc("/update", updateStubbed)
	http.HandleFunc("/help", help)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World - GWS")
}
func jsonTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{"message": "Hello World - GWS"}`)
}


// Test JSON file
//go:embed gws.json
var jsonFiles embed.FS


type Number struct {
	Num int `json:"number"`
}
//returns GET request for syllabus.json file
func embedGWS(w http.ResponseWriter, r *http.Request){
	// Read the embedded JSON file
	gws, err := jsonFiles.ReadFile("gws.json")
	if err != nil {
		http.Error(w, "Error reading JSON file", http.StatusInternalServerError)
		return
	}

	// Unmarshal the JSON data into a slice of Item
	var numbers []Number
	err = json.Unmarshal(gws, &numbers)
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusInternalServerError)
		return
	}

	apiURL := "http://localhost:8080/testAPI"
	response, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Error fetching data from external API", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Read the response body
	apiData, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Error reading response from external API", http.StatusInternalServerError)
		return
	}
	fmt.Println("API Response:", string(apiData))
	// Unmarshal the API data into a slice of Item
	var apiItems []Number
	err = json.Unmarshal(apiData, &apiItems)
	if err != nil {
		http.Error(w, "Error decoding data from external API", http.StatusInternalServerError)
		return
	}

	// Merge the embedded data with the data from the external API
	numbers = append(numbers, apiItems...)

	// Marshal the merged items back to JSON
	mergedResponse, err := json.Marshal(numbers)
	if err != nil {
		http.Error(w, "Error encoding merged JSON data", http.StatusInternalServerError)
		return
	}
	
	// Set the content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(mergedResponse)

}




func html(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Hello World — GWS")
}

//go:embed syllabus.json
var content embed.FS

// Item represents a simple struct for the JSON data
type Syl struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Class string `json:"class"`
}


func testAPI(w http.ResponseWriter, r *http.Request) {
	data, err := content.ReadFile("syllabus.json")
	if err != nil {
		http.Error(w, "Error reading JSON file", http.StatusInternalServerError)
		return
	}

	// Unmarshal the JSON data into a slice of Item
	var syls []Syl
	err = json.Unmarshal(data, &syls)
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusInternalServerError)
		return
	}

	// Marshal the items back to JSON
	response, err := json.Marshal(syls)
	if err != nil {
		http.Error(w, "Error encoding JSON data", http.StatusInternalServerError)
		return
	}

	// Set the content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(response)
}

func deleteStubbed(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleted - stubbed")
	fmt.Fprintf(w, "deleted - stubbed")
}

func createStubbed(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create - stubbed")
	fmt.Fprintf(w, "create - stubbed")
}

func updateStubbed(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update - stubbed")
	fmt.Fprintf(w, "update - stubbed")
}
func help(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Below is a list of all APIs in this web server:")
	fmt.Println("/json ; json test")
	fmt.Println("/ ; default page")
	fmt.Println("/embedGWS ; embed json file test/get request from test API")
	fmt.Println("/html ; html page")
	fmt.Println("/testAPI ; test syllabus API")
	fmt.Println("/delete ; stubbed delete request")
	fmt.Println("/create ; stubbed create request")
	fmt.Println("/update ; stubbed update request")

	fmt.Fprintf(w, "Below is a list of all APIs in this web server: \n")
	fmt.Fprintf(w, "/json ; json test \n")
	fmt.Fprintf(w, "/ ; default page \n")
	fmt.Fprintf(w, "/embedGWS ; embed json file test/get request from test API \n")
	fmt.Fprintf(w, "/html ; html page \n")
	fmt.Fprintf(w, "/testAPI ; test syllabus API \n" )
	fmt.Fprintf(w,"/delete ; stubbed delete request \n")
	fmt.Fprintf(w, "/create ; stubbed create request \n")
	fmt.Fprintf(w, "/update ; stubbed update request \n")
}

/*export const _INSTRUCTOR_ERIC_POGUE = { 
	name:'Eric Pogue', 
	officeHours:'Tuesdays 1-3 PM and Thursdays 10-11 AM CT by appointment',
	office:'AS-124-A', 
	appointmentRequests:'Appointments can be requested via email',
	lewisPhone:'(815) 836-5015',
	lewisEmail:'epogue@lewisu.edu' 
} */

//embed json response in go
//embed binary file in go
//returning json response that was embedded