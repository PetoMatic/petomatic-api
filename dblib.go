package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

var masterDB *sql.DB

func RegisterEventDB(event Event) error {
    var err error
    var str *string
    err = masterDB.QueryRow(
        `SELECT create_event($1, $2, $3)`,
        event.Event,
        event.PetId,
        event.Weight,
    ).Scan(
        &str,)

    return err
}

func AddPet(pet Pet) error {
    var err error
    res, err := masterDB.Exec("INSERT INTO pets VALUES (nextval('pets_pet_id_seq'), $1, $2, $3)", pet.Breed, pet.Name, pet.DispenserId)
    if res != nil {
        return err
    }
    return nil
}

func GetConfig() ([]Pet, error) {
    var err error
    pets := make([]Pet, 0)

    rows, err := masterDB.Query("SELECT * FROM pets")
    defer rows.Close()
    if err != nil {
        return pets, err
    }

    for rows.Next() {
        var pet Pet
        rows.Scan(&pet.PetId, &pet.Breed, &pet.Name, &pet.DispenserId, &pet.DongleId, &pet.Status)
        pets = append(pets, pet)
    }
    return pets, err
}

func DailyStats(PetId int) (Statistics, error) {
    var err error
    var statistics Statistics
    meals := make([]Meal, 0)
    rows, err := masterDB.Query("SELECT (initial_weight - final_weight) AS quantity, lower(valid) AS start_date, upper(valid) AS end_date, EXTRACT(EPOCH FROM (upper(valid) - lower(valid))) AS duration FROM events WHERE pet_id = $1 AND upper(valid) <> 'infinity' and upper(valid) > CURRENT_DATE", PetId)
    defer rows.Close()
    if err != nil {
        return statistics, err
    }

    for rows.Next() {
        var meal Meal
        rows.Scan(&meal.Quantity, &meal.StartDate, &meal.EndDate, &meal.Duration)
        meals = append(meals, meal)
    }

    statistics.Meals = meals
    statistics.PetId = PetId

    return statistics, err
}

func InitDBConn(Host, Port, Name, User string) error {
    constructParams := func(Host, Port, Name, User string) string {
        if User == "" {
            return fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable", Host, Port, Name)
        } else {
            return fmt.Sprintf("host=%s port=%s dbname=%s user=%s sslmode=disable", Host, Port, Name, User)
        }
    }
    masterParams := constructParams(Host, Port, Name, User)
    fmt.Printf("Master database parameters: %s\n", masterParams)
    var err error
    masterDB, err = sql.Open("postgres", masterParams)
    return err
}
