import os.path
leaderboard = ""

if os.path.isfile("temp.txt"):
	temp = open("temp.txt","r")
	leaderboard = temp.read()
	temp.close()

temp = open("temp.txt","w")
info = open("arena/info.txt","r")	
first = info.readline()
second = info.readline()
info.close()
if os.path.isfile("arena/error.txt"):
	print("error")
	error = open("arena/error.txt","r")
	loser = error.read()
	if loser == first:
		if not leaderboard.find(first):
			if not leaderboard.find(second):
				leaderboard += second+"\n"
				leaderboard += first+"\n"
			else:
				leaderboard += first+"\n"
		else:
			place = leaderboard.index(first)
			leaderboard = leaderboard[:place] + second + "\n" + leaderboard[place:]
	elif loser == second:
		if not leaderboard.find(second):
			if not leaderboard.find(first):
				leaderboard += first+"\n"
				leaderboard += second+"\n"
			else:
				leaderboard += second+"\n"
		else:
			place = leaderboard.index(second)
			leaderboard = leaderboard[:place] + first + "\n" + leaderboard[place:]
	temp.write(leaderboard)
	temp.close()
	exit()
data = open("arena/data.txt","r")
about = data.read()
data.close()
if about == "2":
	if leaderboard.find(first) == -1:
		if leaderboard.find(second) == -1:
			leaderboard += second
			leaderboard += first
		else:
			leaderboard += first
	else:
		place = leaderboard.index(first)
		leaderboard = leaderboard[:place] + second + leaderboard[place:]
elif about == "1":
	if leaderboard.find(second) == -1:
		if leaderboard.find(first) == -1:
			leaderboard += first
			leaderboard += second
		else:
			leaderboard += second
	else:
		place = leaderboard.index(second)
		leaderboard = leaderboard[:place] + first + leaderboard[place:]
temp.write(leaderboard)
temp.close()
