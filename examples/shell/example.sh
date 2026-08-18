#!/usr/bin/env bash
set -euo pipefail
: "${EKYCPRO_API_KEY:?Set the EKYCPRO_API_KEY environment variable}"
curl --fail-with-body -X POST "https://api.ekycpro.com/v1/check" -H "X-API-Key: ${EKYCPRO_API_KEY}" -H 'Content-Type: application/json' -d '{"service_type":"netflix","identifier":"alice@example.com"}'
