package services

import (
	"math/rand"
	"time"
)

func Encrypter(msg string) string{
    rand.Seed(time.Now().UnixNano())
    outputMsg := ""
    for _, char := range msg {
        outputMsg += string(char)
        if (rand.Intn(2) % 2 == 0){
            outputMsg += "#"
        }
    }

    return outputMsg
}

