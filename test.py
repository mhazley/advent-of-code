import unittest

from fuel import calculate_fuel_for_mass


class TestFuel(unittest.TestCase):
    def test_fuel_one(self):
        data = [1969]
        result = calculate_fuel_for_mass(data)
        self.assertEqual(result, 966)

    def test_fuel_two(self):
        data = [100756]
        result = calculate_fuel_for_mass(data)
        self.assertEqual(result, 50346)


if __name__ == '__main__':
    unittest.main()