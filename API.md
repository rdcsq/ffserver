# API

This is a JSON API. Please note that all endpoints that require a video/audio URL, require a **direct link** to a file (for example: `https://example.com/videofile.mp4`)

## Type of respones

- 200: success
- 400: bad request (no/incorrect body)
- 401: unauthorized (bad JWT)
- 500: ffmpeg/ffprobe error

## Authentication

All endpoints require authentication. Add a JWT in the `Authorization` header.

Example:

```
POST /thumbnail

Headers:
    Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuZXZlciI6Imdvbm5hIiwiZ2l2ZSI6InlvdSIsInVwIjoiOikifQ.UtLmrUsHJpQcEyHaDmfcXf8JI0OysimvXAX_rQj37Jo
```

## Endpoints

### `GET /`

Get the version of FFmpeg that's available on the server

- Example response:

```json
{
  "ffmpegVersion": "4.4.2-0ubuntu0.22.04.1"
}
```

### `POST /streams`

Get all streams of a file using FFprobe.

Equivalent FFprobe command: `ffprobe -v quiet -print_format json -show_streams <URL/file>`

- Request body:

```json
{
  "url": "<video URL>"
}
```

- Example response (truncated):

```json
{
    "streams": [
        {
            "index": 0,
            "codec_name": "h264",
            "codec_long_name": "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10",
            "profile": "High",
            "codec_type": "video",
            "codec_tag_string": "avc1",
            "codec_tag": "0x31637661",
            "width": 1280,
            "height": 720,
            ...
        },
        {
            "index": 1,
            "codec_name": "aac",
            "codec_long_name": "AAC (Advanced Audio Coding)",
            "profile": "LC",
            "codec_type": "audio",
            ...
        }
    ]
}

```

### `POST /thumbnail`

Generate a thumbnail with the first frame of a video (URL).

Equivalent FFmpeg command: `ffmpeg -i <URL/file> -vframes 1 <output file>`

- Request body:

```json
{
  "url": "<video URL>",
  "format": "webp" // Options: webp, png, jpg, or jpeg
}
```

- Example response:

```json
{
  "id": "hIx3crTl94j6aisEFraqsRUqYmqgMaxR", // ID is random
  "downloadUrl": "http://localhost:3000/thumbnail/hIx3crTl94j6aisEFraqsRUqYmqgMaxR.webp" // <DOMAIN>/thumbnail/<ID>.<FORMAT>
}
```

### `GET /thumbnail/{id}`

Get the generated thumbnail file. Not really an API endpoint but included here anyways.
