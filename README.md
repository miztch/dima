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

### for quick trial

- In advance, You may have to deploy [sasha](https://github.com/miztch/sasha), data source of dima.
- You can deploy with [AWS Chalice](https://github.com/aws/chalice)

```bash
table=YOUR_SASHA_TABLE_NAME
sed -i -e 's/SASHA_DYNAMODB_TABLE_NAME/$table/g' api/.chalice/config_template.json
mv api/.chalice/config_template.json api/.chalice/config.json

cd api
chalice deploy
curl http://${your-api-endpoint}/api/
```

### for permanent deployment

- You can use AWS SAM template (`template.yaml`)
  - Custom domain with ACM / Route53 Hostzone is implemented. Please remove it if you don't need.

```bash
sam build
sam deploy --guided --capabilities CAPABILITY_IAM

curl http://${your-custom-domain-endpoint}/
```