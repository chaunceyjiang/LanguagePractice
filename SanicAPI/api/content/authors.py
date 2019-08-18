from sanic import Blueprint
from sanic.response import json

authors = Blueprint('content_authors', url_prefix='/authors', version=1)


@authors.get('/hello')
async def test(request):
    return json({"hello": "world"})
