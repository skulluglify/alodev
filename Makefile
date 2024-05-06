GITHUB_URL := git@github.com:skulluglify/alodev-backend.git

git-init:
	git init .
	git config --local core.autocrlf false
	git config --local core.eol lf
	git branch -m main
	git add .
	git commit -am '[fix] initial commit'
	git remote add origin ${GITHUB_URL}
	git fetch origin main
	git branch --set-upstream-to=origin/main main
	git push -f origin main

git-pull:
	git pull
	git submodule update --init --recursive

git-commit:
	git add .
	git commit -am '[fix] initial commit'

git-push:
	git push -u origin main

run: build
	build/app.exe

build: clean
	mkdir build
	go build -o build/app.exe .

clean:
	pwsh Scripts/Removal.ps1
