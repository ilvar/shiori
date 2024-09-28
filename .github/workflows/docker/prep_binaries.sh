#!/bin/sh

rm -rf ./binaries
mkdir binaries
cp -r ../../../dist/shiori_linux_* binaries/
#cp -r ../../../shiori_linux_* binaries/
mv binaries/shiori_linux_arm_7 binaries/shiori_linux_arm
mv binaries/shiori_linux_amd64_v1 binaries/shiori_linux_amd64
gzip -d -S binaries/.gz__  -r .
chmod 755 binaries/shiori_linux_*/shiori
echo "Done"