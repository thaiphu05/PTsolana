// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package viewCommon

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func ClickAvt() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_ClickAvt_f79b`,
		Function: `function __templ_ClickAvt_f79b(){let profileMenu = document.getElementById("profileMenu");
	console.log("Clicked");
	profileMenu.classList.toggle("hidden");
}`,
		Call:       templ.SafeScript(`__templ_ClickAvt_f79b`),
		CallInline: templ.SafeScriptInline(`__templ_ClickAvt_f79b`),
	}
}

func Navbar() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"w-full bg-primary pt-5 pb-3 px-16 text-xl text-foreground flex justify-between fixed top-0 left-0 min-h-2.5\"></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
