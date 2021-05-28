package measurement

import "github.com/ben-eh/cocktailTime/database"

type Measurement struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllMeasurements() []Measurement {
	db := database.DBConnection()
	defer db.Close()

	results, err2 := db.Query("SELECT measurement_id, name FROM measurements")
	if err2 != nil {
		panic(err2.Error())
	}

	var measurements []Measurement

	for results.Next() {
		var measurement Measurement
		err := results.Scan(&measurement.ID, &measurement.Name)
		if err != nil {
			panic(err.Error())
		}

		measurements = append(measurements, measurement)
	}

	return measurements
}
