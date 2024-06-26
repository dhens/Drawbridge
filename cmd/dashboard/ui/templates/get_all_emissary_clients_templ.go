// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "dhens/drawbridge/cmd/drawbridge/emissary"
import "fmt"

func GetAllEmissaryClients(clients []*emissary.EmissaryClient, latestClientEvents map[string]emissary.Event) templ.Component {
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
		if len(clients) == 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p>No Fleet Devices created yet. Create an Emissary Bundle to start!</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul id=\"device-fleet-list\" hx-get=\"/emissary/get/clients\" hx-trigger=\"every 10s\" hx-target=\"this\" hx-swap=\"outerHTML\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, client := range clients {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li id=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("fleet-device-%s", client.ID))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/dashboard/ui/templates/get_all_emissary_clients.templ`, Line: 12, Col: 61}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"fleet-device\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if client.Revoked == 1 {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var3 string
					templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(client.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/dashboard/ui/templates/get_all_emissary_clients.templ`, Line: 14, Col: 39}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" (Revoked)</span> <svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\" width=\"36\" height=\"36\" fill=\"currentColor\"><path d=\"M13 18V20H17V22H7V20H11V18H2.9918C2.44405 18 2 17.5511 2 16.9925V4.00748C2 3.45107 2.45531 3 2.9918 3H21.0082C21.556 3 22 3.44892 22 4.00748V16.9925C22 17.5489 21.5447 18 21.0082 18H13Z\"></path></svg> ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					if latestClientEvents[client.ID].Timestamp == "" {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>Last Seen: Never</span> <span>IP Address: N/A</span>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					} else {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>Last Seen: ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var4 string
						templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(latestClientEvents[client.ID].Timestamp)
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/dashboard/ui/templates/get_all_emissary_clients.templ`, Line: 20, Col: 78}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span>IP Address: ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var5 string
						templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(latestClientEvents[client.ID].ConnectionIP)
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/dashboard/ui/templates/get_all_emissary_clients.templ`, Line: 21, Col: 82}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <button value=\"Restore Access\" hx-post=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var6 string
					templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("emissary/post/client/%s/unrevoke_certificate", client.ID))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/dashboard/ui/templates/get_all_emissary_clients.templ`, Line: 23, Col: 131}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var7 string
					templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("#fleet-device-%s", client.ID))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/dashboard/ui/templates/get_all_emissary_clients.templ`, Line: 23, Col: 187}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"outerHTML\" class=\"emissary-restore-btn\">Restore Access</button>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var8 string
					templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(client.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/dashboard/ui/templates/get_all_emissary_clients.templ`, Line: 25, Col: 39}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\" width=\"36\" height=\"36\" fill=\"currentColor\"><path d=\"M13 18V20H17V22H7V20H11V18H2.9918C2.44405 18 2 17.5511 2 16.9925V4.00748C2 3.45107 2.45531 3 2.9918 3H21.0082C21.556 3 22 3.44892 22 4.00748V16.9925C22 17.5489 21.5447 18 21.0082 18H13Z\"></path></svg> ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					if latestClientEvents[client.ID].Timestamp == "" {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>Last Seen: Never</span> <span>IP Address: N/A</span>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					} else {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>Last Seen: ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var9 string
						templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(latestClientEvents[client.ID].Timestamp)
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/dashboard/ui/templates/get_all_emissary_clients.templ`, Line: 31, Col: 78}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span>IP Address: ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var10 string
						templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(latestClientEvents[client.ID].ConnectionIP)
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/dashboard/ui/templates/get_all_emissary_clients.templ`, Line: 32, Col: 82}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <button value=\"Revoke Access\" hx-post=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var11 string
					templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("emissary/post/client/%s/revoke_certificate", client.ID))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/dashboard/ui/templates/get_all_emissary_clients.templ`, Line: 34, Col: 128}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var12 string
					templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("#fleet-device-%s", client.ID))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/dashboard/ui/templates/get_all_emissary_clients.templ`, Line: 34, Col: 184}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"outerHTML\" class=\"emissary-revoke-btn\">Revoke Access</button>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
