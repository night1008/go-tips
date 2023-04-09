package device_info

import (
	"io"
	"net/http"
	"os"

	"github.com/simpleframeworks/jobsd"
)

type DeviceInfoCrawler struct {
	httpClient *http.Client
}

func NewDeviceInfoCrawler() *DeviceInfoCrawler {
	return &DeviceInfoCrawler{
		httpClient: &http.Client{},
	}
}

func (c *DeviceInfoCrawler) createResourceJob(info jobsd.RunInfo, id uint) error {
	return nil
}

func (c *DeviceInfoCrawler) aa() error {
	resp, err := http.Get("https://browser.geekbench.com/v5/cpu/21020281")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := os.WriteFile("a.html", body, 0644); err != nil {
		return err
	}
	return nil
}
