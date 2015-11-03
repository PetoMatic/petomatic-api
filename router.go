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

type Pet struct {
    PetId       int     `json:"pet_id"`
    Breed       string  `json:"breed"`
    Name        string  `json:"name"`
    DispenserId int     `json:"dispenser_id"`
}

type Event struct {
    Event       string  `json:"event"`
    Id          int     `json:"id,omitempty"`
    PetId       int     `json:"pet_id"`
    Timestamp   int     `json:"timestamp,omitempty"`
    Valid       string  `json:"valid,omitempty"`
    Weight      int     `json:"weight"`
}

type Meal struct {
    EndDate     time.Time  `json:"end_date"`
    Duration    float32 `json:"duration"`
    Quantity    int      `json:"quantity"` 
    StartDate   time.Time  `json:"start_date"`
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
        fmt.Println("Error: ", err)
        enc.Encode(err)
        return
    }

    err = RegisterEventDB(event)
    if err != nil {
        fmt.Println("Error: ", err)
        enc.Encode(err)
        return
    }

    enc.Encode(event)
    fmt.Println(event)
}

func PetDailyStats (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    petId, _ := strconv.Atoi(vars["petId"])

    stats, err := DailyStats(petId) 
    if err != nil {
        fmt.Println("Error: ", err)
        json.NewEncoder(w).Encode(err)
        return
    }
    json.NewEncoder(w).Encode(stats)
}

func Config (w http.ResponseWriter, r *http.Request) {
    pets, err := GetConfig()
    if err != nil {
        fmt.Println("Error: ", err)
        json.NewEncoder(w).Encode(err)
        return
    }
    json.NewEncoder(w).Encode(pets)
}

func LovePet (w http.ResponseWriter, r *http.Request) {
    var pet Pet

    enc := json.NewEncoder(w)
    dec := json.NewDecoder(r.Body)
    err := dec.Decode(&pet)
    if err != nil {
        fmt.Println("Error: ", err)
        enc.Encode(err)
        return
    }

    err = AddPet(pet)
    if err != nil {
        fmt.Println("Error: ", err)
        enc.Encode(err)
        return
    }
    enc.Encode(pet)
    fmt.Println(pet)
}
