package main

import "fmt"

const (
	CpuVendorIntel int8 = 1
	CpuVendorArm        = CpuVendorIntel << 1
	CpuVendorHygon      = CpuVendorArm << 1
)

func SupportCpuVendor(imageCpuVendor int8, cpuVendor int8) bool {
	return (imageCpuVendor & cpuVendor) == cpuVendor
}

func main() {
	fmt.Println(CpuVendorArm)   // 2
	fmt.Println(CpuVendorHygon) // 4

	fmt.Println(3 & 1)
	fmt.Println(SupportCpuVendor(CpuVendorIntel+CpuVendorArm+CpuVendorHygon, CpuVendorIntel))
	fmt.Println(SupportCpuVendor(CpuVendorIntel+CpuVendorArm, CpuVendorHygon))
}
