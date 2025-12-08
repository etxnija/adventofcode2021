import unittest
import sys
import os
from pathlib import Path
from typing import NamedTuple, List

sys.path.insert(0, str(Path(__file__).parent.parent))

from read_data import read_input
import mojo.importer
import kernel

class TestCase(NamedTuple):
    name: str
    file: str
    expected: int


TEST_CASES: List[TestCase] = [
    TestCase(
        name="example",
        file="example.txt",
        expected=142
    )
]
# kernel.calculate_distance_py()


class TestDay01(unittest.TestCase):
    pass
def create_test_method(case: TestCase):

    def test_method(self):
        print("Creating test:", case.name)
        input_data = read_input(case.file)
        if not input_data:
            self.skipTest(f"Skipping test '{case.name}': Input data not found")

        actual_result = kernel.calculate_distance_py(input_data)

        self.assertEqual(
            actual_result,
            case.expected,
            f"Test '{case.name}' failed. Expected {case.expected}, got {actual_result}"
        )
    test_method.__name__ = f"test_{case.name}"
    return test_method

for case in TEST_CASES:
    test_func = create_test_method(case)
    setattr(TestDay01, test_func.__name__, test_func)


if __name__ == '__main__':
    unittest.main(argv=['first-arg-is-ignored'], exit=False)
