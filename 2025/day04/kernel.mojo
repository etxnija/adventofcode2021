from python import PythonObject
from python.bindings import PythonModuleBuilder
from os import abort, Atomic
from algorithm import parallelize
from int_help.int_manipulation import int_pow, length_of_int, int_at


@export
fn PyInit_kernel() -> PythonObject:
    try:
        var m = PythonModuleBuilder("kernel")
        m.def_function[movable_py]("movable_py", docstring="Parallel Solver")
        return m.finalize()
    except e:
        return abort[PythonObject](String("error creating python module:", e))


fn movable_py(py_obj: PythonObject) raises -> PythonObject:
    var mojo_data = List[String]()
    for py_str_item in py_obj:
        mojo_data.append(String(py_str_item))
    return movable(mojo_data)


fn movable(data: List[String]) raises -> Int:
    # Will try to make it into one long string and then access it flat insteasd of a two dimentional
    # 1. Record the length of a line
    L = len(data[0]) + 1
    S = ""
    # 2. Make it flat.
    for s in data:
        S += " "
        S += s
    length = len(S)
    var total = Atomic[DType.int](0)
    var found_count = Atomic[DType.int](0)
    # 3. Check suroundings,
    length = len(S)
    x = ord("x")

    var remove: List[Int] = List[Int](capacity=length)
    Sb = List[Int](length=length, fill=0)
    for i in range(length):
        Sb[i] = ord(S[i])

    while True:

        @parameter
        fn worker(idx: Int):
            roll_count = check_surondings(Sb, L, length, idx)
            if roll_count < 4:
                remove_loc = found_count.fetch_add(1)
                remove[remove_loc] = idx
                total.fetch_add(1)

        parallelize[worker](length)

        found = found_count.load()
        # print("found count: ", found)
        if found < 1:
            break

        for re in range(found):
            # print("removing: ", re)
            Sb[remove[re]] = x
        # ss = ""
        # for e in Sb:
        #     ss += String(chr(e))
        # print(ss)
        found_count = Atomic[DType.int](0)
        found = found_count.load()
        # print("After reset: ", found)

    return Int(total.load())


fn check_surondings(Sb: List[Int], L: Int, length: Int, i: Int) -> Int:
    var codePointAt = ord("@")
    # print("i: ", i)
    roll_count = 0
    if Sb[i] != codePointAt:
        #     print("Skipping : ", S[i])
        return 1000
    # 3.1. left
    loc = i - 1
    if loc > -1 and loc < length and Sb[loc] == codePointAt:
        #     print("Adding left: ", loc, length, S[loc])
        roll_count += 1
    # 3.2. right
    loc = i + 1
    if loc > -1 and loc < length and Sb[loc] == codePointAt:
        #     print("Adding right: ", loc, length, S[loc])
        roll_count += 1
    # 3.3. above
    loc = i - L
    if loc > -1 and loc < length and Sb[loc] == codePointAt:
        #     print("Adding above: ", loc, length, S[loc])
        roll_count += 1
    # 3.4. below
    loc = i + L
    if loc > -1 and loc < length and Sb[loc] == codePointAt:
        #     print("Adding below: ", loc, length, S[loc])
        roll_count += 1
    # 3.5. left above
    loc = i - L - 1
    if loc > -1 and loc < length and Sb[loc] == codePointAt:
        #     print("Adding l above: ", loc, length, S[loc])
        roll_count += 1
    # 3.6. right above
    loc = i - L + 1
    if loc > -1 and loc < length and Sb[loc] == codePointAt:
        #     print("Adding r above: ", loc, length, S[loc])
        roll_count += 1
    # 3.7. left below
    loc = i + L - 1
    if loc > -1 and loc < length and Sb[loc] == codePointAt:
        #     print("Adding l below: ", loc, length, S[loc])
        roll_count += 1
    # 3.8. right below
    loc = i + L + 1
    if loc > -1 and loc < length and Sb[loc] == codePointAt:
        #     print("Adding r below: ", loc, length, S[loc])
        roll_count += 1

    # 4. If less then 4 true add to count
    # print("rollcount: ", roll_count)
    return roll_count
