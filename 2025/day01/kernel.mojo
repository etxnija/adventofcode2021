from python import PythonObject
from python.bindings import PythonModuleBuilder
from os import abort


@export
fn PyInit_kernel() -> PythonObject:
    try:
        var m = PythonModuleBuilder("kernel")
        m.def_function[calculate_distance_py](
            "calculate_distance_py", docstring="Calculate Distance"
        )
        return m.finalize()
    except e:
        return abort[PythonObject](String("error creating python module:", e))


fn calculate_distance_py(py_obj: PythonObject) raises -> PythonObject:
    mojo_data = List[String]()

    for py_str_item in py_obj:
        mojo_str = String(py_str_item)
        mojo_data.append(mojo_str)

    return calculate_distance(mojo_data)


fn calculate_distance(data: List[String]) raises -> Int:
    var total_sum: Int32 = 0
    zero = ord("0")
    print(zero)
    # var data = List[String](py_obj)
    for line in data:
        print("line: ", line)
        var first: Int32 = -1
        var last: Int32 = -1

        for cp in line.codepoints():
            if cp.is_ascii_digit():
                n = Int(cp) - zero
                if first == -1:
                    first = n

                last = n

        if first != -1 and last != -1:
            print("adding: ", first * 10, " and ", last, "to ", total_sum)
            total_sum += (first * 10) + last

    return total_sum.__int__()
