package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var hostnameRegex = regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
var userRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_-]*[$]?$`)

func SplitTargetAndPort(target string) (string, string, error) {

	// Runs only when user enters the target flag with hostname with Port e.x '--target google.com:80'
	if strings.Contains(target, ":") {
		parts := strings.Split(target, ":")
		if !IsHostnameValid(parts[0]) {
			return "", "", fmt.Errorf("invalid hostname format: %s", parts[0])
		}
		if !IsPortValid(parts[1]) {
			return "", "", fmt.Errorf("invalid port number: %s", parts[1])
		}
		return parts[0], parts[1], nil
	}

	// Runs only when user enters the target flag with hostname only e.x '--target google.com'
	if !IsHostnameValid(target) {
		return "", "", fmt.Errorf("invalid hostname format: %s", target)
	}
	return target, "", nil
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

func IsHostnameValid(hostname string) bool {
	return hostnameRegex.MatchString(hostname)
}

func GetValidUserInput(user string) (string, error) {
	if user == "" {
		return "root", nil
	}
	length := len(user)
	if length > 32 {
		return "", fmt.Errorf("invalid user length: %d", length)
	}
	if !userRegex.MatchString(user) {
		return "", fmt.Errorf("invalid user format: %s", user)
	}
	return user, nil
}
