package util

import (
    "hash/crc32"
    "fmt"
    "strings"
    "strconv"
)


func GetHash(id int, ids []int) string {
    stringIds := make([]string, len(ids) + 1)
    stringIds[0] = strconv.Itoa(id)
    for i, id := range ids {
        stringIds[i+1] = strconv.Itoa(id)
    }

    data := strings.Join(stringIds, "-")
    hash := crc32.ChecksumIEEE([]byte(data))

    return fmt.Sprintf("%x", hash)
}
