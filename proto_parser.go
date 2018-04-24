package main

type Proto struct {
	syntax   string     `"syntax" "=" String  ";"`
	imprt    string     `["import" String ";"]`
	option   string     `[{"option" Ident "=" {String | Ident} ";"}]`
	pkg      string     `"package" Ident ";" `
	Services []*Service `{ @@ }`
	Message  string     `{ "message" Ident "{" { Ident Ident {Ident} "=" Int ";" } "}" }`
}

type Service struct {
	Name    string    `"service" @Ident "{"`
	Methods []*Method `{ @@ }`
	Closing string    `"}"`
}

type Method struct {
	Name    string `"rpc" @Ident `
	Input   string `"(" @Ident ")"`
	Output  string `"returns" "(" @Ident ")"`
	Closing string `"{""}"`
}

type Input struct {
	Identifier string `@Ident`
}

type Output struct {
	Identifier string `@Ident`
}