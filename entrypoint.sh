#!/bin/bash
set -e
git clone https://github.com/hermosa-circulo/docs.git /docs
if [ -e /resources/src/assets/docs ]; then
    rm -rf /resource/src/assets/docs
fi
mv /docs /resources/src/assets/docs
/resources/node_modules/.bin/ng serve --host 0.0.0.0  --port 80 --open --public ${DOMAIN_NAME}
