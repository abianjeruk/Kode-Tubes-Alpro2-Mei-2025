package main

import("fmt";"strings")

const nmax int = 10

type konten struct {
	IDkonten string
	daftarKonten string
	tglPosting string
	kategori string
	jamPosting string
	jmlInteraksi int
}

type tabKonten [nmax]konten

func inputID(A *tabKonten, index int) {
    var id string
    var idValid bool = false
	var i int
    
    for !idValid {
        fmt.Scan(&id)
		idValid = true
		for i = 0; i < index; i++ {
			if A[i].IDkonten == id {
				fmt.Println("ID konten tidak boleh sama")
				idValid = false
			}
		}
	}
    
    A[index].IDkonten = id
}

func hapusdata(A *tabKonten, n *int) {
	var idx, i int
	
	fmt.Println("Masukkan indeks data yang ingin dihapus (0 sampai", *n-1, "):")
	fmt.Scan(&idx)
	
	if idx < 0 || idx >= *n {
		fmt.Println("indeks tidak ada")
		return
	}
	
	for i = idx; i < *n-1; i++ {
		A[i] = A[i + 1]
	}
	*n = *n - 1
	fmt.Println("Data dihapus")
}

func editdata(A *tabKonten, n int){
	var idx int
	
	fmt.Println("Masukkan indeks data yang ingin diedit (0 sampai", n-1, "):")
	fmt.Scan(&idx)
	
	if idx < 0 || idx >= n {
		fmt.Println("indeks tidak ada")
		return
	}
	fmt.Println("\nEdit Data Konten", idx)
	fmt.Println("ID konten (sebelumnya:", A[idx].IDkonten, "):")
	fmt.Scan(&A[idx].IDkonten)
	fmt.Println("Judul konten (sebelumnya:", A[idx].daftarKonten, "):")
	fmt.Scan(&A[idx].daftarKonten)
	fmt.Print("Kategori (sebelumnya:", A[idx].kategori, "): ")
	fmt.Scan(&A[idx].kategori)
	fmt.Println("Berhasil diedit")
}

func bacaDaftarKonten(A *tabKonten, n *int) {
	var i, pilihan  int
	
	fmt.Println("Mau buat berapa konten? (max = 10)")
	fmt.Scan(n)
	
	if *n > 10 {
		fmt.Println("TIDAK BISA LEBIH DARI 10")
		*n = 0 
		return
	}
	for i = 0; i < *n; i++ {
		fmt.Println("ID konten")
		inputID(A, i)
		
		fmt.Println("\nJudul konten (jangan menggunakan spasi)", i+1, ":")
		fmt.Scan(&A[i].daftarKonten)
		fmt.Print("Kategori: ")
		fmt.Scan(&A[i].kategori)
		A[i].jmlInteraksi = 0
	}
	
	for {
		fmt.Println("\nMenu")
		fmt.Println("1. Hapus data")
		fmt.Println("2. Edit data")
		fmt.Println("3. Keluar")
		fmt.Println("Pilihan: ")
		fmt.Scan(&pilihan)
		
		if pilihan == 1 {
			hapusdata(A, n)
		} else if pilihan == 2 {
			editdata(A, *n)
		} else if pilihan == 3 {
			return
		} else {
			fmt.Println("HANYA PILIH 1 - 3")
		}
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

func binarySearch(A *tabKonten, n int){
	var idx, left, right, mid int
	var found bool = false
	var id, judul, kategori, tgl, jam string
	var interaksi int
	
	left = 0
	right = n - 1
	
	fmt.Println("Masukan indeks konten yang telah diinput sebelumnya", n-1, ":")
	fmt.Scan(&idx)
	
	if idx < 0 || idx >= n {
		fmt.Println("indeks tidak ada")
		return
	}
	
	for left <= right && !found {
		mid = (left + right) / 2
		if mid == idx {
			found = true
			id = A[mid].IDkonten
			judul = A[mid].daftarKonten
			kategori = A[mid].tglPosting
			jam = A[mid].jamPosting
			interaksi = A[mid].jmlInteraksi
		} else if mid > idx {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if found {
		fmt.Println("\nDetail Konten:")
		fmt.Println("ID Konten:", id)
		fmt.Println("Judul Konten:", judul)
		fmt.Println("Kategori:", kategori)
		fmt.Println("Tanggal Posting:", tgl)
		fmt.Println("Jam Posting:", jam)
		fmt.Println("Jumlah Interaksi:", interaksi)
    }
}

func tampilkanKonten(A *tabKonten, n int) {
    var i int
    
    fmt.Println("Daftar Konten:")
	
    for i = 0; i < n; i++ {
        fmt.Println(i+1, ".", A[i].daftarKonten)
        fmt.Println("kategori: ", A[i].kategori)
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
		fmt.Println("3. Cari Ide Konten (sequential search)")
		fmt.Println("4. Cari konten (binary search)")
		fmt.Println("5. Tampilkan Semua Konten anda")
		fmt.Println("6. Kelola Engagement video")
		fmt.Println("7. Keluar")
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
			binarySearch(&data, n)
        case 5:
            tampilkanKonten(&data, n)
		case 6:
            engagement()
        case 7:
            fmt.Println("Anda berhasil keluar")
            return
        default:
            fmt.Println("PILIH 1 - 7 SAJA!!!")
        }
    }
}
