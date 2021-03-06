#!/bin/bash
# Downloads and 'installs' PureCSS for montesquieu

mkdir tmp
cd tmp
wget https://github.com/pure-css/pure-release/archive/v1.0.1.zip
unzip v1.0.1.zip
cd ..
mkdir -p html/css/pure
cp tmp/pure-release-1.0.1/*min.css html/css/pure/.

# don't forget to include copy of PureCSS' license
cp tmp/pure-release-1.0.1/LICENSE.md html/css/pure/LICENSE.md

rm -r tmp
