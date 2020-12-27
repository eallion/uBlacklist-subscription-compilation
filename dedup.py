f = open("blacklist.txt", "r")
res = set()
for i in f.readlines():
    if len(i) > 1:
        res.add(i)
f.close()
print("loaded", len(res), "records")

f = open("blacklist.txt", "w")
res = list(res)
res.sort()
for i in res:
    f.write(i)
print("saved", len(res), "records")
f.close()