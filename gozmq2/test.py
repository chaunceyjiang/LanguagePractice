#!/usr/bin/env python
# -*- coding: utf-8 -*-
import sys
import json
import cPickle
reload(sys)
sys.setdefaultencoding('utf-8')


def test(argv):
    print argv
    res = {
        'price': 0.01,
        'name': "test",
    }
    final = json.dumps(res)
    return final


def test2(x):
    return x*3


def py_eval_expr(expr, context):

    return eval(compile(expr, '<string>', 'eval'), {}, context)


if __name__ == "__main__":
    print py_eval_expr("c", {"a": 1, "b": 2, "c": "222"})
