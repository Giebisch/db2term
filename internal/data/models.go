package data

type TripRequest struct {
	AbfahrtsHalt                      string     `json:"abfahrtsHalt"`
	AnfrageZeitpunkt                  string     `json:"anfrageZeitpunkt"`
	AnkunftsHalt                      string     `json:"ankunftsHalt"`
	AnkunftSuche                      string     `json:"ankunftSuche"`
	Klasse                            string     `json:"klasse"`
	Produktgattungen                  []string   `json:"produktgattungen"`
	Reisende                          []Reisende `json:"reisende"`
	SchnelleVerbindungen              bool       `json:"schnelleVerbindungen"`
	SitzplatzOnly                     bool       `json:"sitzplatzOnly"`
	BikeCarriage                      bool       `json:"bikeCarriage"`
	ReservierungsKontingenteVorhanden bool       `json:"reservierungsKontingenteVorhanden"`
	NurDeutschlandTicketVerbindungen  bool       `json:"nurDeutschlandTicketVerbindungen"`
	DeutschlandTicketVorhanden        bool       `json:"deutschlandTicketVorhanden"`
}

type Reisende struct {
	Typ            string           `json:"typ"`
	Ermaessigungen []Ermaessigungen `json:"ermaessigungen"`
	Alter          []any            `json:"alter"`
	Anzahl         int              `json:"anzahl"`
}

type Ermaessigungen struct {
	Art    string `json:"art"`
	Klasse string `json:"klasse"`
}

