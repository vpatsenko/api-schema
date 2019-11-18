package internal

import (
	"bytes"
	"go/format"
	"log"
	"strings"
	"text/template"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	plugin_go "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
)

func Generate(req *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	// Begin by allocating a generator. The request and response structures are stored there
	// so we can do error handling easily - the response structure contains the field to
	// report failure.
	g := generator.New()
	g.Request = req

	g.CommandLineParameters(g.Request.GetParameter())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	g.WrapTypes()

	g.SetPackageNames()
	g.BuildTypeNameMap()

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	clientsImpl := NewImplementation()
	clientsImpl.Header.AddCustomImport("github.com/gogo/protobuf/types")

	for _, file := range filesForProcess(g) {
		genClient(g, file, clientsImpl)
	}

	source := getClientFileData(clientsImpl)
	formatted, err := format.Source(source)
	if err != nil {
		log.Fatal("failed to format output file", err, "\n", string(source))
	}

	clientsImpl.Header.FileName = NewFileName("sdk.pb.go")

	g.Response.File = []*plugin_go.CodeGeneratorResponse_File{&plugin_go.CodeGeneratorResponse_File{
		Name:    StringPtr(clientsImpl.Header.FullPath()),
		Content: StringPtr(string(formatted)),
	}}

	return g.Response
}

func getClientFileData(clientsImpl *Implementation) []byte {

	buf := bytes.NewBuffer(nil)
	if err := headerTemplate.Execute(buf, clientsImpl); err != nil {
		log.Fatal("failed to write header", err)
	}

	if err := dialogSDKTemplate.Execute(buf, clientsImpl); err != nil {
		log.Fatal("failed to write object description", err)
	}

	type MethodDesc struct {
		Method
		Service string
	}

	for _, s := range clientsImpl.Services {
		for _, m := range s.Methods {
			prop := MethodDesc{
				Method:  m,
				Service: s.ServiceName,
			}
			if (!m.ClientStreaming && m.ServerStreaming) || (m.ClientStreaming && !m.ServerStreaming) {
				if err := dialogSDKMethodTemplateStream.Execute(buf, prop); err != nil {
					log.Fatal("failed to write object method", err, m)
				}
			} else if m.ClientStreaming && m.ServerStreaming {
				if err := dialogSDKMethodTemplateClientServerStream.Execute(buf, prop); err != nil {
					log.Fatal("failed to write object method", err, m)
				}
			} else {
				if err := dialogSDKMethodTemplateRPC.Execute(buf, prop); err != nil {
					log.Fatal("failed to write object method", err, m)
				}
			}
		}
	}

	return buf.Bytes()
}

func genClient(g *generator.Generator, file *descriptor.FileDescriptorProto, clientsImpl *Implementation) {

	if len(file.Service) == 0 {
		return
	}

	imports, importRules := getImportsInfo(g, file)
	header := NewFileHeader(file)

	for _, svc := range file.Service {
		svcInfo := NewService(file, svc)

		for _, method := range svc.Method {
			inputImport, inputPkg, inputName := getType(file.GetPackage(), method.GetInputType(), imports, importRules)
			outputImport, outputPkg, outputName := getType(file.GetPackage(), method.GetOutputType(), imports, importRules)

			inputTypeName := inputName
			if inputPkg != "" {
				inputTypeName = inputPkg + "." + inputName
			}

			outputTypeName := outputName
			if outputPkg != "" {
				outputTypeName = outputPkg + "." + outputName
			}

			header.AddCustomImport(inputImport)
			header.AddCustomImport(outputImport)

			svcInfo.Methods = append(svcInfo.Methods, Method{
				Name:            method.GetName(),
				InType:          inputTypeName,
				OutType:         outputTypeName,
				ClientStreaming: method.GetClientStreaming(),
				ServerStreaming: method.GetServerStreaming(),
			})
		}

		if err := clientsImpl.AddService(svcInfo); err != nil {
			log.Fatal(err, "\n", svcInfo.ServiceName)
		}
	}

	if err := clientsImpl.Header.Merge(header); err != nil {
		log.Fatal(err, "\n", file.GetName())
	}
}

