# dima

- An Unofficial REST API for [vlr.gg](https://www.vlr.gg/) with AWS Chalice.
- Major design pattern was abstracted from [axsddlr/vlrggapi](https://github.com/axsddlr/vlrggapi)

## Endpoints

### `/news`

- Method: `GET`
- Response:
  ```json
  {
    "data": {
      "status": 200,
      "segments": [
        {
          "title": "TenZ returns to Sentinels active roster",
          "description": "He's back.",
          "date": "May 9, 2023",
          "author": "eutalyx",
          "url_path": "/209209/tenz-returns-to-sentinels-active-roster"
        },...
      ]
    }
  }
  ```

### `/matches/results`

- Method: `GET`
- Response:
  ```json
  {
    "data": {
      "status": 200,
      "segments": [
        {
          "team1": "NRG Esports",
          "team2": "LOUD",
          "score1": "2",
          "score2": "0",
          "flag1": "flag_us",
          "flag2": "flag_br",
          "time_completed": "1d 14h ago",
          "round_info": " Regular Season-Week 7",
          "tournament_name": "Champions Tour 2023: Americas League",
          "match_page": "/183815/nrg-esports-vs-loud-champions-tour-2023-americas-league-w7",
          "tournament_icon": "https://owcdn.net/img/640f5ab71dfbb.png"
        },...
      ]
    }
  }
  ```

### `/matches/upcoming`

- Method: `GET`
- Response:
  ```json
  {
    "data": {
      "status": 200,
        "segments": [
        {
          "team1": "FNATIC",
          "team2": "Natus Vincere",
          "flag1": "flag_eu",
          "flag2": "flag_eu",
          "score1": "–",
          "score2": "–",
          "time_until_match": "1d 5h from now",
          "round_info": "Regular Season–Week 8",
          "tournament_name": "Champions Tour 2023: EMEA League",
          "match_page": "/184057/fnatic-vs-natus-vincere-champions-tour-2023-emea-league-w8",
          "tournament_icon": "https://owcdn.net/img/640f5accac10a.png"
        },...
      ]
    }
  }
  ```

### `/events`

- Method: `GET`
- Response:
  ```json
  {
    "data": {
      "status": 200,
      "segments": [
        {
          "event_name": "Champions Tour 2023: EMEA League",
          "status": "ongoing",
          "dates": "Mar 27-May 28",
          "flag": "flag_de",
          "prize": "TBD",
          "event_page": "/event/1190/champions-tour-2023-emea-league",
          "event_icon": "https://owcdn.net/img/640f5accac10a.png"
        },...
      ]
    }
  }
  ```

### `/rankings/<region>`

- Method: `GET`
- Region: `na`(North America), `eu`(Europe), `br`(Brazil), `ap`(Asia-Pacific), `kr`(Korea), `cn`(China), `jp`(Japan), `la-s`(LATAM-South), `la-n`(LATAM-North), `oce`(Oceania), `mn`(MENA), `gc`(Game Changers)
- Response:
  ```json
  {
    "status": 200,
    "data": [
      {
        "rank": "1",
        "team": "FNATIC",
        "country": "Europe",
        "last_played": "4d ago",
        "last_played_team": "vs. Karmine C",
        "last_played_team_logo": "//owcdn.net/img/627403a0d9e48.png",
        "record": "10–0",
        "earnings": "$427,285",
        "logo": "//owcdn.net/img/62a40cc2b5e29.png"
      },...
    ]
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