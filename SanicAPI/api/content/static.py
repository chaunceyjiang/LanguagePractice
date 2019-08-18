from sanic import Blueprint
from sanic.response import json

static = Blueprint('content_static', url_prefix='/static', version=1)


@static.get("/static")
async def test2(request):
    return json({"data": {}, "code": 0, "msg:": "success"})
