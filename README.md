# axis-webhook-receiver

Lightweight HTTP server that receives VAPIX event notifications from Axis cameras and saves snapshot images to disk.

## How it works

Axis cameras send a multipart HTTP POST to `/webhook` containing a JSON event metadata part and a JPEG image part. The server logs all incoming event data and saves image files to the configured output directory.

## Running

```bash
GOWORK=off go run .
```

**Environment variables:**

| Variable | Default | Description |
|---|---|---|
| `PORT` | `8081` | Port to listen on |
| `OUTPUT_DIR` | `./images` | Directory to save images |

## Endpoints

- `POST /webhook` — receives Axis camera events
- `GET /health` — returns `{"status":"ok"}`

## Axis configuration

In the Axis device web UI (System → Events), create an HTTP recipient pointing to:

```
http://<your-ip>:<PORT>/webhook
```

Configure an action rule to send the notification (with image) on your desired trigger.
