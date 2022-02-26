# sunset

Just a Golang version of an existing workflow to generate sunrise and sunset times at a specific location and date.

package is sunset, but function is sunrise :-)

### Usage

sunset.Sunrise(latitude, longitude, day, month, year, zenith, localOffset, sunrise)

. everything is decimal except sunrise, which is boolean.
. localOffset is your current time zone, including daylight savings time
zenith is the angle of the Sun from vertical, as different twilights exist in civil, nautical and astronomic contexts.

### based on algorithm at: https://edwilliams.org/sunrise_sunset_algorithm.htm
The page explains everything better than I can. It in turn is based on:
	Almanac for Computers, 1990
	published by Nautical Almanac Office
	United States Naval Observatory
	Washington, DC 20392
    


