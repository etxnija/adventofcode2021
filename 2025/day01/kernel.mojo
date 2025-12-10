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
    # var atZero: Bool = False
    for line in data:
        direction = String(line[0])
        steps = Int32(line[1:])
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

        if next_possition < 0:
            total_crossings += 1
            var crossings: Int32 = -next_possition / 100
            print("corssing R: ", crossings)
            total_crossings += crossings
            var nloc: Int32 = next_possition % 100
            print(nloc)
            possition = nloc
        else:
            var corssings: Int32 = next_possition / 100
            print("corssings L: ", corssings)
            total_crossings += corssings
            possition = next_possition % 100

        # if possition == 0:
        #     total_crossings += 1

        # var full: Int32 = Int32(steps) / 100
        # total_crossings += full
        # var n: Int32 = Int32(steps) % 100
        # print("full:", full, " rest: ", n)
        # print(
        #     "starting at: ",
        #     possition,
        #     " moding to the ",
        #     direction,
        #     " ",
        #     n,
        #     " steps",
        #     " AtZero ",
        #     atZero,
        # )
        # if direction == "L":
        #     n = -1 * n

        # possition = possition + n

        # print("new possition", possition)
        # if possition == 0 or possition == 100:
        #     total_crossings += 1
        #     possition = 0
        #     atZero = True
        #     continue

        # if possition < 0:
        #     print("minus", total_crossings)
        #     if not atZero:
        #         total_crossings += 1
        #     possition = 100 + possition
        #     continue

        # if possition > 100:
        #     total_crossings += 1
        #     possition = possition % 100

        # atZero = False

    return total_crossings.__int__()


fn floor_div(a: Int32, b: Int32) -> Int32:
    if a >= 0:
        return a / b
    else:
        return (a - b + 1) / b
