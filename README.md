# gobackend-test
Learning a bit about building backends in golang

## First stage: no ORM, no Database, just in-memory

[no-orm branch](https://github.com/marcolussetti/gobackend-test/tree/no-orm), based on several articles which were simply the highest results in google for `rest api golang`, perhaps not the best approach to selecting tutorials:
1. https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da was the initial article, whose code is severely broken and non-functional as is
2. https://tutorialedge.net/golang/creating-restful-api-with-golang/ which doesn't really deal with the update/POST which is actually the main think broken with the previous tutorial so....
3. I can't remember where I got the actually working code, had to clean it up a bit, but the `json.NewDecoder(r.Body).Decode(&newArticle)` approach works where the `Unmarshal` failed on the partial. No actual idea why.

Basically it's two files:
- `cmd/test/main.go` is where the code if you will is
- `pkg/models/setup.go` is where the model if you will is
