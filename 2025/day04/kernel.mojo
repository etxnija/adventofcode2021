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
    var codePointAt = ord("@")
    total = 0
    # 3. Check suroundings,
    for i in range(length):
        print("i: ", i)
        roll_count = 0
        if ord(S[i]) != codePointAt:
            print("Skipping : ", S[i])
            continue
        # 3.1. left
        loc = i - 1
        if loc > -1 and loc < length and ord(S[loc]) == codePointAt:
            print("Adding left: ", loc, length, S[loc])
            roll_count += 1
        # 3.2. right
        loc = i + 1
        if loc > -1 and loc < length and ord(S[loc]) == codePointAt:
            print("Adding right: ", loc, length, S[loc])
            roll_count += 1
        # 3.3. above
        loc = i - L
        if loc > -1 and loc < length and ord(S[loc]) == codePointAt:
            print("Adding above: ", loc, length, S[loc])
            roll_count += 1
        # 3.4. below
        loc = i + L
        if loc > -1 and loc < length and ord(S[loc]) == codePointAt:
            print("Adding below: ", loc, length, S[loc])
            roll_count += 1
        # 3.5. left above
        loc = i - L - 1
        if loc > -1 and loc < length and ord(S[loc]) == codePointAt:
            print("Adding l above: ", loc, length, S[loc])
            roll_count += 1
        # 3.6. right above
        loc = i - L + 1
        if loc > -1 and loc < length and ord(S[loc]) == codePointAt:
            print("Adding r above: ", loc, length, S[loc])
            roll_count += 1
        # 3.7. left below
        loc = i + L - 1
        if loc > -1 and loc < length and ord(S[loc]) == codePointAt:
            print("Adding l below: ", loc, length, S[loc])
            roll_count += 1
        # 3.8. right below
        loc = i + L + 1
        if loc > -1 and loc < length and ord(S[loc]) == codePointAt:
            print("Adding r below: ", loc, length, S[loc])
            roll_count += 1
        # 4. If less then 4 true add to count
        print("rollcount: ", roll_count)
        if roll_count < 4:
            print("Adding to tltal: ", total)
            total += 1
    return total


fn bank_movable(
    s: String, start_i: Int128, num_batteries: Int, mut cache: List[Int]
) raises -> Int:
    # print("s: ", s, " start_i: ", start_i, "num_b: ", num_batteries)
    if num_batteries == 0:
        return 0
    if len(s) - start_i < num_batteries:
        return -1

    l = len(s)
    cache_idx = start_i * 13 + num_batteries
    if cache[cache_idx] != -1:
        # print("found in cache: ", cache[cache_idx], "at index: ", cache_idx)
        return cache[cache_idx]

    multiplier = 1
    for _ in range(num_batteries - 1):
        multiplier *= 10
    end = l - num_batteries + 1
    var p: Int = 0
    for loc in range(start_i, end):
        num = ord(s[loc]) - ord("0")
        # print("num: ", num)

        remaning = bank_movable(s, loc + 1, num_batteries - 1, cache)
        tp = num * multiplier + remaning
        # print("value in path: ", tp)
        if tp > p:
            p = tp

    cache[cache_idx] = p

    return p
