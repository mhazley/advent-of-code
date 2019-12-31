import logging

RANGE_MIN = 138307
RANGE_MAX = 654504


def test_password(password: int) -> bool:
    num_list = [int(x) for x in str(password)]
    logging.debug(f'is_six_digit: {num_list}')
    if len(num_list) != 6:
        return False

    prev_num = None
    adjacents = []
    dead = []
    for num in num_list:
        logging.debug(f'walking list: {prev_num}, {num}')
        if prev_num is not None and num < prev_num:
            logging.debug(f'did decrease')
            return False
        if num == prev_num and \
           num not in adjacents and \
           num not in dead:
            logging.debug(f'{num} is adjacent')
            adjacents.append(num)
        elif num == prev_num and \
           num in adjacents:
            logging.debug(f'{num} is in adjacents list already so adding to dead list')
            adjacents.remove(num)
            dead.append(num)
        prev_num = num

    return True if len(adjacents) > 0 else False


if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    num_valid = 0
    for num in range(RANGE_MIN, RANGE_MAX+1):
        num_valid += 1 if test_password(num) else 0
    logging.info(f'Number of valid passwords: {num_valid}')
    # logging.basicConfig(level=logging.DEBUG)
    # logging.info(test_password(123333))
