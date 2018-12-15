package main

//kod pisatj tut

import (
	"go/token"
	"go/parser"
	"go/ast"
)

type field struct {
	FieldName string
	tag
}

type handler struct {
	HandlerMethod string
	handler_meta
	ParamIn       string
	ResultOut     string
	ParamInStruct []field
}

type apigenApi struct {
	Srv         map[string][]handler
	Validator   map[string][]field // StructName ==> ["FieldName, tags"]
	IsInt       bool
	IsParamName bool
}




func genFunction(f *ast.FuncDecl, apigenapi *apigenApi) {
	var srv string
	h := handler{}

	if f.Doc == nil {
		return
	}
	for _,comment := range f.Doc.List {
		if strings.HasPrefix(comment.Text, "// apigen:api") {

			hm := handler_meta{}

			h.HandlerMethod = strings.ToLower(f.Name.Name)

			apigenDoc := strings.TrimLeft(comment, "// apigen:api")
			if len(s) > 1 {

			}

		

		}



	}





}


func parseStruct(currTypeName, tagValue, fieldName string, isInt bool, 
	apigenapi *apigenapi ){
	f := field{}
	v := strings.Split(strings.Replace(strings.Trim(tagValue, "/`"), "\"", "", -1), ",")
	for _,v := range v {
		if isInt {
			f.IsInt = true
		} else {
		    f.IsInt = false
		}
		if value == "required" {
			f.Required = true
			f.FieldName = fieldName
		}
		s := strings.Split(value, "=")
		if len(s) > 1 {
			//check for enum
			enums := strings.Split(s[1], "|")
			if len(enums) > 1 {
				for _, enum := range enums{
					f.FieldName = fieldName
					f.Enum = append(enums, enum)
				}

			}
		
			switch s[0] {
				case "min":
					f.FieldName = fieldName
					f.Min = s[1]
				case "max":
					f.FieldName = fieldName
					f.Max = s[1]
				case "paramname":
					f.FieldName = fieldName
					f.ParamName = s[1]
				case "default":
					f.FieldName = fieldName
					f.Default = s[1]
				
			}

		}
	}
	apigenapi.Validator[currTypeName] = append(apigenapi.Validator[currTypeName], f)

}
	




func iterateStruct(currStruct *ast.StructType, currType *ast.TypeSpec, apigenapi *apigenApi ) {
    var isInt bool
	for _, field := range currStruct.Fields.List {
		if field.Tag != nil {
			if strings.HasPrefix(field.Tag.Value,"`apivalidator:") {
				tagValue := strings.TrimLeft(field.Tag.Value,"`apivalidator:")
			}
			if field.Type.(*ast.Ident).Name == "int" {
				apigenapi.IsInt = true
				isInt = true
			} else {
				isInt = false
			}
			if strings.Contains(field.Tag.Value, "paramname") {
				apigenapi.IsParamName = true
			}
			parseStruct(currType.Name.Name, 
				tagValue, field.Names[0].Name, isInt, apigenapi)
		}
	}
}



func genStruct(f *ast.GenDecl, apigenapi *apigenApi) {
	for _, spec := range f.Specs {
		currType, ok:= spec.(*ast.TypeSpec)
		if !ok {
			continue
		}
		currStruct, ok:= currType.Type.(*ast.StructType)
		if !ok {
			continue
		}
		iterateStruct(currStruct, currType, apigenapi)
	}

}


func main() {
	fset := token.NewFileSet();
	node, _ := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments);

	out, _ := os.Create(os.Args[2])
	fmt.Fprintln(out, `package `+node.Name.Name)

	apigenapi := apigenApi{make(map[string][]handler), make(map[string][]field), false, false}
   
	for _, f  := range  node.Decls {
		switch a := f.(type) {
		case *ast.GenDecl:
			genStruct(a, &api)
		case *ast.FuncDecl:
			genFunction(a, &apigenapi)
		default:
			continue
		}
	}
	

	
}