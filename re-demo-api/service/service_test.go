package service

import (
	"re-demo-api/structs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPacksResponse(t *testing.T) {
	// case 1
	req := structs.PackRequest{Items: 1, PackSizes: []int{250, 500, 1000, 2000, 5000}}
	res, _ := GetPacksResponse(req)
	exp := &structs.PackResponse{NumOfPacks: 1, Packs: []string{"1x250"}}
	assert.Equal(t, res, exp)

	// case 2
	req = structs.PackRequest{Items: 250, PackSizes: []int{250, 500, 1000, 2000, 5000}}
	res, _ = GetPacksResponse(req)
	exp = &structs.PackResponse{NumOfPacks: 1, Packs: []string{"1x250"}}
	assert.Equal(t, res, exp)

	// case 3
	req = structs.PackRequest{Items: 251, PackSizes: []int{250, 500, 1000, 2000, 5000}}
	res, _ = GetPacksResponse(req)
	exp = &structs.PackResponse{NumOfPacks: 1, Packs: []string{"1x500"}}
	assert.Equal(t, res, exp)

	// case 4
	req = structs.PackRequest{Items: 501, PackSizes: []int{250, 500, 1000, 2000, 5000}}
	res, _ = GetPacksResponse(req)
	exp = &structs.PackResponse{NumOfPacks: 2, Packs: []string{"1x500", "1x250"}}
	assert.Equal(t, res, exp)

	// TODO: fix case 5
	// case 5
	// req := structs.PackRequest{Items: 12001, PackSizes: []int{250, 500, 1000, 2000, 5000}}
	// res, _ := GetPacksResponse(req)
	// exp := &structs.PackResponse{NumOfPacks: 4, Packs: []string{"2x5000", "1x2000", "1x250"}}
	// assert.Equal(t, res, exp)
}
