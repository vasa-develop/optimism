#/bin/bash
set -eu

CONTRACTS_PATH="../packages/contracts-bedrock"

if [ "$#" -ne 2 ]; then
	echo "This script takes 2 arguments - CONTRACT_NAME PACKAGE"
	exit 1
fi

need_cmd() {
  if ! command -v "$1" > /dev/null 2>&1; then
    echo "need '$1' (command not found)"
    exit 1
  fi
}

need_cmd forge
need_cmd abigen

NAME=$1
# This can only handle fully qualified syntax.
# Fully qualified: path-to-contract-file:contract-name
FULLY_QUALIFIED_PATH=$(echo "$NAME" | cut -d ':' -f1)
CONTRACT_NAME=$(echo "$NAME" | cut -d ':' -f2)
PACKAGE=$2

# Convert to lower case to respect golang package naming conventions
CONTRACT_NAME_LOWER=$(echo ${CONTRACT_NAME} | tr '[:upper:]' '[:lower:]')

mkdir -p bin
TEMP=$(mktemp -d)

# Pull ABI and bytecode from foundry artifacts.
ARTIFACT=$(cat ${CONTRACTS_PATH}/forge-artifacts/${CONTRACT_NAME}.sol/${CONTRACT_NAME}.json)
echo $ARTIFACT | jq '.abi' > ${TEMP}/${CONTRACT_NAME}.abi
echo $ARTIFACT | jq '.bytecode.object' | tr -d '"' > ${TEMP}/${CONTRACT_NAME}.bin
echo $ARTIFACT | jq '.deployedBytecode.object' | tr -d '"' > $(pwd)/bin/${CONTRACT_NAME_LOWER}_deployed.hex

# Run ABIGEN
abigen \
	--abi ${TEMP}/${CONTRACT_NAME}.abi \
	--bin ${TEMP}/${CONTRACT_NAME}.bin \
	--pkg ${PACKAGE} \
	--type ${CONTRACT_NAME} \
	--out ./${PACKAGE}/${CONTRACT_NAME_LOWER}.go
