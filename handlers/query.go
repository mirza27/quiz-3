package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"fmt"
	"strconv"
)

// endpoint segitiga
func QuerySegitiga(req *gin.Context){
	var result interface{}

	// hitung luas
	hitung := req.Query("hitung")
	Qalas := req.Query("alas")
	Qtinggi := req.Query("tinggi")

	alas, err := strconv.ParseFloat(Qalas, 64)
	if err != nil {
		panic(err)
	}

	tinggi, err := strconv.ParseFloat(Qtinggi, 64)
	if err != nil {
		panic(err)
	}

	if hitung == "luas"{
		result = alas*tinggi/2
	} else {
		result = alas*3
	}

	req.JSON(http.StatusOK, gin.H{
			"jenis hitung": hitung,
			"alas": alas,
			"tinggi": tinggi,
			"hasil" : result, 
		},
	)
}


// endpoint persegi
func QueryPersegi(req *gin.Context){
	var result interface{}

	// hitung luas
	hitung := req.Query("hitung")
	Qsisi := req.Query("sisi")
	

	sisi, err := strconv.ParseFloat(Qsisi, 64)
	if err != nil {
		panic(err)
	}

	if hitung == "luas"{
		result = sisi*sisi
	} else {
		result = sisi*4
	}

	req.JSON(http.StatusOK, gin.H{
			"jenis hitung": hitung,
			"sisi": sisi,
			"hasil" : result, 
		},
	)
}


// endpoint persegi panjang
func QueryPersegiPanjang(req *gin.Context){
	var result interface{}

	// hitung luas
	hitung := req.Query("hitung")
	Qpanjang := req.Query("panjang")
	Qlebar := req.Query("lebar")

	panjang, err := strconv.ParseFloat(Qpanjang, 64)
	if err != nil {
		panic(err)
	}

	lebar, err := strconv.ParseFloat(Qlebar, 64)
	if err != nil {
		panic(err)
	}

	if hitung == "luas"{
		result = panjang*lebar
	} else {
		result = (2*panjang)+(2*lebar)
	}

	req.JSON(http.StatusOK, gin.H{
			"jenis hitung": hitung,
			"panjang": panjang,
			"lebar": lebar,
			"hasil" : result, 
		},
	)
}

// endpoint lingkaran
func QueryLingkaran(req *gin.Context){
	var result interface{}

	// hitung luas
	hitung := req.Query("hitung")
	QjariJari := req.Query("jariJari")
	

	jariJari, err := strconv.ParseFloat(QjariJari, 64)
	if err != nil {
		panic(err)
	}

	if hitung == "luas"{
		result = jariJari*jariJari*3.14
	} else {
		result = jariJari*2*3.14
	}

	req.JSON(http.StatusOK, gin.H{
			"jenis hitung": hitung,
			"sisi": jariJari,
			"hasil" : result, 
		},
	)
}


