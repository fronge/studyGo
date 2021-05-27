package netrpc

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"studyGo/rpcT/protoc-gen-go-netrpc/generator"
	// "github.com/golang/protobuf/protoc-gen-go/generator"
)

type netrpcPlugin struct {
	*generator.Generator
}

func (p *netrpcPlugin) Name() string {
	return "netrpc"
}

func (p *netrpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor, imports map[generator.GoImportPath]generator.GoPackageName) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}

func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P(`import "net/rpc"`)
}

func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	p.P("//TODO : service code, name = " + svc.GetName())
}

func init() {
	generator.RegisterPlugin(new(netrpcPlugin))
}
