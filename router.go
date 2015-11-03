package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "html"
    "net/http"
    "strconv"
    "time"
)

type Action int

type Event struct {
    Action      string
    Timestampo  time.Time `json:"timestamp"`
}

type Meal struct {
    EndDate     string  `json:"end_date"`
    Quantity    float32 `json:"quantity"` 
    StartDate   string  `json:"start_date"`
}

type Statistics struct {
    PetId int
    Meals []Meal
}

func NewRouter() (*mux.Router) {
    router := mux.NewRouter().StrictSlash(false)
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)
        
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}

func hello (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func RegisterEvent (w http.ResponseWriter, r *http.Request) {
    var event Event

    enc := json.NewEncoder(w)

    dec := json.NewDecoder(r.Body)
    err := dec.Decode(&event)
    if err != nil {
        return
    }

    enc.Encode(event)
}

func PetStats (w http.ResponseWriter, r *http.Request) {
    var statistics Statistics
    var meal Meal
    
    vars := mux.Vars(r)
    petId := vars["petId"]
    meal.Quantity = 100
    startDate := time.Now()
    meal.StartDate = startDate.Format("02 Jan 2006 15:04:05") 
    delta := 5 * time.Second
    endDate := startDate.Add(delta)
    meal.EndDate = endDate.Format("02 Jan 2006 15:04:05")
    statistics.PetId, _ = strconv.Atoi(petId)
    statistics.Meals = make([]Meal, 0, 1)
    statistics.Meals = append(statistics.Meals, meal)
    
    json.NewEncoder(w).Encode(statistics)
}