func getImportsInfo(g *generator.Generator, file *descriptor.FileDescriptorProto) (imports, importRules map[string]string) {

	imports = make(map[string]string)
	for _, item := range file.GetDependency() {
		importPkg, _ := splitByLast(item, "/")
		importPkgAlias := strings.ReplaceAll(importPkg, "/", ".")
		imports[importPkgAlias] = item
	}

	importRules = make(map[string]string)
	for _, item := range strings.Split(g.Request.GetParameter(), ",") {
		i := strings.Index(item, "=")
		if i == -1 {
			continue
		}

		key := strings.TrimSpace(item[:i])
		if !strings.HasPrefix(key, "M") {
			continue
		}
		key = strings.TrimPrefix(key, "M")

		importRules[key] = strings.TrimSpace(item[i+1:])
	}

	return
}

func getType(filePackage, protoType string, imports, importRules map[string]string) (newImport Import, typePkg, typeName string) {

	typePkg, typeName = SplitTypeName(generator.CamelCase(protoType))

	if typePkg != filePackage {
		pkg := imports[typePkg]
		if val, ok := importRules[pkg]; ok {
			pkg = val
		}

		newImport = Import(pkg)
		_, typePkg = splitByLast(pkg, "/")
	} else {
		typePkg = ""
	}

	return
}

func filesForProcess(g *generator.Generator) []*descriptor.FileDescriptorProto {

	countFiles := len(g.Request.FileToGenerate)
	retval := make([]*descriptor.FileDescriptorProto, 0, countFiles)

	for _, item := range g.Request.GetProtoFile() {
		name := item.GetName()
		exist := false
		for _, file := range g.Request.FileToGenerate {
			if file == name {
				exist = true
				break
			}
		}
		if !exist {
			continue
		}

		retval = append(retval, item)
	}

	return retval
}

var headerTemplate = template.Must(template.New("header").Parse(`
// Code generated by protoc-gen-gogo. DO NOT EDIT.

package {{.Package}}
import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	{{ range $key := .Header.GetImports }}{{ $key | printf "%s\n" }}{{end}}
)
`))

var dialogSDKTemplate = template.Must(template.New("dialog-sdk").Parse(`

type DialogSDK struct {
	Token              string
	internalContext    context.Context
	cancel             context.CancelFunc
	conn               *grpc.ClientConn
	{{range $i := .Services}}{{ $i.ServiceName | printf "client%[1]s %[1]sClient\n" }}{{end}}
}

func NewDialogSDK(cc *grpc.ClientConn) (*DialogSDK, error) {

	d := &DialogSDK{
		conn: cc,
		{{range $i := .Services}}{{ $i.ServiceName | printf "client%[1]s: New%[1]sClient(cc),\n" }}{{end}}
	}

	if err := d.initializeContext(); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *DialogSDK) initializeContext() error {
	ctx, cancel := context.WithCancel(context.Background())
	res, err := d.clientRegistration.RegisterDevice(ctx, &RequestRegisterDevice{})
	if err != nil {
		return err
	}
	d.internalContext = metadata.AppendToOutgoingContext(ctx, "x-auth-ticket", res.Token)
	d.Token = res.Token
	d.cancel = cancel
	return nil
}

func (d *DialogSDK) Close() {
	d.cancel()
	d.conn.Close()
}
`))

var dialogSDKMethodTemplateRPC = template.Must(template.New("dialog-sdk-method-rpc").Parse(`
func (d *DialogSDK) {{.Service}}{{.Name}}(in *{{.InType}}, opts ...grpc.CallOption) (*{{.OutType}}, error) {
	return d.client{{.Service}}.{{.Name}}(d.internalContext, in, opts...)
}
`))

var dialogSDKMethodTemplateStream = template.Must(template.New("dialog-sdk-method-stream").Parse(`
func (d *DialogSDK) {{.Service}}{{.Name}}(in *{{.InType}}, opts ...grpc.CallOption) ({{.Service}}_{{.Name}}Client, error) {
	return d.client{{.Service}}.{{.Name}}(d.internalContext, in, opts...)
}
`))

var dialogSDKMethodTemplateClientServerStream = template.Must(template.New("dialog-sdk-method-client-server").Parse(`
func (d *DialogSDK) {{.Service}}{{.Name}}(opts ...grpc.CallOption) ({{.Service}}_{{.Name}}Client, error) {
	return d.client{{.Service}}.{{.Name}}(d.internalContext, opts...)
}
`))
