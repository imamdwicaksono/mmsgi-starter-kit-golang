package generators

type GenerateOptions struct {
	UseAPI  bool
	UseAuth bool
	UseCRUD bool
	UI      string // "tailwind" | "bootstrap" | "none"
}
