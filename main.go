package main

import (
	"encoding/json"
	"github.com/larryhou/appstoreconnect"
	"github.com/larryhou/appstoreconnect/auth"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	{
		Cwd, _ := os.Getwd()
		name := filepath.Join(Cwd, `access.json`)
		if f, err := os.Open(name); err == nil {
			defer f.Close()
			auth.Import(f)
		}
	}
	app := &appstoreconnect.AppStoreConnect{Id: `P4764AZ9HG`}
	{
		req := new(appstoreconnect.BundleIdCreateRequest)
		req.Data.Attributes = appstoreconnect.BundleIdAttributes{
			Identifier: `com.larryhou.apitest`,
			Name:       `apitest`,
			Platform:   appstoreconnect.BundleIdPlatformIOS,
		}
		rsp, err := app.BundleIdRegister(req)
		log.Printf(`BUNDLEID %+v %+v`, rsp, err)
	}

	//{
	//	rsp, err := app.AppList(nil)
	//	log.Printf(`APP %+v %+v`, rsp, err)
	//	j := json.NewEncoder(os.Stdout)
	//	j.SetIndent("", "    ")
	//	j.Encode(rsp)
	//	return
	//}

	p := regexp.MustCompile(`com.larryhou.g\d+`)

	{
		params := make(url.Values)
		//params.Set(`limit`, strconv.Itoa(100))
		rsp, err := app.BundleIdList(params)
		//log.Printf(`%+v %+v`, rsp, err)
		//if err == nil {
		//	for _, b := range rsp.Data {
		//		log.Printf(`%#v`, b)
		//	}
		//}
		if err == nil {
			var g sync.WaitGroup
			for v := range rsp.Paginate(app) {
				log.Printf(`%+v`, v)
				for _, d := range v.Data {
					if p.MatchString(d.Attributes.Identifier) {
						g.Add(1)
						go func(id string) {
							err := app.BundleIdDelete(id)
							log.Printf(`%s %+v`, id, err)
							g.Done()
						}(d.Id)
					}
				}
			}

			g.Wait()
		}
	}

	{
		rsp, err := app.DeviceList(nil)
		log.Printf(`%+v %+v`, rsp, err)
		json.NewEncoder(os.Stdout).Encode(rsp)
	}
}
