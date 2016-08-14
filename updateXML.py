from lxml import etree

# Leaderboard
page = etree.Element('players')
doc = etree.ElementTree(page)
myfile = open('players.txt','r')
for x in myfile:
	item = etree.SubElement(page,'name')
	item.text = x[:-1]
myfile.close()
XMLfile = open('/var/www/html/info/leaderboard.xml','w')
doc.write(XMLfile)
XMLfile.close()
