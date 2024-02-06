package geobackend

import (
	"fmt"
	"testing"
)

// var privatekey = ""
// var publickey = ""
// var encode = ""
//  var dbname = "petalar"
// var collname = "petalar"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "petalar")
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{107.2945485891317, -6.3274172711820285},
				{107.29456548761118, -6.327547436591658},
				{107.29461618304731, -6.327544637336743},
				{107.29464012255858, -6.327414471925223},
				{107.2945485891317, -6.3274172711820285},
			},
		},
	}
	datagedung := GeoWithin(mconn, "petalarnew", coordinates)
	fmt.Println(datagedung)
}

// func TestGeoWithin(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "petalar")
// 	coordinates := Polygon{
// 		Coordinates: [][][]float64{
// 			{
// 				{107.2945485891317, -6.3274172711820285},
// 				{107.29456548761118, -6.327547436591658},
// 				{107.29461618304731, -6.327544637336743},
// 				{107.29464012255858, -6.327414471925223},
// 				{107.2945485891317, -6.3274172711820285},
// 			},
// 		},
// 	}
// 	datagedung := GeoWithin(mconn, "petalara", coordinates)
// 	fmt.Println(datagedung)
// }

// func TestNear(t *testing.T) {
// 	mconn := SetConnection2dsphere("MONGOSTRING", "petalar", "petalarnew")
// 	coordinates := Point{
// 		Coordinates: []float64{
// 			107.29589918959925, -6.329130142705225,
// 		},
// 	}
// 	datagedung := Near(mconn, "petalarnew", coordinates)
// 	fmt.Println(datagedung)
// }

// func TestNearSphere(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "petalar")
// 	coordinates := Point{
// 		Coordinates: []float64{
// 			107.29589918959925, -6.329130142705225,
// 		},
// 	}
// 	datagedung := NearSphere(mconn, "petalara", coordinates)
// 	fmt.Println(datagedung)
// }

// func TestBox(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "petalarnew")
// 	coordinates := Polyline{
// 		Coordinates: [][]float64{
// 			{95.32345678901234, 5.567890123456789},
// 			{95.32355678901234, 5.567990123456789},
// 		},
// 	}
// 	datagedung := Box(mconn, "petalarnew", coordinates)
// 	fmt.Println(datagedung)
// }
