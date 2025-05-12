.PHONY: ex00-up ex00-down ex00-test ex01-up ex01-down ex01-test ex02-up ex02-down ex02-test ex03-up ex03-down ex03-test

ex00-up:
	cd ex00 && docker compose up -d
ex00-down:
	cd ex00 && docker compose down -v
ex00-test:
	cd ex00 && go run test.go
ex01-up:
	cd ex01 && docker compose up -d
ex01-down:
	cd ex01 && docker compose down -v
ex01-test:
	cd ex01 && mv web/nginx/conf.d/maintenance_on web/nginx/conf.d/maintenance_off && go run test.go && mv web/nginx/conf.d/maintenance_off web/nginx/conf.d/maintenance_on && go run test.go mainte
ex02-up:
	cd ex02 && docker compose up -d
ex02-down:
	cd ex02 && docker compose down -v
ex02-test:
	cd ex02 && mv web/nginx/conf.d/maintenance_on web/nginx/conf.d/maintenance_off && go run test.go && mv web/nginx/conf.d/maintenance_off web/nginx/conf.d/maintenance_on && go run test.go mainte
ex03-up:
	cd ex03 && docker compose up -d
ex03-down:
	cd ex03 && docker compose down -v
ex03-test:
	cd ex03 && mv web/nginx/conf.d/maintenance_on web/nginx/conf.d/maintenance_off && go run test.go && mv web/nginx/conf.d/maintenance_off web/nginx/conf.d/maintenance_on && go run test.go mainte
