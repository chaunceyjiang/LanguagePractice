import asyncio
import os
from signal import signal, SIGINT

import uvloop

from sanic import Sanic

from api import api

app = Sanic(__name__)

app.blueprint(api)

if __name__ == "__main__":
    model = os.getenv("MODEL", "ASYNC")
    if model == "ASYNC":
        asyncio.set_event_loop(uvloop.new_event_loop())
        server = app.create_server(
            host="0.0.0.0",
            port=8000,
            return_asyncio_server=True)
        loop = asyncio.get_event_loop()
        task = asyncio.ensure_future(server)
        signal(SIGINT, lambda s, f: loop.stop())
        try:
            loop.run_forever()
        except Exception:
            loop.stop()
    else:
        app.run(host="0.0.0.0", port=8000, workers=4)
