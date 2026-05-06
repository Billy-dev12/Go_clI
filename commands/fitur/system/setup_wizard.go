package system

import (
	"bill/commands/fitur/version"
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// ======================================
//        🎨 ASCII ART & BANNER
// ======================================

func printRobot() {
	// Dihapus sesuai permintaan: no more design aneh
}

func printInstalledBanner() {
	fmt.Println("\n✅ INSTALASI BERHASIL!")
	fmt.Println("Billy CLI sudah terpasang di sistemmu.")
	fmt.Println("Silakan buka terminal baru untuk mulai menggunakan dari mana saja.")
}

func printLoading(text string) {
	fmt.Printf("> %s...\n", text)
}

func PrintHelpBox() {
	fmt.Println("\n📖 DAFTAR PERINTAH:")
	fmt.Println("  bill buat laravel [nama]  -> Buat project Laravel")
	fmt.Println("  bill push                 -> Push ke GitHub")
	fmt.Println("  bill build                -> Build project")
	fmt.Println("  bill help                 -> Menu bantuan")
	fmt.Println("  bill install              -> Pasang ke sistem")
	fmt.Println("\n🌐 Info: https://github.com/Billy-dev12")
}

// IsInstalled mengecek apakah aplikasi sudah terpasang
func IsInstalled() bool {
	exePath, err := os.Executable()
	if err != nil {
		return false
	}
	exeName := filepath.Base(exePath)
	cmdName := strings.TrimSuffix(exeName, filepath.Ext(exeName))

	// 1. Cek apakah ada di PATH
	if _, err := exec.LookPath(cmdName); err == nil {
		// Pastikan bukan sedang jalan dari folder build (agar tetap bisa install)
		if strings.Contains(exePath, "app_build") {
			return false
		}
		return true
	}

	return false
}

// ShowWelcomeInInteractiveMode menampilkan menu selamat datang dan loop petunjuk
func ShowWelcomeInInteractiveMode(callback func([]string)) {
	if !IsInstalled() {
		fmt.Println("\n🚀 Bill CLI belum terpasang di sistem kamu!")
		fmt.Println("Ketik 'install' untuk memasang agar bisa dipanggil dari mana saja.")
	}

	PrintHelpBox()
	fmt.Println("\n🩺 Tip: Ketik 'info' untuk mengecek kesehatan lingkungan (PHP, Git, dll).")

	fmt.Println("\n💡 Ketik perintah (Contoh: 'help' atau 'exit' untuk keluar).")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("bill > ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		args := strings.Fields(input)
		callback(args)
	}
}

// ======================================
//        🔧 LOGIKA INSTALASI
// ======================================

// Install mendeteksi OS dan melakukan instalasi yang sesuai
func Install() {
	fmt.Printf("\n🛠️  Memulai instalasi modular %s...\n", version.Version)

	switch runtime.GOOS {
	case "windows":
		installWindows()
	case "linux", "darwin":
		installUnix()
	default:
		fmt.Printf("❌ OS %s belum didukung untuk instalasi otomatis.\n", runtime.GOOS)
	}
}

func installWindows() {
	newExePath, _ := os.Executable()

	// Paksa nama target selalu 'bill.exe' agar perintahnya selalu 'bill'
	// Gimanapun nama file build-nya (misal: jawa.exe), pas di-install jadi bill.exe
	exeName := "bill.exe"
	cmdName := "bill"

	localAppData := os.Getenv("LOCALAPPDATA")
	targetDir := filepath.Join(localAppData, "bill-tool")
	targetPath := filepath.Join(targetDir,
		exeName)

	// 0. Proteksi: Jangan sampai menghapus diri sendiri kalau dijalankan dari folder target
	if strings.Contains(strings.ToLower(newExePath), "bill-tool") {
		fmt.Printf("\n⚠️  Info: Kamu menjalankan %s langsung dari folder instalasi.\n", cmdName)
		fmt.Println("💡 Untuk update, silakan jalankan file hasil build yang baru (bukan yang di PATH).")
		return
	}

	fmt.Println("\n🔍 Mengecek instalasi lama...")

	// 1. Logika Pembersihan Folder (Trik Rename untuk bypass Locking)
	if _, err := os.Stat(targetDir); err == nil {
		printLoading("Menyingkirkan versi lama")
		tempOldDir := targetDir + "_old_" + fmt.Sprint(os.Getpid())

		// Ganti nama folder lama (Windows biasanya izinkan rename meski file di dalam dikunci)
		errRename := os.Rename(targetDir, tempOldDir)
		if errRename != nil {
			fmt.Printf("❌ Gagal update: %v\n", errRename)
			fmt.Println("💡 Tip: Pastikan tidak ada terminal lain yang sedang menjalankan perintah ini.")
			return
		}
		// Hapus folder lama yang sudah di-rename di background
		go os.RemoveAll(tempOldDir)
	}

	// 2. Buat ulang folder target & Salin binary
	os.MkdirAll(targetDir, 0755)
	printLoading("Menyalin binary terbaru ke lokasi sistem")
	if err := copyFile(newExePath, targetPath); err != nil {
		fmt.Printf("❌ Gagal menyalin binary: %v\n", err)
		return
	}

	// 3. Logika PATH Sakti (PowerShell dengan ExecutionPolicy Bypass)
	printLoading("Memperbarui PATH Environment Windows")

	setupScript := fmt.Sprintf(`
		$target = "%s"
		$oldPath = [Environment]::GetEnvironmentVariable("Path", "User")
		
		# Bersihkan path lama jika ada agar tidak duplikat
		$pathList = $oldPath.Split(';', [System.StringSplitOptions]::RemoveEmptyEntries) | Where-Object { $_.TrimEnd('\') -ne $target.TrimEnd('\') }
		
		# Tambahkan ke AWAL path supaya jadi prioritas utama
		$newPath = ($target + ";" + ($pathList -join ';'))
		
		# Simpan ke Registry User
		[Environment]::SetEnvironmentVariable("Path", $newPath, "User")
		
		# Broadcast sinyal perubahan ke Windows (SendMessage)
		$definition = '[DllImport("user32.dll", SetLastError = true, CharSet = CharSet.Auto)] public static extern IntPtr SendMessageTimeout(IntPtr hWnd, uint Msg, UIntPtr wParam, string lParam, uint fuFlags, uint uTimeout, out UIntPtr lpdwResult);'
		$type = Add-Type -MemberDefinition $definition -Name "Win32" -Namespace "Native" -PassThru
		$type::SendMessageTimeout(0xffff, 0x001A, [UIntPtr]::Zero, "Environment", 0x02, 1000, [ref][UIntPtr]::Zero)
	`, targetDir)

	// Jalankan PowerShell dengan Bypass agar tidak diblokir di VM/Laptop temen
	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-Command", setupScript)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("⚠️  Gagal mendaftarkan PATH: %s\n", string(output))
	} else {
		fmt.Println("✅ PATH sistem berhasil diperbarui.")
	}

	fmt.Printf("\n🚀 %s BERHASIL DI-INSTALL/UPDATE!\n", strings.ToUpper(cmdName))
	fmt.Println("---------------------------------------------------------")
	fmt.Println("💡 PENTING: TUTUP TERMINAL INI DAN BUKA TERMINAL BARU")
	fmt.Printf("💡 Ketik '%s' untuk mulai menggunakan.\n", cmdName)
	fmt.Println("---------------------------------------------------------")
}

