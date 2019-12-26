import unittest

from fuel import calculate_fuel_for_mass
from intcode import run_intcode
from wires import day_three


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

    def test_wire_one(self):
        data = [['R75', 'D30', 'R83', 'U83', 'L12', 'D49', 'R71', 'U7', 'L72'],
                ['U62', 'R66', 'U55', 'R34', 'D71', 'R55', 'D58', 'R83']]
        dist, steps = day_three(data)
        self.assertEqual(dist, 159)
        self.assertEqual(steps, 610)

    def test_wire_two(self):
        data = [['R98','U47','R26','D63','R33','U87','L62','D20','R33','U53','R51'],
                ['U98','R91','D20','R16','D67','R40','U7','R15','U6','R7']]
        dist, steps = day_three(data)
        self.assertEqual(dist, 135)
        self.assertEqual(steps, 410)

if __name__ == '__main__':
    unittest.main()