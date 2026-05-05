#!/bin/bash

# =================================================================
#  Bill CLI Installer for Linux
#  Repository: https://github.com/Billy-dev12/Go_clI
# =================================================================

set -e

REPO="Billy-dev12/Go_clI"
BINARY_NAME="bill"
INSTALL_DIR="$HOME/.local/bin"

# Detect OS and Architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$ARCH" == "x86_64" ]]; then
    ARCH="amd64"
elif [[ "$ARCH" == "aarch64" || "$ARCH" == "arm64" ]]; then
    ARCH="arm64"
else
    echo "❌ Architecture $ARCH not supported yet."
    exit 1
fi

if [[ "$OS" != "linux" && "$OS" != "darwin" ]]; then
    echo "❌ OS $OS not supported by this script. Use the Windows installer for Windows."
    exit 1
fi

echo "🚀 Finding latest version of Bill CLI..."

# Get latest release tag from GitHub API
LATEST_RELEASE=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_RELEASE" ]; then
    echo "❌ Could not find latest release. Please check the repository."
    exit 1
fi

echo "📦 Latest version: $LATEST_RELEASE"
echo "🛠️ Platform: $OS-$ARCH"

# Determine download URL (adjusting to common release naming pattern)
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/bill-$OS-$ARCH"

# If it's macOS, it might be named differently or just 'bill-darwin-amd64'
if [[ "$OS" == "darwin" ]]; then
    DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/bill-darwin-$ARCH"
fi

# Create install directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

echo "📥 Downloading binary..."
curl -L -o "$INSTALL_DIR/$BINARY_NAME" "$DOWNLOAD_URL"

# Make it executable
chmod +x "$INSTALL_DIR/$BINARY_NAME"

echo ""
echo "✅ Bill CLI $LATEST_RELEASE has been installed to $INSTALL_DIR/$BINARY_NAME"

# Check if INSTALL_DIR is in PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo ""
    echo "⚠️  $INSTALL_DIR is not in your PATH."
    echo "   To fix this, add the following line to your ~/.bashrc or ~/.zshrc:"
    echo "   export PATH=\"\$PATH:$INSTALL_DIR\""
    echo ""
    echo "   Then restart your terminal or run: source ~/.bashrc (or ~/.zshrc)"
fi

echo "🚀 Try running: bill help"
