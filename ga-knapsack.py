# coding: utf-8

import random
import math
import copy
import operator
import pandas as pd

N_ITEMS = 50
ITEMS_PER_POP = 5
N_POP = 50
N_GEN = 25
MUTATE_PROB = 0.1
ELITE_RATE = 0.5

class GA:
    def __init__(self):
        self.items = {}
        self.fitness_master = {}

    def main(self): 
        pop = [{'param': p} for p in self.make_population()]

        # pop (N_POPの数だけ個体をもつ)
        # => [{'param': [(14, 7), (12, 4), (31, 7), (22, 10), (7, 8)]}, ... , {'param': [(2, 7), (36, 9), (20, 9), (28, 5), (2, 7)]}]
        
        for g in range(N_GEN):
            # print('Generation%3s:') % str(g), 

            # Get elites
            fitness = self.evaluate(pop)
            elites = fitness[:int(len(pop)*ELITE_RATE)]

            # Cross and mutate
            pop = elites[:]
            while len(pop) < N_POP:
                if random.random() < MUTATE_PROB:
                    m = random.randint(0, len(elites)-1)
                    print(elites[m]['param'])
                    exit()
                    child = self.mutate(elites[m]['param'])
                else:
                    c1 = random.randint(0, len(elites)-1)
                    c2 = random.randint(0, len(elites)-1)
                    child = self.crossover(elites[c1]['param'], elites[c2]['param'])
                pop.append({'param': child})
            
            # Evaluate indivisual
            fitness = self.evaluate(pop)
            pop = fitness[:]

            print(pop[0]['score0'], pop[0]['score1'], pop[0]['param'])

            
    def make_population(self):
        weights = [
            9, 7, 8, 2, 10, 7, 7, 8, 5, 4, 7, 5, 7, 5, 9, 9, 9, 8, 8, 2, 7, 7, 9, 8, 4, 7,
            3, 9, 7, 7, 9, 5, 10, 7, 10, 10, 7, 10, 10, 10, 3, 8, 3, 4, 2, 2, 5, 3, 9, 2
        ]
        prices = [
            20, 28, 2, 28, 15, 28, 21, 7, 28, 12, 21, 4, 31, 28, 24, 36, 33, 2, 25, 21, 35, 14, 36, 25, 12,
            14, 40, 36, 2, 28, 33, 40, 22, 2, 18, 22, 14, 22, 15, 22, 40, 7, 4, 21, 21, 28, 40, 4, 24, 21
        ]

        for i in xrange(N_ITEMS):
            self.items[i] = (prices[i], weights[i])

        # self.items
        # => {0: (20, 9), 1: (28, 7), ... , 48: (24, 9), 49: (21, 2)}

        pop = []
        for i in range(N_POP):
            # itemsからランダムに5つのアイテムを選択して，それを1つの母集団としている
            # この選び方が，最終結果を左右するかもしれない
            # 重さの上限が 60 なので，選択するアイテムの数をもう少し増やしてもいいかも
            # REVIEW: indexではなく，アイテム情報を保持するのはどうなのか？
            item = [self.items[k] for k in random.sample(range(N_ITEMS), ITEMS_PER_POP)]
            pop.append(item)

        return pop


    def clac_score(self, indivisual):
        dic = {}
        dic['score0'] = 0  # value
        dic['score1'] = 0  # weight
        for ind in indivisual:
            dic['score0'] += ind[0]
            dic['score1'] += ind[1]
            
        return dic


    def evaluate(self, pop):
        fitness = []
        for p in pop:
            if not p.has_key('score0'):
                # The indivisual made by crossover or mutation existed before
                if self.fitness_master.has_key(str(p['param'])):
                    p.update(self.fitness_master[str(p['param'])])
                # The indivisual is the first
                else:
                    p.update(self.clac_score(p['param']))
                fitness.append(p)
            else:
                fitness.append(p)

        # Save fitness to all genaration dictinary
        for fit in fitness:
            param = fit['param']
            self.fitness_master[str(param)] = {k:v for k,v in fit.items() if k!='param'}

        # This generation fitness
        df = pd.DataFrame(fitness)
        df = df.sort_values(['score0', 'score1'], ascending=[False, True])

        fitness = df.to_dict('records')
        
        return fitness


    def mutate(self, parent):
        ind_idx = int(math.floor(random.random()*len(parent)))
        item_idx = random.choice(range(N_ITEMS))
        child = copy.deepcopy(parent)
        child[ind_idx] = self.items[item_idx]

        return child


    def crossover(self, parent1, parent2):
        length = len(parent1)
        r1 = int(math.floor(random.random()*length))
        r2 = r1 + int(math.floor(random.random()*(length-r1)))
        
        child = copy.deepcopy(parent1)
        child[r1:r2] = parent2[r1:r2]

        return child


if __name__ == "__main__":
    GA().main()