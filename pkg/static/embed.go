package static

import "embed"

//go:embed assets/*
var StaticFS embed.FS
