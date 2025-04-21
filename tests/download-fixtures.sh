#!/usr/bin/env bash
cd ..
# git submodule update --init --depth 1 --recursive
echo "skipping submodule update"
cd tests
rm -rf spec-tests && mkdir spec-tests && cd spec-tests
wget https://github.com/ronin-chain/execution-spec-tests/releases/download/v1.0.2/fixtures_stable.tar.gz
tar xzf fixtures_stable.tar.gz && rm -f fixtures_stable.tar.gz
