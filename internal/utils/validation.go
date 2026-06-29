package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SplitTargetAndPort(target string) (string, string, error) {
	if strings.Contains(target, ":") {
		parts := strings.Split(target, ":")
		if len(parts) != 2 {
			return "", "", fmt.Errorf("invalid target format: %s", target)
		}
		if !IsPortValid(parts[1]) {
			return "", "", fmt.Errorf("invalid port number: %s", parts[1])
		}
		return parts[0], parts[1], nil
	}
	return target, "80", nil
}

func IsPortValid(port string) bool {
	_port, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	if _port < 1 || _port > 65535 {
		return false
	}
	return true
}

func GetValidUserInput(user string) (string, error) {
	if user == "" {
		return "root", nil
	}
	length := len(user)
	if length > 32 {
		return "", fmt.Errorf("invalid user length: %d", length)
	}
	var userRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_-]*[$]?$`)
	if !userRegex.MatchString(user) {
		return "", fmt.Errorf("invalid user format: %s", user)
	}
	return user, nil
}
