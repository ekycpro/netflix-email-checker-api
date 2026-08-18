const key = process.env.EKYCPRO_API_KEY;
if (!key) throw new Error("Set the EKYCPRO_API_KEY environment variable");

fetch("https://api.ekycpro.com/v1/check", {
  method: "POST",
  headers: { "X-API-Key": key, "Content-Type": "application/json" },
  body: JSON.stringify({ service_type: "netflix", identifier: "alice@example.com" }),
})
  .then(async (r) => {
    if (!r.ok) throw new Error(`request failed: ${r.status} ${await r.text()}`);
    return r.json();
  })
  .then(console.log);
