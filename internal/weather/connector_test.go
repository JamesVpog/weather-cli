package weather

import "testing"

func TestParseWeatherData(t *testing.T) {
	jsonData := []byte(`{
	  "coord": {
	    "lon": 10.99,
	    "lat": 44.34
	  },
	  "weather": [
	    {
	      "id": 804,
	      "main": "Clouds",
	      "description": "overcast clouds",
	      "icon": "04n"
	    }
	  ],
	  "base": "stations",
	  "main": {
	    "temp": 279.07,
	    "feels_like": 279.07,
	    "temp_min": 277.18,
	    "temp_max": 279.86,
	    "pressure": 1023,
	    "humidity": 97,
	    "sea_level": 1023,
	    "grnd_level": 955
	  },
	  "visibility": 1536,
	  "wind": {
	    "speed": 0.91,
	    "deg": 28,
	    "gust": 1.1
	  },
	  "clouds": {
	    "all": 100
	  },
	  "dt": 1766182946,
	  "sys": {
	    "type": 2,
	    "id": 2004688,
	    "country": "IT",
	    "sunrise": 1766126865,
	    "sunset": 1766158714
	  },
	  "timezone": 3600,
	  "id": 3163858,
	  "name": "Zocca",
	  "cod": 200
	}`)

	got, err := ParseWeather(jsonData)
	if err != nil {
		t.Fatalf("ParseWeatherData() returned an error: %v", err)
	}

	// check if assertions match for important fields temp, and weather[0].description
	if got.Main.Temp != 279.07 {
		t.Errorf("got %v, want %f", got, 279.07)
	}

	if got.Main.Humidity != 97 {
		t.Errorf("got %d, want %d:", got.Main.Humidity, 97)
	}

	if len(got.Weather) == 0 {
		t.Fatalf("Expected at least 1 weather description, got 0")
	}

	if got.Weather[0].Description != "overcast clouds" {
		t.Errorf("got %s, want %s", got.Weather[0].Description, "overcast clouds")
	}
}
