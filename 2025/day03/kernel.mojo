from python import PythonObject
from python.bindings import PythonModuleBuilder
from os import abort
from algorithm import parallelize
from int_help.int_manipulation import int_pow, length_of_int, int_at


@export
fn PyInit_kernel() -> PythonObject:
    try:
        var m = PythonModuleBuilder("kernel")
        m.def_function[joltage_py]("joltage_py", docstring="Parallel Solver")
        return m.finalize()
    except e:
        return abort[PythonObject](String("error creating python module:", e))


fn joltage_py(py_obj: PythonObject) raises -> PythonObject:
    var mojo_data = List[String]()
    for py_str_item in py_obj:
        mojo_data.append(String(py_str_item))
    return joltage(mojo_data)


fn joltage(data: List[String]) raises -> Int:
    var total: Int = 0
    for line in data:
        # var i: Int128 = Int128(line)
        total += bank_joltage(line)

    return total


fn bank_joltage(s: String) raises -> Int:
    l = len(s)
    var p: Int = 0
    for loc in range(l - 1):
        num = Int(s[loc])
        # print("num: ", num, " at loc: ", loc, "l: ", l)
        for sub_loc in range(loc + 1, l):
            sub_num = Int(s[sub_loc])
            var tp: Int = 10 * num + sub_num
            # print(
            #     "tp: ",
            #     tp,
            #     " num: ",
            #     num,
            #     "at: ",
            #     loc,
            #     " sub_num:",
            #     sub_num,
            #     "at: ",
            #     sub_loc,
            # )
            if tp > p:
                p = tp
    return p
