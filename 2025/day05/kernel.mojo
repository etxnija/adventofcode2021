from python import PythonObject
from python.bindings import PythonModuleBuilder
from os import abort, Atomic
from algorithm import parallelize
from int_help.int_manipulation import int_pow, length_of_int, int_at


@export
fn PyInit_kernel() -> PythonObject:
    try:
        var m = PythonModuleBuilder("kernel")
        m.def_function[available_py](
            "available_py", docstring="Parallel Solver"
        )
        return m.finalize()
    except e:
        return abort[PythonObject](String("error creating python module:", e))


fn available_py(py_obj: PythonObject) raises -> PythonObject:
    var mojo_data = List[String]()
    for py_str_item in py_obj:
        mojo_data.append(String(py_str_item))
    return available(mojo_data)


fn available(data: List[String]) raises -> Int:
    # Extract ranges from file and get the lowest and highest value
    # Then create an idex used to look up
    total = 0
    var ranges = List[Range]()
    iter = iter(data)
    line = iter.__next__()
    min = Int.MAX
    max = Int.MIN
    while True:
        clean = line.strip()
        if len(clean) == 0:
            break
        # print("Range: ", line)
        pair = line.split("-")
        low = atol(pair[0])
        high = atol(pair[1])
        ranges.append(Range(low, high))
        if low < min:
            min = low
        if high > max:
            max = high

        line = iter.__next__()

    print(" min: ", min, "max: ", max, " ranges: ", len(ranges))

    # Merge ranges
    print("Ranges before: ", len(ranges))
    merge_ranges(ranges)
    print("Ranges after: ", len(ranges))

    for r in ranges:
        total += r.count()
    # index = List[Bool](length=length, fill=False)
    # for r in ranges:
    #     mi = r[0] - min
    #     mx = r[1] - min
    #     for idx in range(mi, mx):
    #         index[idx] = True

    # while iter.__has_next__():
    #     line = iter.__next__()
    #     # print("number", line)
    #     num = atol(line)
    #     for r in ranges:
    #         if r.inRanage(num):
    #             total += 1
    #             break

    return total


@fieldwise_init
struct Range(Comparable, Copyable, Movable, Representable, Stringable):
    var start: Int
    var end: Int

    fn __lt__(self, other: Range) -> Bool:
        return self.start < other.start

    fn __eq__(self, other: Range) -> Bool:
        return self.start == other.start and self.end == other.end

    fn __repr__(self) -> String:
        return "[" + String(self.start) + " - " + String(self.end) + "]"

    fn __str__(self) -> String:
        return "[" + String(self.start) + " - " + String(self.end) + "]"

    fn inRanage(self, i: Int) -> Bool:
        return i >= self.start and i <= self.end

    fn count(self) -> Int:
        return self.end - self.start + 1


fn merge_ranges(mut ranges: List[Range]):
    # First sort on the start so it's get easier to find overlap
    # print("before sort: ", ranges.__str__())
    sort(ranges)
    # print("after sort: ", ranges.__str__())

    merged = List[Range]()
    # Need to start with one
    current_idx = 0

    # Compare with the rest to see if there is overlap

    for next_idx in range(1, len(ranges)):
        # print(
        #     "current: ",
        #     ranges[current_idx].__str__(),
        #     " next: ",
        #     ranges[next_idx].__str__(),
        # )
        # print(ranges.__str__())

        # Senario, Overlap next starts in current range, then update the current to the new end
        if ranges[next_idx].start <= ranges[current_idx].end + 1:
            ranges[current_idx].end = max(
                ranges[current_idx].end, ranges[next_idx].end
            )
        # No overlap, check then next and copy over the range that we porentially had modified above
        else:
            current_idx += 1
            ranges[current_idx] = ranges[next_idx].copy()

        # Now we have merged the ranges and can trim the range, ad default rangages to the end, but here should not be any empty slots
    ranges.resize(current_idx + 1, Range(0, 0))
