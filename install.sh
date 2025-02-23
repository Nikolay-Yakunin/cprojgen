# chmod +x install.sh

# if [ "$EUID" -ne 0 ]; then
#   echo "Please run as root"
#   exit
# fi

BINARY_NAME="cprojgen" # Change this to the name of your binary (if you want)
DEST_DIR="/usr/local/bin"

echo "Building the binary..."
make build-linux

if [ ! -f "build/$BINARY_NAME" ]; then
  echo "Failed to build the binary."
  exit 1
fi

echo "Copying binary to $DEST_DIR..."
cp "build/$BINARY_NAME" "$DEST_DIR/"

if [ $? -eq 0 ]; then
  echo "Binary successfully installed to $DEST_DIR"
else
  echo "Failed to copy binary to $DEST_DIR"
  exit 1
fi