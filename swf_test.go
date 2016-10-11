package swfinfo

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	s, err := Open(`C:\Temp\vds_files\ae8ca7bf-3c16-43ac-9847-3f652a0bbf9f\media\2016\10\11\f8b0d38a-ed98-49c6-b422-44828776b5af_flash_flash.swf`)
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Printf("Compression: %s\n", s.Compression.String())
	fmt.Printf("Version: %d\n", s.Version)
	fmt.Printf("Frame size: %f %f %f %f\n", s.FrameSize.Xmin.Pixels(), s.FrameSize.Ymin.Pixels(), s.FrameSize.Xmax.Pixels(), s.FrameSize.Ymax.Pixels())
	fmt.Printf("Frame rate: %f\n", s.FrameRate)
	fmt.Printf("Frame count: %d\n", s.FrameCount)
}
