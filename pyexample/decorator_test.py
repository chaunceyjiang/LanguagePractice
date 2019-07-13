import functools
import time
import traceback
import sys


def log(func):
    @functools.wraps(func)
    def wapper(*args, **kwargs):
        try:
            return func(*args, **kwargs)
        except Exception:
            print(
                "{} {}".format(
                    time.strftime(
                        "%Y-%m-%d %H:%M:%S",
                        time.localtime()),
                    traceback.format_exc()),
                file=sys.stderr)
    return wapper


@log
def get():
    return 2 / 0


get()
