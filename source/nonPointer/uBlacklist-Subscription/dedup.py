f = open("blacklist.txt", "r")
res = set()
for i in f.readlines():
    res.add(i)
f.close()
print("loaded", len(res), "records")

f = open("blacklist.txt", "w")
for i in res:
    f.write(i)
f.close()