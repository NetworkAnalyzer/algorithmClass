# -*- coding: utf-8 -*-
import random
import math
import copy
import numpy as np

MAX_WEIGHT = 60
INITIAL_INDIVIDUAL_COUNT = 14
POPULATION_COUNT = 50
GENERATION_COUNT = 1000
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

def make_population():
    populations = []

    for i in range(0, POPULATION_COUNT):
        individual = random.sample(range(len(weights)), INITIAL_INDIVIDUAL_COUNT)
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

def select_parent_by_elite(populations, values):
    values = np.array(values)
    elite_index = np.argmax(values)

    return populations[elite_index]

def select_parent_by_roulette(populations, values):
    values = np.array(values)
    # 価値の大小関係を反転する
    values = np.abs(values - np.max(values))
    total = np.sum(values)

    threshold = random.uniform(0.0, total)

    sum = 0.0
    for index, value in enumerate(values):
        sum += value
        if sum >= threshold:
            parent_index = index
            break

    return populations[parent_index]

def crossover(parent_1, parent_2):
    child_1 = parent_1
    child_2 = parent_2

    for i in range(len(child_1)):
        if random.random() < 0.5:
            child_1[i], child_2[i] = child_2[i], child_1[i]

    return random.choice([child_1, child_2])

def mutate(parent):
    mutated_index = int(math.floor(random.random() * len(parent)))
    item_index = random.choice(range(len(weights)))

    child = copy.deepcopy(parent)
    child[mutated_index] = item_index

    return child

def change_generation(populations, values, child):
    values = np.array(values)
    bottom_index = np.argmin(values)

    populations[bottom_index] = child
    
    return populations

def print_populations(populations, count):
    values = []
    for individual in populations:
        weight = calc_total_value(individual, TYPE_WEIGHT)
        price  = calc_total_value(individual, TYPE_PRICE)
        values.append([weight, price])

    print('\nGeneration : %s' % count)

    for i in range(0, POPULATION_COUNT, 10):
        print(values[i:i + 10])

    return values

if __name__ == "__main__":
    populations = make_population()
    print_populations(populations, 0)

    for i in range(GENERATION_COUNT):
        values = evaluate(populations)
        parent_1 = select_parent_by_elite(populations, values)
        parent_2 = select_parent_by_roulette(populations, values)
        if random.random() < MUTATION_RATE:
            child = crossover(parent_1, parent_2)
        else:
            child = mutate(parent_1)

        populations = change_generation(populations, values, child)

    values = print_populations(populations, i + 1)

    weights = []
    prices = []

    for value in values:
        weights.append(value[0])
        prices.append(value[1])

    np_weights = np.array(weights)
    np_prices  = np.array(prices)

    np_weights = [weight if weight <= MAX_WEIGHT else 0 for weight in np_weights]
    max_weight_indices = [i for i, x in enumerate(np_weights) if x == max(np_weights)]
    max_price_indices = [i for i, x in enumerate(np_prices)  if x == max(np_prices)]
    
    print(max_weight_indices)
    print(max_price_indices)

    try:
        best_individual_index = list(set(max_price_indices) and set(max_weight_indices))[0]
        print(best_individual_index)
        print('The best individual is', values[best_individual_index])
    except:
        print('The best individual is not exists')