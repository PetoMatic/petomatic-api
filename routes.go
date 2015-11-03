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
        "PetStats",
        "GET",
        "/stats/{petId}",
        PetStats,
    },
    Route{
        "CreateEvent",
        "POST",
        "/event",
        RegisterEvent,
    },
}