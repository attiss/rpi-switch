.PHONY: build
build:
	GOOS=linux GOARCH=arm64 go build -o rpi-switch .

.PHONY: clean
clean:
	rm -rf rpi-switch

.PHONY: install
install: build
	sudo sh -c "mkdir -p /opt/rpi-swtich && \
		cp rpi-switch /opt/rpi-swtich/rpi-switch && \
		cp config.yaml /opt/rpi-swtich/config.yaml && \
		cp misc/rpi-switch.service /etc/systemd/system/rpi-switch.service && \
		systemctl daemon-reload && \
		systemctl start rpi-swtich && \
		systemctl enable rpi-swtich"
