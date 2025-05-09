.PHONY: ex00-up ex00-down ex01-up ex01-down ex02-up ex02-down ex03-up ex03-down

ex00-up:
	cd ex00 && docker compose up -d
ex00-down:
	cd ex00 && docker compose down -v
ex01-up:
	cd ex01 && docker compose up -d
ex01-down:
	cd ex01 && docker compose down -v
ex02-up:
	cd ex02 && docker compose up -d
ex02-down:
	cd ex02 && docker compose down -v
ex03-up:
	cd ex03 && docker compose up -d
ex03-down:
	cd ex03 && docker compose down -v
