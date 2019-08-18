from sanic import Blueprint

from .static import static
from .authors import authors
_api_list = [
    static,
    authors
]
content = Blueprint.group(*_api_list, url_prefix='/content')
