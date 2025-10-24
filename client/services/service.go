package services

func Decryptor(msg string) string {
    outputMsg := ""
    for _, char := range(msg){
        if (char == '#'){
            println("decrypting ...")
        } else {
            outputMsg += string(char)
        }
    }
    return outputMsg
}
