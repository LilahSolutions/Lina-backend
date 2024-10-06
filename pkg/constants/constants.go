package constants

import "os"

var LOGIN_BASE_PATH = os.Getenv("LOGIN_BASE_PATH")

const LOGIN_API_URL = "/api/login"

const APP_ID = 2

const LILA_LOGIN_ACCESS = "5dd79e61-918f-4215-8358-6241b5d7ecfc"

const TEMPLATE_ROLE = 2

// Gemini
var (
	GEMINI_API_KEY     = os.Getenv("GEMINI_API_KEY")
	GEMINI_MODEL       = "gemini-pro"
	GEMINI_TEMPERATURE = 0.1
)

/**
  * Example of a user in JSON
**/
var USER_DATA = `{
  "user": {
    "name": "John Doe",
    "location": {
      "city": "Santa Rosa",
      "province": "Mendoza",
      "country": "Argentina",
      "coordinates": {
        "latitude": -33.4388,
        "longitude": -68.5285
      }
    },
    "farm": {
      "size": "100 hectares",
      "crops": [
        {
          "type": "potato",
          "areaCultivated": "40 hectares"
        },
        {
          "type": "tomato",
          "areaCultivated": "30 hectares"
        },
        {
          "type": "lettuce",
          "areaCultivated": "30 hectares"
        }
      ],
      "irrigationType": "drip",
      "soilType": "loamy",
      "averageRainfallPerYear": "250 mm"
    }
  }
}
`

/**
  * Summary of tomato, potato and lettuce needs
**/
var CROP_DATA = `{
  "crops": {
    "tomato": {
      "temperatureRange": {
        "optimalDay": [21, 27],
        "optimalNight": [13, 18],
        "thresholds": {
          "max": 35,
          "min": 10
        }
      },
      "waterNeeds": {
        "weekly": 25,
        "unit": "mm"
      },
      "humidityRange": [65, 85],
      "sunlightHours": 6,
      "actions": {
        "hotWeather": "applyShadeCloth, increaseWatering",
        "rainyWeather": "ensureGoodDrainage, avoidOverheadIrrigation",
        "coldWeather": "coverWithFrostCloth",
        "windyWeather": "useWindbreaks"
      }
    },
    "potato": {
      "temperatureRange": {
        "optimal": [15, 22],
        "thresholds": {
          "max": 30,
          "min": 10
        }
      },
      "waterNeeds": {
        "weekly": 30,
        "unit": "mm"
      },
      "humidityRange": [60, 80],
      "sunlightHours": 6,
      "actions": {
        "hotWeather": "applyMulch, chooseHeatTolerantVarieties",
        "rainyWeather": "ensureGoodDrainage",
        "coldWeather": "coverWithStrawMulch",
        "windyWeather": "stakeTallerVarieties"
      }
    },
    "lettuce": {
      "temperatureRange": {
        "optimal": [15, 21],
        "thresholds": {
          "max": 24
        }
      },
      "waterNeeds": {
        "weekly": 30,
        "unit": "mm"
      },
      "humidityRange": [65, 85],
      "sunlightHours": 4,
      "actions": {
        "hotWeather": "applyShadeCloth, increaseWatering",
        "rainyWeather": "ensureGoodDrainage, avoidWetLeaves",
        "coldWeather": "coverWithFrostCloth",
        "windyWeather": "useWindbreaks"
      }
    }
  }
}`

