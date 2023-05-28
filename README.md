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
        "match_page": "/213198/fnatic-vs-team-liquid-champions-tour-2023-emea-league-gf",
        "bestOf": "5",
        "id": "213198",
        "eventCountryFlag": "de"
    }
]
```


## usage

- In advance, You may have to deploy [sasha](https://github.com/miztch/sasha), data source of dima.
- You can deploy with [AWS Chalice](https://github.com/aws/chalice)
```bash
$ cd chalice
$ chalice deploy
$ curl endpoint/api
```
