import sys
arguments = sys.argv

#Adds AI to players.txt if it does not exist

players = open('players.txt','r')
doesExist = False
hello = str(arguments[1])
for x in players:
	if x[:-1] == hello[7:len(arguments[1])-3]:
		doesExist = True
players.close()
if not doesExist:
	players = open('players.txt','a')
	players.write(hello[7:len(arguments[1])-3])
	players.close()

#Checks if AI has illegal attributes

AI = open(hello,"r")
info = AI.read()


print(hello[7:len(arguments[1])-3]
