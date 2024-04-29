package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

type TicketPair struct {
    Source      string `json:"source"`
    Destination string `json:"destination"`
}

func main() {
    e := echo.New()

    e.POST("/flights", func(c echo.Context) error {
        var tickets [][]string
        if err := c.Bind(&tickets); err != nil {
            return err
        }

        itinerary, err := reconstructItinerary(tickets)
        if err != nil {
            return err
        }

        return c.JSON(http.StatusOK, itinerary)
    })

    e.Logger.Fatal(e.Start(":8080"))
}

func reconstructItinerary(tickets [][]string) ([]string, error) {
    var tempSources []string
    var tempDestinations []string

    for _, ticket := range tickets {
        if len(ticket) != 2 {
            return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid pair format")
        }
        tempSources = append(tempSources, ticket[0])
        tempDestinations = append(tempDestinations, ticket[1])
    }

    var startingPoint string
    for _, source := range tempSources {
        found := false
        for _, dest := range tempDestinations {
            if source == dest {
                found = true
                break
            }
        }
        if !found {
            startingPoint = source
            break
        }
    }

    var itinerary []string
    current := startingPoint
    for i := 0; i < len(tickets); i++ {
        itinerary = append(itinerary, current)
        for j := 0; j < len(tickets); j++ {
            if tickets[j][0] == current {
                current = tickets[j][1]
                break
            }
        }
    }

    itinerary = append(itinerary, current)

    return itinerary, nil
}
