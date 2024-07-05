package utils

import (
	"log"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Erro ao converter string para int: %v", err)
	}
	return i
}

func Atof(s string) float64 {
	f, err := strconv.ParseFloat(strings.Replace(s, ",", ".", 1), 64)
	if err != nil {
		log.Printf("Erro ao converter string para float: %v", err)
	}
	return f
}
