//go:generate go run github.com/noho-digital/genny -in ../collection/set.go -out string_set_gen.go gen Type=string
//go:generate go run github.com/noho-digital/genny -in ../collection/type_slice.go -out string_set_slice_gen.go gen Type=StringSet
//go:generate go run github.com/noho-digital/genny -in ../collection/type_slice.go -out string_slice_gen.go gen Type=string
//go:generate go run github.com/noho-digital/genny -in ../collection/type_slice_distinct.go -out string_slice_distinct_gen.go gen Type=string
//go:generate go run github.com/noho-digital/genny -in ../collection/func_slice.go -out string_slice_slice_gen.go gen Func=StringSlice
package stringx
