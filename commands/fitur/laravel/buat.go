package laravel

import (
	"fmt"
	"os"
	"os/exec"
)

func BuatLaravel(name string) {
	fmt.Printf("🚀 Sedang menyiapkan project Laravel: %s...\n", name)
	fmt.Println("⏳ Mohon tunggu, ini mungkin memakan waktu beberapa menit (pake composer)...")

	// Menjalankan command composer create-project laravel/laravel <name>
	cmd := exec.Command("composer", "create-project", "laravel/laravel", name)

	// Menampilkan output composer ke layar terminal user
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("\n❌ Gagal membuat project: %v\n", err)
		return
	}

	fmt.Printf("\n✅ Selesai! Project '%s' sudah siap digunakan.\n", name)
	fmt.Printf("Ketik 'cd %s' lalu 'php artisan serve' buat mulai.\n", name)
}
