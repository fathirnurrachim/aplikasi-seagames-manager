package main

const NMAX int = 10

type Negara struct {
	// id     int
	nama   string
	medali Medali
}

type Medali struct {
	emas     int
	perak    int
	perunggu int
}

type Data [NMAX]Negara