type TripResponse struct {
	Verbindungen []struct {
		TripID                string `json:"tripId"`
		CtxRecon              string `json:"ctxRecon"`
		VerbindungsAbschnitte []struct {
			RisNotizen []struct {
				Key          string `json:"key"`
				Value        string `json:"value"`
				RouteIdxFrom int    `json:"routeIdxFrom"`
				RouteIdxTo   int    `json:"routeIdxTo"`
			} `json:"risNotizen"`
			HimMeldungen          []any `json:"himMeldungen"`
			PriorisierteMeldungen []struct {
				Prioritaet string `json:"prioritaet"`
				Text       string `json:"text"`
			} `json:"priorisierteMeldungen"`
			ExterneBahnhofsinfoIDOrigin      string  `json:"externeBahnhofsinfoIdOrigin"`
			ExterneBahnhofsinfoIDDestination string  `json:"externeBahnhofsinfoIdDestination"`
			AbfahrtsZeitpunkt                string  `json:"abfahrtsZeitpunkt"`
			AbfahrtsOrt                      string  `json:"abfahrtsOrt"`
			AbfahrtsOrtExtID                 string  `json:"abfahrtsOrtExtId"`
			AbschnittsDauer                  float32 `json:"abschnittsDauer"`
			AbschnittsAnteil                 float32 `json:"abschnittsAnteil"`
			AnkunftsZeitpunkt                string  `json:"ankunftsZeitpunkt"`
			AnkunftsOrt                      string  `json:"ankunftsOrt"`
			AnkunftsOrtExtID                 string  `json:"ankunftsOrtExtId"`
			Auslastungsmeldungen             []struct {
				Klasse string `json:"klasse"`
				Stufe  int    `json:"stufe"`
			} `json:"auslastungsmeldungen"`
			EzAbfahrtsZeitpunkt        string `json:"ezAbfahrtsZeitpunkt"`
			EzAbschnittsDauerInSeconds int    `json:"ezAbschnittsDauerInSeconds"`
			EzAnkunftsZeitpunkt        string `json:"ezAnkunftsZeitpunkt"`
			Halte                      []struct {
				ID                   string `json:"id"`
				AbfahrtsZeitpunkt    string `json:"abfahrtsZeitpunkt,omitempty"`
				Auslastungsmeldungen []struct {
					Klasse string `json:"klasse"`
					Stufe  int    `json:"stufe"`
				} `json:"auslastungsmeldungen"`
				EzAbfahrtsZeitpunkt   string `json:"ezAbfahrtsZeitpunkt,omitempty"`
				Gleis                 string `json:"gleis"`
				HaltTyp               string `json:"haltTyp"`
				Name                  string `json:"name"`
				RisNotizen            []any  `json:"risNotizen"`
				BahnhofsInfoID        string `json:"bahnhofsInfoId"`
				ExtID                 string `json:"extId"`
				HimMeldungen          []any  `json:"himMeldungen,omitempty"`
				RouteIdx              int    `json:"routeIdx"`
				PriorisierteMeldungen []any  `json:"priorisierteMeldungen"`
				AnkunftsZeitpunkt     string `json:"ankunftsZeitpunkt,omitempty"`
				EzAnkunftsZeitpunkt   string `json:"ezAnkunftsZeitpunkt,omitempty"`
			} `json:"halte"`
			Idx            int    `json:"idx"`
			JourneyID      string `json:"journeyId"`
			Verkehrsmittel struct {
				ProduktGattung string `json:"produktGattung"`
				Kategorie      string `json:"kategorie"`
				LinienNummer   string `json:"linienNummer"`
				Name           string `json:"name"`
				Nummer         string `json:"nummer"`
				Richtung       string `json:"richtung"`
				Typ            string `json:"typ"`
				Zugattribute   []struct {
					Kategorie           string `json:"kategorie"`
					Key                 string `json:"key"`
					Value               string `json:"value"`
					TeilstreckenHinweis string `json:"teilstreckenHinweis,omitempty"`
				} `json:"zugattribute"`
				KurzText   string `json:"kurzText"`
				MittelText string `json:"mittelText"`
				LangText   string `json:"langText"`
			} `json:"verkehrsmittel"`
		} `json:"verbindungsAbschnitte"`
		UmstiegsAnzahl            int  `json:"umstiegsAnzahl"`
		VerbindungsDauerInSeconds int  `json:"verbindungsDauerInSeconds"`
		IsAlternativeVerbindung   bool `json:"isAlternativeVerbindung"`
		Auslastungsmeldungen      []struct {
			Klasse string `json:"klasse"`
			Stufe  int    `json:"stufe"`
		} `json:"auslastungsmeldungen"`
		Auslastungstexte []struct {
			Klasse      string `json:"klasse"`
			Stufe       int    `json:"stufe"`
			KurzText    string `json:"kurzText"`
			AnzeigeText string `json:"anzeigeText,omitempty"`
			LangText    string `json:"langText,omitempty"`
		} `json:"auslastungstexte"`
		HimMeldungen []any `json:"himMeldungen"`
		RisNotizen   []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"risNotizen"`
		PriorisierteMeldungen []struct {
			Prioritaet string `json:"prioritaet"`
			Text       string `json:"text"`
		} `json:"priorisierteMeldungen"`
		ReservierungsMeldungen          []any `json:"reservierungsMeldungen"`
		IsAngebotseinholungNachgelagert bool  `json:"isAngebotseinholungNachgelagert"`
		IsAlterseingabeErforderlich     bool  `json:"isAlterseingabeErforderlich"`
		ServiceDays                     []struct {
			LastDateInPeriod    string   `json:"lastDateInPeriod"`
			Regular             string   `json:"regular"`
			Irregular           string   `json:"irregular"`
			PlanningPeriodBegin string   `json:"planningPeriodBegin"`
			PlanningPeriodEnd   string   `json:"planningPeriodEnd"`
			Weekdays            []string `json:"weekdays"`
		} `json:"serviceDays"`
		EreignisZusammenfassung struct {
			Prioritaet string `json:"prioritaet"`
		} `json:"ereignisZusammenfassung,omitempty"`
		HasTeilpreis                                bool   `json:"hasTeilpreis"`
		ReiseAngebote                               []any  `json:"reiseAngebote"`
		VerbundCode                                 string `json:"verbundCode"`
		HinRueckPauschalpreis                       bool   `json:"hinRueckPauschalpreis"`
		IsReservierungAusserhalbVorverkaufszeitraum bool   `json:"isReservierungAusserhalbVorverkaufszeitraum"`
		GesamtAngebotsbeziehungList                 []any  `json:"gesamtAngebotsbeziehungList"`
		EzVerbindungsDauerInSeconds                 int    `json:"ezVerbindungsDauerInSeconds,omitempty"`
	} `json:"verbindungen"`
	VerbindungReference struct {
		Earlier string `json:"earlier"`
		Later   string `json:"later"`
	} `json:"verbindungReference"`
}
