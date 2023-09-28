package main

import (
	"context"
	"log"
)

func main() {
	mainContext := context.Background()

	// saya akan membuat child dari main context
	contextKendaraan := context.WithValue(mainContext, "jenis", "roda empat")
	contextManusia := context.WithValue(mainContext, "kelamin", "laki-laki")

	// saya akan membuat child dari context kendaraan
	contextMobilSupra := context.WithValue(contextKendaraan, "merk mobil", "supra")
	contextMobilNisan := context.WithValue(contextKendaraan, "merk mobil", "nisan gtr")

	// saya akan membuat child dari context manusia
	contextBagusRamadan := context.WithValue(contextManusia, "nama", "bagus ramadan")
	contextPaijo := context.WithValue(contextManusia, "nama", "paijo")

	log.Println("main context : ", mainContext)
	log.Println("kendaraan context : ", contextKendaraan)
	log.Println("manusia context : ", contextManusia)
	log.Println("mobil supra context : ", contextMobilSupra)
	log.Println("mobil nisan context : ", contextMobilNisan)
	log.Println("bagus ramadan context : ", contextBagusRamadan)
	log.Println("paijo context : ", contextPaijo)

	// mengambil data dari context
	// jika data tidak ada dicontext saat itu maka context tersebut
	// akan mengambil data dari parentnya
	log.Println(contextManusia.Value("kelamin"))  // ada di context manusia
	log.Println(contextMobilSupra.Value("jenis")) // ada di context parent nya yaitu context kendaraan
	log.Println(contextManusia.Value("nama"))     // tidak bisa mengabil data dari child, nil
	log.Println(contextPaijo.Value("merk mobil")) // tidak ada, nil

}
