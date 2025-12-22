from python import PythonObject
from python.bindings import PythonModuleBuilder
from os import abort
from math import math
from algorithm import parallelize
from os import Atomic


@export
fn PyInit_kernel() -> PythonObject:
    try:
        var m = PythonModuleBuilder("kernel")
        m.def_function[find_invalid_py](
            "find_invalid_py", docstring="Parallel Solver"
        )
        return m.finalize()
    except e:
        return abort[PythonObject](String("error creating python module:", e))


fn find_invalid_py(py_obj: PythonObject) raises -> PythonObject:
    var mojo_data = List[String]()
    for py_str_item in py_obj:
        mojo_data.append(String(py_str_item))
    return find_invalid(mojo_data)


fn find_invalid(data: List[String]) raises -> Int:
    var total = Atomic[DType.int64](0)

    for line in data:
        var id_ranges = line.split(",")
        for id_range in id_ranges:
            var r = id_range.split("-")
            var start = Int64(r[0])
            var end = Int64(r[1])
            var size = Int(end - start + 1)

            @parameter
            fn worker(idx: Int):
                # We capture 'start' from the outer scope
                var i = start + Int64(idx)
                if repeted2(i):
                    total.fetch_add(i)

            parallelize[worker](size)

    return Int(total.load())


fn repeted2(i: Int64) -> Bool:
    var l = length_of_int(i)
    if l < 2:
        return False
    # Check factors of length for tiling
    for j in range(1, (Int(l) // 2) + 1):
        if l % j == 0:
            if is_repeated(i, l, Int64(j)):
                return True
    return False


fn is_repeated(i: Int64, total_len: Int64, chunk_len: Int64) -> Bool:
    var splitter = int_pow(chunk_len)
    var temp = i
    var first = temp % splitter
    temp /= splitter
    for _ in range(1, Int(total_len / chunk_len)):
        if temp % splitter != first:
            return False
        temp /= splitter
    return True


fn int_pow(exp: Int64) -> Int64:
    var res: Int64 = 1
    for _ in range(exp):
        res *= 10
    return res


fn length_of_int(i: Int64) -> Int64:
    if i == 0:
        return 1
    var length = 0
    var n = i
    while n > 0:
        n /= 10
        length += 1
    return length
