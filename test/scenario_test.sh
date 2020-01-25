#!/usr/bin/env bash

set -e
MASTER_PASSWORD="test1234"
NEW_MASTER_PASSWORD="test12345"
VERSION="v0.8.0"

pushd ../
 make build-darwin
popd

export PM_STORAGE_FILE_PATH=./testDB
test -f ./testDB && rm ./testDB
test -f ./export-data.csv  && rm ./export-data.csv

../target/darwin/${VERSION}/password-manager init -m ${MASTER_PASSWORD}
../target/darwin/${VERSION}/password-manager add test -u test.com -p test12345 -l "fb,gmail" -m ${MASTER_PASSWORD}
../target/darwin/${VERSION}/password-manager get test -m ${MASTER_PASSWORD} -s
echo "===Searching password==="
../target/darwin/${VERSION}/password-manager search test -s -m ${MASTER_PASSWORD}
echo "===Searching password with a label==="
../target/darwin/${VERSION}/password-manager search -l "fb" -s -m ${MASTER_PASSWORD}

echo "===Entering new Master password==="
../target/darwin/${VERSION}/password-manager change-master-password -m ${MASTER_PASSWORD} -n ${NEW_MASTER_PASSWORD}
echo "===Switching master password to new master password==="
MASTER_PASSWORD=${NEW_MASTER_PASSWORD}
../target/darwin/${VERSION}/password-manager search test -s -m ${MASTER_PASSWORD}

echo "===Importing password from a CSV file==="
../target/darwin/${VERSION}/password-manager import --csv-file ./mock-data/data.csv -m ${MASTER_PASSWORD}
../target/darwin/${VERSION}/password-manager get ryendalll@latimes.com -m ${MASTER_PASSWORD} -s

echo "===Exporting passwords to a CSV file==="
../target/darwin/${VERSION}/password-manager export --csv-file ./export-data.csv -m ${MASTER_PASSWORD}

echo "===Remove a password==="
../target/darwin/${VERSION}/password-manager remove test -m ${MASTER_PASSWORD}

test -f ./testDB && rm ./testDB
test -f ./export-data.csv  && rm ./export-data.csv
