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
    var total_crossings: Int32 = 0
    var possition: Int32 = 50

    for line in data:
        direction = String(line[0])
        steps = Int32(line[1:])
        var full_rev: Int32 = steps / 100
        total_crossings += full_rev
        steps = steps % 100
        if direction == "L":
            steps = -steps

        var starting_at_zero: Bool = possition == 0
        next_possition = possition + steps

        print(
            "Starting at ",
            possition,
            " moving ",
            steps,
            " in direction ",
            direction,
            "end at ",
            next_possition,
            " current count: ",
            total_crossings,
        )

        # we cross zero going left
        if next_possition < 0:
            if not starting_at_zero:
                total_crossings += 1
            possition = 100 + next_possition
        else:
            if next_possition != 100:
                var corssings: Int32 = next_possition / 100
                total_crossings += corssings
            possition = next_possition % 100

        if possition == 0:
            total_crossings += 1

    return total_crossings.__int__()
