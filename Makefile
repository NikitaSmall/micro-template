BIN_DIR=bin
BYNARY=master-service

BYNARY_OUTPUT= ${BIN_DIR}/${BYNARY}

all:
	go build -o ${BYNARY_OUTPUT} 
