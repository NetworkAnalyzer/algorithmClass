# -*- coding: utf-8 -*-
import random
import numpy as np

MAX_WEIGHT = 60
POPULATION_COUNT = 50
GENERATION_COUNT = 30000
CROSSING_RATE = 0.8
MUTATION_RATE = 0.3

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
        individual = np.random.randint(0, 2, 50)
        populations.append(individual) 

    return populations

if __name__ == "__main__":
    populations = make_population()