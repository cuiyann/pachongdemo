package main

import (
	"fmt"
	"net/http"
)

func main() {
	//urls := "https://www.zhenai.com/zhenghun"
	//engine.Run(engine.Request{urls, parser.ParseCityList})
	resp, _ := http.Get("https://album.zhenai.com/api/profile/getObjectProfile.do?objectID=1325911806&_=1661791740467&ua=h5%2F1.0.0%2F1%2F0%2F0%2F0%2F0%2F0%2F%2F0%2F0%2Fba437a96-96b9-4dd8-a896-b8c4d12b1e56%2F0%2F0%2F1199553966&data=eyJ2IjoiSEJOTWgrMTM3Ym9qMkJQL0NaZS93Zz09Iiwib3MiOiJ3ZWIiLCJpdCI6MzUwLCJ0IjoiYTFQUXdGZU4wemZ6UTNkUVJSYkh1RDRBSU03VllhOS9GMTNTWjhhZ29QNXVLVWlxcTFoaVladC8ydU5aRFJQTHlaUFZIVUNYUHFac0NBVE0xeHFmSFE9PSJ9&MmEwMD=5HX4__T7bxwVxuXM3JadDnZyOJJPoCasFEJIflSyMuIUICSEIkpz0ocvYzfY61bFN3fAPaBeANEzMYgB2b0uptafWQVtquRuu39rWTVQYMtI5B2h6mupNZfNZ3uC0xgE9JE01ibhwmDwpV9EzhqqnuEtddmky5g5fWz.jZzZleU_RlsBLWGwZbJOQac9ncKlBYdSg2xr_VKeulSbIVkTNWh2lmPXqHECRe.E2h8_ncrYbXskFS4wXDJAmsjT3QfUCPShqSbGKjhYksxvaxoOOxmj0oszHTEluVGwaTSwAlpYWWr19Yj_V2gvKvnI6Baaq")

	fmt.Println(resp)
	resp.Body.Close()

}
