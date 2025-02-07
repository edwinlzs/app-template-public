start-infra:
  cd ./infra && podman compose up --detach

stop-infra:
  cd ./infra && podman compose stop
