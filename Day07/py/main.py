import sys
import re
from collections import defaultdict, Counter

# Open the file and read its content
D = open(sys.argv[1]).read().strip()
# Split the content by newline
L = D.split('\n')

# Function to calculate the strength of a hand
def strength(hand, part2):
  # Replace card characters with their corresponding values
  hand = hand.replace('T',chr(ord('9')+1))
  hand = hand.replace('J',chr(ord('2')-1) if part2 else chr(ord('9')+2))
  hand = hand.replace('Q',chr(ord('9')+3))
  hand = hand.replace('K',chr(ord('9')+4))
  hand = hand.replace('A',chr(ord('9')+5))

  # Count the occurrence of each card
  C = Counter(hand)
  
  # If part2 is true, perform additional operations
  if part2:
    target = list(C.keys())[0]
    for k in C:
      if k!='1':
        if C[k] > C[target] or target=='1':
          target = k
    assert target != '1' or list(C.keys()) == ['1']
    if '1' in C and target != '1':
      C[target] += C['1']
      del C['1']
    assert '1' not in C or list(C.keys()) == ['1'], f'{C} {hand}'

  # Check the hand type and return its strength
  if sorted(C.values()) == [5]:
    return (10, hand)
  elif sorted(C.values()) == [1,4]:
    return (9, hand)
  elif sorted(C.values()) == [2,3]:
    return (8, hand)
  elif sorted(C.values()) == [1,1,3]:
    return (7, hand)
  elif sorted(C.values()) == [1,2,2]:
    return (6, hand)
  elif sorted(C.values()) == [1,1,1,2]:
    return (5, hand)
  elif sorted(C.values()) == [1,1,1,1,1]:
    return (4, hand)
  else:
    assert False, f'{C} {hand} {sorted(C.values())}'

# Loop over part2 values
for part2 in [False, True]:
  H = []
  # Loop over each line in L
  for line in L:
    hand,bid = line.split()
    H.append((hand,bid))
  # Sort H based on the strength of the hand
  H = sorted(H, key=lambda hb:strength(hb[0], part2))
  ans = 0
  # Calculate the total score
  for i,(h,b) in enumerate(H):
    ans += (i+1)*int(b)
  print(ans)