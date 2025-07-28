package data

import "time"

type Trip struct {
	Abschnitte []Abschnitt
	Dauer      int
}

type Abschnitt struct {
	AbfahrtsOrt       string
	AnkunftsOrt       string
	AbfahrtsZeitpunkt time.Time
	AnkunftsZeitpunkt time.Time
	Art               string
	Text              string
	Dauer             int
}

func (r *TripResponse) Format() []Trip {
	res := []Trip{}

	for _, verbindung := range r.Verbindungen {
		abschnitte := []Abschnitt{}
		totalTime := float32(0)
		for _, abschnitt := range verbindung.VerbindungsAbschnitte {
			abfahrt, err := time.Parse("2006-01-02T15:04:05", abschnitt.AbfahrtsZeitpunkt)
			if err != nil {
				abfahrt = time.Now()
			}
			ankunft, err := time.Parse("2006-01-02T15:04:05", abschnitt.AnkunftsZeitpunkt)
			if err != nil {
				abfahrt = time.Now()
			}
			abschnitte = append(abschnitte, Abschnitt{
				AbfahrtsOrt:       abschnitt.AbfahrtsOrt,
				AnkunftsOrt:       abschnitt.AnkunftsOrt,
				AbfahrtsZeitpunkt: abfahrt,
				AnkunftsZeitpunkt: ankunft,
				Text:              abschnitt.Verkehrsmittel.MittelText,
				Art:               abschnitt.Verkehrsmittel.KurzText,
				Dauer:             int(abschnitt.AbschnittsDauer),
			})
			totalTime += abschnitt.AbschnittsDauer
		}
		res = append(res, Trip{Abschnitte: abschnitte, Dauer: int(totalTime)})
	}

	return res
}
