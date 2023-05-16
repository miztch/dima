from chalice import Chalice
from chalicelib.api.scrape import Vlr

app = Chalice(app_name='dima')

vlr = Vlr()


@app.route('/')
def index():
    return {'message': 'Remember, bullets hurt.'}


@app.route('/news')
def VLR_news():
    request = app.current_request
    page = request.query_params.get('page') if request.query_params else ''
    return vlr.vlr_recent(page)


@app.route('/matches/results')
def VLR_scores():
    request = app.current_request
    page = request.query_params.get('page') if request.query_params else ''
    return vlr.vlr_score(page)


@app.route('/stats/{region}/{timespan}')
def VLR_stats(region, timespan):
    """
    region shortnames:
        "na" -> "north-america",
        "eu" -> "europe",
        "ap" -> "asia-pacific",
        "sa" -> "latin-america",
        "jp" -> "japan",
        "oce" -> "oceania",
        "mn" -> "mena",
    timespan:
        "30" -> 30 days,
        "60" -> 60 days,
        "90" -> 90 days,
    """
    return vlr.vlr_stats(region, timespan)


@app.route('/rankings/{region}')
def VLR_ranks(region):
    """
    region shortnames:\n
        "na" -> "north-america",\n
        "eu" -> "europe",\n
        "ap" -> "asia-pacific",\n
        "la" -> "latin-america",\n
        "la-s" -> "la-s",\n
        "la-n" -> "la-n",\n
        "oce" -> "oceania",\n
        "kr" -> "korea",\n
        "mn" -> "mena",\n
        "gc" -> "game-changers",\n
        "br" -> "Brazil",\n
        "cn" -> "china",\n
    """
    return vlr.vlr_rankings(region)


@app.route('/matches/upcoming')
def VLR_upcoming():
    request = app.current_request
    page = request.query_params.get('page') if request.query_params else ''
    return vlr.vlr_upcoming(page)


@app.route('/events')
def VLR_events():
    request = app.current_request
    page = request.query_params.get('page') if request.query_params else ''
    return vlr.vlr_events(page)


@app.route('/health')
def health():
    return {'status': 'OK'}
