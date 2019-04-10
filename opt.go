package goetc

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"math/rand"
)

func InSlice(v string, sl []string) int {
	for i, vv := range sl {
		if vv == v {
			return i
		}
	}
	return -1
}

// SliceDiff returns diff slice of slice1 - slice2.
func SliceDiff(slice1, slice2 []string) (diffslice []string) {
	for _, v := range slice1 {
		if InSlice(v, slice2) != -1 {
			diffslice = append(diffslice, v)
		}
	}
	return
}

// SliceIntersect returns slice that are present in all the slice1 and slice2.
func SliceIntersect(slice1, slice2 []string) (diffslice []string) {
	for _, v := range slice1 {
		if InSlice(v, slice2) == -1 {
			diffslice = append(diffslice, v)
		}
	}
	return
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

//before you call func RandStirng, you must use rand.Seed(time.Now().UTC().UnixNano()) to set seed
func RandString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//a string queue,you can change aguments type to apply else type
func EnqueueString(slice []string, element string, queueSize int) []string {
	n := len(slice)
	if n >= queueSize {
		slice = slice[1:]
	}
	return append(slice, element)
}

func SliceRemoveDuplicate(elements []string) []string {
	keys := make(map[string]bool)
	result := make([]string, 0, len(elements))
	for _, element := range elements {
		if _, value := keys[element]; !value {
			keys[element] = true
			result = append(result, element)
		}
	}
	return result
}

func SliceReverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func StringRepeate(elemen string, times int) string {
	var buffer bytes.Buffer
	for i := 0; i < times; i++ {
		buffer.WriteString(elemen)
	}
	return buffer.String()
}

func SliceGetLastElement(s []string, lastInex int) []string {
	return s[len(s)-lastInex:]
}

func SliceInsert(s []string, element string, index int) []string {
	s = append(s, "0")
	copy(s[index+1:], s[index:])
	s[index] = element
	return s
}

func SliceDeleteElementByIndex(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func SliceDeleteElementByValue(s []string, value string) []string {
	for index, v := range s {
		if v == value {
			return SliceDeleteElementByIndex(s, index)
		}
	}
	return s
}

func SliceAdd(s1, s2 []string) []string {
	return append(s1, s2...)
}

func Md5(str string) []byte {
	m := md5.New()
	m.Write([]byte(str))
	return m.Sum([]byte(""))
}

func Sha1(str string) []byte {
	s := sha1.New()
	s.Write([]byte(str))
	return s.Sum([]byte(""))
}
