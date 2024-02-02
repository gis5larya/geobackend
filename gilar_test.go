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

// func TestGeoIntersects(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", dbname)
// 	coordinates := Point{
// 		Coordinates: []float64{
// 			103.60768133536988, -1.628526295003084,
// 		},
// 	}
// 	datagedung := GeoIntersects(mconn, "petalaryak", coordinates)
// 	fmt.Println(datagedung)
// }

// func TestGeoWithin(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", dbname)
// 	coordinates := Polygon{
// 		Coordinates: [][][]float64{
// 			{
// 				{95.31123456789012, 5.553210987654321},
// 				{95.31133456789011, 5.553210987654321},
// 				{95.31133456789011, 5.553310987654321},
// 				{95.31123456789012, 5.553310987654321},
// 				{95.31123456789012, 5.553210987654321},
// 			},
// 		},
// 	}
// 	datagedung := GeoWithin(mconn, "petalaryak", coordinates)
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
