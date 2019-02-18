import storm
import logging
import logging.config


class StormLogHandler(logging.Handler):
    CRITICAL = 50
    FATAL = CRITICAL
    ERROR = 40
    WARNING = 30
    WARN = WARNING
    INFO = 20
    DEBUG = 10
    NOTSET = 0

    _levelToName = {
        CRITICAL: 'CRITICAL',
        ERROR: 'ERROR',
        WARNING: 'WARNING',
        INFO: 'INFO',
        DEBUG: 'DEBUG',
        NOTSET: 'NOTSET',
    }
    _nameToLevel = {
        'CRITICAL': 0,
        'ERROR': 4,
        'WARN': 3,
        'WARNING': 3,
        'INFO': 2,
        'DEBUG': 1,
        'NOTSET': 0,
    }

    def __init__(self):
        super(StormLogHandler, self).__init__()

    def emit(self, record):
        if isinstance(record, LogRecord):
            msg = record.getMessage()
        else:
            msg = repr(record)
        self.level = self._nameToLevel[self._levelToName[self.level]]
        storm.sendMsgToParent({
            "command": "log",
            "msg": msg,
            "level": self.level
        })


LOGGING_ROTATINGFILE_OPTIONS = {
    'maxBytes': 33554432,
    'backupCount': 4,
    'largeBackupCount': 8
}
LOGGING_FILTERS = {}
LOGGING_FORMATTERS = {
    'brief': {
        'format': '[%(asctime)s] %(levelname)s %(message)s'
    },
    'message_only': {
        'format': '%(message)s'
    },
    'brief_center': {
        'format': '[%(asctime)s] level=%(levelname)s %(message)s'
    }
}
VATICANPLUS_LOGGING = {
    'version': 1,
    'formatters': LOGGING_FORMATTERS,
    'filters': LOGGING_FILTERS,
    'handlers': {
        'null': {
            'class': 'django.utils.log.NullHandler',
            'level': 'DEBUG'
        },
        'console': {
            'class': 'logging.StreamHandler',
            'formatter': 'brief',
            'level': 'DEBUG'
        },
    },
    'loggers': {
        'default': {
            'handlers': [],
            'level': 'DEBUG',
            'propagate': False
        },
    }
}
logging.config.dictConfig(VATICANPLUS_LOGGING)
handler = StormLogHandler()
handler.setLevel('DEBUG')

logger = logging.getLogger('default')
logger.info('sdsds')
