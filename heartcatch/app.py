from flask import Flask, request, json, make_response
import os
from googleapiclient.discovery import build
from googleapiclient.http import HttpError

api = Flask(__name__)

@api.route("/images", methods=['GET'])
def images():

    api_key = os.getenvb(b'API_KEY', b'')
    cse_id = os.getenvb(b'CSE_ID', b'')

    default_query = os.getenvb(b'DEFAULT_QUERY', str('可愛い女の子').encode('utf-8'))
    query = request.args.get('q', default=default_query.decode('utf-8'), type=str)

    svc = build("customsearch", "v1", developerKey=api_key.decode('utf-8'))
    try:
        res = svc.cse().list(
            q=query,
            cx=cse_id.decode('utf-8'),
            lr='lang_ja',
            searchType='image',
            start=1,
            num=10,
        ).execute()
        api.logger.info("Result: %s", res)
        return make_response(json.jsonify(res))
    except HttpError as err:
        api.logger.error("Content: %s", err.content)
        api.logger.error("Detail: %s", err.error_details)
        api.logger.error("Response: %s", err.resp)
        return make_response(err.content, 500)

if __name__ == "__main__":

    h = os.getenvb(b'API_HOST', str('localhost').encode('utf-8'))
    p = os.getenvb(b'API_PORT', str('5000').encode('utf-8'))

    api.run(host=h.decode('utf-8'), port=p.decode('utf-8'))
