#! /usr/bin/env python3

with open(".data/day1.txt") as f:
  pos = 50
  times = 0

  for line in f:
    lr = line[0]
    number = int(line[1:])

    if lr == 'L':
      pos = (pos - number) % 100
    elif lr == 'R':
      pos = (pos + number) % 100

    if pos == 0:
      times += 1

  print(times)
