package service

import (
	"errors"
	"fmt"
	"re-demo-api/structs"
	"sort"
)

/*
 * GetPacksResponse: Gets request from controller, calculates packs count and content and returns the response struct
 */
func GetPacksResponse(preq structs.PackRequest) (*structs.PackResponse, error) {
	// check for valid number of items
	if preq.Items == 0 {
		return nil, errors.New("no items to pack")
	}

	// check for empty pack sizes, if so, use the predefined pack sizes in the problem statement
	if len(preq.PackSizes) == 0 {
		preq.PackSizes = []int{250, 500, 1000, 2000, 5000}
	}

	// sort pack sizes in ascending order
	sort.Slice(preq.PackSizes, func(i, j int) bool {
		return preq.PackSizes[i] < preq.PackSizes[j]
	})

	// create multipliers slice
	// set first multiplier
	muls := make([]int, len(preq.PackSizes))
	muls[0] = (preq.Items / preq.PackSizes[0])

	if preq.Items%preq.PackSizes[0] > 0 {
		muls[0]++
	}

	// slide muls by replacing smaller with bigger
	for i := 0; i < len(preq.PackSizes)-1; i++ {
		if muls[i] == 0 {
			break
		}

		div := preq.PackSizes[i+1] / preq.PackSizes[i]
		muls[i+1] = muls[i] / div
		muls[i] %= div
	}

	// define response variables
	numOfPacks := 0
	packs := []string{}

	// set response variables
	for i := len(muls) - 1; i >= 0; i-- {
		m := muls[i]

		if m > 0 {
			numOfPacks += m
			ps := preq.PackSizes[i]
			packStr := fmt.Sprintf("%dx%d", m, ps)
			packs = append(packs, packStr)
		}
	}

	// build response
	pres := &structs.PackResponse{
		NumOfPacks: numOfPacks,
		Packs:      packs,
	}

	return pres, nil
}
