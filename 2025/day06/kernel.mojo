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
    # fdata = ""
    num_lines = len(data)
    # Assume index 0 on last row has the first operator
    start_idx = 0
    end_idx = 0
    operator_row = data[num_lines - 1]
    len_oper_row = len(operator_row)
    for i in range(0, len_oper_row):
        operator = operator_row[i]
        if operator == "*" or operator == "+" or i == len_oper_row - 1:
            end_idx = i + 1 if len_oper_row - 1 == i else i - 1
            problem_total = 0
            var oper = plus
            if operator_row[start_idx] == "*":
                #     problem_total = 1
                oper = multi
            for idx in range(start_idx, end_idx):
                sig = 1
                num = 0
                for row in range(num_lines - 2, -1, -1):
                    # Start from the bottom since the lease significat digit is at the bottom
                    # ranage from second from bottom to 0 with steps of -1
                    d = data[row][idx]  # get the digit on row and index
                    if d.is_ascii_digit():
                        num += sig * atol(d)
                        sig *= 10

                problem_total = oper(problem_total, num)
            start_idx = i
            total += problem_total

    return total


fn plus(a: Int, b: Int) -> Int:
    return a + b


fn multi(a: Int, b: Int) -> Int:
    if a == 0:
        return b
    return a * b


struct Problem:
    var oper: fn (Int, Int) -> Int
    var inputs: List[String]
    var start: Int

    fn __init__(out self, oper: fn (Int, Int) -> Int, start: Int):
        self.oper = oper
        self.inputs = List[String]()
        self.start = start

    fn add_input(mut self, input: String):
        self.inputs.append(input)

    fn total(self) raises -> Int:
        total = self.start
        for input in self.inputs:
            n = atol(input)
            total += self.oper(total, n)

        return total
