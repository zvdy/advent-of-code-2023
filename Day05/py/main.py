import sys
import re
from collections import defaultdict

# Read the input file
D = open(sys.argv[1]).read().strip()
L = D.split('\n')

parts = D.split('\n\n')
seed, *others = parts
# Parse the seed values
seed = [int(x) for x in seed.split(':')[1].split()]

class Function:
  def __init__(self, S):
    # Parse the function tuples
    self.tuples: list[tuple[int,int,int]] = [[int(x) for x in line.split()] for line in S.split('\n')[1:]]

  def apply_one(self, x: int) -> int:
    # Apply the function to a single value
    for (dst, src, sz) in self.tuples:
      if src<=x<src+sz:
        return x+dst-src
    return x

  def apply_range(self, R):
    # Apply the function to a range of values
    A = []
    for (dest, src, sz) in self.tuples:
      src_end = src+sz
      # NR is the new range after applying the function
      NR = []
      while R:
        (st,ed) = R.pop()
        before = (st,min(ed,src))
        inter = (max(st, src), min(src_end, ed))
        after = (max(src_end, st), ed)
        if before[1]>before[0]:
          NR.append(before)
        if inter[1]>inter[0]:
          A.append((inter[0]-src+dest, inter[1]-src+dest))
        if after[1]>after[0]:
          NR.append(after)
      R = NR
    return A+R

# Parse the functions
Fs = [Function(s) for s in others]

P1 = []
# Apply each function to each seed value and store the results
for x in seed:
  for f in Fs:
    x = f.apply_one(x)
  P1.append(x)
# Print the minimum result
print(min(P1))

P2 = []
pairs = list(zip(seed[::2], seed[1::2]))
# Apply each function to each pair of seed values and store the results
for st, sz in pairs:
  R = [(st, st+sz)]
  for f in Fs:
    R = f.apply_range(R)
  P2.append(min(R)[0])
# Print the minimum result
print(min(P2))