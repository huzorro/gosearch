package main

func main() {
	mtn := martini.Classic()
	mtn.Use(render.Renderer())
	mtn.Use(gosearch.SearcherMiddle)
	mtn.Get("/", func(r render.Render) {
		r.HTML(200, "search", "search")
	})
	mtn.Get("/search", gosearch.SearchHandler)
	mtn.Run()
}
