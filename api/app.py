from chalice import Chalice
from chalicelib import database, date_format

app = Chalice(app_name="dima")


@app.route("/")
def index():
    """
    return fixed response.
    """
    return {"status": 200, "message": "Remember, bullets hurt."}


@app.route("/matches")
def matches():
    """
    return upcoming matches list.
    """
    request = app.current_request
    app.log.debug("Request: {}".format(request.to_dict()))

    # default date for search: today
    date = date_format.get_default()

    if request.query_params:
        param = request.query_params.get("date")

        # validate date format of query parameter (YYYY-mm-dd)
        # if format is valid, overwrite date with parameter
        if param and date_format.validate(param):
            date = param

    matches = database.query(date)
    app.log.debug("Return {} items.".format(len(matches)))

    return matches
