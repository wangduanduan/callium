.PHONY: fast-commit
fast-commit:
	git add -A
	git commit -m "wip"
	git push origin master
