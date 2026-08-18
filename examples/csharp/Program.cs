using System;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;

class Program
{
    static async Task Main()
    {
        var apiKey = Environment.GetEnvironmentVariable("EKYCPRO_API_KEY");
        if (string.IsNullOrEmpty(apiKey))
        {
            throw new Exception("Set the EKYCPRO_API_KEY environment variable");
        }

        using var client = new HttpClient();
        client.DefaultRequestHeaders.Add("X-API-Key", apiKey);

        var resp = await client.PostAsync(
            "https://api.ekycpro.com/v1/check",
            new StringContent(@"{""service_type"":""netflix"",""identifier"":""alice@example.com""}", Encoding.UTF8, "application/json"));
        var body = await resp.Content.ReadAsStringAsync();
        if (!resp.IsSuccessStatusCode)
        {
            throw new Exception($"request failed: {(int)resp.StatusCode} {body}");
        }
        Console.WriteLine(body);
    }
}
