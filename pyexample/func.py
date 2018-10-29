#!/opt/python27/bin/python 
"""
build in functions for protocl config expr
"""

import re
import cStringIO as StringIO
import logging
# logger = logging.getLogger('default')
# try:
#     from lxml import etree as ET
# except:
#     from xml.etree import cElementTree as ET

def choose(cond, first, second):
    """
    >>> choose(True, "a", "b")
    'a'
    >>> choose(1==0, "a", "b")
    'b'
    """
    if cond:
        return first
    else:
        return second


# def xpath(xml, xpath, attr = None):
#     """
#     >>> xpath("<a><b>BBB</b></a>", "b")
#     'BBB'
#     >>> xpath("<a><b c='d'>BBB</b></a>", "b", "c")
#     'd'
#     """
#     if type(xml) not in (str, unicode):
#         return
#     else:
#         tree = ET.parse(StringIO.StringIO(xml))
#         node = tree.find(xpath)
#         if node is not None:
#             if attr is None:
#                 return node.text
#             else:
#                 return node.attrib.get(attr, None)
#         else:
#             return
        return


def regrex(val, expr, attr = None):
    if not val:
        return
    else:
        try:
            ret = None
            m = re.search(expr, val)
            if m:
                if attr:
                    if type(attr) == type(0):
                        ret = m.group(attr)
                    else:
                        ret = m.groupdict.get(attr, None)
                else:
                    ret = m.group()
        except Exception:
            import traceback
            # logger.error(traceback.format_exc())
            return

        return ret


__all__ = ['choose','regrex']