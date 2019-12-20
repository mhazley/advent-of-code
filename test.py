import unittest

from fuel import calculate_fuel_for_mass
from intcode import run_intcode


class TestFuel(unittest.TestCase):
    def test_fuel_one(self):
        data = [1969]
        result = calculate_fuel_for_mass(data)
        self.assertEqual(result, 966)

    def test_fuel_two(self):
        data = [100756]
        result = calculate_fuel_for_mass(data)
        self.assertEqual(result, 50346)

    def test_intcode_one(self):
        data = [1,0,0,0,99]
        result = run_intcode(data)
        self.assertEqual(result, [2,0,0,0,99])

    def test_intcode_two(self):
        data = [1,1,1,4,99,5,6,0,99]
        result = run_intcode(data)
        self.assertEqual(result, [30,1,1,4,2,5,6,0,99])

if __name__ == '__main__':
    unittest.main()