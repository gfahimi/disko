// +build linux

package linux

import "github.com/anuvu/disko"

// VolumeManager returns the linux implementation of disko.VolumeManager interface.
func VolumeManager() disko.VolumeManager {
	return &linuxLVM{}
}

type linuxLVM struct {
}

func (ls *linuxLVM) ScanPVs(filter disko.PVFilter) (disko.PVSet, error) {
	return nil, nil
}

func (ls *linuxLVM) ScanVGs(filter disko.VGFilter) (disko.VGSet, error) {
	return nil, nil
}

func (ls *linuxLVM) CreatePV(name string) (disko.PV, error) {
	return disko.PV{}, nil
}

func (ls *linuxLVM) DeletePV(pv disko.PV) error {
	return nil
}

func (ls *linuxLVM) HasPV(name string) bool {
	return false
}

func (ls *linuxLVM) CreateVG(name string, pvs ...disko.PV) (disko.VG, error) {
	return disko.VG{}, nil
}

func (ls *linuxLVM) ExtendVG(vgName string, pvs ...disko.PV) error {
	return nil
}

func (ls *linuxLVM) RemoveVG(vgName string) error {
	return nil
}

func (ls *linuxLVM) HasVG(vgName string) bool {
	return false
}

func (ls *linuxLVM) CryptFormat(lvName string, key string) error {
	return nil
}

func (ls *linuxLVM) CryptOpen(lvName string, key string) error {
	return nil
}

func (ls *linuxLVM) CryptClose(lvName string, key string) error {
	return nil
}

func (ls *linuxLVM) CreateLV(vgName string, name string, size uint64,
	lvType disko.LVType) (disko.LV, error) {
	return disko.LV{}, nil
}

func (ls *linuxLVM) RemoveLV(lvName string) error {
	return nil
}

func (ls *linuxLVM) ExtendLV(newSize uint64) error {
	return nil
}

func (ls *linuxLVM) HasLV(name string) bool {
	return false
}
