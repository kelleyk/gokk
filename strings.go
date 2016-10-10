package gokk

import "strings"

func Partition(s, sep string) (string, string, string) {
	i := strings.Index(s, sep)
	if i == -1 {
		return "", "", s
	}
	return s[0:i], s[i : i+1], s[i+1:]
}

func PartitionLast(s, sep string) (string, string, string) {
	i := strings.LastIndex(s, sep)
	if i == -1 {
		return s, "", ""
	}
	return s[0:i], s[i : i+1], s[i+1:]
}
