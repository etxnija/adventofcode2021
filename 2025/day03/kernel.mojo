from python import PythonObject
from python.bindings import PythonModuleBuilder
from os import abort, Atomic
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
    var total = Atomic[DType.int64](0)
    # for line in data:
    #     total += bank_joltage(line, 2)
    num_batteries = 12

    @parameter
    fn worker(i: Int):
        try:
            line = data[i]
            cache = List[Int]()
            cache_size = (len(line) + 1) * (num_batteries + 1)
            for _ in range(cache_size):
                cache.append(-1)

            result = bank_joltage(line, 0, 12, cache)
            # print("result : ", result, "from line:", line)
            total.fetch_add(Int64(result))
        except:
            pass

    parallelize[worker](len(data))

    return Int(total.load())


fn bank_joltage(
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

        remaning = bank_joltage(s, loc + 1, num_batteries - 1, cache)
        tp = num * multiplier + remaning
        # print("value in path: ", tp)
        if tp > p:
            p = tp

    cache[cache_idx] = p

    return p
