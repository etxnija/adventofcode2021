from python import PythonObject
from python.bindings import PythonModuleBuilder
from os import abort
from math import math


@export
fn PyInit_kernel() -> PythonObject:
    try:
        var m = PythonModuleBuilder("kernel")
        m.def_function[find_invalid_py](
            "find_invalid_py", docstring="Calculate Distance"
        )
        return m.finalize()
    except e:
        return abort[PythonObject](String("error creating python module:", e))


fn find_invalid_py(py_obj: PythonObject) raises -> PythonObject:
    mojo_data = List[String]()

    for py_str_item in py_obj:
        mojo_str = String(py_str_item)
        mojo_data.append(mojo_str)

    return find_invalid(mojo_data)


fn find_invalid(data: List[String]) raises -> Int:
    var total_crossings: Int64 = 0

    for line in data:
        id_ranges = line.split(",")
        for id_range in id_ranges:
            r = id_range.split("-")
            # print(r[0], r[1])
            start = Int64(r[0])
            end = Int64(r[1])
            range_of_ids = range(start, end + 1)
            for i in range_of_ids:
                l = length_of_int(i)
                if l % 2 == 0:
                    splitter = int_pow(l / 2)
                    # print("pow: ", splitter)
                    var s: Int64 = i / splitter
                    var e: Int64 = i % splitter
                    a = s & e
                    # print(
                    #     "num ",
                    #     i,
                    #     "has length ",
                    #     l,
                    #     " start:",
                    #     s,
                    #     " end:",
                    #     e,
                    #     " and: a",
                    # )
                    if s == e:
                        # print("----- this ----: ", i)
                        total_crossings += i

    return total_crossings.__int__()


fn int_pow(exponent: Int64) -> Int64:
    var result: Int64 = 1
    for _ in range(exponent):
        result *= 10
    return result


fn length_of_int(i: Int64) raises -> Int64:
    length = 0
    if i == 0:
        length = 1
    else:
        var current_num: Int64 = i
        while current_num > 0:
            current_num /= 10
            length += 1
    return length
