package main

import (
	"fmt"
	"goJson"

	_ "../"
)

func main() {
	json := `{
				"count": 2, 
				"videos": {
					"total":168133,
					"perPage":5,
					"results":
					[
						{	
							"youtubeUrl":"http://www.youtube.com/watch?v=kXYiU_JCYtU",
							"thumbnailUrl":"http://video-tub-ru.yandex.net/i?id=18672620-02-96",
							"title":"Linkin Park - Numb - YouTube",
							"duration":187,
							"text":"Linkin Park - Numb - YouTube."
						},
						{
							"youtubeUrl":"http://www.youtube.com/watch?v=oIwWqYSbzGA",
							"thumbnailUrl":"http://video-tub-ru.yandex.net/i?id=19614220-22-96",
							"title":"Linkin Park - In The End - YouTube",
							"duration":218,
							"text":"...linkinpark.com/ Like Linkin Park at http..."
						}
					]
				}
			}`
	js, _ := goJson.ParseJson([]byte(json))
	total := js.Int("count")
	videos := js.Get("videos").GetAll("results")
	fmt.Printf("Total count:\t%v\n\n", total)
	for _, element := range videos {
		title := element.String("title")
		url := element.String("youtubeUrl")
		duration := element.Int("duration")
		fmt.Printf("Title:\t%s\nVideo url:\t%s\nDuration:\t%v\n\n", title, url, duration)
	}

}
