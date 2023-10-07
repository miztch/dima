import re
from datetime import datetime
from chalice import Chalice
from chalicelib import database

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
    date = datetime.strftime(datetime.now(), "%Y-%m-%d")

    if request.query_params:
        param = request.query_params.get("date")

        # validate date format of query parameter (YYYY-mm-dd)
        # if invalid, parameter date -> today.
        pattern = "^(?!([02468][1235679]|[13579][01345789])00-02-29)(([0-9]{4}-(01|03|05|07|08|10|12)-(0[1-9]|[12][0-9]|3[01]))|([0-9]{4}-(04|06|09|11)-(0[1-9]|[12][0-9]|30))|([0-9]{4}-02-(0[1-9]|1[0-9]|2[0-8]))|([0-9]{2}([02468][048]|[13579][26])-02-29))$"

        try:
            if re.fullmatch(pattern, param) is not None:
                app.log.debug(
                    "Date format in query parameter is valid. input: {}".format(param)
                )
                date = param
            else:
                app.log.debug(
                    "Date format in query parameter is invalid or empty. input: {}".format(
                        param
                    )
                )
        except TypeError:
            app.log.debug(
                "query parameter with invalid format given. input: {}".format(param)
            )

    matches = database.query(date)
    app.log.debug("Return {} items.".format(len(matches)))

    return matches
