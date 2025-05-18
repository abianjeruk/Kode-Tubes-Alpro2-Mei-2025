package main

import("fmt";"strings")

const nmax int = 10

type konten struct {
	daftarKonten string
	tglPosting string
	kategori string
	jamPosting string
	jmlInteraksi int
}

type tabKonten [nmax]konten

func bacaDaftarKonten(A *tabKonten, n *int) {
	var i int
	
	fmt.Println("Mau buat berapa konten? (max = 10)")
	fmt.Scan(n)
	for i = 0; i < *n; i++ {
		fmt.Println("\nJudul konten (jangan menggunakan spasi)", i+1, ":")
		fmt.Scan(&A[i].daftarKonten)
		fmt.Print("Kategori: ")
		fmt.Scan(&A[i].kategori)
		A[i].jmlInteraksi = 0
	}
}

func jadwalkanKonten(A *tabKonten, n int) {
	
	var idx int
	fmt.Print("\nMasukkan indeks konten yang ingin dijadwalkan (0 hingga ", n-1, "): ")
	fmt.Scan(&idx)

	if idx >= 0 && idx < n {
		fmt.Print("Masukkan tanggal posting : ")
		fmt.Scan(&A[idx].tglPosting)
		fmt.Print("Masukkan jam posting : ")
		fmt.Scan(&A[idx].jamPosting)
		fmt.Println("Konten berhasil dijadwalkan.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func ideKonten() {
	var input string
	var i int
	var ide string
	var ditemukan bool
	var selesai bool
	var daftarIde = []string {
	"Tutorial Photoshop untuk pemula", "Review smartphone terbaru", "Tutorial membuat makeup natural", "Vlog liburan ke Bali", "Tips fotografi dengan HP", "Panduan editing video dasar"}

	fmt.Println("Cari ide untuk konten kamu yuk:")
	fmt.Println("(ketik 'selesai' bila sudah selesai)")

	selesai = false
	for !selesai {
		fmt.Print("\nCari ide: ")
		fmt.Scanln(&input)

		if input == "selesai" {
			selesai = true
		} else {
			fmt.Println(" ")
			ditemukan = false

			for i = 0; i < len(daftarIde); i++ {
				ide = daftarIde[i]
				if strings.Contains(ide, input) {
					fmt.Println("-", ide)
					ditemukan = true
				}
			}

			if !ditemukan {
				fmt.Println("Tidak ditemukan ide yang cocok.")
			}
		}
	}
}

func tampilkanKonten(A *tabKonten, n int) {
    var i int
    
    fmt.Println("Daftar Konten:")
    for i = 0; i < n; i++ {
        fmt.Println(i+1, ".", A[i].daftarKonten)
        fmt.Println(A[i].kategori)
        fmt.Println(A[i].tglPosting, A[i].jamPosting)
        fmt.Println("Jumlah interaksi:", A[i].jmlInteraksi)
        fmt.Println(" ")
    }
}

func engagement(){
	var jumlah, i, j int
	var konten [5]struct{
		nama string
		like int
		comment int
		share int
		total int
	}

	fmt.Println("Berapa jumlah konten yang akan diinput engagementnya? ")
	fmt.Scan(&jumlah)

	fmt.Println("masukan judul(string) dan like, comment, share (integer)")
	for i = 0; i < jumlah; i++{
		fmt.Scan(&konten[i].nama, &konten[i].like, &konten[i].comment, &konten[i].share)
		konten[i].total = konten[i].like + konten[i].comment + konten[i].share
	}

	for i = 0; i < jumlah-1; i++ {
		for j = 0; j < jumlah-i-1; j++ {
			if konten[j].total < konten[j+1].total {
				konten[j], konten[j+1] = konten[j+1], konten[j]
            }
        }
    }

	fmt.Println("\nUrutan berdasarkan engagement terbanyak")
	for i = 0; i < jumlah; i++{
		fmt.Println("-", konten[i].nama, ":", konten[i].total)
	}
}

func main() {
	var data tabKonten
	var n int
	var pilihan int

for {
        fmt.Println("\nAplikasi Manajemen Konten")
        fmt.Println("1. Tambah Konten")
        fmt.Println("2. Jadwalkan Konten")
        fmt.Println("3. Cari Ide Konten")
        fmt.Println("4. Tampilkan Semua Konten anda")
		fmt.Println("5. Kelola Engagement video")
        fmt.Println("6. Keluar")
        fmt.Print("Pilihan: ")
        fmt.Scan(&pilihan)

        switch pilihan {
        case 1:
            bacaDaftarKonten(&data, &n)
        case 2:
            jadwalkanKonten(&data, n)
        case 3:
            ideKonten()
        case 4:
            tampilkanKonten(&data, n)
		case 5:
            engagement()
        case 6:
            fmt.Println("Anda berhasil keluar")
            return
        default:
            fmt.Println("PILIH 1 - 6 SAJA!!!")
        }
    }
}