#!/bin/sh

echo "Creating /task/bin/ directory"
mkdir bin
export PATH=$PATH:/task/bin

echo "Installing pngout..."
curl http://shrinkrayio.s3.amazonaws.com/worker-libs/pngout > bin/pngout
chmod +x bin/pngout

echo "Installing jfifremove..."
curl http://shrinkrayio.s3.amazonaws.com/worker-libs/jfifremove.c > bin/jfifremove.c
gcc -o bin/jfifremove bin/jfifremove.c

git clone https://github.com/mwarkentin/colourio.git
imgopt .
