package helpers

import qrcode "github.com/skip2/go-qrcode"

func QrCodeImageGenerator(data string, path_file string) error {

	err := qrcode.WriteFile(data, qrcode.Medium, 256, path_file)
	return err

}
