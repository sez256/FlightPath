
## Flight Path Tracking API

This microservice API is designed to help track a person's flight path based on a list of flights provided. The API accepts a request that includes a list of flights, each defined by a source and destination airport code. These flights may not be listed in order and will be sorted to find the total flight paths starting and ending at airports.

### Endpoint

The API endpoint for tracking a person's flight path is:
```POST /calculate```

### Request Body

The request body should contain a JSON array of flight segments, where each segment is represented by an array of two airport codes (source and destination).

#### Example

```json
{
  "SFO": "EWR",
  "ATL": "EWR",
  "SFO": "ATL",
  "IND": "EWR",
  "GSO": "IND",
  "ATL": "GSO"
}
```

### Response

The API will respond with a JSON array representing the sorted total flight path starting and ending at airports.

#### Example

```json
{
    "start": "SFO",
    "end": "EWR"
}
```