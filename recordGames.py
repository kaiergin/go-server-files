from lxml import etree

games = open("games.xml","r")
stuff = etree.parse(games)
games.close()

games = open("games.xml","w")

page = etree.Element('games')
doc = etree.ElementTree(page)
play = etree.SubElement(page,"play")

item1 = etree.SubElement(play,"playerone")
item2 = etree.SubElement(play,"playertwo")

info = open("arena/info.txt","r")
item1.text = info.readline()[:-1]
item2.text = info.readline()[:-1]
info.close()

counter = 0
for x in stuff.xpath("//play"):
	if counter > 100:
		break
	page.append(x)
	counter += 1
doc.write(games)
games.close()