/**
  * Weather data for the week starting on 2024-10-03 in Santa Rosa, Mendoza, Argentina
**/
var WEATHER_DATA = `{ 
  "weatherData": {
    "weekStarting": "2024-10-03",
    "dailyData": [
      {
        "date": "2024-10-03",
        "temperature": {
          "dailyAverage": 26.1,
          "minTemperature": 18.2,
          "maxTemperature": 32.5
        },
        "rainfall": {
          "dailyTotal": 10.4
        },
        "humidity": {
          "dailyAverage": 67.1
        },
        "wind": {
          "averageSpeed": 11.3,
          "maxGustSpeed": 22.5,
          "prevailingDirection": "NW"
        },
        "solarRadiation": {
          "dailySunlightHours": 8.2
        },
        "growingDegreeDays": 22,
        "evapotranspirationRate": 4.8,
        "barometricPressure": 1013.1,
        "dewPoint": 16.3
      },
      {
        "date": "2024-10-04",
        "temperature": {
          "dailyAverage": 27.4,
          "minTemperature": 19.0,
          "maxTemperature": 33.8
        },
        "rainfall": {
          "dailyTotal": 5.7
        },
        "humidity": {
          "dailyAverage": 65.4
        },
        "wind": {
          "averageSpeed": 10.8,
          "maxGustSpeed": 21.9,
          "prevailingDirection": "N"
        },
        "solarRadiation": {
          "dailySunlightHours": 8.6
        },
        "growingDegreeDays": 24,
        "evapotranspirationRate": 5.0,
        "barometricPressure": 1012.7,
        "dewPoint": 17.1
      },
      {
        "date": "2024-10-05",
        "temperature": {
          "dailyAverage": 25.8,
          "minTemperature": 17.5,
          "maxTemperature": 31.9
        },
        "rainfall": {
          "dailyTotal": 0.0
        },
        "humidity": {
          "dailyAverage": 60.2
        },
        "wind": {
          "averageSpeed": 12.1,
          "maxGustSpeed": 23.7,
          "prevailingDirection": "NE"
        },
        "solarRadiation": {
          "dailySunlightHours": 9.1
        },
        "growingDegreeDays": 20,
        "evapotranspirationRate": 4.9,
        "barometricPressure": 1014.2,
        "dewPoint": 15.8
      },
      {
        "date": "2024-10-06",
        "temperature": {
          "dailyAverage": 26.3,
          "minTemperature": 18.6,
          "maxTemperature": 32.4
        },
        "rainfall": {
          "dailyTotal": 12.9
        },
        "humidity": {
          "dailyAverage": 68.7
        },
        "wind": {
          "averageSpeed": 13.4,
          "maxGustSpeed": 24.8,
          "prevailingDirection": "W"
        },
        "solarRadiation": {
          "dailySunlightHours": 7.8
        },
        "growingDegreeDays": 23,
        "evapotranspirationRate": 4.7,
        "barometricPressure": 1011.6,
        "dewPoint": 16.5
      },
      {
        "date": "2024-10-07",
        "temperature": {
          "dailyAverage": 27.9,
          "minTemperature": 19.3,
          "maxTemperature": 34.0
        },
        "rainfall": {
          "dailyTotal": 8.3
        },
        "humidity": {
          "dailyAverage": 66.0
        },
        "wind": {
          "averageSpeed": 10.9,
          "maxGustSpeed": 22.0,
          "prevailingDirection": "NW"
        },
        "solarRadiation": {
          "dailySunlightHours": 9.3
        },
        "growingDegreeDays": 25,
        "evapotranspirationRate": 5.2,
        "barometricPressure": 1010.3,
        "dewPoint": 17.6
      },
      {
        "date": "2024-10-08",
        "temperature": {
          "dailyAverage": 28.6,
          "minTemperature": 20.1,
          "maxTemperature": 35.2
        },
        "rainfall": {
          "dailyTotal": 4.5
        },
        "humidity": {
          "dailyAverage": 63.8
        },
        "wind": {
          "averageSpeed": 12.7,
          "maxGustSpeed": 25.3,
          "prevailingDirection": "NE"
        },
        "solarRadiation": {
          "dailySunlightHours": 9.5
        },
        "growingDegreeDays": 27,
        "evapotranspirationRate": 5.1,
        "barometricPressure": 1009.7,
        "dewPoint": 18.1
      },
      {
        "date": "2024-10-09",
        "temperature": {
          "dailyAverage": 27.3,
          "minTemperature": 19.2,
          "maxTemperature": 33.7
        },
        "rainfall": {
          "dailyTotal": 9.7
        },
        "humidity": {
          "dailyAverage": 64.9
        },
        "wind": {
          "averageSpeed": 11.6,
          "maxGustSpeed": 23.1,
          "prevailingDirection": "SE"
        },
        "solarRadiation": {
          "dailySunlightHours": 8.9
        },
        "growingDegreeDays": 26,
        "evapotranspirationRate": 4.9,
        "barometricPressure": 1010.8,
        "dewPoint": 17.4
      }
    ]
  }
}`
