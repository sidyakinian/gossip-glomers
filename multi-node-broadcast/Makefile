BIN_PATH?=~/go/bin/maelstrom-broadcast
NODE_COUNT?=5
TIME_LIMIT?=20
RATE?=10
MAELSTROM_EXEC_PATH=../maelstrom/maelstrom test

test:
	go install .
	@echo "Package installed"
	$(MAELSTROM_EXEC_PATH) -w broadcast \
	--bin $(BIN_PATH) \
	--node-count $(NODE_COUNT) \
	--time-limit $(TIME_LIMIT) \
	--rate $(RATE)