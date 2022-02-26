package sunset

import (
	"math"
)

/*
Source:
Almanac for Computers, 1990
published by Nautical Almanac Office
United States Naval Observatory
Washington, DC 20392

Inputs:
day, month, year:      date of sunrise/sunset
latitude, longitude:   location for sunrise/sunset
zenith:                Sun's zenith for sunrise/sunset
  offical      = 90 degrees 50'
  civil        = 96 degrees
  nautical     = 102 degrees
  astronomical = 108 degrees

NOTE: longitude is positive for East and negative for West
	NOTE: the algorithm assumes the use of a calculator with the
	trig functions in "degree" (rather than "radian") mode. Most
	programming languages assume radian arguments, requiring back
	and forth convertions. The factor is 180/pi. So, for instance,
	the equation RA = atan(0.91764 * tan(L)) would be coded as RA
	= (180/pi)*atan(0.91764 * tan((pi/180)*L)) to give a degree
	answer with a degree input for L.


1. first calculate the day of the year

N1 = floor(275 * month / 9)
N2 = floor((month + 9) / 12)
N3 = (1 + floor((year - 4 * floor(year / 4) + 2) / 3))
N = N1 - (N2 * N3) + day - 30

2. convert the longitude to hour value and calculate an approximate time

lngHour = longitude / 15

if rising time is desired:
  t = N + ((6 - lngHour) / 24)
if setting time is desired:
  t = N + ((18 - lngHour) / 24)

3. calculate the Sun's mean anomaly

M = (0.9856 * t) - 3.289

4. calculate the Sun's true longitude

L = M + (1.916 * sin(M)) + (0.020 * sin(2 * M)) + 282.634
NOTE: L potentially needs to be adjusted into the range [0,360) by adding/subtracting 360

5a. calculate the Sun's right ascension

RA = atan(0.91764 * tan(L))
NOTE: RA potentially needs to be adjusted into the range [0,360) by adding/subtracting 360

5b. right ascension value needs to be in the same quadrant as L

Lquadrant  = (floor( L/90)) * 90
RAquadrant = (floor(RA/90)) * 90
RA = RA + (Lquadrant - RAquadrant)

5c. right ascension value needs to be converted into hours

RA = RA / 15

6. calculate the Sun's declination

sinDec = 0.39782 * sin(L)
cosDec = cos(asin(sinDec))

7a. calculate the Sun's local hour angle

cosH = (cos(zenith) - (sinDec * sin(latitude))) / (cosDec * cos(latitude))

if (cosH >  1)
  the sun never rises on this location (on the specified date)
if (cosH < -1)
  the sun never sets on this location (on the specified date)

7b. finish calculating H and convert into hours

if if rising time is desired:
  H = 360 - acos(cosH)
if setting time is desired:
  H = acos(cosH)

H = H / 15

8. calculate local mean time of rising/setting

T = H + RA - (0.06571 * t) - 6.622

9. adjust back to UTC

UT = T - lngHour
NOTE: UT potentially needs to be adjusted into the range [0,24) by adding/subtracting 24

10. convert UT value to local time zone of latitude/longitude

localT = UT + localOffset
*/



func radians(x float64) float64 {
	return x / 180.0 * math.Pi
}


func Sunrise(latitude,longitude,day,month,year,zenith,localOffset float64, sunrise bool) float64{
	// 1. first calculate the day of the year

	N1 := math.Floor(275.0 * month / 9.0)
	N2 := math.Floor((month + 9.0) / 12.0)
	N3 := (1.0 + math.Floor((year - 4.0 * math.Floor(year / 4.0) + 2.0) / 3.0))
	N  := N1 - (N2 * N3) + day - 30.0
	
	// 2. convert the longitude to hour value and calculate an approximate time
	
	lngHour := longitude / 15.0
	t:=0.0
	// if rising time is desired:
	if sunrise{t = N + ((6.0 - lngHour) / 24.0)}else{
		t = N + ((18.0 - lngHour) / 24.0)
	}
	
	// 3. calculate the Sun's mean anomaly
	
	M := (0.9856 * t) - 3.289
	
	// 4. calculate the Sun's true longitude
	
	L := M + (1.916 * math.Sin(radians(M))) + (0.020 * math.Sin(2 * radians(M))) + 282.634
	for L>360{
		L-=360
	}
	for L<0{
		L+=360
	}
	
	// 5a. calculate the Sun's right ascension
	
	RA := math.Atan(0.91764 * math.Tan(radians(L)))*180/math.Pi
	for RA>360{
		RA-=360
	}
	for RA<0{
		RA+=360
	}
	
	// 5b. right ascension value needs to be in the same quadrant as L
	
	Lquadrant  := (math.Floor( L/90)) * 90
	RAquadrant := (math.Floor(RA/90)) * 90
	RA = RA + (Lquadrant - RAquadrant)
	
	// 5c. right ascension value needs to be converted into hours
	
	RA = RA / 15.0
	
	// 6. calculate the Sun's declination
	
	sinDec := 0.39782 * math.Sin(radians(L))
	cosDec := math.Cos(math.Asin(sinDec))
	
	// 7a. calculate the Sun's local hour angle
	
	cosH := (math.Cos(radians(zenith)) - (sinDec * math.Sin(radians(latitude)))) / (cosDec * math.Cos(radians(latitude)))
	
	if (cosH >  1){return -999.0}
	//   the sun never rises on this location (on the specified date)
	if (cosH < -1){return 999.0}
	//   the sun never sets on this location (on the specified date)
	
	// 7b. finish calculating H and convert into hours
	
	H:=0.0
	if sunrise{
		H = 360 - math.Acos(cosH)*180/math.Pi
		}else{
			H = math.Acos(cosH)*180/math.Pi}
		
	H = H / 15.0
	
	// 8. calculate local mean time of rising/setting
	
	T := H + RA - (0.06571 * t) - 6.622
	
	// 9. adjust back to UTC
	
	UT := T - lngHour
	for UT<0{UT+=24}
	for UT>24{UT-=24}
	// NOTE: UT potentially needs to be adjusted into the range [0,24) by adding/subtracting 24
	
	// 10. convert UT value to local time zone of latitude/longitude
	
	localT := UT + localOffset
	return localT
}