from flask import Flask, request, json, make_response
import os
from googleapiclient.discovery import build
import pprint

api = Flask(__name__)

@api.route("/images", methods=['GET'])
def images():

    api_key = os.getenvb(b'API_KEY', str(''))
    custom_engine_id = os.getenvb(b'CSE_ID', str(''))

    default_query = os.getenvb(b'DEFAULT_QUERY', str('可愛い女の子'))
    query = request.args.get('q', default_query)

    svc = build("customsearch", "v1", developerKey=api_key)

    try:
        res = svc.cse().list(
            q=query,
            cx=custom_engine_id,
            lr='lang_ja',
            searchType='image',
            start=1,
            num=10,
        ).execute()
    except:
        pprint.pprint(res)

    return make_response(json.jsonify(res))


if __name__ == "__main__":

    h = os.getenvb(b'API_HOST', str('0.0.0.0'))
    p = os.getenvb(b'API_PORT', str('5000'))
    api.run(host=h, port=p)
