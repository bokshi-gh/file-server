#!/usr/bin/env bash
set -e  # Exit immediately if a command fails

REPO="https://github.com/bokshi-gh/file-server.git"
TMP_DIR="/tmp/file-server-build"
EXEC_NAME="goserve"

rm -rf "$TMP_DIR"

git clone "$REPO" "$TMP_DIR"

cd "$TMP_DIR/cmd/goserve"
go build -o "$EXEC_NAME"

sudo mv "$EXEC_NAME" /usr/local/bin/

cd ~
rm -rf "$TMP_DIR"

echo "Build complete! You can now run '$EXEC_NAME' from anywhere."
