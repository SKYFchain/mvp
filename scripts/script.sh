#!/bin/bash

echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "End-to-end test"
echo
CHANNEL_NAME="$1"
DELAY="$2"
LANGUAGE="$3"
TIMEOUT="$4"
VERBOSE="$5"
: ${CHANNEL_NAME:="mychannel"}
: ${DELAY:="3"}
: ${LANGUAGE:="golang"}
: ${TIMEOUT:="10"}
: ${VERBOSE:="false"}
LANGUAGE=`echo "$LANGUAGE" | tr [:upper:] [:lower:]`
COUNTER=1
MAX_RETRY=5

CC_SRC_PATH="gitlab.qdlt.io/skyf/skyfchain/chaincode/"

echo "Channel name : "$CHANNEL_NAME

# import utils
. scripts/utils.sh

createChannel() {
	setGlobals 0 "Customer1"

	if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
                set -x
		peer channel create -o orderer.skyfchain.io:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx >&log.txt
		res=$?
                set +x
	else
				set -x
		peer channel create -o orderer.skyfchain.io:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA >&log.txt
		res=$?
				set +x
	fi
	cat log.txt
	verifyResult $res "Channel creation failed"
	echo "===================== Channel '$CHANNEL_NAME' created ===================== "
	echo
}

joinChannel () {
	declare -a arr=("Customer1" "Operator1" "Admin1" "Monitor1")
	for org in "${arr[@]}"; do
	    for peer in 0 1; do
		joinChannelWithRetry $peer $org
		echo "===================== peer${peer}.${org} joined channel '$CHANNEL_NAME' ===================== "
		sleep $DELAY
		echo
	    done
	done
}

## Create channel
echo "Creating channel..."
createChannel

## Join all the peers to the channel
echo "Having all peers join the channel..."
joinChannel

## Set the anchor peers for each org in the channel
echo "Updating anchor peers for customer1..."
updateAnchorPeers 0 "Customer1"
echo "Updating anchor peers for operator1..."
updateAnchorPeers 0 "Operator1"
echo "Updating anchor peers for admin1..."
updateAnchorPeers 0 "Admin1"
echo "Updating anchor peers for monitor1..."
updateAnchorPeers 0 "Monitor1"


## Go get package
go get "github.com/google/uuid"
## Install chaincode on peer0 of each org
echo "Installing chaincode on peer0.customer1..."
installChaincode 0 "Customer1"
echo "Install chaincode on peer0.operator1..."
installChaincode 0 "Operator1"
echo "Installing chaincode on peer0.admin1..."
installChaincode 0 "Admin1"
echo "Install chaincode on peer0.monitor1..."
installChaincode 0 "Monitor1"


# Instantiate chaincode on peer0.operator1
echo "Instantiating chaincode on peer0.operator1..."
instantiateChaincode 0 "Operator1"

# Query chaincode on peer0.customer1
echo "Querying drones on peer0.customer1..."
queryDrones 0 "Customer1"

echo "Querying drone on peer0.customer1..."
queryDrone 0 "Customer1" 1

# Invoke chaincode on peer0.customer1 and peer0.operator1
#echo "Sending invoke transaction on peer0.customer1 peer0.operator1..."
#chaincodeInvoke 0 "Customer1" 0 "Operator1" 0 "Admin1" 0 "Monitor1"

# Install chaincode on peer1.operator1
echo "Installing chaincode on peer1.operator1..."
installChaincode 1 "Operator1"

# Query on chaincode on peer1.operator1, check if the result is 90
#echo "Querying chaincode on peer1.operator1..."
#chaincodeQuery 1 "Operator1" 90

echo
echo "========= All GOOD, Test execution completed =========== "
echo

echo
echo " _____   _   _   ____   "
echo "| ____| | \ | | |  _ \  "
echo "|  _|   |  \| | | | | | "
echo "| |___  | |\  | | |_| | "
echo "|_____| |_| \_| |____/  "
echo

exit 0
