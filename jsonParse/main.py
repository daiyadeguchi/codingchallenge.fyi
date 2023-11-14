import typer

def isValid(content: str):
	if not content:
		return False 

	stack = []
	for c in content:
		if c == '{':
			stack.append(c)
		else:
			if not stack or (c == '}' and stack[-1] != '{'):
				return False 
	return True

def boolify(s: str):
	if s.lower() == "true":
		return True
	elif s.lower() == "false":
		return False
	
def parseJson(file: str):
	file = open(file, "r")
	content = file.read()
	# check if the file is valid json format
	if not isValid(content):
		return 1

  # clean the content. It sucks.
	content = content.replace('{', '')
	content = content.replace('}', '')
	content = content.replace('\n', '')
	content = content.replace(' ', '')
	#content = content.replace('"', '')
	content = content.split(',')

	json = {}	
	for c in content:
		c = c.split(":")
		c[0] = c[0].replace('"', '')

		if c[1] in {"true", "false"}:
			c[1] = boolify(c[1])

		if type(c[1]) != bool and c[1].isdigit():
			c[1] = int(c[1])

		if c[1] == "null":
			c[1] = None

		json[c[0]] = c[1]

	print(json)

	return 0

typer.run(parseJson)
