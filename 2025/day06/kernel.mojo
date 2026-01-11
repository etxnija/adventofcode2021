from python import PythonObject
from python.bindings import PythonModuleBuilder
from os import abort, Atomic
from algorithm import parallelize
from int_help.int_manipulation import int_pow, length_of_int, int_at


@export
fn PyInit_kernel() -> PythonObject:
    try:
        var m = PythonModuleBuilder("kernel")
        m.def_function[grand_total_py](
            "grand_total_py", docstring="Parallel Solver"
        )
        return m.finalize()
    except e:
        return abort[PythonObject](String("error creating python module:", e))


fn grand_total_py(py_obj: PythonObject) raises -> PythonObject:
    var mojo_data = List[String]()
    for py_str_item in py_obj:
        mojo_data.append(String(py_str_item))
    return grand_total(mojo_data)


fn grand_total(data: List[String]) raises -> Int:
    total = 0
    # Will make a flat list and then start from the bottom of each colum
    #
    # 1. Flatten
    fdata = List[String]()
    num_lines = 0
    for line in data:
        num_lines += 1
        num = line.split(" ")
        for n in num:
            if not n == "":
                fdata.append(String(n))
    # Check
    length = len(fdata)
    if not length % num_lines == 0:
        return 0
    var num_col: Int = length // num_lines

    for c in range(num_col):
        # Decide operator
        var oper: fn (Int, Int) -> Int
        problem_total = 0
        if fdata[num_col * (num_lines - 1) + c] == "*":
            oper = multi
            problem_total = 1
        else:
            oper = plus

        for r in range(num_lines - 1):
            snum = fdata[num_col * r + c]
            n = atol(snum)
            problem_total = oper(problem_total, n)
        total += problem_total

    return total


fn plus(a: Int, b: Int) -> Int:
    return a + b


fn multi(a: Int, b: Int) -> Int:
    return a * b
