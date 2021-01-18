all:
	go build -o pgn

install:
	cp ./pgn /usr/local/bin
	if [ ! -d "/etc/pgn" ]; then \
		mkdir /etc/pgn; \
		cp ./conf/pgn.conf.template /etc/pgn/pgn.conf; \
	fi
