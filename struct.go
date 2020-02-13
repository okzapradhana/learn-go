package main

import "fmt"

type mahasiswa struct {
	Nama     string
	Umur     int32
	Fakultas string
	Matkul   matakuliah //embedded struct
	alamat
}

type alamat struct {
	NamaJalan string
	KodePos   int32
}

type matakuliah struct {
	NamaMatkul string
	Kode       string
}

//nested struct
type student struct {
	person struct {
		name string
		age  int
	}
	grade int
}

func main() {
	var mhs mahasiswa
	mhs.Nama = "Kirito"
	mhs.Umur = 20
	mhs.Fakultas = "Comuter Science"

	fmt.Println(mhs.Nama)
	fmt.Println(mhs.Umur)
	fmt.Println(mhs.Fakultas)

	var m1 = mahasiswa{}
	m1.Nama = "harry"
	m1.Umur = 19

	var m3 = mahasiswa{Nama: "mary"}

	fmt.Println("mhs 1", m1.Nama)
	fmt.Println("mhs 3", m3.Nama)

	//pointer struct
	var mtkl = matakuliah{NamaMatkul: "Math", Kode: "CIF450"}
	var mm1 = mahasiswa{Nama: "hendro", Fakultas: "Science", Matkul: mtkl}
	var mm2 *mahasiswa = &mm1

	fmt.Println(&mm1)
	fmt.Println(mm2)

	mm1.Nama = "jonathan"
	fmt.Println(&mm1)
	fmt.Println(mm2)

	var m = mahasiswa{}
	m.Matkul.NamaMatkul = "Fisika"
	m.Matkul.Kode = "CF123"
	m.NamaJalan = "Tenosius"
	m.KodePos = 14356

	fmt.Println(m.Matkul.NamaMatkul)
	fmt.Println(m.Matkul.Kode)
	fmt.Println(m.NamaJalan)
	fmt.Println(m.alamat.KodePos)

	var mhs1 = struct {
		alamat
		hobi string
	}{} //empty initialization

	var mhs2 = struct {
		alamat
		hobi string
	}{
		alamat: alamat{"Johor", 14566},
		hobi:   "Melukis",
	}

	mhs1.alamat = alamat{"Kapuari", 123455}
	mhs1.hobi = "Makan"

	fmt.Println("Mhs1 anonymous struct", mhs1)
	fmt.Println("Mhs1 anonymous struct", mhs2)

	var courses = []matakuliah{
		{NamaMatkul: "English", Kode: "ENG12"},
		{NamaMatkul: "Astronomy", Kode: "AS12"},
		{NamaMatkul: "Sport", Kode: "SP12"},
	}

	fmt.Println("List Mata kuliah: ", courses)
	for _, course := range courses {
		fmt.Println(course.NamaMatkul, course.Kode)
	}

	//anonymous slice struct
	var phones = []struct {
		number   int64
		location string
	}{
		{number: 121221312, location: "Indo"},
		{number: 3123215, location: "Malay"},
		{number: 23123127, location: "Bangladesh"},
	}

	for _, phone := range phones {
		fmt.Println(phone.number, phone.location)
	}

	type Number = int
	var num Number = 20
	fmt.Println(num)

}
