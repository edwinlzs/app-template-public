start-infra:
  cd ./infra && docker compose up --detach

stop-infra:
  cd ./infra && docker compose stop
