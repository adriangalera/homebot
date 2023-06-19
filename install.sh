TMP_FOLDER=/tmp/homebot
LOCAL_PATH="${HOME}/.local/bin"
has_installed() {
  type "$1" >/dev/null 2>&1
}
check_requirements() {
  if ! has_installed "git"; then
    echo "Please install git"
    exit 1
  fi

  if ! has_installed "go"; then
    echo "Please install golang from https://go.dev/dl"
    exit 1
  fi

  if ! has_installed "make"; then
    echo "Please install make"
    exit 1
  fi
}
download_homebot() {
  rm -rf $TMP_FOLDER || exit 0
  if ! git clone --quiet https://github.com/adriangalera/homebot -b install-script $TMP_FOLDER; then
    echo "Can't clone repository"
    exit 1
  fi
}

build() {
  PLATFORM=$(arch)
  VERSION=$(git describe --tags)
  PLATFORM=$PLATFORM VERSION=$VERSION make clean build >$TMP_FOLDER/install.log
  cp "build/homebot-${PLATFORM}-${VERSION}" "$LOCAL_PATH/homebot"
}

service() {
  sudo cp $TMP_FOLDER/homebot.service /etc/systemd/system/homebot.service
  sudo chmod 644 /etc/systemd/system/homebot.service
  sudo chown root:root /etc/systemd/system/homebot.service
  sudo systemctl enable homebot
  sudo systemctl start homebot
}

ensure_folders_and_files() {
  mkdir -p "$LOCAL_PATH"
  sudo mkdir -p /etc/homebot
  sudo mkdir -p /etc/homebot/commands
  sudo cp $TMP_FOLDER/internal/config/examples/empty-config.yml /etc/homebot/config.yml
}
install() {
  check_requirements
  ensure_folders_and_files
  echo "Downloading ..."
  download_homebot
  cd $TMP_FOLDER || exit
  echo "Building ..."
  build
  service
  echo "Service installed, check if it's running fine with:"
  echo "journalctl -u homebot -f"
  echo "If you need to restart it:"
  echo "sudo systemctl restart homebot"
}

install
