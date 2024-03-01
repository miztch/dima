import os

import boto3
from boto3.dynamodb.conditions import Key, Attr
from chalicelib import log

logger = log.getLogger()


def _get_table():
    dynamodb = boto3.resource("dynamodb")
    table = dynamodb.Table(os.environ["TABLE_NAME"])

    return table


def query(date):
    """
    scan DynamoDB table by following condition.
    - 'startTime' begins with $date
    """
    table = _get_table()

    logger.info("Query table. date: {}".format(date))

    filter_expression = Attr("startTime").begins_with(date)
    result = []

    response = table.scan(FilterExpression=filter_expression)

    items = response.get("Items", [])
    for item in items:
        item["id"] = int(item["id"])
        item["bestOf"] = int(item["bestOf"])
    result.extend(items)

    while "LastEvaluatedKey" in response:
        response = table.scan(
            FilterExpression=filter_expression,
            ExclusiveStartKey=response["LastEvaluatedKey"],
        )
        for item in items:
            item["id"] = int(item["id"])
            item["bestOf"] = int(item["bestOf"])
        result.extend(items)

    logger.info("Found {} items.".format(len(result)))
    return result
