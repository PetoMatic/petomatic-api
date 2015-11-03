package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Main",
        "GET",
        "/",
        hello,
    },
    Route{
        "Config",
        "GET",
        "/config",
        Config,
    },
    Route{
        "PetStats",
        "GET",
        "/stats/{petId}/daily",
        PetDailyStats,
    },
    Route{
        "CreateEvent",
        "POST",
        "/event",
        RegisterEvent,
    },
    Route{
        "AddPet",
        "POST",
        "/love/pet",
        LovePet,
    },
}
