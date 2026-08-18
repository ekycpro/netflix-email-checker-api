<?php
$apiKey = getenv("EKYCPRO_API_KEY");
if (!$apiKey) {
    fwrite(STDERR, "Set the EKYCPRO_API_KEY environment variable\n");
    exit(1);
}

$ch = curl_init("https://api.ekycpro.com/v1/check");
curl_setopt_array($ch, [
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_POST => true,
    CURLOPT_HTTPHEADER => ["X-API-Key: $apiKey", "Content-Type: application/json"],
    CURLOPT_POSTFIELDS => '{"service_type":"netflix","identifier":"alice@example.com"}',
]);
$body = curl_exec($ch);
if ($body === false) {
    fwrite(STDERR, "request error: " . curl_error($ch) . "\n");
    exit(1);
}
$status = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);
if ($status >= 300) {
    fwrite(STDERR, "request failed: $status $body\n");
    exit(1);
}
echo "$body\n";
