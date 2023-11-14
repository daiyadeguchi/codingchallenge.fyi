import typer

def isValid(content: str) -> int:
	stack = []
	if not content:
		return 1

	for c in content:
		if c in '{':
			stack.append(c)
		else:
			if not stack or (c == '}' and stack[-1] != '{'):
				return 1
			stack.pop()
			
	return 0
	
def parseJson(file: str) -> int:
	file = open(file, "r")
	print(isValid(file.read()))
	return isValid(file.read())

typer.run(parseJson)
