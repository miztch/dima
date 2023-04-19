# dima

- An Unofficial REST API for [vlr.gg](https://www.vlr.gg/) with AWS Chalice.
- Major design pattern was abstracted from [axsddlr/vlrggapi](https://github.com/axsddlr/vlrggapi)

## Endpoints

### `/news`

- Method: `GET`
- Response:
  ```python
  {
      "data": {
          "status": 200,
          'segments': [
              {
                  'title': str,
                  'description': str,
                  'date': str,
                  'author': str,
                  'url_path': str
              }
          ],
      }
  }
  ```

### `/matches/results`

- Method: `GET`
- Response:
  ```python
  {
      "data": {
          "status": 200,
          'segments': [
              {
                "team1": str,
                "team2": str,
                "score1": str,
                "score2": str,
                "time_completed": str,
                "round_info": str,
                "tournament_name": str,
                "match_page": str,
                "tournament_icon": str
              }
          ],
      }
  }
  ```

### `/matches/upcoming`

- Method: `GET`
- Response:
  ```python
  {
      "data": {
          "status": 200,
          'segments': [
              {
                "team1": str,
                "team2": str,
                "flag1": str,
                "flag2": str,
                "score1": str,
                "score2": str,
                "time_until_match": str,
                "round_info": str,
                "tournament_name": str,
                "match_page": str,
                "tournament_icon": str
              }
          ],
      }
  }
  ```

### `/rankings/<region>`

- Method: `GET`
- Region: `na`(North America), `eu`(Europe), `br`(Brazil), `ap`(Asia-Pacific), `kr`(Korea), `cn`(China), `jp`(Japan), `la-s`(LATAM-South), `la-n`(LATAM-North), `oce`(Oceania), `mn`(MENA), `gc`(Game Changers)
- Response:
  ```python
  {
      "data": {
          "status": 200,
          'segments': [
              {
                  'rank': str,
                  'team': str,
                  'country': str,
                  'last_played': str,
                  'last_played_team_logo': str,
                  'record': str,
                  'earnings': str,
                  'logo': str
              }
          ],
      }
  }
  ```

## usage

- You can deploy with [AWS Chalice](https://github.com/aws/chalice)
```bash
$ cd chalice
$ chalice deploy
$ curl https://endpoint/api
```

## License

MIT. [axsddlr/vlrggapi is also distributed with MIT License](https://github.com/axsddlr/vlrggapi/blob/master/LICENSE). 
- see also: [LICENSE](./LICENSE)