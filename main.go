package main
import "fmt"


func convertKeyModeIntToKey(key, mode int) string {
	keyIntArr := []string{"C","Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B" }

	fmt.Printf("Key: %v, Mode: %v\n", key, mode)

	if key > 11 || key < 0 || mode > 1 || mode < 0 {
		return ""
	}

	foundKey := keyIntArr[key]

	if mode == 0 {
		foundKey += "m"
	}

	//check for enharmonic
	if foundKey == "Gbm" { return "F#m" }
	if foundKey == "Abm" { return "G#m" }
	if foundKey == "Dbm" { return "C#m" }


	return foundKey
}

func main() {
	fmt.Println(convertKeyModeIntToKey(0,0))
}
