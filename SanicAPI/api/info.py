from sanic import Blueprint
from sanic.exceptions import NotFound
from sanic.response import text

info = Blueprint('info', url_prefix='/info')


@info.exception(NotFound)
async def ignore_404s(request, exception):
    return text("dssd")
