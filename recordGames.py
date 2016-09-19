from lxml import etree
import time

games = open("/var/www/html/games.xml","r")
stuff = etree.parse(games)
games.close()

games = open("/var/www/html/games.xml","w")

page = etree.Element('all')
doc = etree.ElementTree(page)
play = etree.SubElement(page,"game")

item1 = etree.SubElement(play,"playerone")
item2 = etree.SubElement(play,"playertwo")
item3 = etree.SubElement(play,"winner")
item4 = etree.SubElement(play,"date")

info = open("arena/info.txt","r")
one = info.readline()[:-1]
two = info.readline()[:-1]

item1.text = one
item2.text = two
info.close()

info = open("arena/data.txt")
winner = info.readline()[:-1]
if winner == "1":
	item3.text = one
elif winner == "2":
	item3.text = two

item4.text = time.strftime("%d/%m/%Y")

counter = 0
for x in stuff.xpath("//play"):
	if counter > 100:
		break
	page.append(x)
	counter += 1
doc.write(games)
games.close()
