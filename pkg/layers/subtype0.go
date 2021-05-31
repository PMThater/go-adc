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

package layers

type Subtype0Header struct {
	DeviceSerial uint32
	UserDefBytes uint8
	EventNum uint32
}

//ms.DeviceSerial = binary.LittleEndian.Uint32(data[8:12])
//ms.UserDefBytes = data[12]
//ms.EventNum = binary.LittleEndian.Uint32(data[13:16])
