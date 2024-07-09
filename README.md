# dima

- An unofficial & simple API for [vlr.gg](https://www.vlr.gg/) upcoming matches
  - get matches information from DynamoDB
    - source data is provided by [sasha](https://github.com/miztch/sasha), scraper for upcoming matches
  - named after the real name of Valorant agent Neon 🩹⚡

## Example

### Request

```bash
$ curl endpoint/api/matches?date=2023-05-28
```

### Response

```json
[
  {
    "matchName": "Playoffs: Grand Final",
    "startTime": "2023-05-28T15:00:00+0000",
    "eventName": "Champions Tour 2023: EMEA League",
    "teams": [
      {
        "title": "FNATIC"
      },
      {
        "title": "Team Liquid"
      }
    ],
    "matchPage": "/213198/fnatic-vs-team-liquid-champions-tour-2023-emea-league-gf",
    "bestOf": "5",
    "id": "213198",
    "eventCountryFlag": "de"
  }
]
```

## usage

- You can use AWS SAM template (`template.yaml`)
  - Custom domain with ACM / Route53 Hostzone is implemented. Please remove it if you don't need.

```bash
sam build
sam deploy --guided --capabilities CAPABILITY_IAM

curl http://${your-custom-domain-endpoint}/
```
