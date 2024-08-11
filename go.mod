module main

go 1.22.5

replace github.com/balagrivine/linky/config => ./config

replace github.com/balagrivine/linky/logic => ./logic

replace github.com/balagrivine/linky/utils => ./utils

require (
	github.com/balagrivine/linky/config v0.0.0-00010101000000-000000000000
	github.com/balagrivine/linky/logic v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/balagrivine/linky v0.0.0-20240811170939-5fd22fdc7bcb // indirect
	github.com/balagrivine/linky/utils v0.0.0-00010101000000-000000000000 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.22.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
