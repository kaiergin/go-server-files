import os
import fileinput

listfiles = os.listdir("arena")
y = 0
info = open("arena/info.txt","w")

for x in listfiles:
	isLegal = False
	if x != "engine.go":
		y+=1
		temp = open("arena/"+x,"r")
		new = open("arena/"+x[0:len(x)-4]+".go","w")
		for line in temp:
			if line.find("func returnMove(") != -1:
				isLegal = True
			line = line.replace("returnMove", "returnMove"+str(y))
			new.write(line)
		temp.close()
		new.close()
		if not isLegal:
			lose = open("arena/error.txt","w")
			lose.write(x[0:len(x)-4])
			lose.close()
			break
		info.write(x[0:len(x)-4]+"\n")
info.close()
