clean:
	@rm -rf dist
	@mkdir -p dist

dependencies:
	go get github.com/veandco/go-sdl2

build: clean
	go build -o dist/main