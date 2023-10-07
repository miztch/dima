import logging


def getLogger():
    """
    init logger.
    """
    logger = logging.getLogger()
    logger.setLevel(logging.INFO)

    return logger
