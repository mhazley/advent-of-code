import logging
from fuel import get_data

def do_path(data):
    cur_x = cur_y = steps = 0
    points = {}
    directions = {'U': (0,1),
                  'D': (0,-1),
                  'R': (1,0),
                  'L': (-1,0)}

    for instr in data:
        dir = instr[0]
        dis = int(instr[1:])

        dx, dy = directions[dir]

        for _ in range(dis):
            cur_x += dx
            cur_y += dy
            steps += 1
            if (cur_x, cur_y) not in points:
                points[(cur_x, cur_y)] = steps

    return points


def intersect_paths(path_one, path_two):
    return [value for value in path_one if value in path_two]


def day_three(data):
    path_one = do_path(data[0])
    path_two = do_path(data[1])

    logging.info(f'Intersecting...')

    intersections = intersect_paths(path_one, path_two)

    logging.info(f'Intersections: {intersections}')

    distance = 0xFFFF
    steps = 0xFFFF

    for point in intersections:
        d = compute_manhattan_distance(point)
        distance = d if d < distance and d is not 0 else distance

        point_steps = path_one[point] + path_two[point]
        steps = point_steps if point_steps < steps and point_steps is not 0 else steps

    return distance, steps


def compute_manhattan_distance(point):
    return abs(point[0]) + abs(point[1])


def get_wire_data():
    data = get_data(3)
    for i in range(0, len(data)):
        data[i] = data[i].split(sep=',')
    return data


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    data = get_wire_data()
    # data = [['R75', 'D30', 'R83', 'U83', 'L12', 'D49', 'R71', 'U7', 'L72'],
    #         ['U62', 'R66', 'U55', 'R34', 'D71', 'R55', 'D58', 'R83']]
    dist, steps = day_three(data)
    logging.info(f'Smallest Manhattan Distance: {dist}')
    logging.info(f'Smallest Step Distance: {steps}')