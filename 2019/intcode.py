import logging
from fuel import get_data
from typing import List
import time

LEN_OF_LINE = 4
OPERAND_SUM = 1
OPERAND_MUL = 2
OPERAND_EXIT = 99

REQ_RESULT = 19690720

ops = {
    OPERAND_SUM: lambda x, y: x + y,
    OPERAND_MUL: lambda x, y: x * y
}


def get_intcode_data():
    data = get_data(2, ',')
    for i, elem in enumerate(data):
        data[i] = int(elem)
        logging.debug(f'Intcode data: {elem}')
    return data

def find_vals(intcode):
    for x in range(0,100):
        for y in range(0,100):
            if test_intcode(list(intcode), x, y) == REQ_RESULT:
                logging.info((100 * x + y))
                return


def test_intcode(data, x, y):
    data[1] = x
    data[2] = y
    result = run_intcode(data)
    logging.debug(f'Testing {x}, {y}: {result[0]}')
    return result[0]


def run_intcode(program: List[int]):
    index = 0
    # OPCODE POS1 POS2 POSRES
    while index < len(program):
        if program[index] == OPERAND_EXIT:
            break
        else:
            program[program[index+3]] = ops[program[index]](program[program[index+1]], program[program[index+2]])
        index += LEN_OF_LINE
    return program


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    intcode = get_intcode_data()
    find_vals(intcode)
