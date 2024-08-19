package greetings

import (
    "errors"
    "fmt"
)
func Hello(name string) (string, error) {
    // 如果没有给出名字，返回一个带有消息的错误。
    if name == "" {
        return "", errors.New("empty name")
    }

    
    var message string
    message = fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
