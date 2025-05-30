# 🔗 URL Shortener

A lightweight and simple URL shortener built with Go. It converts long URLs into short, unique identifiers and allows redirection back to the original URL. This is a basic implementation for learning purposes and demonstrates the use of Go's standard libraries for building web services.

## 📌 Features

- Generate a short URL using an MD5 hash of the original URL.
- Redirect short URLs to their original destination.
- Basic REST API with JSON support.
- In-memory storage using Go maps.

## 📂 Repository

GitHub: [samiaabid06/URL-Shortner](https://github.com/samiaabid06/URL-Shortner)

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher installed

### Run Locally

```bash
git clone https://github.com/samiaabid06/URL-Shortner.git
cd URL-Shortner
go run main.go
````

Server starts on: `http://localhost:3000`

## 📬 API Endpoints

### `POST /shorten`

Creates a short URL.

#### Request Body

```json
{
  "url": "https://www.example.com"
}
```

#### Response

```json
{
  "short_url": "1a79a4d6"
}
```

---

### `GET /redirect/{short_url}`

Redirects to the original URL associated with the given short ID.

#### Example

`GET http://localhost:3000/redirect/1a79a4d6`
→ Redirects to `https://www.example.com`

---

## 🧪 Example (Using `curl`)

```bash
curl -X POST http://localhost:3000/shorten \
-H "Content-Type: application/json" \
-d '{"url":"https://www.google.com"}'
```

## ⚠️ Limitations

* Uses in-memory storage (data resets on server restart).
* No URL validation or duplication checks.
* Not suitable for production without persistence and security layers.

## 🛠️ To Do / Future Enhancements

* Add persistent storage (e.g., Redis, PostgreSQL)
* URL validation and expiration
* Click tracking and analytics
* User authentication and dashboard

## 👩‍💻 Author

Made with 💙 by [Samia Abid](https://github.com/samiaabid06)

## 📄 License

MIT License

```
Let me know if you want to add deployment instructions (e.g., for Heroku, Render, or Docker), or improve formatting for a portfolio.
```

