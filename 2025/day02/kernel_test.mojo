from kernel import repeted2
from testing import (
    assert_equal,
    assert_false,
    assert_raises,
    assert_true,
    TestSuite,
)


def test_repeted2():
    assert_true(repeted2(1188511881))


def main():
    TestSuite.discover_tests[__functions_in_module()]().run()
