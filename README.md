# Netflix Email Checker API — official examples | eKYC Pro

This repository documents the eKYC Pro Netflix Email Checker API integration contract, verified against the live product page at https://ekycpro.com/tools/netflix.

- Product page: https://ekycpro.com/tools/netflix
- API base URL: `https://api.ekycpro.com`
- Authentication: `X-API-Key`
- API key environment variable: `EKYCPRO_API_KEY`
- License: MIT

## Processing model

This is a synchronous single-identifier JSON API. Send `service_type` and one valid email address; a determined result is returned in the response. An undetermined result is not a confirmed negative and is not billed.

```bash
curl -X POST "https://api.ekycpro.com/v1/check" -H "X-API-Key: ${EKYCPRO_API_KEY}" -H "Content-Type: application/json" -d '{"service_type":"netflix","identifier":"alice@example.com"}'
```

## Products and limits

| Code | Name | Input | Limits | Result fields | Documentation |
| --- | --- | --- | --- | --- | --- |
| `netflix` | Netflix Email Checker | email | 1–1 identifier | `id, identifier, registered, billed, cost, cost_usd` | [docs](https://docs.ekycpro.com/netflix-number-checker) |

## Response and boundaries

- Treat only an explicit determined result as a positive or negative classification; `undetermined`/`exists=false` is not a negative.
- Authentication keys must come from environment variables or a secret manager; never commit keys or call private/internal endpoints.
- Product prices can change. Check the official pricing page at runtime; this repository intentionally does not hard-code a price.
- Use only identifiers you are authorized to process and follow applicable privacy, platform, and data-protection requirements.

## Runnable examples

Seven self-contained examples are provided under `examples/`: Python, Node.js, Go, Java, C#, PHP, and Shell each call the synchronous `POST /v1/check` endpoint with HTTP error handling. Set the server-side API-key environment variable before running them.

## Official resources

- `netflix`: https://docs.ekycpro.com/netflix-number-checker
- Pricing: https://ekycpro.com/pricing
- OpenAPI contract: [openapi.yaml](openapi.yaml)

Last reviewed: 2026-08-18.
