package main

import (
	"log"

	"github.com/dialogs/api-schema/utils/protoc-gen-sdk/internal"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Lmicroseconds)

	req := command.Read()
	files := req.GetProtoFile()

	vanity.ForEachFile(files, vanity.TurnOffGogoImport)
	vanity.ForEachFile(files, vanity.TurnOnMarshalerAll)
	vanity.ForEachFile(files, vanity.TurnOnSizerAll)
	vanity.ForEachFile(files, vanity.TurnOnUnmarshalerAll)

	resp := internal.Generate(req)
	command.Write(resp)
}
