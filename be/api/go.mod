module gapp.api

go 1.25.4

replace gapp.models.sys => ../models/sys

replace gapp.sys => ../sys

require gapp.sys v0.0.0-00010101000000-000000000000

require (
	github.com/joho/godotenv v1.5.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
