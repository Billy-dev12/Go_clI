package kode

import (
	"bill/commands/fitur/build"
	"bill/commands/fitur/github"
	fitur "bill/commands/fitur/help"
	"bill/commands/fitur/laravel"
	"bill/commands/fitur/system"
	"fmt"
	"os"
)

func Jalankan() {
	// Jika dipanggil tanpa argumen (bisa jadi Double-Click)
	if len(os.Args) < 2 {
		system.ShowWelcomeInInteractiveMode(prosesPerintah)
		return
	}

	prosesPerintah(os.Args[1:])
}

func prosesPerintah(args []string) {
	if len(args) == 0 {
		return
	}

	var customName string
	if len(args) > 1 {
		customName = args[1]
	}

	subCommand := args[0]

	switch subCommand {
	case "buat":
		if len(args) < 3 || args[1] != "laravel" {
			fmt.Println("    ❌ Format salah. Gunakan: bill buat laravel [nama]")
			return
		}
		laravel.BuatLaravel(args[2])

	case "cek":
		if len(args) < 2 || args[1] != "laravel" {
			fmt.Println("    ❌ Format salah. Gunakan: bill cek laravel")
			return
		}
		laravel.CekDev()

	case "setup":
		if len(args) < 2 {
			fmt.Println("    ❌ Gunakan: bill setup [lingkungan|laravel]")
			return
		}
		if args[1] == "lingkungan" && len(args) > 2 && args[2] == "laravel" {
			laravel.SetupDev()
		} else if args[1] == "laravel" {
			laravel.SetupProject()
		} else {
			fmt.Println("    ❌ Perintah tidak dikenal. Gunakan: bill setup lingkungan laravel ATAU bill setup laravel")
		}

	case "ser":
		port := ""
		if len(args) > 1 {
			port = args[1]
		}
		laravel.Serve(port)

	case "cleanup":
		laravel.Cleanup()

	case "push":
		github.PushToGithub(args[1:])

	case "arch":
		if len(args) < 2 || args[1] != "push" {
			fmt.Println("    ❌ Format salah. Gunakan: bill arch push [pesan]")
			return
		}
		github.PushToGithub(args[2:])

	case "install":
		system.Install()

	case "help":
		fitur.PrintHelp()

	case "exit", "keluar":
		fmt.Println("    👋 Sampai jumpa!")
		os.Exit(0)

	case "build-win":
		build.HandleBuild("windows", customName)
	case "build-linux":
		build.HandleBuild("linux", customName)
	case "build-mac":
		build.HandleBuild("darwin", customName)
	case "go":
		if len(args) < 2 || args[1] != "build" {
			fmt.Println("    ❌ Format salah. Gunakan: bill go build [nama]")
			return
		}
		name := ""
		if len(args) > 2 {
			name = args[2]
		}
		build.InteractiveBuild(name)

	case "info":
		build.ShowCurrentEnv()

	default:
		fmt.Printf("    ❌ Perintah '%s' tidak dikenali.\n", subCommand)
		fitur.PrintHelp()
	}
}
