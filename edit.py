import os
import fileinput

listfiles = os.listdir("arena")
y = 0

for x in listfiles:
	if x != "engine.go":
		y+=1
		temp = open("arena/"+x,"r")
		new = open("arena/"+x[0:len(x)-4]+".go","w")
		for line in temp:
			line = line.replace("returnMove", "returnMove"+str(y))
			new.write(line)
		temp.close()
		new.close()
