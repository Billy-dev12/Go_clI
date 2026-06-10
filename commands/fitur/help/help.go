package fitur

import (
	"bill/commands/fitur/version"
	"fmt"
	"strings"
)

func PrintHelp() {
	bold := "\033[1m"
	reset := "\033[0m"
	cyan := "\033[36m"
	yellow := "\033[33m"
	green := "\033[32m"

	fmt.Printf("\n%s🚀 BILL CLI %s - Asisten CLI Terkeren kamu!%s\n", bold+cyan, version.Version, reset)
	fmt.Println(strings.Repeat("-", 50))

	fmt.Printf("\n%s🐙 GITHUB & GIT%s\n", yellow, reset)
	fmt.Println("  bill push                -> Smart push (auto init/commit/sync)")
	fmt.Println("  bill push release [tag]  -> Buat git tag lokal & GitHub Release")
	fmt.Println("  bill branch              -> Manage git branches (list/switch)")
	fmt.Println("  bill arch push           -> Smart push (khusus pengguna Arch)")
	fmt.Println("  bill push delete         -> Reset konfigurasi Git (.git)")

	fmt.Printf("\n%s🔨 LARAVEL%s\n", yellow, reset)
	fmt.Println("  bill buat laravel [nama] -> Buat project Laravel baru")
	fmt.Println("  bill cek laravel          -> Cek status environment")
	fmt.Println("  bill setup laravel        -> Auto-setup repo hasil clone")
	fmt.Println("  bill setup lingkungan laravel -> Setup dev env (PHP, Composer, dll)")
	fmt.Println("  bill ser [port]           -> Jalankan server artisan")
	fmt.Println("  bill cleanup              -> Bersihkan cache & log Laravel")

	fmt.Printf("\n%s📦 BUILD & SYSTEM%s\n", yellow, reset)
	fmt.Println("  bill go build [nama]     -> Build project Go (Interaktif)")
	fmt.Println("  bill info                -> Tampilkan info environment")
	fmt.Println("  bill help                -> Tampilkan menu bantuan ini")
	fmt.Println("  bill exit / keluar       -> Keluar dari mode interaktif")

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Printf("%s✨ Bill CLI %s %s| Gunakan tanpa argumen untuk Mode Interaktif\n", green, version.Version, reset)
	fmt.Printf("%s🌐 Info & Update: https://github.com/Billy-dev12/Go_clI%s\n\n", cyan, reset)
}
