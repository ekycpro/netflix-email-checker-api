// Browser code must never hold an API key. This calls a same-origin backend
// proxy (/api/check-proxy) that attaches the real API key server-side and
// forwards to https://api.ekycpro.com/v1/check.
fetch("/api/check-proxy", {
  method: "POST",
  headers: { "Content-Type": "application/json" },
  body: JSON.stringify({ service_type: "netflix", identifier: "alice@example.com" }),
})
  .then(async (r) => {
    if (!r.ok) throw new Error(await r.text());
    return r.json();
  })
  .then(console.log);
