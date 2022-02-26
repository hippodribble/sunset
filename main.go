package main
import "github.com/hippodribble/sunset/sunset"

func main(){
	latitude := -31.5
	longitude:=13.0
	year:=2022.0
	month:=2.0
	day:=26.0
	zenith:=0.0
	localOffset:=0.0
	sunrise:=false

	t:=sunset.sunrise(latitude,longitude,day,month,year,zenith,localOffset,sunrise)
	
}