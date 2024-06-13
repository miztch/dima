import os

import boto3
from boto3.dynamodb.conditions import Key
from chalicelib import log

logger = log.getLogger()


def _get_table():
    dynamodb = boto3.resource("dynamodb")
    table = dynamodb.Table(os.environ["TABLE_NAME"])

    return table


def query(date):
    """
    query DynamoDB table by following condition.
    - 'startDate' equals to $date
    """
    table = _get_table()

    logger.info("Query table. date: {}".format(date))

    result = []

    response = table.query(KeyConditionExpression=Key("startDate").eq(date))

    items = response.get("Items", [])
    for item in items:
        item["id"] = int(item["id"])
        item["bestOf"] = int(item["bestOf"])
    result.extend(items)

    while "LastEvaluatedKey" in response:
        response = table.query(
            KeyConditionExpression=Key("startDate").eq(date),
            ExclusiveStartKey=response["LastEvaluatedKey"],
        )
        for item in items:
            item["id"] = int(item["id"])
            item["bestOf"] = int(item["bestOf"])
        result.extend(items)

    logger.info("Found {} items.".format(len(result)))
    return result
