from python import PythonObject
from python.bindings import PythonModuleBuilder
from os import abort, Atomic
from algorithm import parallelize
from int_help.int_manipulation import int_pow, length_of_int, int_at
from collections import Set


@export
fn PyInit_kernel() -> PythonObject:
    try:
        var m = PythonModuleBuilder("kernel")
        m.def_function[splits_py]("splits_py", docstring="Parallel Solver")
        return m.finalize()
    except e:
        return abort[PythonObject](String("error creating python module:", e))


fn splits_py(py_obj: PythonObject) raises -> PythonObject:
    var mojo_data = List[String]()
    for py_str_item in py_obj:
        mojo_data.append(String(py_str_item))
    return splits(mojo_data)


fn splits(data: List[String]) raises -> Int:
    # total = 0
    debth = len(data)

    # Find the start node.
    row1 = data[0]
    ## flaten
    fdata = ""
    for r_idx in range(1, debth):
        fdata += data[r_idx]
    length = len(row1)
    start_col = row1.find("S")
    # print(start_col)
    # print(fdata)

    cache = Dict[Int, Int]()
    total = count_nodes(fdata, start_col, length, cache)
    # print(cache.__str__())
    return total + 1


fn count_nodes(
    s: String, idx: Int, rl: Int, mut cache: Dict[Int, Int]
) raises -> Int:
    i = idx

    target = ord("^")
    # move down the col until we find a splitter
    while ord(s[i]) != target:
        i += rl
        if i >= len(s):
            return 0

    if i in cache:
        return cache[i]

    # Now split.
    # left and right
    count = 1
    count += count_nodes(s, i - 1, rl, cache)
    count += count_nodes(s, i + 1, rl, cache)
    cache[i] = count
    return count
