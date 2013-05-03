#!/bin/sh

echo "Creating /task/bin/ directory"
mkdir bin
export PATH=$PATH:/task/bin

# "Installing pngout..."
curl http://shrinkrayio.s3.amazonaws.com/worker-libs/pngout > bin/pngout
chmod +x bin/pngout

# "Installing jfifremove..."
curl http://shrinkrayio.s3.amazonaws.com/worker-libs/jfifremove.c > bin/jfifremove.c
gcc -o bin/jfifremove bin/jfifremove.c

# Clone and optimize repo
if [ $# -lt 4 ]
then
  echo "Usage: `basename $0` owner repo oauth_key branch_name <file1 file2 file3...>"
  exit $E_BADARGS
fi

FULL_REPO="$1/$2"
CLONE_URL="https://$3:x-oauth-basic@github.com/$FULL_REPO.git"

echo "Cloning $FULL_REPO ($CLONE_URL)"
git clone "$CLONE_URL"

echo "Running imgopt on full repo"
imgopt $2

cd $2
git config --global user.name "shrinkray.io"
git config --global user.email michael@shrinkray.io
git checkout -b $4
git status
git commit -am 'Bleep bloop.. optimizing your images.'
git push origin $4
