import logging
import os
import boto3
from boto3.dynamodb.conditions import Key, Attr

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def _get_table():
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table(os.environ['TABLE_NAME'])

    return table


def query(date):
    '''
    scan DynamoDB table by following condition.
    - 'startTime' begins with $date
    '''
    table = _get_table()

    logger.info('Query table. date: {}'.format(date))
    response = table.scan(
        FilterExpression=Attr('startTime').begins_with(date)
    )

    if response['Items']:
        for item in response['Items']:
            item['id'] = str(item['id'])

    logger.info('Found {} items.'.format(len(response['Items'])))
    return response['Items']
