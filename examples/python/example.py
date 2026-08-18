import os
import requests

key = os.environ.get("EKYCPRO_API_KEY")
if not key:
    raise SystemExit("Set the EKYCPRO_API_KEY environment variable")

r = requests.post(
    "https://api.ekycpro.com/v1/check",
    headers={"X-API-Key": key},
    json={"service_type": "netflix", "identifier": "alice@example.com"},
    timeout=60,
)
r.raise_for_status()
print(r.json())
