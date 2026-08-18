import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

public class ApiExample {
    public static void main(String[] args) throws Exception {
        String key = System.getenv("EKYCPRO_API_KEY");
        if (key == null || key.isBlank()) {
            throw new IllegalStateException("Set the EKYCPRO_API_KEY environment variable");
        }

        HttpRequest req = HttpRequest.newBuilder(URI.create("https://api.ekycpro.com/v1/check"))
                .header("X-API-Key", key)
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofString("{\"service_type\":\"netflix\",\"identifier\":\"alice@example.com\"}"))
                .build();
        HttpResponse<String> resp = HttpClient.newHttpClient().send(req, HttpResponse.BodyHandlers.ofString());
        if (resp.statusCode() >= 300) {
            throw new IllegalStateException("request failed: " + resp.statusCode() + " " + resp.body());
        }
        System.out.println(resp.body());
    }
}
