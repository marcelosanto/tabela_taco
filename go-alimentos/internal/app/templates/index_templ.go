// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.646
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "go-alimentos/internal/app/templates/components"

func Index(contents templ.Component, erroSearch bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><script src=\"https://unpkg.com/htmx.org@1.9.8\"></script><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/bulma@1.0.0/css/bulma.min.css\"><link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css\"><style>\n    /* Estilo para esconder os detalhes inicialmente */\n    .taco-details {\n      display: none;\n    }\n    /* Estilo para mostrar os detalhes quando o item da lista é clicado */\n    .taco-item.expanded .taco-details {\n      display: block;\n    }\n  </style></head><body><nav class=\"navbar is-primary\" role=\"navigation\" aria-label=\"main navigation\"><div class=\"navbar-brand\"><a class=\"navbar-item has-text-centered\" href=\"#\"><strong>Tabela de Lista de Alimentos</strong></a></div></nav><div class=\"box m-1\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Search().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if erroSearch {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"icon-text has-text-danger\"><span class=\"icon has-text-danger\"><i class=\"fas fa-ban\"></i></span> <span>Alimento procurado não encontrado</span></span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"content columns is-mobile is-centered\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = contents.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
