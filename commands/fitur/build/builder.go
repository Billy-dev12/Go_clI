package build

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func HandleBuild(targetOS string, customName string) {
	// Jika user tidak kasih nama, pakai default "bill"
	if customName == "" {
		customName = "bill"
	}

	// 1. Tentukan folder build
	buildFolder := "app_build"
	if _, err := os.Stat(buildFolder); os.IsNotExist(err) {
		fmt.Printf("📂 Membuat folder %s...\n", buildFolder)
		os.MkdirAll(buildFolder, 0755)
	}

	outputName := customName
	if targetOS == "windows" && !strings.HasSuffix(strings.ToLower(outputName), ".exe") {
		outputName += ".exe"
	}

	// 2. Tentukan path output lengkap
	outputPath := filepath.Join(buildFolder, outputName)

	cmd := exec.Command("go", "build", "-ldflags", "-s -w", "-o", outputPath)

	cmd.Env = append(os.Environ(),
		"GOOS="+targetOS,
		"GOARCH=amd64",
	)

	fmt.Printf("🚀 Merakit [%s] ke %s/%s untuk %s...\n", outputName, buildFolder, outputName, targetOS)

	err := cmd.Run()
	if err != nil {
		fmt.Printf("❌ Gagal: %v\n", err)
		return
	}

	fmt.Printf("✅ Selesai! File: %s\n", outputPath)
}

func InteractiveBuild(customName string) {
	fmt.Println("\n🏗️  Pilih Target Platform:")
	fmt.Println("1. Windows (64-bit)")
	fmt.Println("2. Linux (64-bit)")
	fmt.Println("3. Mac (64-bit)")

	fmt.Print("\nMasukkan pilihan (1-3): ")
	var choice string
	fmt.Scanln(&choice)

	targetOS := ""
	switch choice {
	case "1":
		targetOS = "windows"
	case "2":
		targetOS = "linux"
	case "3":
		targetOS = "darwin"
	default:
		fmt.Println("❌ Pilihan tidak valid.")
		return
	}

	HandleBuild(targetOS, customName)
}

func ShowCurrentEnv() {
	fmt.Printf("Kamu sekarang lagi di: %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