func installUnix() {
	exePath, _ := os.Executable()
	exeName := filepath.Base(exePath)
	home, _ := os.UserHomeDir()

	// 1. Bersihkan versi lama di folder sistem jika ada (misal hasil install sudo sebelumnya)
	systemPaths := []string{"/usr/bin/" + exeName, "/usr/local/bin/" + exeName}
	for _, p := range systemPaths {
		if _, err := os.Stat(p); err == nil {
			printLoading("Menghapus versi sistem lama (" + p + ")")
			// Coba hapus, jika gagal (permission denied) biarkan saja tapi beri info
			err := os.Remove(p)
			if err != nil {
				fmt.Printf("⚠️  Bisa jadi ada versi lama di %s yang tidak bisa dihapus (butuh sudo).\n", p)
			}
		}
	}

	// 2. Gunakan ~/.local/bin sebagai standar user binary di Linux/Mac
	targetDir := filepath.Join(home, ".local", "bin")
	targetPath := filepath.Join(targetDir, exeName)

	printLoading("Mengecek folder ~/.local/bin")
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		os.MkdirAll(targetDir, 0755)
	}

	// Cek apakah file target sudah ada, jika ya hapus dulu agar bersih
	if _, err := os.Stat(targetPath); err == nil {
		printLoading("Menghapus versi user lama")
		os.Remove(targetPath)
	}

	printLoading("Menyalin binary terbaru")
	if err := copyFile(exePath, targetPath); err != nil {
		fmt.Printf("❌ Gagal menyalin: %v\n", err)
		return
	}

	// Beri izin eksekusi
	exec.Command("chmod", "+x", targetPath).Run()

	printInstalledBanner()

	// Berikan saran jika path belum terdaftar
	if !strings.Contains(os.Getenv("PATH"), targetDir) {
		fmt.Printf("\n💡 Tip: Pastikan %s ada di PATH kamu.\n", targetDir)
		fmt.Println("   Tambahkan baris berikut ke .bashrc atau .zshrc:")
		fmt.Printf("   export PATH=\"$PATH:%s\"\n", targetDir)
	}
}

func Pause() {
	fmt.Println("    Tekan Enter untuk keluar...")
	fmt.Scanln()
}

func equalsIgnoreCase(a, b string) bool {
	return len(a) == len(b) && filepath.Clean(a) == filepath.Clean(b)
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
