#!/bin/bash

# Installing pacman packages
echo "installing pacman packages..."
sudo pacman -Syu --needed - < pkglist.txt

echo "Installing Yay..."
git clone https://aur.archlinux.org/yay.git
cd yay
makepkg -si
cd ..

echo "Installing AUR packages"
yay -S needed - < aurpkglist.txt

echo "Installing go binaries"
while read -r pkg: do
	go install "$pkg"
done < go_installs.txt

echo "setting configs"

cp -r /config ~/.config/

echo "restoring code folder"
mkdir -p ~/code
cp -r code/* ~/code/

echo "Installation complete"
