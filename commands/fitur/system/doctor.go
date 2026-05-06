package system

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func CheckEverything() {
	fmt.Println("\n🩺 BILL CLI DOCTOR - Mengecek kesehatan sistem kamu...")
	fmt.Println(strings.Repeat("-", 50))

	missing := []string{}

	// 1. OS & Architecture
	fmt.Printf("💻 OS: %s | Arch: %s\n", runtime.GOOS, runtime.GOARCH)

	// List tools to check
	tools := map[string][]string{
		"Git":      {"git", "--version"},
		"PHP":      {"php", "-v"},
		"Composer": {"composer", "--version"},
		"Node.js":  {"node", "-v"},
		"MySQL":    {"mysql", "--version"},
	}

	for name, args := range tools {
		if !checkTool(name, args[0], args[1]) {
			missing = append(missing, args[0])
		}
	}

	fmt.Println(strings.Repeat("-", 50))
	
	if len(missing) > 0 {
		fmt.Printf("⚠️  Ada %d komponen yang belum terpasang: %s\n", len(missing), strings.Join(missing, ", "))
		fmt.Print("🤔 Mau aku bantu pasang otomatis? (y/n): ")
		var choice string
		fmt.Scanln(&choice)
		if strings.ToLower(choice) == "y" {
			AutoFix(missing)
		}
	} else {
		fmt.Println("✅ Semua komponen terdeteksi! Lingkungan kamu siap tempur. 🚀")
	}
}

func checkTool(name, cmdName, versionArg string) bool {
	path, err := exec.LookPath(cmdName)
	if err == nil {
		version := "Unknown"
		out, _ := exec.Command(cmdName, versionArg).Output()
		if len(out) > 0 {
			version = strings.TrimSpace(strings.Split(string(out), "\n")[0])
		}
		fmt.Printf("✅ %-10s: Terdeteksi (%s) -> %s\n", name, path, version)
		return true
	} else {
		fmt.Printf("❌ %-10s: Tidak ditemukan di PATH.\n", name)
		return false
	}
}

func AutoFix(missing []string) {
	fmt.Println("\n🛠️  Memulai proses perbaikan otomatis...")
	
	switch runtime.GOOS {
	case "linux":
		fixLinux(missing)
	case "windows":
		fixWindows(missing)
	default:
		fmt.Println("❌ OS kamu belum didukung untuk auto-fix.")
	}
}

func fixLinux(missing []string) {
	// Deteksi Distro (Sederhana)
	distro := "unknown"
	if _, err := exec.LookPath("pacman"); err == nil {
		distro = "arch"
	} else if _, err := exec.LookPath("apt"); err == nil {
		distro = "debian"
	}

	fmt.Printf("🐧 Deteksi Distro: %s\n", distro)

	for _, tool := range missing {
		fmt.Printf("📦 Menginstall %s...\n", tool)
		var cmd *exec.Cmd
		
		pkgName := tool
		if tool == "node" {
			pkgName = "nodejs"
		}
		if tool == "mysql" && distro == "arch" {
			pkgName = "mariadb"
		}

		if distro == "arch" {
			cmd = exec.Command("sudo", "pacman", "-S", "--noconfirm", pkgName)
		} else if distro == "debian" {
			cmd = exec.Command("sudo", "apt", "install", "-y", pkgName)
		}

		if cmd != nil {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		}
	}
}

func fixWindows(missing []string) {
	fmt.Println("🪟 Mencoba install via winget...")
	for _, tool := range missing {
		fmt.Printf("📦 Menginstall %s...\n", tool)
		// Map package name to winget ID
		id := tool
		switch tool {
		case "php": id = "PHP.PHP"
		case "composer": id = "Composer.Composer"
		case "git": id = "Git.Git"
		case "node": id = "OpenJS.NodeJS"
		case "mysql": id = "Oracle.MySQL"
		}
		
		cmd := exec.Command("winget", "install", id, "--silent")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
