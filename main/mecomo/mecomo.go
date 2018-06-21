package mecomo

import (
	"fmt"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	url2 "net/url"
)

type GetDeviceResponse struct {
	Data struct {
		Device Device `xml:"RESULT"`
	} `xml:"urn:schemas-microsoft-com:xml-diffgram-v1 diffgram"`
}

type GetDevicesResponse struct {
	Data struct{
		Devices []DeviceHeader `xml:"RESULT>DEVICE"`
	} `xml:"urn:schemas-microsoft-com:xml-diffgram-v1 diffgram"`
}

type MecomoAPI struct {
	apiRoot string
	username string
	password string
}

const page_size  = 10

func CreateClient(apiUrl string, login string, pass string) MecomoAPI {
	return MecomoAPI{apiUrl, login, pass}
}

func (api* MecomoAPI) getDevicesPage (from int, size int) []DeviceHeader {
	var url, _ = url2.Parse(api.apiRoot + "/Devices.asmx/DevicesList")
	var params = url.Query()
	params.Add("startIndex", strconv.Itoa(from))
	params.Add("pageSize", strconv.Itoa(size))
	url.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(api.username, api.password)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return make([]DeviceHeader, 0)
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	result := GetDevicesResponse{}
	xml.Unmarshal(bytes, &result)
	return result.Data.Devices
}

func (api* MecomoAPI) getDevice(id int) (Device, error) {
	var url, _ = url2.Parse(api.apiRoot + "/Devices.asmx/DevicesGet")
	var params = url.Query()
	params.Add("deviceId", strconv.Itoa(id))
	url.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(api.username, api.password)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return Device{}, err
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	result := GetDeviceResponse{}
	xml.Unmarshal(bytes, &result)
	return result.Data.Device, nil
}

func (api* MecomoAPI) GetDevices(limit int) (result []Device) {
	from:=0
	for true {
		batch:= api.getDevicesPage(from, page_size)
		for _,id:= range batch {
			device,err := api.getDevice(id.Id)
			if err==nil {
				result = append(result, device)
			}
		}
		if len(batch) < page_size {
			return
		}
		fmt.Printf("%d devices were processed\n", len(batch))
		from+=len(batch)
		if limit>0 && len(result)>=limit {
			return
		}
	}
	return
}
