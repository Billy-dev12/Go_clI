package laravel

import (
	"fmt"
	"os/exec"
	"strings"
)

func CekDev() {
	fmt.Println("🔍 Mengecek lingkungan pengembangan Laravel...")

	// 1. Cek PHP
	phpPath, err := exec.LookPath("php")
	if err == nil {
		versionCmd := exec.Command("php", "-v")
		output, _ := versionCmd.Output()
		firstLine := strings.Split(string(output), "\n")[0]
		fmt.Printf("✅ PHP terdeteksi: %s (%s)\n", firstLine, phpPath)
	} else {
		fmt.Println("❌ PHP tidak ditemukan di PATH.")
	}

	// 2. Cek Composer
	composerPath, err := exec.LookPath("composer")
	if err == nil {
		versionCmd := exec.Command("composer", "--version")
		output, _ := versionCmd.Output()
		fmt.Printf("✅ Composer terdeteksi: %s (%s)\n", strings.TrimSpace(string(output)), composerPath)
	} else {
		fmt.Println("❌ Composer tidak ditemukan di PATH.")
	}

	// 3. Cek XAMPP (Biasanya di C:\xampp)
	xamppFound := false
	if _, err := exec.Command("cmd", "/c", "dir", "C:\\xampp").Output(); err == nil {
		fmt.Println("✅ XAMPP terdeteksi di C:\\xampp")
		xamppFound = true
	} else {
		fmt.Println("❓ XAMPP tidak ditemukan di lokasi standar (C:\\xampp).")
	}

	if xamppFound && err == nil {
		fmt.Println("\n👍 Lingkungan kamu sudah siap buat tempur!")
	} else {
		fmt.Println("\n⚠️  Beberapa komponen belum lengkap. Gunakan 'bill setup dev' buat install otomatis.")
	}
}
