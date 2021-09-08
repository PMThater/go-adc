/*
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package command

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
	"jinr.ru/greenlab/go-adc/pkg/command/ifc"
	"jinr.ru/greenlab/go-adc/pkg/config"
	"jinr.ru/greenlab/go-adc/pkg/srv/control"
)

type ApiClient struct {
	*config.Config
	ApiPrefix string
}

func NewApiClient(cfg *config.Config) ifc.ApiClient {
	return &ApiClient{
		Config: cfg,
		ApiPrefix: fmt.Sprintf("http://%s:%d/api", cfg.IP, control.ApiPort),
	}
}

func (c *ApiClient) regReadUrl(device, addr string) string {
	return fmt.Sprintf("%s/reg/r/%s/%s", c.ApiPrefix, device, addr)
}

func (c *ApiClient) regWriteUrl(device string) string {
	return fmt.Sprintf("%s/reg/w/%s", c.ApiPrefix, device)
}

func (c *ApiClient) mstreamUrl(action, device string) string {
	return fmt.Sprintf("%s/mstream/%s/%s", c.ApiPrefix, action, device)
}

// RegRead sends request to get the value of a register of a device
func (c *ApiClient) RegRead(device, addr string) (string, error) {
	r, err := req.Get(c.regReadUrl(device, addr))
	if err != nil {
		return "", err
	}
	if r.Response().StatusCode != 200 {
		return "", errors.New(r.Response().Status)
	}
	reg := &control.RegHex{}
	err = r.ToJSON(reg)
	if err != nil {
		return "", err
	}
	return reg.Value, nil
}

// RegReadAll sends request to get values of all registers of a device
func (c *ApiClient) RegReadAll(device string) (map[string]string, error) {
	r, err := req.Get(c.regReadUrl(device, "all"))
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, errors.New(r.Response().Status)
	}
	var regs []*control.RegHex
	result := make(map[string]string)
	err = r.ToJSON(regs)
	if err != nil {
		return nil, err
	}
	for _, reg := range regs {
		result[reg.Addr] = reg.Value
	}
	return result, nil
}

// RegWrite sends request to write the value to a register of a device
func (c *ApiClient) RegWrite(device, addr, value string) error {
	reg := &control.RegHex{
		Addr: addr,
		Value: value,
	}
	r, err := req.Post(c.regWriteUrl(device), req.BodyJSON(reg))
	if err != nil {
		return err
	}
	if ! (r.Response().StatusCode != 200) {
		return errors.New(r.Response().Status)
	}
	return nil
}

// MStreamStart sends request to start streaming for a device
func (c *ApiClient) MStreamStart(device string) error {
	r, err := req.Get(c.mstreamUrl("start", device))
	if err != nil {
		return err
	}
	if ! (r.Response().StatusCode != 200) {
		return errors.New(r.Response().Status)
	}
	return nil
}

// MStreamStop sends request to stop streaming for a device
func (c *ApiClient) MStreamStop(device string) error {
	r, err := req.Get(c.mstreamUrl("stop", device))
	if err != nil {
		return err
	}
	if ! (r.Response().StatusCode != 200) {
		return errors.New(r.Response().Status)
	}
	return nil
}

// MStreamStartAll sends request to start streaming for all devices
func (c *ApiClient) MStreamStartAll() error {
	r, err := req.Get(c.mstreamUrl("start", ""))
	if err != nil {
		return err
	}
	if ! (r.Response().StatusCode != 200) {
		return errors.New(r.Response().Status)
	}
	return nil
}

// MStreamStop sends request to stop streaming for all devices
func (c *ApiClient) MStreamStopAll() error {
	r, err := req.Get(c.mstreamUrl("stop", ""))
	if err != nil {
		return err
	}
	if ! (r.Response().StatusCode != 200) {
		return errors.New(r.Response().Status)
	}
	return nil
}
