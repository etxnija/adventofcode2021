from kernel import bank_joltage
from testing import (
    assert_equal,
    assert_false,
    assert_raises,
    assert_true,
    TestSuite,
)


def test_bank_joltage():
    assert_equal(74, bank_joltage("6634733134", 0, 2))
    assert_equal(98, bank_joltage("987654321111111", 0, 2))
    assert_equal(987654321111, bank_joltage("987654321111111", 0, 12))

    assert_equal(89, bank_joltage("811111111111119", 0, 2))
    assert_equal(811111111119, bank_joltage("811111111111119", 0, 12))
    assert_equal(78, bank_joltage("234234234234278", 0, 2))
    assert_equal(434234234278, bank_joltage("234234234234278", 0, 12))
    assert_equal(92, bank_joltage("818181911112111", 0, 2))
    assert_equal(888911112111, bank_joltage("818181911112111", 0, 12))


def main():
    TestSuite.discover_tests[__functions_in_module()]().run()
