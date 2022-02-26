package main

import (
	"fmt"
	"math"

	"github.com/hippodribble/sunset/sunset"
)

func main() {
	latitude := -31.166
	longitude := 12.5
	year := 2022.0
	month := 2.0
	day := 27.0
	zenith := 90.8333333
	localOffset := 2.0
	sunrise:=true

	t := sunset.Sunrise(latitude, longitude, day, month, year, zenith, localOffset, sunrise)
	h:=math.Floor(t)
	m:=math.Floor((t-h)*60.0)
	s:=(t-h-m/60)*3600
	fmt.Printf("%02.0f:%02.0f:%02.0f\n",h,m,s)

	sunrise=false
	t = sunset.Sunrise(latitude, longitude, day, month, year, zenith, localOffset, sunrise)
	h=math.Floor(t)
	m=math.Floor((t-h)*60.0)
	s=(t-h-m/60)*3600
	fmt.Printf("%02.0f:%02.0f:%02.0f\n",h,m,s)
}
