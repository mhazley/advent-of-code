import logging
import requests
from typing import List


def read_cookie():
    with open('cookie', 'r') as f:
        c = f.read()
        logging.debug(f'Read Cookie from file: {c}')
        return c


def get_mass_data(cookie):
    data = get_data(1, cookie)
    for i, elem in enumerate(data):
        data[i] = int(elem)
        logging.debug(f'Mass data: {elem}')
    return data


def get_data(day, cookie):
    response = requests.get(f'https://adventofcode.com/2019/day/{day}/input', headers={'Cookie': f'session={cookie}'})
    logging.debug(f'Response data: {response.content}')
    return response.content.decode().splitlines(keepends=False)


def calculate_fuel_for_mass(masses: List[int]):
    total_fuel = 0
    for mass in masses:
        total_fuel += fuel(mass)
        if fuel(mass) <= 0:
            return 0
        else:
            logging.debug(f'Computed Fuel: {total_fuel}')
            total_fuel += calculate_fuel_for_mass([fuel(mass)])
    return total_fuel


def fuel(x):
    return (x // 3) - 2


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    cookie = read_cookie()
    masses = get_mass_data(cookie)
    total_fuel = calculate_fuel_for_mass(masses)
    logging.info(f'Fuel Required for input Mass: {total_fuel}')
