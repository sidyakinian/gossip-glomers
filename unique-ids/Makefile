BIN_PATH?=~/go/bin/maelstrom-unique-ids
TIME_LIMIT?=30
RATE?=1000
NODE_COUNT?=3
AVAILABILITY?=total
NEMESIS?=partition
MAELSTROM_PATH=../maelstrom/maelstrom

run-test:
	$(MAELSTROM_PATH) test -w unique-ids \
	--bin $(BIN_PATH) \
	--time-limit $(TIME_LIMIT) \
	--rate $(RATE) \
	--node-count $(NODE_COUNT) \
	--availability $(AVAILABILITY) \
	--nemesis $(NEMESIS)