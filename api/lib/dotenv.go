package lib

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const invalidMessage =  "Not a valid .env file format"

func DotenvConfig(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error reading .env file: %s", err.Error())
	}
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			envVar := strings.SplitN(string(line), "=", 2)
			// verify if envVar is valid env var name
				if len(envVar) == 2 {
					key := envVar[0]
					value := envVar[1]

					validateDotenvVarName(key)

					os.Setenv(key, value)
				} else {
					panic(invalidMessage)
				}
			// set env var
		}
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic("Error reading file")
		}
	}

}

// UTF-8. Only linux
func validateDotenvVarName(name string) {
	bytes := []byte(name)	
	lowLineCode := 95
	for _, char := range bytes {
		if char < 65 || char > 90 {
			if (char != byte(lowLineCode)) {
				log.Fatalf("%s\n", invalidMessage)
			}
		}
	}
}
