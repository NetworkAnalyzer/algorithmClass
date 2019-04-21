# -*- coding: utf-8 -*-
import random
import math
import copy
import numpy as np

MAX_WEIGHT = 60
POPULATION_COUNT = 50
GENERATION_COUNT = 30000
CROSSING_RATE = 0.8
MUTATION_RATE = 0.3
TYPE_WEIGHT = 1
TYPE_PRICE = 2
MIN_VALUE = 1

weights = [
    9, 7, 8, 2, 10, 7, 7, 8, 5, 4, 7, 5, 7, 5, 9, 9, 9, 8, 8, 2, 7, 7, 9, 8, 4, 7,
    3, 9, 7, 7, 9, 5, 10, 7, 10, 10, 7, 10, 10, 10, 3, 8, 3, 4, 2, 2, 5, 3, 9, 2
]
prices = [
    20, 28, 2, 28, 15, 28, 21, 7, 28, 12, 21, 4, 31, 28, 24, 36, 33, 2, 25, 21, 35, 14, 36, 25, 12,
    14, 40, 36, 2, 28, 33, 40, 22, 2, 18, 22, 14, 22, 15, 22, 40, 7, 4, 21, 21, 28, 40, 4, 24, 21
]

# 母集団を生成する
def make_population():
    populations = []

    for i in range(0, POPULATION_COUNT):
        individual = random.sample(range(len(weights)), 7)
        populations.append(individual)

    return populations

def evaluate(populations):
    values = []

    for individual in populations:
        total_weight = calc_total_value(individual, TYPE_WEIGHT)
        total_value = calc_total_value(individual, TYPE_PRICE)
        
        if total_weight <= MAX_WEIGHT:    
            values.append(total_value)
        else:
            values.append(MIN_VALUE)

    return values

def calc_total_value(array, type):
    sum = 0

    if type == TYPE_WEIGHT:
        for i in array:
            sum += weights[i]
    else:
        for i in array:
            sum += prices[i]

    return sum

def select_parents(populations, values):
    values = np.array(values)
    # 価値の大小関係を反転する
    values = np.abs(values - np.max(values))

    total = np.sum(values)

    parent_indices = []
    for _ in range(2):
        threshold = random.uniform(0.0, total)

        sum = 0.0
        for index, value in enumerate(values):
            sum += value
            if sum >= threshold:
                parent_indices.append(index)
                # 選択した個体が次のforで選択されないように
                values[index] = 0.
                total -= value
                break

    return populations[parent_indices[0]], populations[parent_indices[1]]

def crossover(parent_1, parent_2):
    length = len(parent_1)
    r1 = int(math.floor(random.random() * length))
    r2 = r1 + int(math.floor(random.random() * (length - r1)))
    
    child = copy.deepcopy(parent_1)
    child[r1:r2] = parent_2[r1:r2]

    return child

if __name__ == "__main__":
    populations = make_population()
    values = evaluate(populations)
    parent_1, parent_2 = select_parents(populations, values)
    print(parent_1)
    print(parent_2)
    child = crossover(parent_1, parent_2)
    print(child)