package controllers

import (
	"fmt"
	"math/rand"
	"strconv"

	"bitbucket.org/sjbog/go-dbscan"
	"github.com/masatana/go-textdistance"
)

type Information struct {
	Points     []dbscan.ClusterablePoint
	BaseString string
	BaseNumber float64
}

var World map[uint]Information = make(map[uint]Information)

func GetClusters(points []dbscan.ClusterablePoint) [][]dbscan.ClusterablePoint {
	clusterer := dbscan.NewDBSCANClusterer(0.1, 2)

	clusterer.MinPts = 2
	clusterer.SetEps(2.0)

	// Automatic discovery of dimension with max variance
	clusterer.AutoSelectDimension = false
	// Set dimension manually
	clusterer.SortDimensionIndex = 1

	var result [][]dbscan.ClusterablePoint = clusterer.Cluster(points)

	return result
}

func InitUserInWorld(userId uint) {
	World[userId] = Information{Points: []dbscan.ClusterablePoint{}, BaseString: "", BaseNumber: 0}
}

func AddPointToUser(userId uint, point dbscan.ClusterablePoint) {
	ourWorld, ok := World[userId]
	if !ok {
		World[userId] = Information{Points: []dbscan.ClusterablePoint{point}, BaseString: ""}
	} else {
		ourWorld.Points = append(ourWorld.Points, point)
		World[userId] = ourWorld
	}
}

func Setup(record Record) {
	userId, err := strconv.ParseUint(record.User_id, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	ourWorld := World[uint(userId)]

	var x int
	var y int

	if record.Param[0].Data_type == "STRING" {
		if ourWorld.BaseString == "" {
			ourWorld.BaseString = record.Param[0].Value
		}
		x = textdistance.LevenshteinDistance(ourWorld.BaseString, record.Param[0].Value)
	} else {
		if ourWorld.BaseNumber == 0 {
			ourWorld.BaseNumber, err = strconv.ParseFloat(record.Param[0].Value, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		x, err = strconv.Atoi(record.Param[0].Value)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if record.Param[1].Data_type == "STRING" {
		if ourWorld.BaseString == "" {
			ourWorld.BaseString = record.Param[1].Value
		}
		y = textdistance.LevenshteinDistance(ourWorld.BaseString, record.Param[1].Value)
	} else {
		if ourWorld.BaseNumber == 0 {
			ourWorld.BaseNumber, err = strconv.ParseFloat(record.Param[1].Value, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		y, err = strconv.Atoi(record.Param[1].Value)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	numberOfPoints := len(ourWorld.Points)

	numberOfPointsInString := strconv.Itoa(numberOfPoints)

	point := &dbscan.NamedPoint{Name: numberOfPointsInString, Point: []float64{float64(x), float64(y)}}

	ourWorld.Points = append(ourWorld.Points, point)
}

func GetClustersForUser(userId uint) [][]dbscan.ClusterablePoint {
	ourWorld := World[userId]
	return GetClusters(ourWorld.Points)
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func TestSetup() {
	World = make(map[uint]Information)

	InitUserInWorld(1)

	for i := 0; i < 100; i++ {
		intValueString := strconv.Itoa(rand.Intn(1000))
		Setup(Record{User_id: "1", Param: []ParamType{{Data_type: "STRING", Value: RandomString(rand.Intn(30))}, {Data_type: "INT", Value: intValueString}}})
	}

	ourWorld := World[1]

	// Print points
	for _, point := range ourWorld.Points {
		fmt.Println(point)
	}

}
