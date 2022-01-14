import random
import math

"""
$ python git:(python) ✗ time python3 thing.py
  13.326018
  7.393233140093008
  python3 thing.py  90.08s user 0.86s system 94% cpu 1:36.01 total
"""

print('hello team')

def play_itr() -> int:
    deck = list(range(52))
    random.shuffle(deck)
    hand = []
    for card in deck:
        hand.append(card)
        hand = try_reduce(hand)
    return len(hand)
    

def try_reduce(hand: list[int]) -> list[int]:
    if len(hand) < 4:
        return hand
    top = hand[-1]
    fourth = hand[-4]
    if same_rank(top, fourth):
        return try_reduce(hand[:-4])
    if same_suit(top, fourth):
        a_tmp = hand[:-3]
        a_tmp.append(hand[-1])
        return try_reduce(a_tmp)
    return hand

def same_rank(a: int, b: int) -> bool:
    return a%13 == b%13

def same_suit(a: int, b: int) -> bool:
    return int(a/13) == int(b/13)

def avg(a: list[int]) -> float:
    return sum(a) / len(a)

def stdev(a: list[int]) -> float:
    mean = avg(a)
    s = 0.0
    for num in a:
        s += (num - mean) ** 2
    return math.sqrt(s / len(a))

def main():
    iters = 1_000_000
    results = [play_itr() for i in range(iters)]
    print(avg(results))
    print(stdev(results))

if __name__ == '__main__':
    main()
