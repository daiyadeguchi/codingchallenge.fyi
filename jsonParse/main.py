import typer

def isValid(content: str) -> bool:
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
	
def parseJson(file: str) -> int:
	file = open(file, "r")
	content = file.read()
	# check if the file is valid json format
	if not isValid(content):
		return 1

	print(content)
	return 0

typer.run(parseJson)
