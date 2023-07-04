package helper

import "github.com/speps/go-hashids/v2"

func UUIDEncode(userID int) string {
	hd := hashids.NewData()
	hd.Salt = "saltituasin"
	hd.MinLength = 30
	newWithData, _ := hashids.NewWithData(hd)
	encode, err := newWithData.Encode([]int{userID})
	PanicIfError(err)

	return encode
}

func UUIDDecode(uuid string) int {
	hd := hashids.NewData()
	hd.Salt = "saltituasin"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	decodeWithError, err := h.DecodeWithError(uuid)
	PanicIfError(err)

	return decodeWithError[0]
}
