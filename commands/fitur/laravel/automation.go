package laravel

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// SetupProject menangani inisialisasi project Laravel setelah clone
func SetupProject() {
	fmt.Println("\n🚀 Memulai Inisialisasi Project Laravel...")

	// 1. Cek .env.example -> .env
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		if _, errEx := os.Stat(".env.example"); errEx == nil {
			fmt.Println("📄 Menyalin .env.example ke .env...")
			copyFile(".env.example", ".env")
		} else {
			fmt.Println("⚠️  .env.example tidak ditemukan. Lewati langkah ini.")
		}
	} else {
		fmt.Println("✅ File .env sudah ada.")
	}

	// 2. Composer Install
	fmt.Println("\n📦 Menjalankan composer install...")
	runCommand("composer", "install")

	// 3. Generate Key
	fmt.Println("\n🔑 Menghasilkan application key...")
	runCommand("php", "artisan", "key:generate")

	// 4. Smart Database Creation (Opsional & Eksperimental)
	setupDatabase()

	// 5. Migrate
	fmt.Println("\n🗄️  Menjalankan migrasi database...")
	runCommand("php", "artisan", "migrate")

	fmt.Println("\n✨ Project Laravel siap digunakan! Ketik 'bill ser' untuk menyalakan server.")
}

// Serve menyalakan server Laravel dan membuka browser
func Serve(port string) {
	if port == "" {
		port = "8000"
	}

	url := fmt.Sprintf("http://localhost:%s", port)

	fmt.Printf("\n🚀 Menyalakan server di %s...\n", url)

	// Jalankan perintah di background agar tidak memblokir pembukaan browser (atau sebaliknya)
	cmd := exec.Command("php", "artisan", "serve", "--port="+port)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Buka browser setelah jeda sebentar biar server nyala dulu
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Printf("🌐 Membuka browser ke %s...\n", url)
		openBrowser(url)
	}()

	err := cmd.Run()
	if err != nil {
		fmt.Printf("❌ Server berhenti dengan error: %v\n", err)
	}
}

// Cleanup membersihkan cache dan log Laravel
func Cleanup() {
	fmt.Println("\n🧹 Membersihkan sampah Laravel...")

	commands := [][]string{
		{"php", "artisan", "optimize:clear"},
		{"php", "artisan", "view:clear"},
		{"php", "artisan", "config:clear"},
	}

	for _, c := range commands {
		fmt.Printf("⚙️  Menjalankan: %s %s...\n", c[0], c[1])
		runCommand(c[0], c[1:]...)
	}

	// Hapus isi storage/logs
	logPath := filepath.Join("storage", "logs")
	files, err := os.ReadDir(logPath)
	if err == nil {
		fmt.Println("📁 Menghapus isi storage/logs...")
		for _, f := range files {
			if f.Name() != ".gitignore" {
				os.Remove(filepath.Join(logPath, f.Name()))
			}
		}
	}

	fmt.Println("✅ Bersih-bersih selesai!")
}

// Helper: Jalankan perintah shell
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Helper: Copy file sederhana
func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

// Helper: Buka browser sesuai OS
func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		fmt.Printf("⚠️  Gagal membuka browser: %v\n", err)
	}
}

// Helper: Cek .env dan coba buat database jika pakai MySQL
func setupDatabase() {
	file, err := os.Open(".env")
	if err != nil {
		return
	}
	defer file.Close()

	var dbName, dbUser, dbPass, dbHost string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "DB_DATABASE=") {
			dbName = strings.TrimPrefix(line, "DB_DATABASE=")
		}
		if strings.HasPrefix(line, "DB_USERNAME=") {
			dbUser = strings.TrimPrefix(line, "DB_USERNAME=")
		}
		if strings.HasPrefix(line, "DB_PASSWORD=") {
			dbPass = strings.TrimPrefix(line, "DB_PASSWORD=")
		}
		if strings.HasPrefix(line, "DB_HOST=") {
			dbHost = strings.TrimPrefix(line, "DB_HOST=")
		}
	}

	// Bersihkan quotes jika ada
	re := regexp.MustCompile(`^"(.+)"$`)
	dbName = re.ReplaceAllString(dbName, "$1")

	if dbName != "" {
		fmt.Printf("\n🗄️  Mendeteksi database: %s\n", dbName)
		fmt.Println("❓ Apakah kamu ingin saya mencoba membuat database ini di MySQL? (y/n)")

		var answer string
		fmt.Print("   > ")
		fmt.Scanln(&answer)

		if strings.ToLower(answer) == "y" {
			// Coba konek ke mysql tanpa password dulu
			createDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName)
			args := []string{"-h", dbHost, "-u", dbUser, "-e", createDB}
			if dbPass != "" {
				args = append(args, "-p"+dbPass)
			}

			cmd := exec.Command("mysql", args...)
			err := cmd.Run()
			if err != nil {
				fmt.Printf("⚠️  Gagal membuat database otomatis: %v\n", err)
				fmt.Println("   Pastikan MySQL sudah jalan dan perintah 'mysql' ada di PATH.")
			} else {
				fmt.Println("✅ Database berhasil dibuat/sudah ada!")
			}
		}
	}
}
