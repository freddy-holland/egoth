// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Auth() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div x-data=\"{\n        tabSelected: 1,\n        tabId: $id(&#39;tabs&#39;),\n        tabButtonClicked(tabButton){\n            this.tabSelected = tabButton.id.replace(this.tabId + &#39;-&#39;, &#39;&#39;);\n            this.tabRepositionMarker(tabButton);\n        },\n        tabRepositionMarker(tabButton){\n            this.$refs.tabMarker.style.width=tabButton.offsetWidth + &#39;px&#39;;\n            this.$refs.tabMarker.style.height=tabButton.offsetHeight + &#39;px&#39;;\n            this.$refs.tabMarker.style.left=tabButton.offsetLeft + &#39;px&#39;;\n        },\n        tabContentActive(tabContent){\n            return this.tabSelected == tabContent.id.replace(this.tabId + &#39;-content-&#39;, &#39;&#39;);\n        }\n    }\" x-init=\"tabRepositionMarker($refs.tabButtons.firstElementChild);\" class=\"relative w-full max-w-sm\"><div x-ref=\"tabButtons\" class=\"relative inline-grid items-center justify-center w-full h-10 grid-cols-2 p-1 text-gray-500 bg-gray-100 rounded-lg select-none\"><button :id=\"$id(tabId)\" @click=\"tabButtonClicked($el);\" type=\"button\" class=\"relative z-20 inline-flex items-center justify-center w-full h-8 px-3 text-sm font-medium transition-all rounded-md cursor-pointer whitespace-nowrap\">Login</button> <button :id=\"$id(tabId)\" @click=\"tabButtonClicked($el);\" type=\"button\" class=\"relative z-20 inline-flex items-center justify-center w-full h-8 px-3 text-sm font-medium transition-all rounded-md cursor-pointer whitespace-nowrap\">Register</button><div x-ref=\"tabMarker\" class=\"absolute left-0 z-10 w-1/2 h-full duration-300 ease-out\" x-cloak><div class=\"w-full h-full bg-white rounded-md shadow-sm\"></div></div></div><div class=\"relative w-full mt-2 content\"><div :id=\"$id(tabId + &#39;-content&#39;)\" x-show=\"tabContentActive($el)\" class=\"relative\"><!-- Tab Content 1 - Replace with your content --><div class=\"border rounded-lg shadow-sm bg-card text-neutral-900\"><div class=\"flex flex-col space-y-1.5 p-6\"><h3 class=\"text-lg font-semibold leading-none tracking-tight\">Login</h3><p class=\"text-sm text-neutral-500\">Make changes to your account here. Click save when you're done.</p></div><div class=\"p-6 pt-0 space-y-2\"><div class=\"space-y-1\"><label class=\"text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"name\">Email</label><input type=\"text\" placeholder=\"Email\" id=\"name\" value=\"\" class=\"flex w-full h-10 px-3 py-2 text-sm bg-white border rounded-md peer border-neutral-300 ring-offset-background placeholder:text-neutral-400 focus:border-neutral-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed disabled:opacity-50\"></div><div class=\"space-y-1\"><label class=\"text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"username\">Password</label><input type=\"email\" placeholder=\"Password\" id=\"password\" value=\"\" class=\"flex w-full h-10 px-3 py-2 text-sm bg-white border rounded-md peer border-neutral-300 ring-offset-background placeholder:text-neutral-400 focus:border-neutral-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed disabled:opacity-50\"></div></div><div class=\"flex items-center p-6 pt-0\"><button type=\"button\" class=\"inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide text-white transition-colors duration-200 rounded-md bg-neutral-950 hover:bg-neutral-900 focus:ring-2 focus:ring-offset-2 focus:ring-neutral-900 focus:shadow-outline focus:outline-none\">Sign in</button></div></div><!-- End Tab Content 1 --></div><div :id=\"$id(tabId + &#39;-content&#39;)\" x-show=\"tabContentActive($el)\" class=\"relative\" x-cloak><!-- Tab Content 2 - Replace with your content --><div class=\"border rounded-lg shadow-sm bg-card text-neutral-900\"><div class=\"flex flex-col space-y-1.5 p-6\"><h3 class=\"text-lg font-semibold leading-none tracking-tight\">Register</h3><p class=\"text-sm text-neutral-500\">Change your password here. After saving, you'll be logged out.</p></div><div class=\"p-6 pt-0 space-y-2\"><div class=\"space-y-1\"><label class=\"text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"password\">Email</label><input type=\"email\" placeholder=\"Email\" id=\"password\" class=\"flex w-full h-10 px-3 py-2 text-sm bg-white border rounded-md peer border-neutral-300 ring-offset-background placeholder:text-neutral-400 focus:border-neutral-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed disabled:opacity-50\"></div><div class=\"space-y-1\"><label class=\"text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"password_new\">New Password</label><input type=\"password\" placeholder=\"New Password\" id=\"password_new\" class=\"flex w-full h-10 px-3 py-2 text-sm bg-white border rounded-md border-neutral-300 ring-offset-background placeholder:text-neutral-400 focus:border-neutral-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed disabled:opacity-50\"></div></div><div class=\"flex items-center p-6 pt-0\"><button type=\"button\" class=\"inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide text-white transition-colors duration-200 rounded-md bg-neutral-950 hover:bg-neutral-900 focus:ring-2 focus:ring-offset-2 focus:ring-neutral-900 focus:shadow-outline focus:outline-none\">Create account</button></div></div><!-- End Tab Content 2 --></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
