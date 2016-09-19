from lxml import etree

temp = open("temp.txt","r")
info = temp.read()
temp.close()
players = open("players.txt","w")
players.write(info)
players.close()

# Update XML leaderboard

leader = open("/var/www/html/leaderboard.xml","w")

page = etree.Element('board')
doc = etree.ElementTree(page)
play = etree.SubElement(page,"player")

players = open("players.txt","r")

num = 1

for x in players:
	item1 = etree.SubElement(play,"name")
	item2 = etree.SubElement(play,"rank")
	item1.text = x
	item2.text = num
	num += 1

players.close()
doc.write(leader)
leader.close()
