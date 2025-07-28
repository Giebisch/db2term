package data

import "time"

func NewRequest(abfahrt string, ankunft string, timeOffset int) *TripRequest {
	return &TripRequest{
		AbfahrtsHalt:     abfahrt,
		AnfrageZeitpunkt: get_current_time_with_offset(timeOffset),
		AnkunftsHalt:     ankunft,
		AnkunftSuche:     "ABFAHRT",
		Klasse:           "KLASSE_2",
		Produktgattungen: []string{
			"ICE",
			"EC_IC",
			"IR",
			"REGIONAL",
			"SBAHN",
			"BUS",
			"SCHIFF",
			"UBAHN",
			"TRAM",
			"ANRUFPFLICHTIG",
		},
		Reisende: []Reisende{
			{
				Typ:            "ERWACHSENER",
				Ermaessigungen: []Ermaessigungen{{Art: "KEINE_ERMAESSIGUNG", Klasse: "KLASSENLOS"}},
				Alter:          []any{},
				Anzahl:         1,
			},
		},
		SchnelleVerbindungen:              true,
		SitzplatzOnly:                     false,
		BikeCarriage:                      false,
		ReservierungsKontingenteVorhanden: false,
		NurDeutschlandTicketVerbindungen:  false,
		DeutschlandTicketVorhanden:        false,
	}
}

func get_current_time_with_offset(timeOffset int) string {
	target_time := time.Now().Add(time.Duration(timeOffset) * time.Hour)
	return target_time.Format("2006-01-02T15:04:05")
}
