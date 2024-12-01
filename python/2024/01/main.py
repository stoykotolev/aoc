print(sum([abs(int(a) - int(b)) for a,b in zip(*map(lambda x: sorted(x), zip(*[line.split("  ") for line in open("input.txt").read().splitlines()])) )]))
