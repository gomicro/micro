package dockercompose

var (
	stub = `version: '2.3'

services:
  app:
    image: {{ .Org }}/{{ .Name }}
    ports:
      - 4567:4567
    networks:
      - services
    healthcheck:
      test: ["CMD", "/probe", "-kv", "https://localhost:4567/v1/status"]
      interval: 5s
      timeout: 60s
      retries: 3
    env_file: ./env.test{{ if .Database }}

  db:
    image: postgres:11.7-alpine
    environment:
      POSTGRES_USER: {{ .Name }}
      POSTGRES_DB: {{ .Name }}
      POSTGRES_HOST_AUTH_METHOD: trust
    ports:
      - 5432:5432
    networks:
      services:
        aliases:
          - database
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U {{ .Name }}"]
      interval: 5s
      timeout: 30s
      retries: 3{{ end }}

networks:
  services:
    internal: false
`
)
