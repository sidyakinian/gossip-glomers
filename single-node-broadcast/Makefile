# Default variable values (these can be overridden)
BIN_PATH?=~/go/bin/maelstrom-broadcast
NODE_COUNT?=1
TIME_LIMIT?=20
RATE?=10
MAELSTROM_EXEC_PATH=../maelstrom/maelstrom test

# Target for the maelstrom test with broadcast
test:
	go install .
	@echo "Package installed"
	$(MAELSTROM_EXEC_PATH) -w broadcast \
	--bin $(BIN_PATH) \
	--node-count $(NODE_COUNT) \
	--time-limit $(TIME_LIMIT) \
	--rate $(RATE)