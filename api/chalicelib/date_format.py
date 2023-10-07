import logging
import re
from datetime import datetime

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def validate(date):
    """
    Validate date format (YYYY-mm-dd).
    """
    pattern = "^(?!([02468][1235679]|[13579][01345789])00-02-29)(([0-9]{4}-(01|03|05|07|08|10|12)-(0[1-9]|[12][0-9]|3[01]))|([0-9]{4}-(04|06|09|11)-(0[1-9]|[12][0-9]|30))|([0-9]{4}-02-(0[1-9]|1[0-9]|2[0-8]))|([0-9]{2}([02468][048]|[13579][26])-02-29))$"

    if re.fullmatch(pattern, date) is not None:
        logger.debug("Date format is valid. input: {}".format(date))
        return True
    else:
        logger.debug("Date format is invalid or empty. input: {}".format(date))
        return False


def get_default():
    """
    Get the default date for search: today.
    """
    return datetime.strftime(datetime.now(), "%Y-%m-%d")
