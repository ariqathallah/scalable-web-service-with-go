package helper

import (
	models "assignment-3/model"
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func UpdateData() {
	for {
		data := models.Status{
			Status: models.Data{
				Water: rand.Intn(100) + 1,
				Wind:  rand.Intn(100) + 1,
			},
		}

		// encode json
		bytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		// get current directory
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		// write file
		err = os.WriteFile(filepath.Join(dir, "lib", "data.json"), bytes, 0644)
		if err != nil {
			panic(err)
		}

		// update every 15 seconds
		time.Sleep(15 * time.Second)
	}
}
