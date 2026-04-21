package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct Item ini buat representasi info satu barang di gudang ya
type Item struct {
	ID    int
	Name  string
	Price int
	Stock int
}

// Variabel global buat nyimpen data stok gudang kita nih
var inventory []Item
var nextID = 1

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== TOKO KELONTONG INVENTORY MANAGER ===")
		fmt.Println("\n1. Tambah Barang ke Gudang")
		fmt.Println("2. Lihat Semua Stok Barang")
		fmt.Println("3. Beli Barang")
		fmt.Println("4. Keluar")
		fmt.Print("\nPilih Menu (1-4): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			tambahBarang(reader)
		case "2":
			lihatStok()
		case "3":
			beliBarang(reader)
		case "4":
			fmt.Println("[SISTEM]: Program selesai. Selamat beristirahat, Juragan!")
			return
		default:
			fmt.Println("[SISTEM]: Pilihan tidak valid, silakan coba lagi!")
		}
	}
}

// Fungsi buat nambahin barang baru ke dalam gudang
func tambahBarang(reader *bufio.Reader) {
	fmt.Print("\nMasukkan Nama Barang: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Harga: ")
	hargaStr, _ := reader.ReadString('\n')
	hargaStr = strings.TrimSpace(hargaStr)
	harga, err := strconv.Atoi(hargaStr)
	if err != nil {
		fmt.Println("[SISTEM]: Harga harus berupa angka!")
		return
	}

	fmt.Print("Masukkan Stok Awal: ")
	stokStr, _ := reader.ReadString('\n')
	stokStr = strings.TrimSpace(stokStr)
	stok, err := strconv.Atoi(stokStr)
	if err != nil {
		fmt.Println("[SISTEM]: Stok harus berupa angka!")
		return
	}

	// Bikin item baru terus masukin deh ke dalam slice inventory
	item := Item{
		ID:    nextID,
		Name:  nama,
		Price: harga,
		Stock: stok,
	}
	inventory = append(inventory, item)
	nextID++

	fmt.Println("[SISTEM]: Barang berhasil ditambahkan ke gudang!")
	fmt.Println("\n--------------------------------------------------")
}

// Fungsi buat ngeliatin semua barang yang lagi ada di stok gudang
func lihatStok() {
	fmt.Println("\n=== DAFTAR STOK GUDANG ===")
	if len(inventory) == 0 {
		fmt.Println("Gudang kosong!")
	} else {
		for _, item := range inventory {
			fmt.Printf("ID: %d | Nama: %-15s | Harga: Rp %d | Stok: %d pcs\n", item.ID, item.Name, item.Price, item.Stock)
		}
	}
	fmt.Println("========================================")
	fmt.Printf("Total Jenis Barang: %d\n", len(inventory))
	fmt.Println("\n--------------------------------------------------")
}

// Fungsi buat ngurusin alur pas ada yang mau beli barang
func beliBarang(reader *bufio.Reader) {
	if len(inventory) == 0 {
		fmt.Println("\n[SISTEM]: Gudang kosong, tidak ada barang untuk dibeli!")
		return
	}

	fmt.Println("\n=== MENU PEMBELIAN ===")
	for _, item := range inventory {
		fmt.Printf("ID: %d | Nama: %-15s | Stok: %d | Harga: Rp %d\n", item.ID, item.Name, item.Stock, item.Price)
	}

	fmt.Print("\nPilih ID Barang yang mau dibeli: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("[SISTEM]: ID harus berupa angka!")
		return
	}

	fmt.Print("Jumlah yang mau dibeli: ")
	jumlahStr, _ := reader.ReadString('\n')
	jumlahStr = strings.TrimSpace(jumlahStr)
	jumlah, err := strconv.Atoi(jumlahStr)
	if err != nil {
		fmt.Println("[SISTEM]: Jumlah harus berupa angka!")
		return
	}

	// Cari dulu nih barangnya ada apa nggak
	var selectedItem *Item
	var itemIndex int
	for i := range inventory {
		if inventory[i].ID == id {
			selectedItem = &inventory[i]
			itemIndex = i
			break
		}
	}

	if selectedItem == nil {
		fmt.Println("[SISTEM]: Barang dengan ID tersebut tidak ditemukan!")
		return
	}

	if selectedItem.Stock < jumlah {
		fmt.Println("[SISTEM]: Stok tidak mencukupi untuk jumlah pembelian tersebut!")
		return
	}

	totalHarga := selectedItem.Price * jumlah
	fmt.Printf("\n[SISTEM]: Total Harga: Rp %d\n", totalHarga)
	
	fmt.Print("Masukkan Uang Anda: ")
	uangStr, _ := reader.ReadString('\n')
	uangStr = strings.TrimSpace(uangStr)
	uang, err := strconv.Atoi(uangStr)
	if err != nil {
		fmt.Println("[SISTEM]: Uang harus berupa angka!")
		return
	}

	if uang < totalHarga {
		fmt.Println("[SISTEM]: Uang Anda tidak mencukupi untuk transaksi ini!")
		return
	}

	kembalian := uang - totalHarga
	
	// Kurangin stok barang yang ada di dalem slice
	inventory[itemIndex].Stock -= jumlah

	fmt.Println("\n[SISTEM]: Transaksi Berhasil!")
	fmt.Printf("[SISTEM]: Kembalian Anda: Rp %d\n", kembalian)
	fmt.Printf("[SISTEM]: Stok %s sekarang: %d pcs\n", selectedItem.Name, inventory[itemIndex].Stock)
	fmt.Println("\n--------------------------------------------------")
}
