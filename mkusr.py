import sys
from lxml import etree
arguments = sys.argv
users = open("users.xml","r")
stuff = etree.parse(users)
users.close()

users = open("users.xml","w")

page = etree.Element('invited')
doc = etree.ElementTree(page)
play = etree.SubElement(page,"player")

item1 = etree.SubElement(play,"name")
item2 = etree.SubElement(play,"username")
item3 = etree.SubElement(play,"password")

item1.text = arguments[1]
item2.text = arguments[2]
item3.text = arguments[3]

for x in stuff.xpath("//player"):
	page.append(x)
doc.write(users)
users.close()
