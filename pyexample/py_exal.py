import func

class SafeDict(dict):

    def __missing__(self, k):
        if hasattr(func, k):
            return getattr(func, k)
        else:
            #return __builtins__.get(k, None)
            return None


code_cache = {}

def py_eval_expr(expr, context):
    if expr not in code_cache:
        code_cache[expr] = compile(expr, '<string>', 'eval')
    return eval(code_cache[expr], {}, SafeDict(context))

content={
    "76":"76content",
    "RRA":"resp",
    "FMT":"010",
}

cmd='( E76  if    (0 if FMT in ("196", 196, "192", 192, "194", 194, "024", 024, "941", 941, "034", 034) else (E56 or F01 or F50 or E51 or 76 )) in [76] esle (0 if FMT in ("196", 196, "192", 192, "194", 194, "024", 024, "941", 941, "034", 034) else (E56 or F01 or F50 or E51 or 76 ))) if RRA=="resp" else null'
print py_eval_expr(cmd, content)