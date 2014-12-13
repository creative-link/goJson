goJson
======

Fast and easy json parsing without the need to create data structures

## Install

```shell
go get github.com/creative-link/goJson
```

## Import

```go
import (
  "github.com/creative-link/goJson"
)
```

## Usage
#### Create goJson parser and handling errors
```go
parser, err := goJson.ParseJson([]byte(jsonString))
if err != nil {
	panic(fmt.Sprintf("JSON error: %v", err))
}
```
#### Read values
```go
intValue 	 := js.Int("jsonInteger")  // e.g. 2
floatValue 	 := js.Float("jsonFloat")  // e.g. 2.24
stringValue 	 := js.String("jsonString")  // e.g. "My string"
stringsArray := js.Strings("jsonStringsArray")  // e.g. "collection": ["one", "two" ...]
objectValue 	 := js.Get("jsonSingleObject")  // e.g. "jsonSingleObject" : { /*our object*/ }
objectsArray := js.GetAll("jsonObjects")  // e.g. "jsonObjects" : [{ /*our object*/ }, ...]
```

## Example
```go
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

// Start parsing our JSON 
js, err := goJson.ParseJson([]byte(json))

if err != nil {
	panic(fmt.Sprintf("JSON parsing error: %v", err))
}

// Getting integer value `count`
total := js.Int("count")
// Getting object array `results`
videos := js.Get("videos").GetAll("results")
fmt.Printf("Total count:\t%v\n\n", total)
// Fetch `results` array
for _, element := range videos {
	title := element.String("title")
	url := element.String("youtubeUrl")
	duration := element.Int("duration")
	fmt.Printf("Title:\t%s\nVideo url:\t%s\nDuration:\t%v\n\n", title, url, duration)
}
```

## TODO
* Raed json as array of integers
* Error handling
* Comments in code
* Tests