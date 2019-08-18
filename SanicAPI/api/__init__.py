from sanic import Blueprint

from .content import content
from .info import info
_api_list = [
    content,
    info
]
api = Blueprint.group(*_api_list, url_prefix='/api')
