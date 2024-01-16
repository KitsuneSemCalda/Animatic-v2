#!/usr/bin/env bash

check_dir_owner() {
  dir=$1
  echo $(ls -ld $dir | awk '{print $3}')
}

print_message() {
  local color_code="$1"
  local text="$2"
  echo -e "\033[0;${color_code}m${text}\033[0m"
}

print_green() {
  print_message "32" "$1"
}

print_red() {
  print_message "31" "$1"
}

check_directory() {
  dir=$1
  if [ ! -d "$dir" ]; then
    print_red "$dir was not found, creating it"
    mkdir -p "$dir"
  else
    print_green "[OK]: $dir was found"
  fi
}

check_owner() {
  dir=$1
  if [ ! "$USER_DEFAULT" == $(check_dir_owner "$dir") ]; then
    print_red "Fixing permissions"
    sudo chown "$USER_DEFAULT" "$dir"/*
  else
    print_green "[OK]: The permissions of $dir are correct"
  fi
}

BASE_DIR="/chromeMedia"
USER_DEFAULT=$USER
LOCAL_BIN="$HOME/.local/bin"

if [ ! command -v go &> /dev/null ]; then
  print_red "Go is not installed"
  exit 1
else
  print_green "[OK]: Go is installed"
fi

check_directory "$LOCAL_BIN"

if echo $PATH | grep -qv "$LOCAL_BIN"; then
  print_red "$LOCAL_BIN isn't found in PATH"
  echo "export PATH=$PATH:$LOCAL_BIN" >> $HOME/.profile
else
  print_green "[OK]: $LOCAL_BIN was found in PATH"
fi

GO_PATH=$(which go)
GO_VERSION=$($GO_PATH version | awk '{print $3}' | tr -d "go")

if [ "$GO_VERSION" != "1.21.5" ]; then
  print_red "Go 1.21.5 is not installed, we only have version $GO_VERSION"
  exit 1
else
  print_green "[OK]: Go 1.21.5 is installed"
fi

$($GO_PATH build)

mv "animatic-v2" "$LOCAL_BIN/Animatic"

check_directory "$BASE_DIR"
check_owner "$BASE_DIR"
