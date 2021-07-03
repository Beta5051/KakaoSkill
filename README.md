# kakaoskill
go언어 카카오 i 오픈빌더 스킬 라이브러리

## Install
```shell
go install github.com/Beta5051/kakaoskill
```

## Example
```go
package main

import (
	"github.com/Beta5051/kakaoskill"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	kakaoskill.AddSkill(mux, "/", func(payload kakaoskill.Payload) kakaoskill.Response {
		return kakaoskill.Response{
			Template: &kakaoskill.Template{
				Outputs: []kakaoskill.Component{
					{
						SimpleText: &kakaoskill.SimpleText{
							Text: payload.UserRequest.Utterance,
						},
					},
				},
			},
		}
	})
	log.Fatalln(http.ListenAndServe(":8080", mux))
}
```